package ruleset

import (
	"reflect"
	"testing"

	"github.com/certinia/asist/config"
	"github.com/certinia/asist/parser/options"
	"github.com/certinia/asist/rules"
)

func TestGetAllStdRuleIDs_StandardRuleIds(t *testing.T) {
	//Given
	const STANDARD_RULES_COUNT = 32

	//When
	standardRuleIds := GetAllStdRuleIDs()

	//Then
	if !reflect.DeepEqual(len(standardRuleIds), STANDARD_RULES_COUNT) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Standard Rules count mismatched!", len(standardRuleIds), STANDARD_RULES_COUNT)
	}
}

func TestGetRuleIdsToRun_BaselineEnabledAndConfigNil_returnAllStandardAndCustomRuleIds(t *testing.T) {
	//Given
	opts := options.Options{
		BaselineScan: true,
	}
	expectedStandardRuleIds := GetAllStdRuleIDs()
	expectedCustomRuleIds := []rules.RuleID{}

	//When
	actualStandardRuleIds, actualCustomRuleIds, err := GetRuleIdsToRun(nil, &opts)

	//Then
	if !reflect.DeepEqual(actualStandardRuleIds, expectedStandardRuleIds) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Standard ruleIds are mismatched!", actualStandardRuleIds, expectedStandardRuleIds)
	}
	if !reflect.DeepEqual(actualCustomRuleIds, expectedCustomRuleIds) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Custom ruleIds are mismatched!", actualCustomRuleIds, expectedCustomRuleIds)
	}
	if err != nil {
		t.Errorf("GetRuleIdsToRun method should not return error!")
	}
}

func TestGetRuleIdsToRun_BaselineEnabledAndHasConfigCustomRules_returnAllStandardAndCustomRuleIds(t *testing.T) {
	//Given
	opts := options.Options{
		BaselineScan: true,
	}

	configFile := config.Config{
		CustomRegexRules: map[string]config.CustomRegexRule{
			"customRule1": {},
		},
	}

	expectedStandardRuleIds := GetAllStdRuleIDs()
	expectedCustomRuleIds := []rules.RuleID{"customRule1"}

	//When
	actualStandardRuleIds, actualCustomRuleIds, err := GetRuleIdsToRun(&configFile, &opts)

	//Then
	if !reflect.DeepEqual(actualStandardRuleIds, expectedStandardRuleIds) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Standard ruleIds are mismatched!", actualStandardRuleIds, expectedStandardRuleIds)
	}
	if !reflect.DeepEqual(actualCustomRuleIds, expectedCustomRuleIds) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Custom ruleIds are mismatched!", actualCustomRuleIds, expectedCustomRuleIds)
	}
	if err != nil {
		t.Errorf("GetRuleIdsToRun method should not return error!")
	}
}

func TestGetRuleIdsToRun_BaselineDisabledAndHasSpecificRuleIds_returnSpecificRuleIds(t *testing.T) {
	//Given
	opts := options.Options{
		BaselineScan: false,
		Rules:        "ApexClassNoSharing,XSSTooltip",
	}
	expectedStandardRuleIds := []rules.RuleID{"ApexClassNoSharing", "XSSTooltip"}

	//When
	actualStandardRuleIds, actualCustomRuleIds, err := GetRuleIdsToRun(nil, &opts)

	//Then
	if !reflect.DeepEqual(actualStandardRuleIds, expectedStandardRuleIds) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Standard ruleIds are mismatched!", actualStandardRuleIds, expectedStandardRuleIds)
	}
	if !reflect.DeepEqual(len(actualCustomRuleIds), 0) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Number of custom rule Ids should be 0!", len(actualCustomRuleIds), 0)
	}
	if err != nil {
		t.Errorf("GetRuleIdsToRun method should not return error!")
	}
}

