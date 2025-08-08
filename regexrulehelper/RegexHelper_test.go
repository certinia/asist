package regexrulehelper

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/certinia/asist/files"
	"github.com/certinia/asist/rules"
)

func TestValidateRegexRun_WhenEmptyIncludeAndExcludePatterns_ReturnsTrue(t *testing.T) {
	//Given
	ruleMetadata := rules.RuleMetadata{
		ID:             "SampleRule1",
		Name:           "Sample Rule",
		Description:    "This is a description for the sample rule",
		Severity:       rules.SeverityLow,
		RuleCategory:   rules.CategorySecurity,
		IncludePattern: "",
		ExcludePattern: "",
		Pattern:        "System\\.debug",
	}
	expectedResult := true

	//When
	actualResult := RunIncludeExcludePatternsOnFile("testFile.cls", ruleMetadata)

	//Then
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Given file path is allowed in case of empty IncludePattern and ExcludePattern!", actualResult, expectedResult)
	}
}

func TestValidateRegexRun_WhenHavingValidIncludeAndExcludePatterns_ReturnsTrue(t *testing.T) {
	//Given
	ruleMetadata := rules.RuleMetadata{
		ID:             "SampleRule1",
		Name:           "Sample Rule",
		Description:    "This is a description for the sample rule",
		Severity:       rules.SeverityLow,
		RuleCategory:   rules.CategorySecurity,
		IncludePattern: "\\.cls$",
		ExcludePattern: "Test\\.cls$",
		Pattern:        "System\\.debug",
	}
	expectedResult := true

	//When
	actualResult := RunIncludeExcludePatternsOnFile("CustomObjectController.cls", ruleMetadata)

	//Then
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Given file path is allowed in case of valid conditions of IncludePattern and ExcludePattern!", actualResult, expectedResult)
	}
}

func TestValidateRegexRun_WhenFileNotSupportIncludeAndExcludePatterns_ReturnsFalse(t *testing.T) {
	//Given
	ruleMetadata := rules.RuleMetadata{
		ID:             "SampleRule1",
		Name:           "Sample Rule",
		Description:    "This is a description for the sample rule",
		Severity:       rules.SeverityLow,
		RuleCategory:   rules.CategorySecurity,
		IncludePattern: ".cls$",
		ExcludePattern: "Test.cls$",
		Pattern:        "System\\.debug",
	}
	expectedResult := false

	//When
	actualResult := RunIncludeExcludePatternsOnFile("childComponent.js", ruleMetadata)

	//Then
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Given file path is not allowed in case of IncludePattern and ExcludePattern doesn't match the file extension!", actualResult, expectedResult)
	}
}

func TestFindMatchesForFile_WhereFileLinesGenuine_ReturnsOutput(t *testing.T) {
	//Given
	fileName := "sampleTest-meta.xml"
	line := []files.Line{
		{
			LineNumber:      1,
			Text:            "<linkType>javascript</linkType>",
			IsCommentedLine: false,
		},
		{
			LineNumber:      2,
			Text:            "<!--asist-ignore-begin:[XSSIsRichText]-->",
			IsCommentedLine: true,
		},
		{
			LineNumber:      3,
			Text:            "<type>isRichText</type>",
			IsCommentedLine: false,
		},
		{
			LineNumber:      4,
			Text:            "<!--asist-ignore-end-->",
			IsCommentedLine: true,
		},
	}
	ruleIds := map[string]bool{
		"XSSIsRichText": true,
	}
	ignoreSelected := []files.IgnoreSelected{
		{
			BeginLine: 2,
			EndLine:   4,
			RuleIDs:   ruleIds,
		},
	}
	fileToScan := files.File{
		FileName:        fileName,
		Lines:           line,
		IgnoresSelected: ignoreSelected,
	}

	metadata := rules.RuleMetadata{
		ID:        rules.RuleID("XSSIsRichText"),
		Pattern:   "<linkType>javascript</linkType>",
		Qualifier: "linkType",
	}

	expectedResult := []rules.Occurrence{
		{
			FileName:        "sampleTest-meta.xml",
			LineNumber:      1,
			LineContent:     "<linkType>javascript</linkType>",
			ColumnRange:     []int{0, 31},
			IsFalsePositive: false,
		},
	}
	//When
	actualResult := FindOccurancesForFile(fileToScan, &metadata, true)

	//Then
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Occurrences data mismatched!", actualResult, expectedResult)
	}
}

func TestFindMatchesForFile_WhenQualifierIsValidAndHavingZeroResultsFound_ReturnsEmptyOutput(t *testing.T) {
	//Given
	fileName := "sampleTest-meta.xml"
	line := []files.Line{
		{
			LineNumber:      1,
			Text:            "<type>isRichText</type>",
			IsCommentedLine: false,
		},
	}
	ignoreSelected := []files.IgnoreSelected{}
	fileToScan := files.File{
		FileName:        fileName,
		Lines:           line,
		IgnoresSelected: ignoreSelected,
	}

	metadata := rules.RuleMetadata{
		ID:      rules.RuleID("XSSIsRichText"),
		Pattern: "<linkType>javascript</linkType>",
	}

	expectedResult := []rules.Occurrence{}

	//When
	actualResult := FindOccurancesForFile(fileToScan, &metadata, true)

	//Then
	if !reflect.DeepEqual(len(actualResult), len(expectedResult)) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Occurrences data mismatched!", len(actualResult), len(expectedResult))
	}
}

func TestReplaceEncodePartByHash_ReturnsEncodedPartByHash(t *testing.T) {
	//Given
	line := "{!JSENCODE($Api.Session_ID)"
	encodeRegexp := regexp.MustCompile(`(JSENCODE|HTMLENCODE|JSINHTMLENCODE|URLFOR)\(`)
	expectedResult := "{!########################)"

	//When
	actualResult := ReplaceEncodePartByHash(line, encodeRegexp)

	//Then
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Replace encode strings mismatched!", actualResult, expectedResult)
	}
}
