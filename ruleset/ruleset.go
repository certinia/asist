// This file is dynamically generated using code generator and contains rule mapping

package ruleset

import (
	rules "github.com/certinia/asist/rules"
	codequality "github.com/certinia/asist/rules/standard/codequality"
	security "github.com/certinia/asist/rules/standard/security"
)

var ruleMapping = map[rules.RuleID]rules.Rule{
	codequality.DetectImportJavascriptFromFileRuleID:     codequality.NewDetectImportJavascriptFromFileRule(),
	codequality.DetectMissingAccessibilityModifierRuleID: codequality.NewDetectMissingAccessibilityModifierRule(),
	security.ApexClassNoSharingRuleID:                    security.NewApexClassNoSharingRule(),
	security.ApexClassWithoutSharingRuleID:               security.NewApexClassWithoutSharingRule(),
	security.AuraComponentCssExposedRuleID:               security.NewAuraComponentCssExposedRule(),
	security.EmailInjectionRuleID:                        security.NewEmailInjectionRule(),
	security.ExposedMessageChannelRuleID:                 security.NewExposedMessageChannelRule(),
	security.HardcodedCredentialsRuleID:                  security.NewHardcodedCredentialsRule(),
	security.InsecureCryptoAlgorithmRuleID:               security.NewInsecureCryptoAlgorithmRule(),
	security.InsecureEndpointRuleID:                      security.NewInsecureEndpointRule(),
	security.JSNotInStaticResourceRuleID:                 security.NewJSNotInStaticResourceRule(),
	security.LightningImproperCSSLoadRuleID:              security.NewLightningImproperCSSLoadRule(),
	security.LwcNonStandardPositioningRuleID:             security.NewLwcNonStandardPositioningRule(),
	security.ProtectedCustomSettingRuleID:                security.NewProtectedCustomSettingRule(),
	security.SensitiveInfoInDebugRuleID:                  security.NewSensitiveInfoInDebugRule(),
	security.SessionIDApexRuleID:                         security.NewSessionIDApexRule(),
	security.SessionIDVisualForceRuleID:                  security.NewSessionIDVisualForceRule(),
	security.XSSApexChartRuleID:                          security.NewXSSApexChartRule(),
	security.XSSAuraUnescapedHtmlRuleID:                  security.NewXSSAuraUnescapedHtmlRule(),
	security.XSSCurrentPageParametersRuleID:              security.NewXSSCurrentPageParametersRule(),
	security.XSSDomHtmlRuleID:                            security.NewXSSDomHtmlRule(),
	security.XSSEscapeFalseRuleID:                        security.NewXSSEscapeFalseRule(),
	security.XSSEscapeFalseInJSRuleID:                    security.NewXSSEscapeFalseInJSRule(),
	security.XSSFormActionRuleID:                         security.NewXSSFormActionRule(),
	security.XSSIsRichTextRuleID:                         security.NewXSSIsRichTextRule(),
	security.XSSJavascriptButtonRuleID:                   security.NewXSSJavascriptButtonRule(),
	security.XSSLabelRuleID:                              security.NewXSSLabelRule(),
	security.XSSLocationSearchRuleID:                     security.NewXSSLocationSearchRule(),
	security.XSSLwcDomManualRuleID:                       security.NewXSSLwcDomManualRule(),
	security.XSSMergeFieldRuleID:                         security.NewXSSMergeFieldRule(),
	security.XSSSrcDocRuleID:                             security.NewXSSSrcDocRule(),
	security.XSSTooltipRuleID:                            security.NewXSSTooltipRule(),
}