func TestGetRuleIdsToRun_BaselineDisabledAndHasCICDRuleIds_returnCicdRuleIds(t *testing.T) {
	//Given
	opts := options.Options{
		CICDScan: true,
	}
	configFile := config.Config{
		CICDRules: []string{
			"ApexClassNoSharing",
		},
	}

	expectedCicdRuleIds := []rules.RuleID{"ApexClassNoSharing"}

	//When
	actualCicdRuleIds, actualCustomRuleIds, err := GetRuleIdsToRun(&configFile, &opts)

	//Then
	if !reflect.DeepEqual(actualCicdRuleIds, expectedCicdRuleIds) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "CICD ruleIds are mismatched!", actualCicdRuleIds, expectedCicdRuleIds)
	}
	if !reflect.DeepEqual(len(actualCustomRuleIds), 0) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Number of custom rule Ids should be 0!", len(actualCustomRuleIds), 0)
	}
	if err != nil {
		t.Errorf("GetRuleIdsToRun method should not return error!")
	}
	opts.CICDScan = false
}

func TestGetRuleIdsToRun_BaselineDisabledAndHasCustomRulesInCICDRuleIds_returnCicdRuleIds(t *testing.T) {
	//Given
	opts := options.Options{
		CICDScan: true,
	}
	enable := true

	configFile := config.Config{
		CICDRules: []string{
			"ApexClassNoSharing",
			"CustomRule1",
		},
		CustomRegexRules: map[string]config.CustomRegexRule{
			"CustomRule1": {
				Name:           "customName1",
				Description:    "Please fix this",
				Enabled:        &enable,
				Severity:       "High",
				RuleCategory:   "Security",
				Pattern:        "Label",
				IncludePattern: "\\.component$|\\.page$|\\.cls$|\\.email",
				ExcludePattern: "",
			},
		},
	}

	expectedCicdStandardRuleIds := []rules.RuleID{"ApexClassNoSharing"}
	expectedCicdCustomRuleIds := []rules.RuleID{"CustomRule1"}

	//When
	actualCicdStandardRuleIds, actualCicdCustomRuleIds, err := GetRuleIdsToRun(&configFile, &opts)

	//Then
	if !reflect.DeepEqual(actualCicdStandardRuleIds, expectedCicdStandardRuleIds) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Standard CICD ruleIds are mismatched!", actualCicdStandardRuleIds, expectedCicdStandardRuleIds)
	}
	if !reflect.DeepEqual(actualCicdCustomRuleIds, expectedCicdCustomRuleIds) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Custom CICD ruleIds are mismatched!", actualCicdCustomRuleIds, expectedCicdCustomRuleIds)
	}
	if err != nil {
		t.Errorf("GetRuleIdsToRun method should not return error!")
	}
	opts.CICDScan = false
}

func TestGetRuleIdsToRun_BaselineDisabledAndConfigEnableAllStandardRulesEnabled_returnAllStandardRuleIds(t *testing.T) {
	//Given
	ENABLED_TRUE := true
	opts := options.Options{
		BaselineScan: false,
	}
	configFile := config.Config{
		EnableAllStandardRules: &ENABLED_TRUE,
		CustomRegexRules: map[string]config.CustomRegexRule{
			"customRule1": {},
			"customRule2": {},
		},
	}

	expectedStandardRuleIds := GetAllStdRuleIDs()

	//When
	actualStandardRuleIds, actualCustomRuleIds, err := GetRuleIdsToRun(&configFile, &opts)

	//Then
	if !reflect.DeepEqual(actualStandardRuleIds, expectedStandardRuleIds) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Standard ruleIds are mismatched!", actualStandardRuleIds, expectedStandardRuleIds)
	}
	if !reflect.DeepEqual(len(actualCustomRuleIds), 2) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Number of Custom ruleIds should be 2!", len(actualCustomRuleIds), 2)
	}
	if err != nil {
		t.Errorf("GetRuleIdsToRun method should not return error!")
	}
}

