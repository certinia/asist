package finding

import (
	"testing"

	"github.com/certinia/asist/rules"
)

func TestCreateFindingID(t *testing.T) {
	// Given
	occurrence := rules.Occurrence{
		FileName:        "TestMessageChannel.messageChannel-meta.xml",
		LineContent:     "	<isExposed>true</isExposed>",
		LineNumber:      10,
		ColumnRange:     []int{1, 28},
		IsFalsePositive: false,
	}

	finding := Finding{
		ID:         "ExposedMessageChannel",
		Occurrence: occurrence,
	}
	// When
	actualFindingID := finding.CreateFindingID()

	// Then
	if !(actualFindingID != "") {
		t.Errorf("Finding ID should not be empty!")
	}
}
