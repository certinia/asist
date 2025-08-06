package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var ApexClassWithoutSharingRuleID rules.RuleID = "ApexClassWithoutSharing"

type ApexClassWithoutSharing struct {
	metadata rules.RuleMetadata
}

func NewApexClassWithoutSharingRule() *ApexClassWithoutSharing {
	return &ApexClassWithoutSharing{
		metadata: rules.RuleMetadata{
			ID:             ApexClassWithoutSharingRuleID,
			Name:           "Apex Class Without Sharing",
			Description:    "Use of the ‘without sharing’ clause can break the Salesforce security model, as it doesn't abide by row-level permissions. Always set the most restrictive sharing clause for each class, or provide a justification as to why 'without sharing' is used.",
			Severity:       rules.SeverityMedium,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.cls$",
			ExcludePattern: "(?i)/(force-app-autotest|autotest|systemtest|test)(s)?/|Test\\.cls$",
			Pattern:        "^(?i)\\s*(public|private|global)?\\s*(abstract|virtual)?\\s*(without\\s+sharing)\\s*(abstract|virtual)?\\s+class\\s+\\w+",
		},
	}
}

func (r *ApexClassWithoutSharing) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *ApexClassWithoutSharing) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
