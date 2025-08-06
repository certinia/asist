package security

import (
	"regexp"

	"github.com/certinia/asist/files"
	"github.com/certinia/asist/rules"
)

var XSSMergeFieldRuleID rules.RuleID = "XSSMergeField"

type XSSMergeFieldRule struct {
	metadata rules.RuleMetadata
}

func NewXSSMergeFieldRule() *XSSMergeFieldRule {
	return &XSSMergeFieldRule{
		metadata: rules.RuleMetadata{
			ID:             XSSMergeFieldRuleID,
			Name:           "Potential XSS with mergefield",
			Description:    "Using a merge field inside a script or style tag may introduce an XSS issue. Consider adding an encode method for the merge field; otherwise, ensure that all user inputs are sanitized properly.",
			Severity:       rules.SeverityHigh,
			RuleCategory:   rules.CategorySecurity,
			IncludePattern: "\\.component$|\\.page$",
			ExcludePattern: "",
			Pattern:        "\\{![^\\}]+\\}",
		},
	}
}

func (r *XSSMergeFieldRule) GetMetadata() *rules.RuleMetadata {
	return &r.metadata
}

func (r *XSSMergeFieldRule) Run(fileToScan files.File) []rules.Occurrence {
	return findXssMergeField(fileToScan, &r.metadata)
}

func findXssMergeField(fileToScan files.File, ruleMetadata *rules.RuleMetadata) []rules.Occurrence {
	var output []rules.Occurrence

	salesforceGlobalVariables := "\\$(Label|RemoteAction|Profile|Page|CurrentPage|Permission|Resource|Component|Action|User\\.UITheme|ObjectType[^\\}]*\\.(Name|Createable|Updateable|Deletable|KeyPrefix)[^\\.])"
	encodedFunction := "(((JS|HTML|JSINHTML)ENCODE|URLFOR)\\s*\\()"
	numberPattern := "(((discount|number|limit|Separator|size|count|enabled|length|total|[<>=+\\-*/]\\s*\\d+)\\s*})|({!((Min|Max)\\w*)|(\\w*(Min|Max))}))"
	booleanPattern := "({![a-z.]+\\.(is|can|has|show)\\w*[\\s)]*})"
	booleanCamelCasePattern := "({!((?i)(is|has|show|can|namespace))((_|[A-Z])\\w*)?([\\s=]+(true|false))?})"
	idPattern := "{![^\\}]*(\\.[iI]d|_?I[Dd]s?)[\\s\\)]*}"
	onEventRegexp := `\bon\w+\s*=\s*"(.*?"|.*)`

	patternToExcludeRegexp := regexp.MustCompile("((?i)" + salesforceGlobalVariables + "|" + booleanPattern + "|" + numberPattern + ")" + "|" + encodedFunction + "|" + booleanCamelCasePattern + "|" + idPattern)
	mergefieldRegexp := regexp.MustCompile(ruleMetadata.Pattern)

	for _, line := range findVulnerableLinesBetweenTags(fileToScan, string(ruleMetadata.ID), []string{onEventRegexp}, false) {
		subLineRange1 := 0
		subLineRange2 := 0
		tempLineText := line.LineContent

		if len(line.ColumnRange) != 0 {
			subLineRange1 = line.ColumnRange[0]
			subLineRange2 = line.ColumnRange[1]
			tempLineText = line.LineContent[subLineRange1:subLineRange2]
		}

		allOccurrencesInLine := mergefieldRegexp.FindAllStringIndex(tempLineText, -1)

		for _, mergefield_range := range allOccurrencesInLine {
			mergefield_str := tempLineText[mergefield_range[0]:mergefield_range[1]]

			//if not encoded then Append occurance into the output
			if !patternToExcludeRegexp.MatchString(mergefield_str) {
				output = append(
					output,
					rules.Occurrence{
						FileName:        fileToScan.FileName,
						LineContent:     line.LineContent,
						LineNumber:      line.LineNumber,
						ColumnRange:     []int{mergefield_range[0] + subLineRange1, mergefield_range[1] + subLineRange1},
						IsFalsePositive: line.IsFalsePositive,
					},
				)
			}
		}
	}
	return output
}