func TestGetRuleIdsToRun_BaselineDisabledAndConfigEnableAllStandardRulesDisabled_returnAllStandardRuleIds(t *testing.T) {
	//Given
	ENABLED_FALSE := false
	opts := options.Options{
		BaselineScan: false,
	}
	configFile := config.Config{
		EnableAllStandardRules: &ENABLED_FALSE,
		CustomRegexRules: map[string]config.CustomRegexRule{
			"customRule1": {},
			"customRule2": {},
		},
	}

	//When
	actualStandardRuleIds, actualCustomRuleIds, err := GetRuleIdsToRun(&configFile, &opts)

	//Then
	if !reflect.DeepEqual(len(actualStandardRuleIds), 0) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Number of Standard ruleIds should be 0!", len(actualStandardRuleIds), 0)
	}
	if !reflect.DeepEqual(len(actualCustomRuleIds), 2) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Number of Custom ruleIds should be 2!", len(actualCustomRuleIds), 2)
	}
	if err != nil {
		t.Errorf("GetRuleIdsToRun method should not return error!")
	}
}

func TestGetRuleIdsToRun_OptionsAllDisabledAndConfigFileNil_returnAllStandardRuleIds(t *testing.T) {
	//Given
	opts := options.Options{}

	expectedStandardRuleIds := GetAllStdRuleIDs()

	//When
	actualStandardRuleIds, actualCustomRuleIds, err := GetRuleIdsToRun(nil, &opts)

	//Then
	if !reflect.DeepEqual(actualStandardRuleIds, expectedStandardRuleIds) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Standard ruleIds are mismatched!", actualStandardRuleIds, expectedStandardRuleIds)
	}
	if !reflect.DeepEqual(len(actualCustomRuleIds), 0) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Number of Custom ruleIds should be 0!", len(actualCustomRuleIds), 0)
	}
	if err != nil {
		t.Errorf("GetRuleIdsToRun method should not return error!")
	}
}

func TestGetRuleIdsToRun_SpecificRuleHasInvalidRuleId_ReturnsInvalidRuleId(t *testing.T) {
	//Given
	opts := options.Options{
		Rules: "InvalidId",
	}

	//When
	actualStandardRuleIds, actualCustomRuleIds, err := GetRuleIdsToRun(nil, &opts)

	//Then
	if !reflect.DeepEqual(len(actualCustomRuleIds), 0) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Standard ruleIds are mismatched!", len(actualStandardRuleIds), 0)
	}
	if !reflect.DeepEqual(len(actualCustomRuleIds), 0) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Number of custom rule Ids should be 0!", len(actualCustomRuleIds), 0)
	}
	if err != nil {
		t.Errorf("GetRuleIdsToRun method should not return error!")
	}
}

func TestGetRuleIdsToRun_SpecificRuleHasCustomRuleId_ReturnsCustomRuleId(t *testing.T) {
	//Given
	opts := options.Options{
		Rules: "CustomRule1",
	}
	enable := true
	configFile := config.Config{
		CustomRegexRules: map[string]config.CustomRegexRule{
			"CustomRule1": {
				Name:           "customName1",
				Description:    "Please fix this",
				Enabled:        &enable,
				Severity:       "High",
				RuleCategory:   "Security",
				Pattern:        "Label",
				IncludePattern: "\\.component$|\\.page$|\\.cls$|\\.email",
				ExcludePattern: "",
			},
		},
	}

	expectedCustomRuleIds := []rules.RuleID{"CustomRule1"}

	//When
	actualStandardRuleIds, actualCustomRuleIds, err := GetRuleIdsToRun(&configFile, &opts)

	//Then
	if !reflect.DeepEqual(len(actualStandardRuleIds), 0) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Number of standard rule Ids should be 0!", len(actualCustomRuleIds), 0)
	}
	if !reflect.DeepEqual(actualCustomRuleIds, expectedCustomRuleIds) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Custom ruleIds are mismatched!", actualCustomRuleIds, expectedCustomRuleIds)
	}
	if err != nil {
		t.Errorf("GetRuleIdsToRun method should not return error!")
	}
}

