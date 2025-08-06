package files

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

var fileName string
var line []Line
var file File
var ignoreSelected []IgnoreSelected

func TestMain(m *testing.M) {
	initializeData()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestRead_WhenFileExist_ReturnsFileContent(t *testing.T) {
	//Given
	expectedFileResult := File{
		Lines:           line,
		FileName:        fileName,
		IgnoresSelected: ignoreSelected,
	}

	//When
	actualResult, err := Read(fileName)

	//Then
	if err != nil {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Should not return any error by read method!", err, nil)
	}
	if !reflect.DeepEqual(*actualResult, expectedFileResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "File content mismatched!", *actualResult, expectedFileResult)
	}
}

func TestRead_WhenFileNotExist_ReturnsNil(t *testing.T) {
	//Given
	invalidFileName := "./invalid/file"

	//When
	actualResult, err := Read(invalidFileName)

	//Then
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
	if actualResult != nil {
		t.Errorf("%s Actual: %+v, Expected: %+v", "File content should be nil!!", actualResult, nil)
	}
}

func TestIgnoresFindingOnLine_WhenOccurrenceIsFalsePositive_ReturnsTrue(t *testing.T) {
	//When
	actualResult := file.IsLineMarkedFalsePositive("XSSIsRichText", 3)

	//Then
	if !actualResult {
		t.Errorf("Expected true as finding is false positive on line, returns false")
	}
}

func TestIgnoresFindingOnLine_WhenOccurrencesIsGenuine_ReturnsFalse(t *testing.T) {
	//When
	actualResult := file.IsLineMarkedFalsePositive("XSSIsRichText", 1)

	//Then
	if actualResult {
		t.Errorf("Expected false as finding is positive on line, returns false")
	}
}

func initializeData() {
	fileName, _ = filepath.Abs("./testData/sampleField-meta.xml")
	line = []Line{
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
			Text:            "<!--asist-ignore--><type>isRichText</type>",
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
	ignoreSelected = []IgnoreSelected{
		{
			BeginLine: 2,
			EndLine:   4,
			RuleIDs:   ruleIds,
		},
	}
	file = File{
		Lines:           line,
		FileName:        fileName,
		IgnoresSelected: ignoreSelected,
	}
}
