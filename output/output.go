package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"runtime/debug"
	"sort"

	"time"

	"github.com/certinia/asist/config"
	"github.com/certinia/asist/debugger"
	"github.com/certinia/asist/errorhandler"
	"github.com/certinia/asist/finding"
	"github.com/certinia/asist/message"
	"github.com/certinia/asist/parser/options"
	"github.com/certinia/asist/rules"
	"github.com/certinia/asist/ruleset"
)

type ScanTime struct {
	StartedTime string
	EndingTime  string
}

func ListRules(ruleInstances []*rules.Rule) {
	allRulesMetadata := []*rules.RuleMetadata{}
	if !options.IsListRules() {
		return
	}
	for _, rule := range ruleInstances {
		metadata := (*rule).GetMetadata()
		allRulesMetadata = append(allRulesMetadata, metadata)
	}

	fmt.Printf("%s\n", PrettyPrintJSON(allRulesMetadata))

	debugger.Debug("listed rules")
	os.Exit(int(errorhandler.ExitCodeSuccess))
}

/**
 * countFindingsPerRule counts the findings of each rule.
 */
func countFindingsPerRule(finalResult *finding.Output) map[rules.RuleID]int {
	findingsPerRule := make(map[rules.RuleID]int)
	for _, finding := range finalResult.Results {
		findingsPerRule[finding.ID]++
	}
	return findingsPerRule
}

/**
 * getViolatedRuleIds returns the IDs of rules whose finding count exceeds their
 * configured cicdmaxissues, sorted for deterministic output.
 * Default cicdmaxissues is 0 (no issues allowed).
 */
func getViolatedRuleIds(findingsPerRule map[rules.RuleID]int, configFile *config.Config) []rules.RuleID {
	violatedRuleIds := []rules.RuleID{}
	for ruleId, count := range findingsPerRule {
		if count > configFile.GetRuleCicdMaxIssues(ruleId) {
			violatedRuleIds = append(violatedRuleIds, ruleId)
		}
	}
	sort.Slice(violatedRuleIds, func(i, j int) bool {
		return string(violatedRuleIds[i]) < string(violatedRuleIds[j])
	})
	return violatedRuleIds
}

/**
 * filterFindingsByRules removes findings that do not belong to the given rules
 * and updates the result count accordingly.
 */
func filterFindingsByRules(finalResult *finding.Output, ruleIds []rules.RuleID) {
	includedRuleIds := make(map[rules.RuleID]bool, len(ruleIds))
	for _, ruleId := range ruleIds {
		includedRuleIds[ruleId] = true
	}

	filteredResults := []finding.Finding{}
	for _, finding := range finalResult.Results {
		if includedRuleIds[finding.ID] {
			filteredResults = append(filteredResults, finding)
		}
	}
	finalResult.Results = filteredResults
	finalResult.Count = len(filteredResults)
}

/**
 * CheckThresholdViolations checks if any rule exceeds its configured cicdmaxissues.
 * Default cicdmaxissues is 0 (no issues allowed).
 * Returns true if any threshold is violated, false otherwise.
 */
func CheckThresholdViolations(w io.Writer, finalResult *finding.Output, configFile *config.Config) bool {
	findingsPerRule := countFindingsPerRule(finalResult)
	violatedRuleIds := getViolatedRuleIds(findingsPerRule, configFile)

	if len(violatedRuleIds) == 0 {
		return false
	}

	fmt.Fprintf(w, "\n%s\n", message.GetThresholdViolationHeader())
	for _, ruleId := range violatedRuleIds {
		fmt.Fprintf(w, "%s\n", message.GetThresholdViolation(string(ruleId), findingsPerRule[ruleId], configFile.GetRuleCicdMaxIssues(ruleId)))
	}
	fmt.Fprintf(w, "\n%s\n", message.GetThresholdViolationSummary(len(violatedRuleIds)))

	return true
}

/**
 * DisplayOutput - method used to display the output of scans by type
 */
