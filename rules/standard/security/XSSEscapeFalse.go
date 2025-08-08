package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var XSSEscapeFalseRuleID rules.RuleID = "XSSEscapeFalse"

type XSSEscapeFalseRule struct {
	metadata rules.RuleMetadata
}

func NewXSSEscapeFalseRule() *XSSEscapeFalseRule {
	return &XSSEscapeFalseRule{
		metadata: rules.RuleMetadata{
			ID:             XSSEscapeFalseRuleID,
			Name:           "Potential XSS with escape=false",
			Description:    "Use of escape=\"false\" may introduce an XSS issue. Consider removing it, otherwise ensure that no user input is rendered unescaped.",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.component$|\\.page$|\\.email$",
			ExcludePattern: "",
			Pattern:        "(?i)escape\\s*=\\s*\"false\"",
		},
	}
}

func (r *XSSEscapeFalseRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *XSSEscapeFalseRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
