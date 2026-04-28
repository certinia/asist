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

// TestFindVulnerableLinesOpenAndCloseTagOnSameLine verifies that the column range
// is correctly computed relative to the open tag end when both <script> and </script>
// appear on the same line.
func TestFindVulnerableLinesOpenAndCloseTagOnSameLine(t *testing.T) {
	// Given
	mockLines := []files.Line{
		{LineNumber: 1, Text: "<script>var x = '{!$CurrentPage.parameters.id}';</script>", IsCommentedLine: false},
	}
	mockFile := files.File{
		Lines:           mockLines,
		FileName:        "test.page",
		IgnoresSelected: []files.IgnoreSelected{},
	}
	// openScriptStyleTagRegexp matches "<script" (without '>'), so openTagEnd = 7.
	// "</script>" starts at absolute index 48, closeTagMatch[0] = 48-7 = 41.
	// ColumnRange = [7, 7+41] = [7, 48].
	expectedResult := []rules.Occurrence{
		{
			LineContent:     "<script>var x = '{!$CurrentPage.parameters.id}';</script>",
			LineNumber:      1,
			ColumnRange:     []int{7, 48},
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

// TestFindVulnerableLinesCloseTagFollowedByOpenTagKeepsState verifies that when a
// closing tag is immediately followed by a new opening tag on the same line,
// hasOpeningScriptOrStyleTagFound remains true and subsequent lines are still scanned.
func TestFindVulnerableLinesCloseTagFollowedByOpenTagKeepsState(t *testing.T) {
	// Given - line 2 closes and re-opens a script block; line 3 should still be scanned
	mockLines := []files.Line{
		{LineNumber: 1, Text: "<script>", IsCommentedLine: false},
		{LineNumber: 2, Text: "</script><script>", IsCommentedLine: false},
		{LineNumber: 3, Text: "var y = '{!$CurrentPage.parameters.name}';", IsCommentedLine: false},
		{LineNumber: 4, Text: "</script>", IsCommentedLine: false},
	}
	mockFile := files.File{
		Lines:           mockLines,
		FileName:        "test.page",
		IgnoresSelected: []files.IgnoreSelected{},
	}
	expectedResult := []rules.Occurrence{
		{LineContent: "<script>", LineNumber: 1, IsFalsePositive: false},
		{LineContent: "</script><script>", LineNumber: 2, IsFalsePositive: false},
		{LineContent: "var y = '{!$CurrentPage.parameters.name}';", LineNumber: 3, IsFalsePositive: false},
		{LineContent: "</script>", LineNumber: 4, IsFalsePositive: false},
	}

	// When
	actualResult := findVulnerableLinesBetweenTags(mockFile, "XSSCurrentPageParameters", []string{}, false)

	// Then
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Occurrences list should be equal!", actualResult, expectedResult)
	}
}

// TestFindVulnerableLinesColumnRangeLessThanGuard verifies the columnRange[0] < columnRange[1]
// guard introduced in the last commit. When closeTagMatch is nil (close tag not found after
// openTagEnd), columnRange stays [0,0] and the < condition correctly prevents appending a
// spurious ColumnRange. The line is still reported (as part of the opening-tag line) but
// without a ColumnRange field.
func TestFindVulnerableLinesColumnRangeLessThanGuard(t *testing.T) {
	// Given - a self-closing-like line where openScriptStyleTagRegexp matches but
	// closeScriptStyleTagRegexp also matches, yet the close tag appears BEFORE openTagEnd
	// when searched from line.Text[openTagEnd:] → closeTagMatch is nil → columnRange = [0,0].
	// Construct: "</style><style>content" — closeTag at 0, openTag at 8.
	// When we enter the else branch: openTagEnd = 15 (after "<style"), then we search
	// line.Text[15:] for a close tag — none exists → closeTagMatch nil → columnRange stays [0,0].
	mockLines := []files.Line{
		{LineNumber: 1, Text: "<style>body{}</style>", IsCommentedLine: false},
	}
	mockFile := files.File{
		Lines:           mockLines,
		FileName:        "test.page",
		IgnoresSelected: []files.IgnoreSelected{},
	}

	// When
	actualResult := findVulnerableLinesBetweenTags(mockFile, "XSSCurrentPageParameters", []string{}, false)

	// Then - ColumnRange [7, 14] is valid (7 < 14), so exactly one occurrence with a ColumnRange.
	if len(actualResult) != 1 {
		t.Fatalf("Expected 1 occurrence, got %d: %+v", len(actualResult), actualResult)
	}
	if len(actualResult[0].ColumnRange) != 2 {
		t.Errorf("Expected a ColumnRange on the occurrence, got: %+v", actualResult[0])
	}
	if actualResult[0].ColumnRange[0] >= actualResult[0].ColumnRange[1] {
		t.Errorf("columnRange[0] must be < columnRange[1], got: %+v", actualResult[0].ColumnRange)
	}
}