func DisplayOutput(finalResult *finding.Output, scanTime *ScanTime) {
	if options.IsBaselineScan() {
		debugger.Debug("writing baseline output")
		baselineScanOutput := createBaselineOutput(finalResult, options.GetRepoURL())
		displayOutput(baselineScanOutput)
	} else {
		debugger.Debug("writing regular output")
		scanTime.EndingTime = time.Now().String()
		finalResult.ScanStartedTime = scanTime.StartedTime
		finalResult.ScanEndingTime = scanTime.EndingTime
		finalResult.Count = len(finalResult.Results)

		if options.IsCICDScan() {
			configFile := config.GetConfigInstance()
			violatedRuleIds := getViolatedRuleIds(countFindingsPerRule(finalResult), configFile)

			if len(violatedRuleIds) > 0 {
				filterFindingsByRules(finalResult, violatedRuleIds)
				displayOutput(finalResult)
				CheckThresholdViolations(os.Stdout, finalResult, configFile)
				os.Exit(int(errorhandler.ExitCodeOccurrence))
			}

			displayOutput(finalResult)
			fmt.Printf("\n%s\n", message.GetNoThresholdViolationSummary())
			return
		}

		displayOutput(finalResult)
	}
}

/**
 * extractRepoNameFromURL - method used to extract repoName from a sshUrl of repository
 */
func extractRepoNameFromURL(url string) string {
	if url != "" {
		findRepoNameRegexp := regexp.MustCompile(`(?i)/([a-z0-9-_.]+)\.git$`)
		match := findRepoNameRegexp.FindStringSubmatch(url)
		if len(match) > 1 {
			return match[1]
		}
	}
	return ""
}

/**
 * displayOutput - method used to display the output of scans
 */
func displayOutput(finalResult interface{}) {
	var jsonOutput []byte
	var err error
	if options.IsBaselineScan() {
		jsonOutput, err = json.Marshal(finalResult)
	} else {
		jsonOutput, err = json.MarshalIndent(finalResult, "", " ")
	}
	if err != nil {
		errorhandler.ExitWithCode(message.GetMarshallingOutputError(err), errorhandler.ExitCodeInternalError)
	}
	fmt.Println(string(jsonOutput))
}

/**
 * createBaselineOutput - method used to create the output for baseline scan
 */
func createBaselineOutput(finalResultList *finding.Output, repositoryURL string) []finding.BaselineOutput {
	var baselineOutputContent finding.BaselineOutputContent
	var baselineOutput []finding.BaselineOutput

	repositoryName := extractRepoNameFromURL(repositoryURL)

	for _, result := range finalResultList.Results {
		isCustom := !*ruleset.IsStandardRuleID(result.ID)
		baselineOutputContent = finding.BaselineOutputContent{
			FindingID:       result.CreateFindingID(),
			IsCustom:        isCustom,
			IsFalsePositive: result.Occurrence.IsFalsePositive,
			Id:              result.ID,
			Severity:        result.Severity,
			RuleCategory:    result.RuleCategory,
		}
		baselineOutput = append(baselineOutput, finding.BaselineOutput{
			RepositoryName: repositoryName,
			RepositoryURL:  repositoryURL,
			RecordType:     finding.BaselineFinding,
			Content:        baselineOutputContent,
		})
	}

	// Either add config file content if exists for repository otherwise add blank
	baselineOutput = append(baselineOutput, finding.BaselineOutput{
		RepositoryName: repositoryName,
		RepositoryURL:  repositoryURL,
		RecordType:     finding.BaselineConfig,
		Content:        config.GetConfigInstance(),
	},
	)
	return baselineOutput
}

/**
 * PrettyPrintJSON - method used to prettify the json of Rules metadata
 */
func PrettyPrintJSON(results interface{}) []byte {
	resultsJSON, err := json.MarshalIndent(results, "", "\t")
	if err != nil {
		errorhandler.ExitWithCode(message.GetMarshalIndentError(err), errorhandler.ExitCodeInternalError)
	}
	return resultsJSON
}

/**
 * DisplayVersion - method used to display the version of ASIST binary
 */
func DisplayVersion(version string) {
	if !options.IsVersion() {
		return
	}

	if version != "" {
		fmt.Printf("%s\n", version)
	} else {
		buildInfo, ok := debug.ReadBuildInfo()
		if !ok {
			fmt.Println("Version information not available.")
		}
		if buildInfo.Main.Version != "" {
			fmt.Printf("%s\n", buildInfo.Main.Version)
		} else {
			fmt.Println("unknown")
		}
	}
	os.Exit(int(errorhandler.ExitCodeSuccess))
}
