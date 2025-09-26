package output

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"runtime/debug"
	"strings"
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
		displayOutput(finalResult)

		//Returns exit 1 in case of the CICD rules defined in the configFile
		if finalResult.Count > 0 && options.IsCICDScan() {
			os.Exit(int(errorhandler.ExitCodeOccurrence))
		}
	}
}

/**
 * extractRepoNameFromURL - method used to extract repoFolder/repoName from a sshUrl of repository
 */
func extractRepoNameFromURL(url string) string {
	if url != "" {
		findRepoNameRegexp := regexp.MustCompile(`[a-z0-9-]+/[a-z0-9-_]+\.git$`)
		match := findRepoNameRegexp.FindStringSubmatch(url)
		if len(match) > 0 {
			parts := strings.Split(match[0], ".")
			if len(parts) > 1 {
				return parts[0]
			}
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
