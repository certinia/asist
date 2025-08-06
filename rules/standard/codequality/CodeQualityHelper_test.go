package codequality

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/certinia/asist/files"
)

func TestGetFunctionsList_WhenFileProvided_ReturnsAllFunctionsInFile(t *testing.T) {
	// Given
	mockLines := []files.Line{
		{LineNumber: 1, Text: "public class MyClass {", IsCommentedLine: false},
		{LineNumber: 2, Text: "    public MyClass() {}", IsCommentedLine: false},
		{LineNumber: 3, Text: "public static void myFunction(int a, String b) {", IsCommentedLine: false},
		{LineNumber: 4, Text: "    return;", IsCommentedLine: false},
		{LineNumber: 5, Text: "}", IsCommentedLine: false},
		{LineNumber: 6, Text: "}", IsCommentedLine: false},
	}
	mockFile := files.File{
		Lines:           mockLines,
		FileName:        "sampleClass.cls",
		IgnoresSelected: []files.IgnoreSelected{},
	}
	expectedFunctionsList := []functionsDetail{
		{
			lineIndex:          1,
			functionDefinition: "public MyClass() {",
		},
		{
			lineIndex:          2,
			functionDefinition: "void myFunction(int a, String b) {",
		},
	}

	// When
	actualFunctionsList := getFunctionsList("DetectMissingAccessibilityModifier", mockFile)

	// Then
	if !reflect.DeepEqual(actualFunctionsList, expectedFunctionsList) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Functions list should be equal!", actualFunctionsList, expectedFunctionsList)
	}
}

func TestGetFunctionsList_WhenNoConstructorRegex_ReturnsEmptyFunctionsList(t *testing.T) {
	// Given
	mockLines := []files.Line{
		{LineNumber: 1, Text: "public static void myFunction(int a, String b) {", IsCommentedLine: false},
		{LineNumber: 2, Text: "    return;", IsCommentedLine: false},
		{LineNumber: 3, Text: "}", IsCommentedLine: false},
	}
	mockFile := files.File{
		Lines:           mockLines,
		FileName:        "sampleClass.cls",
		IgnoresSelected: []files.IgnoreSelected{},
	}

	// When
	actualFunctionsList := getFunctionsList("DetectMissingAccessibilityModifier", mockFile)

	// Then
	if actualFunctionsList != nil {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Functions list should be empty!!", actualFunctionsList, nil)
	}
}

func TestFindAllFunctions_WhenStringOfFileContentExist_ReturnsAllFunctions(t *testing.T) {
	// Given
	functionOrConstructor := `(?i)` +
		`(` +
		`((MyClass))\s*` + `|` +
		`(` +
		`\b[a-z0-9_.\[\]]+\s*(<[a-z0-9_.,\s<>]+>)?\s+` +
		`\w+\s*` +
		`)` +
		`)` +
		`\(\s*(((final\s+)?[a-z0-9_.\[\]]+(\s*<[a-z0-9,_.\s<>]+>)?\s+[a-z0-9_\[\]]+)?(\s*,\s*(final\s+)?[a-z0-9_.\[\]]+(\s*<[a-z0-9,_.\s<>]+>)?\s+[a-z0-9_\[\]]+)*)\s*\)` +
		`\s*\{`
	functionRegexp := regexp.MustCompile(functionOrConstructor)
	expectedFunctionsList := []functionsDetail{
		{
			lineIndex:          2,
			functionDefinition: "public MyClass() {",
		},
		{
			lineIndex:          2,
			functionDefinition: "void myFunction(int a, String b) {",
		},
	}

	// When
	actualFunctionsList := findAllFunctions("class MyClass {    public MyClass() {}public static void myFunction(int a, String b) {    return;}}", []int{15, 8, 86, 97, 98, 99}, functionRegexp)

	// Then
	if !reflect.DeepEqual(actualFunctionsList, expectedFunctionsList) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Functions list should be equal!", actualFunctionsList, expectedFunctionsList)
	}
}

func TestGetConstructorRegex_WhenStringOfFileContentExist_ReturnsConstructorRegex(t *testing.T) {
	// Given
	fileContentAsString := []string{"class MyClass {"}

	// When
	constructorRegex := getConstructorRegex(fileContentAsString[0])

	// Then
	if !(constructorRegex != "") {
		t.Errorf("Constructor Regex should not be empty!")
	}

}

func TestConvertFileIntoSingleString_WhenFileExist_ReturnsSingleString(t *testing.T) {
	// Given
	mockLines := []files.Line{
		{LineNumber: 1, Text: "public class MyClass {", IsCommentedLine: false},
		{LineNumber: 2, Text: "public MyClass() {}", IsCommentedLine: false},
		{LineNumber: 3, Text: " // my function", IsCommentedLine: true},
		{LineNumber: 4, Text: "public static void myFunction(int a,/*comment*/ int b) {", IsCommentedLine: false},
		{LineNumber: 5, Text: "return;", IsCommentedLine: false},
		{LineNumber: 6, Text: "}", IsCommentedLine: false},
		{LineNumber: 7, Text: "}", IsCommentedLine: false},
	}
	mockFile := files.File{
		Lines:           mockLines,
		FileName:        "sampleClass.cls",
		IgnoresSelected: []files.IgnoreSelected{},
	}
	expectedStringResult := "public class MyClass {public MyClass() {}public static void myFunction(int a, int b) {return;}}"
	expectedLinesLength := []int{22, 41, 41, 86, 93, 94, 95}

	// When
	actualStringResult, actualLineslength := convertFileIntoSingleString("DetectMissingAccessibilityModifier", mockFile)

	// Then
	if !reflect.DeepEqual(actualStringResult, expectedStringResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "String results should be equal!", actualStringResult, expectedStringResult)
	}
	if !reflect.DeepEqual(actualLineslength, expectedLinesLength) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Strings lengths should be equal!", actualLineslength, expectedLinesLength)
	}
}
