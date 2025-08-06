package codequality

import (
	"regexp"

	"github.com/certinia/asist/files"
	"github.com/certinia/asist/rules"
)

var DetectMissingAccessibilityModifierRuleID rules.RuleID = "DetectMissingAccessibilityModifier"

type DetectMissingAccessibilityModifierRule struct {
	metadata rules.RuleMetadata
}

func NewDetectMissingAccessibilityModifierRule() *DetectMissingAccessibilityModifierRule {
	return &DetectMissingAccessibilityModifierRule{
		metadata: rules.RuleMetadata{
			ID:             DetectMissingAccessibilityModifierRuleID,
			Name:           "Missing Accessibility Modifier in methods and constructors",
			Description:    "Missing accessibility modifiers in methods and constructors in apex classes may reduce the clarity and overall quality of your code and lead to security issues. Ensure that all methods have an explicit accessibility modifier.",
			Severity:       rules.SeverityLow,
			RuleCategory:   rules.CategoryCodeQuality,
			IncludePattern: "\\.cls$",
			ExcludePattern: "(?i)/(force-app-autotest|autotest|systemtest|test)(s)?/|Test\\.cls$",
			Pattern:        "(?i)\\b(public|protected|private|global|webService)\\b",
		},
	}
}

func (r *DetectMissingAccessibilityModifierRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *DetectMissingAccessibilityModifierRule) Run(file files.File) []rules.Occurrence {
	return findMatchesForDetectMissingAccessibilityModifier(file, &r.metadata)
}

func findMatchesForDetectMissingAccessibilityModifier(fileToScan files.File, ruleMetadata *rules.RuleMetadata) []rules.Occurrence {
	var output []rules.Occurrence
	accessibilityModifiersRegexp := regexp.MustCompile(ruleMetadata.Pattern)

	for _, functionInfo := range getFunctionsList(string(ruleMetadata.ID), fileToScan) {
		if !(accessibilityModifiersRegexp.MatchString(fileToScan.Lines[functionInfo.lineIndex].Text)) {
			lineNumber := fileToScan.Lines[functionInfo.lineIndex].LineNumber
			output = append(
				output,
				rules.Occurrence{
					FileName:        fileToScan.FileName,
					LineNumber:      lineNumber,
					LineContent:     fileToScan.Lines[functionInfo.lineIndex].Text,
					ColumnRange:     []int{0, len(fileToScan.Lines[functionInfo.lineIndex].Text) - 1},
					IsFalsePositive: fileToScan.IsLineMarkedFalsePositive(string(ruleMetadata.ID), lineNumber),
				},
			)
		}
	}
	return output
}
