package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var LwcNonStandardPositioningRuleID rules.RuleID = "LwcNonStandardPositioning"

type LwcNonStandardPositioning struct {
	metadata rules.RuleMetadata
}

func NewLwcNonStandardPositioningRule() *LwcNonStandardPositioning {
	return &LwcNonStandardPositioning{
		metadata: rules.RuleMetadata{
			ID:             LwcNonStandardPositioningRuleID,
			Name:           "LWC non-standard positioning",
			Description:    "Use of CSS that avoids the component encapsulation that uses non-standard positioning (e.g., float: left or right, position: absolute or fixed) or its slds class equivalents can break the Salesforce website layout and violate the spirit of Lightning’s security model, where LWCs are strictly sandboxed and guaranteed to stay in their lane. \nReference: https://developer.salesforce.com/blogs/2023/08/the-top-20-vulnerabilities-found-in-the-appexchange-security-review.",
			Severity:       rules.SeverityMedium,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\/lwc\\/.*(\\.css|\\.html|\\.js)$",
			ExcludePattern: "",
			Pattern:        "(float\\s*(:|=)\\s*'?\"?(left|right)\"?'?)|(position\\s*(:|=)\\s*'?\"?(absolute|fixed)\"?'?)|slds-float_(left|right)|slds-is-absolute|slds-is-fixed",
		},
	}
}

func (r *LwcNonStandardPositioning) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *LwcNonStandardPositioning) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
