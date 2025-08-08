package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var XSSLocationSearchRuleID rules.RuleID = "XSSLocationSearch"

type XSSLocationSearchRule struct {
	metadata rules.RuleMetadata
}

func NewXSSLocationSearchRule() *XSSLocationSearchRule {
	return &XSSLocationSearchRule{
		metadata: rules.RuleMetadata{
			ID:             XSSLocationSearchRuleID,
			Name:           "Potential XSS with location.search",
			Description:    "Use of location.search may introduce an XSS issue. Consider removing it, otherwise ensure that no user input is rendered unescaped.",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.html$|\\.js$",
			ExcludePattern: "\\.test\\.js$|\\.min\\.js$|\\.t\\.js$|\\-min\\.js$|\\-debug\\.js$",
			Pattern:        "location\\.search",
		},
	}
}

func (r *XSSLocationSearchRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *XSSLocationSearchRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
