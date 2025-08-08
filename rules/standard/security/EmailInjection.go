package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var EmailInjectionRuleID rules.RuleID = "EmailInjection"

type EmailInjectionRule struct {
	metadata rules.RuleMetadata
}

func NewEmailInjectionRule() *EmailInjectionRule {
	return &EmailInjectionRule{
		metadata: rules.RuleMetadata{
			ID:             EmailInjectionRuleID,
			Name:           "Potential issues with email injection",
			Description:    "Using setHtmlBody() in emails sent using SingleEmailMessage or MassEmailMessage classes can lead to arbitrary markup being included. Consider using setPlainTextBody(); otherwise, ensure that user input is properly escaped.",
			Severity:       rules.SeverityMedium,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.cls$",
			ExcludePattern: "(?i)/(force-app-autotest|autotest|systemtest|test)(s)?/|Test\\.cls$",
			Pattern:        "(?i)\\.setHtmlBody\\(",
		},
	}
}

func (r *EmailInjectionRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *EmailInjectionRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
