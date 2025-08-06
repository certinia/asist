package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var XSSApexChartRuleID rules.RuleID = "XSSApexChart"

type XSSApexChartRule struct {
	metadata rules.RuleMetadata
}

func NewXSSApexChartRule() *XSSApexChartRule {
	return &XSSApexChartRule{
		metadata: rules.RuleMetadata{
			ID:             XSSApexChartRuleID,
			Name:           "Potential XSS with apex:chart",
			Description:    "Use of data property inside apex:chart tag may introduce an XSS issue. Consider removing it, otherwise ensure that no user input is rendered unescaped.",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.page$|\\.component$",
			ExcludePattern: "",
			Pattern:        "(?i)<\\s*apex\\s*:\\s*chart\\s+",
		},
	}
}

func (r *XSSApexChartRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *XSSApexChartRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
