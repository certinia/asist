package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var XSSFormActionRuleID rules.RuleID = "XSSFormAction"

type XSSFormActionRule struct {
	metadata rules.RuleMetadata
}

func NewXSSFormActionRule() *XSSFormActionRule {
	return &XSSFormActionRule{
		metadata: rules.RuleMetadata{
			ID:             XSSFormActionRuleID,
			Name:           "Potential XSS with formaction in JS context",
			Description:    "Use of the formaction attribute in JavaScript context may introduce an XSS issue. Consider removing it, otherwise ensure that no user input is rendered unescaped.",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.cmp$",
			ExcludePattern: "",
			Pattern:        "(?i)formaction\\s*=\\s*\"javascript:",
		},
	}
}

func (r *XSSFormActionRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *XSSFormActionRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
