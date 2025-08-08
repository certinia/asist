package security

import (
	"regexp"

	"github.com/certinia/asist/files"
	"github.com/certinia/asist/parser/options"
	"github.com/certinia/asist/rules"
)

var SensitiveInfoInDebugRuleID rules.RuleID = "SensitiveInfoInDebug"

type SensitiveInfoInDebugRule struct {
	metadata rules.RuleMetadata
}

func NewSensitiveInfoInDebugRule() *SensitiveInfoInDebugRule {
	return &SensitiveInfoInDebugRule{
		metadata: rules.RuleMetadata{
			ID:             SensitiveInfoInDebugRuleID,
			Name:           "Potential info in debug",
			Description:    "Use of sensitive keywords in debug logs is vulnerable to showing sensitive information to the inappropriate user. To get a passed Salesforce Security Review, all the detailed logging for development purposes should be pared down appropriately.",
			Severity:       rules.SeverityLow,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.cls$",
			ExcludePattern: "",
			Pattern:        "(?i)system\\.debug\\(",
		},
	}
}

func (r *SensitiveInfoInDebugRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *SensitiveInfoInDebugRule) Run(fileToScan files.File) []rules.Occurrence {
	result := findSensitiveInfoInDebug(fileToScan, string(SensitiveInfoInDebugRuleID), r.metadata.Pattern)
	return result
}

func findSensitiveInfoInDebug(fileToScan files.File, currentRuleID string, pattern string) []rules.Occurrence {
	var output []rules.Occurrence
	debugMethodRegexp := regexp.MustCompile(pattern)
	findSystemDebugClosingBracket := regexp.MustCompile(`\);`)
	hasDebugLineFound := false
	isCommentedLinesIncluded := false
	sensitiveKeywordsRegex := regexp.MustCompile(`[+\(]+\s*(?i)\w*\.?(secret|credential|phone|email|address|income|gender|ethinicity|education|password|credit\w?card|session|api\w?key|token|user\w?name|account|opportunity|authorization|authentication)\w*[+]?`)
	systemDebugActualLine := ""
	systemDebugColumnRange := []int{-1, -1}
	systemDebugLineNumber := -1
	systemDebugVirtualLine := ""
	for _, line := range fileToScan.Lines {
		isFalsePositive := fileToScan.IsLineMarkedFalsePositive(currentRuleID, line.LineNumber)
		if (!isCommentedLinesIncluded && line.IsCommentedLine) || (isFalsePositive && !options.IsBaselineScan()) {
			continue
		}
		isDebugLine := debugMethodRegexp.MatchString(line.Text)
		hasDebugLineFound = isDebugLine || hasDebugLineFound

		if hasDebugLineFound {
			isDebugClosing := findSystemDebugClosingBracket.MatchString(line.Text)

			if isDebugLine {
				systemDebugLineNumber = line.LineNumber
				systemDebugActualLine = line.Text
				systemDebugColumnRange = debugMethodRegexp.FindStringIndex(line.Text)
			}

			systemDebugVirtualLine += line.Text
			if isDebugClosing {
				isSensitiveKeywordInDebug := sensitiveKeywordsRegex.MatchString(systemDebugVirtualLine)
				systemDebugVirtualLine = ""
				if isSensitiveKeywordInDebug {
					output = append(
						output,
						rules.Occurrence{
							FileName:        fileToScan.FileName,
							LineNumber:      systemDebugLineNumber,
							LineContent:     systemDebugActualLine,
							ColumnRange:     systemDebugColumnRange,
							IsFalsePositive: isFalsePositive,
						},
					)
				}
			}
			hasDebugLineFound = !isDebugClosing
		}
	}
	return output
}
