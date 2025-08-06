package test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/certinia/asist/rules/standard/codequality"
	"github.com/certinia/asist/rules/standard/security"
	"github.com/certinia/asist/scanner"
)

func TestLightningImproperCSSLoadRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 7,
		Results: []PartialFinding{
			{ID: "LightningImproperCSSLoad", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/aura/aura.css"), ColumnRange: []int{14, 25}, LineNumber: 11}},
			{ID: "LightningImproperCSSLoad", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/component/component.component"), ColumnRange: []int{0, 28}, LineNumber: 33}},
			{ID: "LightningImproperCSSLoad", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/component/component.component"), ColumnRange: []int{0, 33}, LineNumber: 35}},
			{ID: "LightningImproperCSSLoad", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/component/component.component"), ColumnRange: []int{0, 59}, LineNumber: 37}},
			{ID: "LightningImproperCSSLoad", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{0, 7}, LineNumber: 40}},
			{ID: "LightningImproperCSSLoad", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{0, 25}, LineNumber: 62}},
			{ID: "LightningImproperCSSLoad", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{12, 23}, LineNumber: 63}},
		},
	}

	createData("./src", security.LightningImproperCSSLoadRuleID, "")
	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestInsecureEndpointRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 6,
		Results: []PartialFinding{
			{ID: "InsecureEndpoint", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/component/component.component"), ColumnRange: []int{0, 41}, LineNumber: 28}},
			{ID: "InsecureEndpoint", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/component/component.component"), ColumnRange: []int{0, 46}, LineNumber: 31}},
			{ID: "InsecureEndpoint", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/component/sessionID.component"), ColumnRange: []int{0, 29}, LineNumber: 5}},
			{ID: "InsecureEndpoint", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{0, 46}, LineNumber: 59}},
			{ID: "InsecureEndpoint", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{0, 41}, LineNumber: 60}},
			{ID: "InsecureEndpoint", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{0, 64}, LineNumber: 61}},
		},
	}

	createData("./src", security.InsecureEndpointRuleID, "")
	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestAuraComponentCssExposedRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 6,
		Results: []PartialFinding{
			{ID: "AuraComponentCssExposed", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/aura/aura.cmp"), ColumnRange: []int{14, 25}, LineNumber: 5}},
			{ID: "AuraComponentCssExposed", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/aura/aura.cmp"), ColumnRange: []int{14, 31}, LineNumber: 6}},
			{ID: "AuraComponentCssExposed", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/aura/aura.css"), ColumnRange: []int{4, 19}, LineNumber: 14}},
			{ID: "AuraComponentCssExposed", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/aura/aura.css"), ColumnRange: []int{4, 16}, LineNumber: 15}},
			{ID: "AuraComponentCssExposed", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/aura/aura.css"), ColumnRange: []int{4, 15}, LineNumber: 27}},
			{ID: "AuraComponentCssExposed", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/aura/aura.css"), ColumnRange: []int{4, 22}, LineNumber: 28}},
		},
	}
	createData("./src", security.AuraComponentCssExposedRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestSensitiveInfoInDebugRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 5,
		Results: []PartialFinding{
			{ID: "SensitiveInfoInDebug", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/emailSending.cls"), ColumnRange: []int{8, 21}, LineNumber: 12}},
			{ID: "SensitiveInfoInDebug", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/emailSending.cls"), ColumnRange: []int{8, 21}, LineNumber: 13}},
			{ID: "SensitiveInfoInDebug", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/sessionID.cls"), ColumnRange: []int{4, 17}, LineNumber: 3}},
			{ID: "SensitiveInfoInDebug", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/sessionID.cls"), ColumnRange: []int{4, 17}, LineNumber: 4}},
			{ID: "SensitiveInfoInDebug", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/sessionID.cls"), ColumnRange: []int{4, 17}, LineNumber: 9}},
		},
	}

	createData("./src", security.SensitiveInfoInDebugRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestJSNotInStaticResourceRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 6,
		Results: []PartialFinding{
			{ID: "JSNotInStaticResource", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/component/component.component"), ColumnRange: []int{0, 17}, LineNumber: 28}},
			{ID: "JSNotInStaticResource", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/component/component.component"), ColumnRange: []int{0, 34}, LineNumber: 29}},
			{ID: "JSNotInStaticResource", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/component/component.component"), ColumnRange: []int{0, 62}, LineNumber: 30}},
			{ID: "JSNotInStaticResource", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{0, 31}, LineNumber: 59}},
			{ID: "JSNotInStaticResource", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{0, 17}, LineNumber: 60}},
			{ID: "JSNotInStaticResource", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{0, 40}, LineNumber: 61}},
		},
	}
	createData("./src", security.JSNotInStaticResourceRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestExposedMessageChannelRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 1,
		Results: []PartialFinding{
			{ID: "ExposedMessageChannel", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/metadata/projectMessageChannel.messageChannel-meta.xml"), ColumnRange: []int{1, 28}, LineNumber: 4}},
		},
	}
	createData("./src", security.ExposedMessageChannelRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestProtectedCustomSettingRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 2,
		Results: []PartialFinding{
			{ID: "ProtectedCustomSetting", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/metadata/ProtectedCustomSettingNewApiProperty__c.object-meta.xml"), ColumnRange: []int{1, 35}, LineNumber: 8}},
			{ID: "ProtectedCustomSetting", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/metadata/ProtectedCustomSettingOldApiProperty__c.object-meta.xml"), ColumnRange: []int{1, 63}, LineNumber: 8}},
		},
	}
	createData("./src", security.ProtectedCustomSettingRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestXssEscapeFalseRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 3,
		Results: []PartialFinding{
			{ID: "XSSEscapeFalse", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{22, 36}, LineNumber: 7}},
			{ID: "XSSEscapeFalse", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{71, 85}, LineNumber: 7}},
			{ID: "XSSEscapeFalse", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{120, 134}, LineNumber: 7}},
		},
	}
	createData("./src", security.XSSEscapeFalseRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestXssLabelRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 8,
		Results: []PartialFinding{
			{ID: "XSSLabel", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/component/component.component"), ColumnRange: []int{82, 92}, LineNumber: 21}},
			{ID: "XSSLabel", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{17, 40}, LineNumber: 37}},
			{ID: "XSSLabel", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{16, 39}, LineNumber: 41}},
			{ID: "XSSLabel", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{61, 83}, LineNumber: 41}},
			{ID: "XSSLabel", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{16, 39}, LineNumber: 42}},
			{ID: "XSSLabel", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{111, 129}, LineNumber: 46}},
			{ID: "XSSLabel", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{32, 42}, LineNumber: 48}},
			{ID: "XSSLabel", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{15, 25}, LineNumber: 49}},
		},
	}
	createData("./src", security.XSSLabelRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestEmailInjectionRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 2,
		Results: []PartialFinding{
			{ID: "EmailInjection", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/emailSending.cls"), ColumnRange: []int{13, 26}, LineNumber: 26}},
			{ID: "EmailInjection", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/emailSending.cls"), ColumnRange: []int{191, 204}, LineNumber: 26}},
		},
	}

	createData("./src", security.EmailInjectionRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestRichTextRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 1,
		Results: []PartialFinding{
			{ID: "XSSIsRichText", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/metadata/meta.field-meta.xml"), ColumnRange: []int{6, 16}, LineNumber: 1}},
		},
	}

	createData("./src", security.XSSIsRichTextRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestXSSDomHtmlRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 5,
		Results: []PartialFinding{
			{ID: "XSSDomHtml", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/component/component.component"), ColumnRange: []int{34, 51}, LineNumber: 23}},
			{ID: "XSSDomHtml", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/javascript/javascript.js"), ColumnRange: []int{13, 24}, LineNumber: 25}},
			{ID: "XSSDomHtml", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/javascript/javascript.js"), ColumnRange: []int{13, 24}, LineNumber: 28}},
			{ID: "XSSDomHtml", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{20, 31}, LineNumber: 53}},
			{ID: "XSSDomHtml", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{16, 27}, LineNumber: 55}},
		},
	}

	createData("./src", security.XSSDomHtmlRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestXssTooltipRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 5,
		Results: []PartialFinding{
			{ID: "XSSTooltip", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/javascript/javascript.js"), ColumnRange: []int{0, 17}, LineNumber: 30}},
			{ID: "XSSTooltip", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/javascript/javascript.js"), ColumnRange: []int{1, 11}, LineNumber: 31}},
			{ID: "XSSTooltip", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/javascript/javascript.js"), ColumnRange: []int{4, 13}, LineNumber: 33}},
			{ID: "XSSTooltip", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/javascript/javascript.js"), ColumnRange: []int{1, 11}, LineNumber: 34}},
			{ID: "XSSTooltip", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/javascript/javascript.js"), ColumnRange: []int{36, 45}, LineNumber: 35}},
		},
	}

	createData("./src", security.XSSTooltipRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)
	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestXssAuraUnescapedHtmlRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 1,
		Results: []PartialFinding{
			{ID: "XSSAuraUnescapedHtml", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/aura/aura.cmp"), ColumnRange: []int{4, 24}, LineNumber: 2}},
		},
	}

	createData("./src", security.XSSAuraUnescapedHtmlRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestXssCurrentPageParameterRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 2,
		Results: []PartialFinding{
			{ID: "XSSCurrentPageParameters", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/component/component.component"), ColumnRange: []int{53, 79}, LineNumber: 25}},
			{ID: "XSSCurrentPageParameters", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{52, 78}, LineNumber: 56}},
		},
	}

	createData("./src", security.XSSCurrentPageParametersRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestXssOccurrenceSearchRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 2,
		Results: []PartialFinding{
			{ID: "XSSLocationSearch", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/html/html.html"), ColumnRange: []int{16, 31}, LineNumber: 5}},
			{ID: "XSSLocationSearch", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/javascript/javascript.js"), ColumnRange: []int{11, 26}, LineNumber: 5}},
		},
	}

	createData("./src", security.XSSLocationSearchRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestXssSrcdocRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 2,
		Results: []PartialFinding{
			{ID: "XSSSrcDoc", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/aura/aura.cmp"), ColumnRange: []int{12, 19}, LineNumber: 4}},
			{ID: "XSSSrcDoc", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/html/html.html"), ColumnRange: []int{12, 19}, LineNumber: 9}},
		},
	}

	createData("./src", security.XSSSrcDocRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestXssMergeField(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 2,
		Results: []PartialFinding{
			{ID: "XSSMergeField", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/component/component.component"), ColumnRange: []int{13, 20}, LineNumber: 26}},
			{ID: "XSSMergeField", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{15, 25}, LineNumber: 57}},
		},
	}

	createData("./src", security.XSSMergeFieldRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestXssFormActionRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 1,
		Results: []PartialFinding{
			{ID: "XSSFormAction", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/aura/aura.cmp"), ColumnRange: []int{47, 70}, LineNumber: 3}},
		},
	}

	createData("./src", security.XSSFormActionRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestXSSApexChartRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 2,
		Results: []PartialFinding{
			{ID: "XSSApexChart", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/component/component.component"), ColumnRange: []int{4, 17}, LineNumber: 3}},
			{ID: "XSSApexChart", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{4, 16}, LineNumber: 9}},
		},
	}

	createData("./src", security.XSSApexChartRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestXssEscapeFalseInJSRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 2,
		Results: []PartialFinding{
			{ID: "XSSEscapeFalseInJS", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/component/component.component"), ColumnRange: []int{7, 19}, LineNumber: 16}},
			{ID: "XSSEscapeFalseInJS", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/page/visualforce.page"), ColumnRange: []int{7, 22}, LineNumber: 22}},
		},
	}

	createData("./src", security.XSSEscapeFalseInJSRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestXSSLwcDomManualRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count:   1,
		Results: []PartialFinding{{ID: "XSSLwcDomManual", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.html"), ColumnRange: []int{37, 53}, LineNumber: 5}}},
	}
	createData("./src", security.XSSLwcDomManualRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestXSSJavascriptButtonRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 1,
		Results: []PartialFinding{
			{ID: "XSSJavascriptButton", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/metadata/abc.webLink-meta.xml"), ColumnRange: []int{1, 32}, LineNumber: 6}},
		},
	}

	createData("./src", security.XSSJavascriptButtonRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//thn
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestApexClassWithoutSharingRuleIDRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 3,
		Results: []PartialFinding{
			{ID: "ApexClassWithoutSharing", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/emailSending.cls"), ColumnRange: []int{0, 49}, LineNumber: 45}},
			{ID: "ApexClassWithoutSharing", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/sampleClass.cls"), ColumnRange: []int{0, 32}, LineNumber: 1}},
			{ID: "ApexClassWithoutSharing", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/sampleClass.cls"), ColumnRange: []int{0, 46}, LineNumber: 6}},
		},
	}
	createData("./src", security.ApexClassWithoutSharingRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestApexClassNoSharingRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 4,
		Results: []PartialFinding{
			{ID: "ApexClassNoSharing", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/emailSending.cls"), ColumnRange: []int{0, 25}, LineNumber: 1}},
			{ID: "ApexClassNoSharing", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/emailSending.cls"), ColumnRange: []int{0, 17}, LineNumber: 39}},
			{ID: "ApexClassNoSharing", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/emailSending.cls"), ColumnRange: []int{0, 26}, LineNumber: 41}},
			{ID: "ApexClassNoSharing", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/emailSending.cls"), ColumnRange: []int{0, 33}, LineNumber: 43}},
		},
	}
	createData("./src", security.ApexClassNoSharingRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestHardcodedCredentialsRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 1,
		Results: []PartialFinding{
			{ID: "HardcodedCredentials", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/emailSending.cls"), ColumnRange: []int{7, 24}, LineNumber: 3}},
		},
	}
	createData("./src", security.HardcodedCredentialsRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestInsecureCryptoAlgorithmRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 2,
		Results: []PartialFinding{
			{ID: "InsecureCryptoAlgorithm", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/sampleClass.cls"), ColumnRange: []int{26, 32}, LineNumber: 3}},
			{ID: "InsecureCryptoAlgorithm", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/sampleClass.cls"), ColumnRange: []int{22, 27}, LineNumber: 5}},
		},
	}
	createData("./src", security.InsecureCryptoAlgorithmRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestDetectMissingAccessibilityModifierRule(t *testing.T) {
	//Given
	expectedResult := PartialOutput{
		Count: 3,
		Results: []PartialFinding{
			{ID: "DetectMissingAccessibilityModifier", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/sampleClass.cls"), ColumnRange: []int{0, 121}, LineNumber: 9}},
			{ID: "DetectMissingAccessibilityModifier", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/sampleClass.cls"), ColumnRange: []int{0, 33}, LineNumber: 13}},
			{ID: "DetectMissingAccessibilityModifier", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/class/sampleClass.cls"), ColumnRange: []int{0, 50}, LineNumber: 23}},
		},
	}
	createData("./src", codequality.DetectMissingAccessibilityModifierRuleID, "")

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestLwcNonStandardPositioningRule(t *testing.T) {
	// Given
	createData("./src", security.LwcNonStandardPositioningRuleID, "")
	expectedResult := PartialOutput{
		Count: 24,
		Results: []PartialFinding{
			// CSS issues
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.css"), LineNumber: 6, ColumnRange: []int{4, 19}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.css"), LineNumber: 7, ColumnRange: []int{4, 16}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.css"), LineNumber: 19, ColumnRange: []int{4, 15}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.css"), LineNumber: 20, ColumnRange: []int{4, 22}}},
			// HTML classes
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.html"), LineNumber: 6, ColumnRange: []int{16, 31}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.html"), LineNumber: 7, ColumnRange: []int{16, 32}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.html"), LineNumber: 8, ColumnRange: []int{16, 32}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.html"), LineNumber: 9, ColumnRange: []int{16, 29}}},
			// JS classes and style
			// classList.add;
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.js"), LineNumber: 4, ColumnRange: []int{13, 26}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.js"), LineNumber: 5, ColumnRange: []int{13, 29}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.js"), LineNumber: 6, ColumnRange: []int{13, 29}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.js"), LineNumber: 7, ColumnRange: []int{13, 28}}},
			// element.style.<styleType> = '<value>';
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.js"), LineNumber: 12, ColumnRange: []int{34, 55}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.js"), LineNumber: 13, ColumnRange: []int{34, 49}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.js"), LineNumber: 14, ColumnRange: []int{37, 51}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.js"), LineNumber: 15, ColumnRange: []int{37, 55}}},
			// CSS classes in getter
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.js"), LineNumber: 19, ColumnRange: []int{16, 29}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.js"), LineNumber: 19, ColumnRange: []int{30, 46}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.js"), LineNumber: 19, ColumnRange: []int{47, 63}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.js"), LineNumber: 19, ColumnRange: []int{64, 79}}},
			// Style in getter
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.js"), LineNumber: 23, ColumnRange: []int{16, 28}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.js"), LineNumber: 23, ColumnRange: []int{30, 41}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.js"), LineNumber: 23, ColumnRange: []int{43, 61}}},
			{ID: "LwcNonStandardPositioning", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/lwc.js"), LineNumber: 23, ColumnRange: []int{63, 79}}},
		},
	}

	// When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	// Then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}

func TestDetectImportJavascriptFromFileRule(t *testing.T) {
	//Given
	createData("./src", codequality.DetectImportJavascriptFromFileRuleID, "")
	expectedResult := PartialOutput{
		Count: 46,
		Results: []PartialFinding{
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 48}, LineNumber: 3}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 46}, LineNumber: 4}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 43}, LineNumber: 7}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 39}, LineNumber: 8}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 41}, LineNumber: 9}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 51}, LineNumber: 10}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 50}, LineNumber: 11}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 50}, LineNumber: 12}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 68}, LineNumber: 13}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 56}, LineNumber: 14}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 64}, LineNumber: 15}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 54}, LineNumber: 16}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 24}, LineNumber: 17}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 43}, LineNumber: 20}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 39}, LineNumber: 21}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 41}, LineNumber: 22}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 51}, LineNumber: 23}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 50}, LineNumber: 24}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 50}, LineNumber: 25}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 68}, LineNumber: 26}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 56}, LineNumber: 27}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 64}, LineNumber: 28}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 54}, LineNumber: 29}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 24}, LineNumber: 30}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 43}, LineNumber: 33}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 39}, LineNumber: 34}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 42}, LineNumber: 35}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 53}, LineNumber: 36}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 51}, LineNumber: 37}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 52}, LineNumber: 38}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 71}, LineNumber: 39}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 57}, LineNumber: 40}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 65}, LineNumber: 41}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 54}, LineNumber: 42}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{0, 24}, LineNumber: 43}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{1, 49}, LineNumber: 46}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{1, 47}, LineNumber: 47}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{1, 49}, LineNumber: 48}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{1, 61}, LineNumber: 49}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{1, 60}, LineNumber: 50}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{1, 59}, LineNumber: 51}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{1, 80}, LineNumber: 52}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{1, 67}, LineNumber: 53}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{1, 74}, LineNumber: 54}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{1, 63}, LineNumber: 55}},
			{ID: "DetectImportJavascriptFromFile", Occurrence: PartialOccurrence{FileName: GetAbsPath("src/lwc/import_javascript/import_javascript.js"), ColumnRange: []int{1, 28}, LineNumber: 56}},
		},
	}

	//When
	actualResult, _ := scanner.RunRulesOnFiles(filePaths, ruleInstances)

	//Then
	if !reflect.DeepEqual(expectedResult, projectOutputToPartial(*actualResult)) {
		actualResultJson, _ := json.MarshalIndent(projectOutputToPartial(*actualResult), "", "  ")
		expectedResultJson, _ := json.MarshalIndent(expectedResult, "", "  ")
		t.Errorf("%s \nActual: %+v, \n\nExpected: %+v", "Actual and expected results are not Equal.", string(actualResultJson), string(expectedResultJson))
	}
}
