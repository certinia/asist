package rules

import (
	"fmt"

	"github.com/certinia/asist/files"
)

type Severity string

const (
	SeverityCritical Severity = "Critical"
	SeverityHigh     Severity = "High"
	SeverityMedium   Severity = "Medium"
	SeverityLow      Severity = "Low"
)

type RuleCategory string

const (
	CategorySecurity    RuleCategory = "Security"
	CategoryPerformance RuleCategory = "Performance"
	CategoryCodeQuality RuleCategory = "Code Quality"
	CategoryUX          RuleCategory = "UX"
)

type RuleID string

// RuleMetadata contains the metadata relevant for processing a specific rules. Defined by each rule
type RuleMetadata struct {
	// ID of the rule
	ID RuleID
	// Name shown in IDE
	Name string
	// Description shown in IDE
	Description string
	// Severity of finding in IDE
	Severity Severity
	// Filename pattern to filter on to run this rule (regex)
	IncludePattern string
	//Filename pattern to filter on to run this rule (regex)
	ExcludePattern string
	//Pattern to find the particular occurrences
	Pattern string
	//Category of rule shown in IDE
	RuleCategory RuleCategory
	//Qualifier - don't search the file if the file does not contain the qualifier
	Qualifier string
}

type Occurrence struct {
	FileName        string `json:"File"`
	LineContent     string `json:"Line"`
	LineNumber      int    `json:"LineNumber"`
	ColumnRange     []int  `json:"ColumnRange"`
	IsFalsePositive bool   `json:"-"`
}

type RuleMetadataOverride struct {
	Severity       string
	ExcludePattern string
	IncludePattern string
	Enabled        *bool
}

/**
 * Override - method used to override the standard rule's metadata properties Severity, IncludePattern, ExcludePattern
 */
func (md *RuleMetadata) Override(overrideRule RuleMetadataOverride, isBaselineScan bool) {
	if overrideRule.Severity != "" && !isBaselineScan {
		md.Severity = Severity(overrideRule.Severity)
	}
	if overrideRule.IncludePattern != "" {
		md.IncludePattern = overrideRule.IncludePattern
	}
	if overrideRule.ExcludePattern != "" {
		md.ExcludePattern = overrideRule.ExcludePattern
	}
}

// Rule is the generic rule interface. All functions in it must be declared for a rule type
type Rule interface {
	Run(fileToScan files.File) []Occurrence
	GetMetadata() *RuleMetadata
}

/**
 * CreateHashableString - method used to create hash string of an occurrence
 */
func (occurrence *Occurrence) CreateHashableString() string {
	contentToHash := fmt.Sprintf(
		"%v-%v-%v-%v-%v",
		occurrence.FileName,
		occurrence.LineContent,
		occurrence.LineNumber,
		occurrence.ColumnRange,
		occurrence.IsFalsePositive,
	)
	return contentToHash
}
