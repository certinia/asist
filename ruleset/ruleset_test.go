package ruleset

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/certinia/asist/config"
	"github.com/certinia/asist/parser/options"
	"github.com/certinia/asist/rules"
)

func TestGetAllRuleIDs_StandardRuleIds(t *testing.T) {
	//Given
	const STANDARD_RULES_COUNT = 32

	//When
	standardRuleIds := GetAllRuleIDs()

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
	expectedStandardRuleIds := GetAllRuleIDs()
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

	expectedStandardRuleIds := GetAllRuleIDs()
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

	expectedStandardRuleIds := GetAllRuleIDs()

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

	expectedStandardRuleIds := GetAllRuleIDs()

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

func TestGetRuleIdsToRun_SpecificRuleHasInvalidRuleId_InvalidRuleIdError(t *testing.T) {
	//Given
	opts := options.Options{
		Rules: "InvalidId",
	}
	expectedStandardRuleIds := "Invalid rule ID: InvalidId"

	//When
	_, _, actualInvalidRuleIdError := GetRuleIdsToRun(nil, &opts)

	//Then
	if strings.TrimSpace(actualInvalidRuleIdError.Error()) != expectedStandardRuleIds {
		t.Errorf("RuleIds mismatched")
	}
}

func TestGetRuleIdsToRun_OverrideRuleHasInvalidRuleId_InvalidRuleIdError(t *testing.T) {
	//Given
	opts := options.Options{}

	configFile := config.Config{
		RuleOverrides: map[string]rules.RuleMetadataOverride{
			"InvalidId": {},
		},
	}
	expectedError := "Invalid rule ID: InvalidId"

	//When
	_, _, actualInvalidRuleIdError := GetRuleIdsToRun(&configFile, &opts)

	//Then
	if strings.TrimSpace(actualInvalidRuleIdError.Error()) != expectedError {
		t.Errorf("RuleIds mismatched")
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

func TestCreateRules_WhenStandardRuleIdInvalid_ReturnInvalidRuleIdError(t *testing.T) {
	//Given
	invalidRuleId := rules.RuleID("xyz")
	standardRuleIds := []rules.RuleID{invalidRuleId}
	customRuleIds := []rules.RuleID{}
	expectedInvalidRuleIdError := fmt.Sprintf("Invalid rule ID: %s", invalidRuleId)

	//When
	_, actualInvalidRuleIdError := CreateAndOverrideRules(standardRuleIds, customRuleIds, nil)

	//Then
	if strings.TrimSpace(actualInvalidRuleIdError.Error()) != expectedInvalidRuleIdError {
		t.Errorf("Actual and expected errors mismatched!\n Actual %s, Expected %s", actualInvalidRuleIdError.Error(), expectedInvalidRuleIdError)
	}
}
