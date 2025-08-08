package customrule

import (
	"reflect"
	"testing"

	"github.com/certinia/asist/config"
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/rules"
)

func TestNewCustomRule(t *testing.T) {
	// Given
	customRule := config.CustomRegexRule{
		Name:           "CustomRule1",
		Description:    "This is sample description for the custom rule",
		Severity:       "High",
		RuleCategory:   "Security",
		IncludePattern: ".*.cls",
		ExcludePattern: ".*test.cls$",
		Pattern:        "pattern",
	}

	customRuleInstance := rules.RuleMetadata{
		ID:             "custom_rule_1",
		Name:           customRule.Name,
		Description:    customRule.Description,
		Severity:       rules.Severity(customRule.Severity),
		RuleCategory:   rules.RuleCategory(customRule.RuleCategory),
		IncludePattern: customRule.IncludePattern,
		ExcludePattern: customRule.ExcludePattern,
		Pattern:        customRule.Pattern,
	}

	expectedcustomRuleInstance := CustomRule{
		metadata: customRuleInstance,
	}

	// When
	actualCustomRuleInstance := *NewCustomRule(customRule, "custom_rule_1")

	// Then
	if !reflect.DeepEqual(actualCustomRuleInstance, expectedcustomRuleInstance) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "CustomRule is mismatched!", actualCustomRuleInstance, expectedcustomRuleInstance)
	}
}

func TestRun(t *testing.T) {
	// Given
	customRule := config.CustomRegexRule{
		Name:           "CustomRule1",
		Severity:       "High",
		RuleCategory:   "Security",
		IncludePattern: ".*.cls",
		ExcludePattern: ".*test.cls$",
		Pattern:        "pattern",
	}

	mockLines := []files.Line{
		{LineNumber: 1, Text: "This is a test file with pattern matching the rules.", IsCommentedLine: false},
		{LineNumber: 2, Text: "This line should be ignored by the rules.", IsCommentedLine: false},
	}

	mockFile := files.File{
		Lines:    mockLines,
		FileName: "testfile.cls",
		IgnoresSelected: []files.IgnoreSelected{
			{
				BeginLine: 2,
				EndLine:   2,
				RuleIDs: map[string]bool{
					"customrule1": true,
				},
			},
		},
	}

	expectedOccurrenceResult := rules.Occurrence{
		FileName:    "testfile.cls",
		LineNumber:  1,
		ColumnRange: []int{25, 32},
		LineContent: "This is a test file with pattern matching the rules.",
	}
	customRuleInstance := NewCustomRule(customRule, "customrule1")

	// When
	actualOccurrenceResult := customRuleInstance.Run(mockFile)

	// Then
	if !reflect.DeepEqual(actualOccurrenceResult[0], expectedOccurrenceResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Occurrences are mismatched!", actualOccurrenceResult[0], expectedOccurrenceResult)
	}
}
