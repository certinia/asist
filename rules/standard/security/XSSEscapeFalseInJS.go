package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var XSSEscapeFalseInJSRuleID rules.RuleID = "XSSEscapeFalseInJS"

type XSSEscapeFalseInJSRule struct {
	metadata rules.RuleMetadata
}

func NewXSSEscapeFalseInJSRule() *XSSEscapeFalseInJSRule {
	return &XSSEscapeFalseInJSRule{
		metadata: rules.RuleMetadata{
			ID:             XSSEscapeFalseInJSRuleID,
			Name:           "Potential XSS with escape:false",
			Description:    "Use of escape:false in JavaScript VisualForce Remoting may introduce an XSS issue. Consider removing it, otherwise ensure that no user input is rendered unescaped.",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.page$|\\.js$|\\.component$",
			ExcludePattern: "/.*sencha\\d*/|\\.test\\.js$|\\.min\\.js$|\\.t\\.js$|\\-min\\.js$|\\-debug\\.js$",
			Pattern:        "escape\\s*:\\s*false",
		},
	}
}

func (r *XSSEscapeFalseInJSRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *XSSEscapeFalseInJSRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
