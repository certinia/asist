package output

import (
	"reflect"
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
			RepositoryName: "test/test",
			RepositoryURL:  sshUrl,
			RecordType:     "Finding",
			Content:        baselineOutputContent,
		},
		{
			RepositoryName: "test/test",
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
