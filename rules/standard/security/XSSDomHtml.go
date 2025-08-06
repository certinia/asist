package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var XSSDomHtmlRuleID rules.RuleID = "XSSDomHtml"

type XSSDomHtmlRule struct {
	metadata rules.RuleMetadata
}

func NewXSSDomHtmlRule() *XSSDomHtmlRule {
	return &XSSDomHtmlRule{
		metadata: rules.RuleMetadata{
			ID:             XSSDomHtmlRuleID,
			Name:           "Potential XSS with .html or .innerHTML",
			Description:    "Use of .html() or .innerHTML may introduce an XSS issue. Consider removing it; otherwise, ensure that all user inputs are sanitized properly.",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.page$|\\.js$|\\.component$",
			ExcludePattern: "/.*deps/|\\.min\\.js$|\\.test\\.js$|\\.t\\.js$|\\-min\\.js$|\\-debug\\.js$",
			Pattern:        "\\.html\\(((.*{!.*)|\\w+)\\)|\\.innerHTML\\s*\\=",
		},
	}
}

func (r *XSSDomHtmlRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *XSSDomHtmlRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
