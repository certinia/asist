package rules

import (
	"reflect"
	"testing"
)

func TestOverride(t *testing.T) {
	// Given
	standardRuleMetadata := RuleMetadata{
		ID:             "SampleRule1",
		Name:           "Potential security failure due to misconfiguration",
		Description:    "This is sample description of rule",
		Severity:       SeverityHigh,
		IncludePattern: ".*\\.cls$",
		ExcludePattern: ".*test.*",
		Pattern:        "public\\s+class\\s+\\w+",
		RuleCategory:   CategorySecurity,
		Qualifier:      "Qualifier",
	}
	ENABLED_TRUE := true
	override := RuleMetadataOverride{
		Severity:       "Critical",
		ExcludePattern: ".*IntegrationTest.cls$",
		Enabled:        &ENABLED_TRUE,
	}

	expectedRuleMetadata := RuleMetadata{
		ID:             "SampleRule1",
		Name:           "Potential security failure due to misconfiguration",
		Description:    "This is sample description of rule",
		Severity:       "Critical",
		IncludePattern: ".*\\.cls$",
		ExcludePattern: ".*IntegrationTest.cls$",
		Pattern:        "public\\s+class\\s+\\w+",
		RuleCategory:   CategorySecurity,
		Qualifier:      "Qualifier",
	}

	// When
	standardRuleMetadata.Override(override, false)

	// Then
	if !reflect.DeepEqual(standardRuleMetadata, expectedRuleMetadata) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Standard rule metadata should match expected metadata!", standardRuleMetadata, expectedRuleMetadata)
	}
}

func TestOverride_WhenCicdMaxIssuesSet_PopulatesRuleMetadata(t *testing.T) {
	// Given
	standardRuleMetadata := RuleMetadata{
		ID:       "SampleRule1",
		Severity: SeverityHigh,
	}
	maxIssues := 42
	override := RuleMetadataOverride{
		CicdMaxIssues: &maxIssues,
	}

	// When
	standardRuleMetadata.Override(override, false)

	// Then
	if standardRuleMetadata.CicdMaxIssues != 42 {
		t.Errorf("Expected CicdMaxIssues to be 42, got %d", standardRuleMetadata.CicdMaxIssues)
	}
}

func TestOverride_WhenCicdMaxIssuesNotSet_DefaultsToZero(t *testing.T) {
	// Given
	standardRuleMetadata := RuleMetadata{
		ID:       "SampleRule1",
		Severity: SeverityHigh,
	}
	override := RuleMetadataOverride{
		Severity: "Critical",
	}

	// When
	standardRuleMetadata.Override(override, false)

	// Then
	if standardRuleMetadata.CicdMaxIssues != 0 {
		t.Errorf("Expected CicdMaxIssues to default to 0, got %d", standardRuleMetadata.CicdMaxIssues)
	}
}

func TestCreateHashableString(t *testing.T) {
	// Given
	occurrence := Occurrence{
		FileName:        "testFile.cls",
		LineContent:     "public class(){",
		LineNumber:      10,
		ColumnRange:     []int{1, 16},
		IsFalsePositive: false,
	}

	// When
	hashableString := occurrence.CreateHashableString()

	// Then
	if !(hashableString != "") {
		t.Errorf("Hashable string should not be empty!")
	}
}
