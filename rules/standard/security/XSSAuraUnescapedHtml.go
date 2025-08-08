package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var XSSAuraUnescapedHtmlRuleID rules.RuleID = "XSSAuraUnescapedHtml"

type XSSAuraUnescapedHtmlRule struct {
	metadata rules.RuleMetadata
}

func NewXSSAuraUnescapedHtmlRule() *XSSAuraUnescapedHtmlRule {
	return &XSSAuraUnescapedHtmlRule{
		metadata: rules.RuleMetadata{
			ID:             XSSAuraUnescapedHtmlRuleID,
			Name:           "Potential XSS with aura:unescapedHtml",
			Description:    "Use of aura:unescapedHtml tag may introduce an XSS issue. Consider removing it, otherwise ensure that no user input is rendered unescaped.",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.cmp$",
			ExcludePattern: "",
			Pattern:        "(?i)<\\s*aura\\s*:\\s*unescapedHtml\\s+",
		},
	}
}

func (r *XSSAuraUnescapedHtmlRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *XSSAuraUnescapedHtmlRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
