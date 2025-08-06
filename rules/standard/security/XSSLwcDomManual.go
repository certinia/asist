package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var XSSLwcDomManualRuleID rules.RuleID = "XSSLwcDomManual"

type XSSLwcDomManualRule struct {
	metadata rules.RuleMetadata
}

func NewXSSLwcDomManualRule() *XSSLwcDomManualRule {
	return &XSSLwcDomManualRule{
		metadata: rules.RuleMetadata{
			ID:             XSSLwcDomManualRuleID,
			Name:           "Potential XSS with lwc:dom=\"manual\"",
			Description:    "Use of lwc:dom=\"manual\" may introduce an XSS issue. Consider removing it, otherwise ensure that no user input is rendered unescaped.",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: ".*(/|^)lwc/.+\\.html$",
			ExcludePattern: "",
			Pattern:        "lwc:dom=\"manual\"",
		},
	}
}

func (r *XSSLwcDomManualRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *XSSLwcDomManualRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
