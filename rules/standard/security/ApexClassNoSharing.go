package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var ApexClassNoSharingRuleID rules.RuleID = "ApexClassNoSharing"

type ApexClassNoSharing struct {
	metadata rules.RuleMetadata
}

func NewApexClassNoSharingRule() *ApexClassNoSharing {
	return &ApexClassNoSharing{
		metadata: rules.RuleMetadata{
			ID:             ApexClassNoSharingRuleID,
			Name:           "Apex Class No Sharing",
			Description:    "Use of Apex classes where the sharing clause is not specified may cause confusion for other developers. Always set the most restrictive sharing clause for each class.",
			Severity:       rules.SeverityMedium,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.cls$",
			ExcludePattern: "(?i)/(force-app-autotest|autotest|systemtest|test)(s)?/|Test\\.cls$",
			Pattern:        "^(?i)\\s*(public|private|global)?\\s*(abstract|virtual)?\\s*\\bclass\\b\\s+\\w+",
		},
	}
}

func (r *ApexClassNoSharing) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *ApexClassNoSharing) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
