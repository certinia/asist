package security

import (
	"github.com/certinia/asist/files"
	"github.com/certinia/asist/regexrulehelper"
	"github.com/certinia/asist/rules"
)

var ExposedMessageChannelRuleID rules.RuleID = "ExposedMessageChannel"

type ExposedMessageChannelRule struct {
	metadata rules.RuleMetadata
}

func NewExposedMessageChannelRule() *ExposedMessageChannelRule {
	return &ExposedMessageChannelRule{
		metadata: rules.RuleMetadata{
			ID:             ExposedMessageChannelRuleID,
			Name:           "Potential security failure due to misconfiguration of LMS",
			Description:    "Lightning Message Services provides access to the Lightning Message Service (LMS) API, which lets you publish and subscribe to messages across the DOM and between Aura, Visualforce, and Lightning Web Components. Setting isExposed in LMS is known to cause Salesforce Security Review failures.\nReference: https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_lightningmessagechannel.htm",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.messageChannel-meta\\.xml$|\\.messageChannel$",
			ExcludePattern: "",
			Pattern:        "<isExposed>(?i)(true)\\s*</isExposed>",
		},
	}
}

func (r *ExposedMessageChannelRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *ExposedMessageChannelRule) Run(fileToScan files.File) []rules.Occurrence {
	return regexrulehelper.FindOccurancesForFile(fileToScan, &r.metadata, false)
}
