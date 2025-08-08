package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var AuraComponentCssExposedRuleID rules.RuleID = "AuraComponentCssExposed"

type AuraComponentCssExposed struct {
	metadata rules.RuleMetadata
}

func NewAuraComponentCssExposedRule() *AuraComponentCssExposed {
	return &AuraComponentCssExposed{
		metadata: rules.RuleMetadata{
			ID:             AuraComponentCssExposedRuleID,
			Name:           "Aura Component CSS Exposed",
			Description:    "Use of CSS that avoids the component encapsulation that uses non-standard positioning (e.g., float: left or right, position: absolute or fixed) can break the Salesforce website layout and violate the spirit of Lightning’s security model, where Aura components are strictly sandboxed and guaranteed to stay in their lane. \nReference: https://developer.salesforce.com/blogs/2023/08/the-top-20-vulnerabilities-found-in-the-appexchange-security-review",
			Severity:       rules.SeverityMedium,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\/aura\\/.*(\\.css|\\.cmp)$",
			ExcludePattern: "",
			Pattern:        "(float\\s*:\\s*(left|right))|(position\\s*:\\s*(absolute|fixed))",
		},
	}
}

func (r *AuraComponentCssExposed) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *AuraComponentCssExposed) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
