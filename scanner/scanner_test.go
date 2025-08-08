package scanner

import (
	"reflect"
	"testing"

	"github.com/certinia/asist/finding"
	"github.com/certinia/asist/rules"
	testrule "github.com/certinia/asist/scanner/testData"
)

func TestRunRulesOnFiles_WhenFileExistAndRulesIsValidToRunOnFile_ReturnsFindings(t *testing.T) {
	//Given
	filepaths := []string{"./testData/testFile.cls"}
	ruleInstances := []*rules.Rule{}

	testRuleMetadata := rules.RuleMetadata{
		ID:             "testId",
		Name:           "test",
		Description:    "test description",
		Severity:       rules.Severity("Medium"),
		RuleCategory:   rules.RuleCategory("Security"),
		IncludePattern: "\\.cls$",
		ExcludePattern: "Test\\.cls$",
		Pattern:        "^(?i)\\s*(public|private|global)?\\s*(abstract|virtual)?\\s*\\bclass\\b\\s+\\w+",
	}
	testRule := testrule.NewTestRule(testRuleMetadata)
	mockData := []rules.Occurrence{{
		FileName:    "./testData/testFile.cls",
		LineContent: "public without sharing class Testclass {",
		ColumnRange: []int{0, 38},
		LineNumber:  1},
	}
	testrule.SetMockData(mockData)
	ruleInstances = append(ruleInstances, &testRule)

	expectedResult := []finding.Finding{
		{
			ID:           "testId",
			Name:         testRuleMetadata.Name,
			Description:  testRuleMetadata.Description,
			Severity:     testRuleMetadata.Severity,
			RuleCategory: testRuleMetadata.RuleCategory,
			Occurrence:   mockData[0],
		},
	}

	//When
	actualResult, err := RunRulesOnFiles(filepaths, ruleInstances)

	//Then
	if actualResult.Count != len(expectedResult) {
		t.Errorf("Occurrences count mismatched.\n Actual %v, Expected %v", actualResult.Count, len(expectedResult))
	}
	if !reflect.DeepEqual(actualResult.Results, expectedResult) {
		t.Errorf("Findings results are not equal.\n Actual %v, Expected %v", actualResult.Results, expectedResult)
	}
	if err != nil {
		t.Errorf("RunRuleOnFiles method should not return error!")
	}
}

func TestRunRulesOnFiles_WhenFileExistAndRulesIsInValidToRunOnFile_ReturnsEmptyResult(t *testing.T) {
	//Given
	filepaths := []string{"./testData/testFile.cls"}
	ruleInstances := []*rules.Rule{}

	testRuleMetadata := rules.RuleMetadata{
		ID:             "testId",
		Name:           "test",
		Description:    "test description",
		Severity:       rules.Severity("Medium"),
		RuleCategory:   rules.RuleCategory("Security"),
		IncludePattern: "\\.js$",
		ExcludePattern: "min\\.js$",
		Pattern:        "^(?i)\\s*(public|private|global)?\\s*(abstract|virtual)?\\s*\\bclass\\b\\s+\\w+",
	}
	testRule := testrule.NewTestRule(testRuleMetadata)
	mockData := []rules.Occurrence{}
	testrule.SetMockData(mockData)
	ruleInstances = append(ruleInstances, &testRule)

	expectedResultCount := 0

	//When
	actualResult, err := RunRulesOnFiles(filepaths, ruleInstances)

	//Then
	if len((*actualResult).Results) != expectedResultCount {
		t.Errorf("Occurrences count mismatched.\n Actual %v, Expected %v", len((*actualResult).Results), expectedResultCount)
	}
	if err != nil {
		t.Errorf("RunRuleOnFiles method should not return error!")
	}
}
