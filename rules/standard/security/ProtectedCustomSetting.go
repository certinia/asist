package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var ProtectedCustomSettingRuleID rules.RuleID = "ProtectedCustomSetting"

type ProtectedCustomSettingRule struct {
	metadata rules.RuleMetadata
}

func NewProtectedCustomSettingRule() *ProtectedCustomSettingRule {
	return &ProtectedCustomSettingRule{
		metadata: rules.RuleMetadata{
			ID:             ProtectedCustomSettingRuleID,
			Name:           "Potential issue in Government Cloud orgs due to Protected Custom Setting",
			Description:    "Protected Custom Settings can only be directly changed in the Salesforce org via Subscriber access, which is not allowed on Government Cloud Orgs. If adding a protected custom setting, consider a way for the customer to modify the custom setting without using Subscriber access. Consider making the custom setting public instead.",
			Severity:       rules.SeverityLow,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "__c\\.object-meta\\.xml$|\\.object$",
			ExcludePattern: "",
			Pattern:        "<visibility>Protected</visibility>|<customSettingsVisibility>Protected</customSettingsVisibility>",
			Qualifier:      "<customSettingsType>",
		},
	}
}

func (r *ProtectedCustomSettingRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *ProtectedCustomSettingRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
