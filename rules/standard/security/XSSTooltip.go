package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var XSSTooltipRuleID rules.RuleID = "XSSTooltip"

type XSSTooltipRule struct {
	metadata rules.RuleMetadata
}

func NewXSSTooltipRule() *XSSTooltipRule {
	return &XSSTooltipRule{
		metadata: rules.RuleMetadata{
			ID:             XSSTooltipRuleID,
			Name:           "Potential XSS with tooltip",
			Description:    "Use of a tooltip may introduce an XSS issue. Consider removing it, otherwise ensure that no user input is rendered unescaped.",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.js$",
			ExcludePattern: "/.*lwc/|\\.test\\.js$|\\.min\\.js$|\\.t\\.js$|\\-min\\.js$|\\-debug\\.js$",
			Pattern:        "(\\.tooltip\\s*=|setTooltip|getLockedTooltip:|tooltip\\s*:)",
		},
	}
}

func (r *XSSTooltipRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *XSSTooltipRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
