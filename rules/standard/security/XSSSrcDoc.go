package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var XSSSrcDocRuleID rules.RuleID = "XSSSrcDoc"

type XSSSrcDocRule struct {
	metadata rules.RuleMetadata
}

func NewXSSSrcDocRule() *XSSSrcDocRule {
	return &XSSSrcDocRule{
		metadata: rules.RuleMetadata{
			ID:             XSSSrcDocRuleID,
			Name:           "Potential XSS with srcdoc",
			Description:    "Use of srcdoc inside the iframe tag may introduce an XSS issue. Consider removing it, otherwise ensure that no user input is rendered unescaped.",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.cmp$|\\.html$",
			ExcludePattern: "",
			Pattern:        "srcdoc\\s*=",
		},
	}
}

func (r *XSSSrcDocRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *XSSSrcDocRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
