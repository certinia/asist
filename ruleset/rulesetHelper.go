package ruleset

import (
	"log"

	"github.com/certinia/asist/config"
	"github.com/certinia/asist/message"
	"github.com/certinia/asist/parser/options"
	"github.com/certinia/asist/rules"
	"github.com/certinia/asist/rules/customrule"
)

var ruleIDs []rules.RuleID

func GetAllRuleIDs() []rules.RuleID {
	if len(ruleIDs) == 0 {
		for ruleID := range ruleMapping {
			ruleIDs = append(ruleIDs, ruleID)
		}
	}
	return ruleIDs
}

/**
* CreateRules - Method will create rules for provided ruleIds
 */
func CreateAndOverrideRules(standardRuleIDs, customRuleIds []rules.RuleID, configFile *config.Config) ([]*rules.Rule, error) {
	rules := []*rules.Rule{}
	if configFile == nil {
		configFile = &config.Config{}
	}
	for _, customRuleID := range customRuleIds {
		customRuleMetadata := configFile.CustomRegexRules[string(customRuleID)]
		rule := createCustomRule(customRuleMetadata, customRuleID)
		rules = append(rules, &rule)
	}
	for _, standardruleID := range standardRuleIDs {
		standardRuleMetadataOverride, isStandardRuleOverride := configFile.RuleOverrides[string(standardruleID)]
		rule, err := createStandardRule(standardruleID)
		if err != nil {
			return nil, err
		}
		if rule != nil {
			if isStandardRuleOverride {
				rule.GetMetadata().Override(standardRuleMetadataOverride, options.IsBaselineScan())
			}
			rules = append(rules, &rule)
		}
	}
	return rules, nil
}

func IsStandardRuleID(ruleId rules.RuleID) *bool {
	_, isexist := ruleMapping[ruleId]
	return &isexist
}

func createStandardRule(ruleId rules.RuleID) (rules.Rule, error) {
	rule := ruleMapping[ruleId]
	if rule == nil {
		log.Println(message.GetInvalidRuleIdWarning(string(ruleId)))
		return nil, nil
	}
	return rule, nil
}

func createCustomRule(ruleMetadata config.CustomRegexRule, ruleID rules.RuleID) rules.Rule {
	rule := customrule.NewCustomRule(ruleMetadata, ruleID)
	return rule
}

/**
 * GetRuleIdsToRun - Returns a map of rule IDs (standard, custom, CI/CD, specific) mapped to a boolean.
 * - If specificRuleIds are provided, returns map[specificRuleIds] = true.
 * - If CI/CD rules are enabled and a config file is provided, returns map[CICDRuleIds] = true.
 * - If a config is provided, overrides standardRuleIds with the provided values and includes custom rule IDs.
 * - Otherwise, returns map[standardRuleIds] = true.
 */
func GetRuleIdsToRun(configFile *config.Config, opts *options.Options) ([]rules.RuleID, []rules.RuleID, error) {
	//To Baseline scan on all ruleIds
	if opts.BaselineScan {
		return GetAllRuleIDs(), configFile.GetCustomRuleIds(), nil
	}
	// If user has specified specific rules, just return those
	if opts.Rules != "" {
		var standardRuleIds []rules.RuleID
		var customRuleIds []rules.RuleID

		for _, ruleID := range opts.SpecificRuleIds() {
			if containsRuleID(configFile.GetCustomRuleIds(), ruleID) {
				customRuleIds = append(customRuleIds, ruleID)
			} else {
				standardRuleIds = append(standardRuleIds, ruleID)
			}
		}

		return standardRuleIds, customRuleIds, nil
	}
	if configFile != nil {
		if opts.CICDScan {
			var standardRuleIds []rules.RuleID
			var customRuleIds []rules.RuleID

			for _, ruleID := range configFile.GetCICDRuleIds() {
				if containsRuleID(configFile.GetCustomRuleIds(), ruleID) {
					customRuleIds = append(customRuleIds, ruleID)
				} else {
					standardRuleIds = append(standardRuleIds, ruleID)
				}
			}

			return standardRuleIds, customRuleIds, nil
		}

		// Warn if overridden rule IDs are not valid
		warnForInvalidRuleIds(configFile.GetOverridedRulesId())

		return configFile.GetEnabledOverridedStandardRuleIds(GetAllRuleIDs()), configFile.GetEnabledCustomRuleIds(), nil
	}

	// If no config file, just include all the standard rules
	return GetAllRuleIDs(), nil, nil
}

func containsRuleID(ruleIDs []rules.RuleID, target rules.RuleID) bool {
	for _, id := range ruleIDs {
		if id == target {
			return true
		}
	}
	return false
}
func warnForInvalidRuleIds(ruleIds []rules.RuleID) {
	for _, ruleId := range ruleIds {
		_, isRuleIdExist := ruleMapping[ruleId]
		if !isRuleIdExist {
			log.Println(message.GetInvalidRuleIdWarning(string(ruleId)))
		}
	}
}
