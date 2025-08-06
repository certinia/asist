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
