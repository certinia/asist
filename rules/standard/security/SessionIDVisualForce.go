package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var SessionIDVisualForceRuleID rules.RuleID = "SessionIDVisualForce"

type SessionIDVisualForceRule struct {
	metadata rules.RuleMetadata
}

func NewSessionIDVisualForceRule() *SessionIDVisualForceRule {
	return &SessionIDVisualForceRule{
		metadata: rules.RuleMetadata{
			ID:             SessionIDVisualForceRuleID,
			Name:           "Use of Session ID in Visualforce",
			Description:    "Direct use of Session ID is forbidden by Salesforce due to risk of exposure and lack of audit trails for the operations performed. This will cause Salesforce Security Reviews to fail. If additional APIs need to be leveraged, implement OAuth using a Connected App.",
			Severity:       rules.SeverityCritical,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.page$|\\.component$",
			Pattern:        "(?i)\\$Api\\.Session_ID",
		},
	}
}

func (r *SessionIDVisualForceRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *SessionIDVisualForceRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
