package security

import (
	"reflect"
	"testing"

	"github.com/certinia/asist/files"
	"github.com/certinia/asist/rules"
)

func TestFindVulnerableLinesWithoutExtraVulnerableTags(t *testing.T) {
	// Given
	mockLines := []files.Line{
		{LineNumber: 1, Text: "<apex:component layout='none' selfClosing='false' controller='CustomFormController'>", IsCommentedLine: false},
		{LineNumber: 2, Text: "<script><apex:outputPanel layout='none' rendered='{!NOT(LEN($CurrentPage.parameters.id) == 0)}'</script>", IsCommentedLine: false},
		{LineNumber: 3, Text: "</apex:component>", IsCommentedLine: false},
	}
	mockFile := files.File{
		Lines:           mockLines,
		FileName:        "component.component",
		IgnoresSelected: []files.IgnoreSelected{},
	}
	expectedResult := []rules.Occurrence{
		{
			LineContent:     "<script><apex:outputPanel layout='none' rendered='{!NOT(LEN($CurrentPage.parameters.id) == 0)}'</script>",
			LineNumber:      2,
			ColumnRange:     []int{7, 95},
			IsFalsePositive: false,
		},
	}

	// When
	actualResult := findVulnerableLinesBetweenTags(mockFile, "XSSCurrentPageParameters", []string{}, false)

	// Then
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Occurrences list should be equal!", actualResult, expectedResult)
	}
}

func TestFindVulnerableLinesWithExtraVulnerableTags(t *testing.T) {
	// Given
	mockLines := []files.Line{
		{LineNumber: 1, Text: "<apex:component layout='none' selfClosing='false' controller='CustomFormController'>", IsCommentedLine: false},
		{LineNumber: 2, Text: "<script>", IsCommentedLine: false},
		{LineNumber: 3, Text: "<apex:outputPanel layout='none' rendered='{!NOT(LEN($CurrentPage.parameters.id) == 0)}'", IsCommentedLine: false},
		{LineNumber: 4, Text: "</script>", IsCommentedLine: false},
		{LineNumber: 5, Text: "</apex:component>", IsCommentedLine: false},
	}
	mockFile := files.File{
		Lines:           mockLines,
		FileName:        "component.component",
		IgnoresSelected: []files.IgnoreSelected{},
	}
	expectedResult := []rules.Occurrence{
		{
			LineContent:     "<apex:component layout='none' selfClosing='false' controller='CustomFormController'>",
			LineNumber:      1,
			ColumnRange:     []int{16, 22},
			IsFalsePositive: false,
		},
		{
			LineContent:     "<script>",
			LineNumber:      2,
			IsFalsePositive: false,
		},
		{
			LineContent:     "<apex:outputPanel layout='none' rendered='{!NOT(LEN($CurrentPage.parameters.id) == 0)}'",
			LineNumber:      3,
			IsFalsePositive: false,
		},
		{
			LineContent:     "</script>",
			LineNumber:      4,
			IsFalsePositive: false,
		},
	}

	// When
	actualResult := findVulnerableLinesBetweenTags(mockFile, "XSSCurrentPageParameters", []string{"layout"}, false)

	// Then
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Occurrences list should be equal!", actualResult, expectedResult)
	}
}
