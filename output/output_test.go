package output

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	"github.com/certinia/asist/config"
	"github.com/certinia/asist/finding"
	"github.com/certinia/asist/rules"
)

func TestCreateBaselineOutput(t *testing.T) {
	//Given
	findingResult := []finding.Finding{
		{
			ID:           "ExposedMessageChannel",
			Name:         "Potential security failure due to misconfiguration of LMS",
			Severity:     "High",
			RuleCategory: "Security",
			Occurrence: rules.Occurrence{
				FileName:    "/src/metadata/projectMessageChannel.messageChannel-meta.xml",
				ColumnRange: []int{1, 28},
				LineNumber:  4,
			},
		},
	}
	testResult := finding.Output{
		Count:           1,
		ScanStartedTime: "0",
		ScanEndingTime:  "2",
		Results:         findingResult,
	}
	sshUrl := "ssh://git@example.com:12345/test/test.git"
	baselineOutputContent := finding.BaselineOutputContent{FindingID: "431d4935962ceb80271200", IsCustom: false, IsFalsePositive: false, Id: "ExposedMessageChannel", Severity: "High", RuleCategory: "Security"}
	expectedBaselineOutput := []finding.BaselineOutput{
		{
			RepositoryName: "test",
			RepositoryURL:  sshUrl,
			RecordType:     "Finding",
			Content:        baselineOutputContent,
		},
		{
			RepositoryName: "test",
			RepositoryURL:  sshUrl,
			RecordType:     "Config",
			Content:        config.GetConfigInstance(),
		},
	}

	//When
	actualBaselineOutput := createBaselineOutput(&testResult, sshUrl)

	//Then
	if !reflect.DeepEqual(actualBaselineOutput, expectedBaselineOutput) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Baseline output is mismatched!", actualBaselineOutput, expectedBaselineOutput)
	}
}

func TestPrettyPrintJSON(t *testing.T) {
	//Given
	var actualResult interface{}
	standardRuleMetadata := []rules.RuleMetadata{{
		ID:             "SampleRule1",
		Name:           "Potential security failure due to misconfiguration",
		Description:    "This is sample description of rule",
		Severity:       "High",
		IncludePattern: ".*\\.cls$",
		ExcludePattern: ".*test.*",
		Pattern:        "public\\s+class\\s+\\w+",
		RuleCategory:   "Security",
		Qualifier:      "Qualifier",
	}}

	//When
	actualResult = PrettyPrintJSON(standardRuleMetadata)

	//Then
	_, isByteType := actualResult.([]byte)
	if !isByteType {
		t.Errorf("Actual result should be of type []byte!")
	}
}

func TestCreateBaselineOutput_WhenInputUrlNil_ReturnsOutputWithoutRepoNameAndUrl(t *testing.T) {
	//Given
	findingResult := []finding.Finding{
		{ID: "ExposedMessageChannel", Name: "Potential security failure due to misconfiguration of LMS", Severity: "High", RuleCategory: "Security", Occurrence: rules.Occurrence{FileName: "/src/metadata/projectMessageChannel.messageChannel-meta.xml", ColumnRange: []int{1, 28}, LineNumber: 4}},
	}
	testResult := finding.Output{
		Count:           1,
		ScanStartedTime: "0",
		ScanEndingTime:  "2",
		Results:         findingResult,
	}
	sshUrl := ""
	baselineOutputContent := finding.BaselineOutputContent{FindingID: "431d4935962ceb80271200", IsCustom: false, IsFalsePositive: false, Id: "ExposedMessageChannel", Severity: "High", RuleCategory: "Security"}
	expectedBaselineOutput := []finding.BaselineOutput{
		{
			RecordType: "Finding",
			Content:    baselineOutputContent,
		},
		{
			RecordType: "Config",
			Content:    config.GetConfigInstance(),
		},
	}

	//When
	actualBaselineOutput := createBaselineOutput(&testResult, sshUrl)

	//Then
	if !reflect.DeepEqual(actualBaselineOutput, expectedBaselineOutput) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Baseline output is mismatched!", actualBaselineOutput, expectedBaselineOutput)
	}
}

// Helper to create an int pointer
func intPtr(v int) *int {
	return &v
}

func TestCheckThresholdViolations_WhenNoFindings_ReturnsFalse(t *testing.T) {
	//Given
	finalResult := &finding.Output{
		Results: []finding.Finding{},
	}
	configFile := &config.Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{},
	}
	var buf bytes.Buffer

	//When
	result := CheckThresholdViolations(&buf, finalResult, configFile)

	//Then
	if result {
		t.Errorf("Expected no violation when there are no findings")
	}
	if buf.Len() != 0 {
		t.Errorf("Expected no output when there are no findings, got: %s", buf.String())
	}
}

func TestCheckThresholdViolations_WhenWithinThreshold_ReturnsFalse(t *testing.T) {
	//Given
	finalResult := &finding.Output{
		Results: []finding.Finding{
			{ID: "XSSTooltip"},
			{ID: "XSSTooltip"},
			{ID: "XSSTooltip"},
		},
	}
	configFile := &config.Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{
			"XSSTooltip": {
				CicdMaxIssues: intPtr(10),
			},
		},
	}
	var buf bytes.Buffer

	//When
	result := CheckThresholdViolations(&buf, finalResult, configFile)

	//Then
	if result {
		t.Errorf("Expected no violation when findings (3) are within threshold (10)")
	}
	if buf.Len() != 0 {
		t.Errorf("Expected no output when within threshold, got: %s", buf.String())
	}
}

