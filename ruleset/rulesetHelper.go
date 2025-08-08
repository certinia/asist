package ruleset

import (
	"github.com/certinia/asist/config"
	"github.com/certinia/asist/errorhandler"
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
		if isStandardRuleOverride {
			rule.GetMetadata().Override(standardRuleMetadataOverride, options.IsBaselineScan())
		}
		rules = append(rules, &rule)
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
		return nil, errorhandler.NewUserError(message.GetInvalidRuleIdError(string(ruleId)))
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
		ruleIds := opts.SpecificRuleIds()
		if err := validateRuleIds(ruleIds); err != nil {
			return nil, nil, err
		}
		return opts.SpecificRuleIds(), nil, nil
	}
	if configFile != nil {
		if opts.CICDScan {
			return configFile.GetCICDRuleIds(), nil, nil
		}
		//Validate overrided rule Ids are valid
		if err := validateRuleIds(configFile.GetOverridedRulesId()); err != nil {
			return nil, nil, err
		}
		return configFile.GetEnabledOverridedStandardRuleIds(GetAllRuleIDs()), configFile.GetEnabledCustomRuleIds(), nil
	}
	// If no config file, just include all the standard rules
	return GetAllRuleIDs(), nil, nil
}

func validateRuleIds(ruleIds []rules.RuleID) error {
	for _, ruleId := range ruleIds {
		_, isRuleIdExist := ruleMapping[ruleId]
		if !isRuleIdExist {
			return errorhandler.NewUserError(message.GetInvalidRuleIdError(string(ruleId)))
		}
	}
	return nil
}