func TestGetRuleIdsToRun_OverrideRuleHasInvalidRuleId_ReturnsAllStandardRuleIds(t *testing.T) {
	//Given
	opts := options.Options{}

	configFile := config.Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{
			"InvalidId": {},
		},
	}

	expectedStandardRuleIds := GetAllStdRuleIDs()

	//When
	actualStandardRuleIds, actualCustomRuleIds, err := GetRuleIdsToRun(&configFile, &opts)
	//Then
	if !reflect.DeepEqual(actualStandardRuleIds, expectedStandardRuleIds) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Standard ruleIds are mismatched!", actualStandardRuleIds, expectedStandardRuleIds)
	}
	if !reflect.DeepEqual(len(actualCustomRuleIds), 0) {
		t.Errorf("%s Actual: %+v, Expected: %+v", "Number of custom rule Ids should be 0!", len(actualCustomRuleIds), 0)
	}
	if err != nil {
		t.Errorf("GetRuleIdsToRun method should not return error!")
	}
}

func TestCreateRules_WhenStandardAndCustomRuleIds_ReturnRules(t *testing.T) {
	//Given
	standardRuleIds := []rules.RuleID{"ApexClassNoSharing"}
	customRuleIds := []rules.RuleID{"CustomRule1"}
	enable := true
	configFile := config.Config{
		CustomRegexRules: map[string]config.CustomRegexRule{
			"CustomRule1": {
				Name:           "customName1",
				Description:    "Please fix this",
				Enabled:        &enable,
				Severity:       "High",
				RuleCategory:   "Security",
				Pattern:        "Label",
				IncludePattern: "\\.component$|\\.page$|\\.cls$|\\.email",
				ExcludePattern: "",
			},
		},
	}
	expectedRulesCount := 2
	//When
	actualRules, err := CreateAndOverrideRules(standardRuleIds, customRuleIds, &configFile)

	//Then
	if len(actualRules) != expectedRulesCount {
		t.Errorf("Rules count mismatched!\n Actual %v, Expected %v", len(actualRules), expectedRulesCount)
	}
	if err != nil {
		t.Errorf("CreateAndOverrideRules should not return error.")
	}
}

func TestCreateRules_WhenStandardRulesOverrided_ReturnOverridedRules(t *testing.T) {
	//Given
	standardRuleIds := []rules.RuleID{"ApexClassNoSharing"}
	customRuleIds := []rules.RuleID{}
	configFile := config.Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{
			"ApexClassNoSharing": {
				Severity: "Low",
			},
		},
	}
	expectedRulesSeverity := rules.Severity("Low")

	//When
	actualRules, err := CreateAndOverrideRules(standardRuleIds, customRuleIds, &configFile)

	//Then
	if (*actualRules[0]).GetMetadata().Severity != expectedRulesSeverity {
		t.Errorf("Rule Override mismatched!!\n Actual %v, Expected %v", (*actualRules[0]).GetMetadata().Severity, expectedRulesSeverity)
	}
	if err != nil {
		t.Errorf("CreateAndOverrideRules method should not return error.")
	}
}

func TestCreateRules_WhenStandardRules_ReturnOverridedRules(t *testing.T) {
	//Given
	standardRuleIds := []rules.RuleID{"ApexClassNoSharing"}
	customRuleIds := []rules.RuleID{}
	expectedRulesCount := 1

	//When
	actualRules, err := CreateAndOverrideRules(standardRuleIds, customRuleIds, nil)

	//Then
	if len(actualRules) != expectedRulesCount {
		t.Errorf("Rule count mismatched!\n Actual %v, Expected %v", len(actualRules), expectedRulesCount)
	}
	if err != nil {
		t.Errorf("CreateAndOverrideRules method should not return error")
	}
}

func TestCreateRules_WhenStandardRuleIdInvalid_ReturnsNoRules(t *testing.T) {
	//Given
	invalidRuleId := rules.RuleID("xyz")
	standardRuleIds := []rules.RuleID{invalidRuleId}
	customRuleIds := []rules.RuleID{}

	//When
	actualRules, err := CreateAndOverrideRules(standardRuleIds, customRuleIds, nil)

	//Then
	if len(actualRules) != 0 {
		t.Errorf("Rule count mismatched!\n Actual %v, Expected %v", len(actualRules), 0)
	}
	if err != nil {
		t.Errorf("CreateAndOverrideRules method should not return error")
	}
}