func TestCheckThresholdViolations_WhenExceedsThreshold_ReturnsTrue(t *testing.T) {
	//Given
	finalResult := &finding.Output{
		Results: []finding.Finding{
			{ID: "XSSTooltip"},
			{ID: "XSSTooltip"},
			{ID: "XSSTooltip"},
			{ID: "XSSTooltip"},
			{ID: "XSSTooltip"},
		},
	}
	configFile := &config.Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{
			"XSSTooltip": {
				CicdMaxIssues: intPtr(2),
			},
		},
	}
	var buf bytes.Buffer

	//When
	result := CheckThresholdViolations(&buf, finalResult, configFile)

	//Then
	if !result {
		t.Errorf("Expected violation when findings (5) exceed threshold (2)")
	}
	output := buf.String()
	if !strings.Contains(output, "Threshold violations:") {
		t.Errorf("Expected header in output, got: %s", output)
	}
	if !strings.Contains(output, "Rule XSSTooltip has 5 issues (max allowed: 2)") {
		t.Errorf("Expected violation detail in output, got: %s", output)
	}
	if !strings.Contains(output, "1 rule(s) exceeded their cicdmaxissues threshold.") {
		t.Errorf("Expected summary in output, got: %s", output)
	}
}

func TestCheckThresholdViolations_WhenExactlyAtThreshold_ReturnsFalse(t *testing.T) {
	//Given
	finalResult := &finding.Output{
		Results: []finding.Finding{
			{ID: "XSSTooltip"},
			{ID: "XSSTooltip"},
			{ID: "XSSTooltip"},
			{ID: "XSSTooltip"},
			{ID: "XSSTooltip"},
		},
	}
	configFile := &config.Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{
			"XSSTooltip": {
				CicdMaxIssues: intPtr(5),
			},
		},
	}
	var buf bytes.Buffer

	//When
	result := CheckThresholdViolations(&buf, finalResult, configFile)

	//Then
	if result {
		t.Errorf("Expected no violation when findings (5) are exactly at threshold (5)")
	}
	if buf.Len() != 0 {
		t.Errorf("Expected no output when at threshold, got: %s", buf.String())
	}
}

func TestCheckThresholdViolations_WhenMultipleRules_ReportsAllViolations(t *testing.T) {
	//Given
	finalResult := &finding.Output{
		Results: []finding.Finding{
			{ID: "RuleA"},
			{ID: "RuleA"},
			{ID: "RuleA"},
			{ID: "RuleB"},
			{ID: "RuleB"},
			{ID: "RuleC"},
		},
	}
	configFile := &config.Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{
			"RuleA": {
				CicdMaxIssues: intPtr(1), // 3 > 1, violation
			},
			"RuleB": {
				CicdMaxIssues: intPtr(5), // 2 <= 5, no violation
			},
			// RuleC has no override, defaults to 0, 1 > 0, violation
		},
	}
	var buf bytes.Buffer

	//When
	result := CheckThresholdViolations(&buf, finalResult, configFile)

	//Then
	if !result {
		t.Errorf("Expected violations for RuleA and RuleC")
	}
	output := buf.String()
	if !strings.Contains(output, "Rule RuleA has 3 issues (max allowed: 1)") {
		t.Errorf("Expected RuleA violation in output, got: %s", output)
	}
	if !strings.Contains(output, "Rule RuleC has 1 issues (max allowed: 0)") {
		t.Errorf("Expected RuleC violation in output, got: %s", output)
	}
	if strings.Contains(output, "RuleB") {
		t.Errorf("RuleB should not appear in violations, got: %s", output)
	}
	if !strings.Contains(output, "2 rule(s) exceeded their cicdmaxissues threshold.") {
		t.Errorf("Expected summary with 2 violations, got: %s", output)
	}
}

func TestCheckThresholdViolations_WhenMaxIssuesNotSet_DefaultsToZero(t *testing.T) {
	//Given
	finalResult := &finding.Output{
		Results: []finding.Finding{
			{ID: "SomeRule"},
		},
	}
	configFile := &config.Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{},
	}
	var buf bytes.Buffer

	//When
	result := CheckThresholdViolations(&buf, finalResult, configFile)

	//Then
	if !result {
		t.Errorf("Expected violation when maxissues defaults to 0 and rule has 1 finding")
	}
	output := buf.String()
	if !strings.Contains(output, "Rule SomeRule has 1 issues (max allowed: 0)") {
		t.Errorf("Expected default threshold violation, got: %s", output)
	}
}

func TestCheckThresholdViolations_WhenConfigNil_DefaultsToZero(t *testing.T) {
	//Given
	finalResult := &finding.Output{
		Results: []finding.Finding{
			{ID: "SomeRule"},
		},
	}
	var configFile *config.Config
	var buf bytes.Buffer

	//When
	result := CheckThresholdViolations(&buf, finalResult, configFile)

	//Then
	if !result {
		t.Errorf("Expected violation when config is nil and rule has findings")
	}
}

func TestCheckThresholdViolations_OutputIsSortedByRuleID(t *testing.T) {
	//Given
	finalResult := &finding.Output{
		Results: []finding.Finding{
			{ID: "ZRule"},
			{ID: "ARule"},
			{ID: "MRule"},
		},
	}
	configFile := &config.Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{},
	}
	var buf bytes.Buffer

	//When
	CheckThresholdViolations(&buf, finalResult, configFile)

	//Then
	output := buf.String()
	aIdx := strings.Index(output, "ARule")
	mIdx := strings.Index(output, "MRule")
	zIdx := strings.Index(output, "ZRule")
	if aIdx == -1 || mIdx == -1 || zIdx == -1 {
		t.Fatalf("Expected all rules in output, got: %s", output)
	}
	if !(aIdx < mIdx && mIdx < zIdx) {
		t.Errorf("Expected rules sorted alphabetically (A < M < Z), got: %s", output)
	}
}
