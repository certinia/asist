package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var SessionIDApexRuleID rules.RuleID = "SessionIDApex"

type SessionIDApexRule struct {
	metadata rules.RuleMetadata
}

func NewSessionIDApexRule() *SessionIDApexRule {
	return &SessionIDApexRule{
		metadata: rules.RuleMetadata{
			ID:             SessionIDApexRuleID,
			Name:           "Use of Session ID in Apex",
			Description:    "Direct use of Session ID is forbidden by Salesforce due to risk of exposure and lack of audit trails for the operations performed. This will cause Salesforce Security Reviews to fail. If additional APIs need to be leveraged, implement OAuth using a Connected App.",
			Severity:       rules.SeverityCritical,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.cls$|\\.trigger$",
			ExcludePattern: "(?i)/(force-app-autotest|autotest|systemtest|test)(s)?/|Test\\.cls$",
			Pattern:        "(?i)UserInfo\\.getSessionID",
		},
	}
}

func (r *SessionIDApexRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *SessionIDApexRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
