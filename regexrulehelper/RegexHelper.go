package regexrulehelper

import (
	"path/filepath"
	"regexp"

	"github.com/certinia/asist/files"
	"github.com/certinia/asist/parser/options"
	"github.com/certinia/asist/rules"
)

/**
*	FindOccurancesForFile - Method will return occurrences found in the file
 */
func FindOccurancesForFile(fileToScan files.File, ruleToScan *rules.RuleMetadata, isCommentedLinesIncluded bool) []rules.Occurrence {
	var occurrences []rules.Occurrence
	compiledPattern := regexp.MustCompile(ruleToScan.Pattern)
	compiledQualifier := regexp.MustCompile(ruleToScan.Qualifier)
	foundQualifier := false

	for _, line := range fileToScan.Lines {
		isFalsePositive := fileToScan.IsLineMarkedFalsePositive(string(ruleToScan.ID), line.LineNumber)
		if (!isCommentedLinesIncluded && line.IsCommentedLine) || (isFalsePositive && !options.IsBaselineScan()) {
			continue
		}

		if len(ruleToScan.Qualifier) > 0 && !foundQualifier {
			foundQualifier = len(compiledQualifier.FindStringIndex(line.Text)) > 0
		}

		allOccurrencesInLine := compiledPattern.FindAllStringIndex(line.Text, -1)
		if len(allOccurrencesInLine) > 0 {
			for _, value := range allOccurrencesInLine {
				occurrences = append(
					occurrences,
					rules.Occurrence{
						FileName:        fileToScan.FileName,
						LineNumber:      line.LineNumber,
						LineContent:     line.Text,
						ColumnRange:     value,
						IsFalsePositive: isFalsePositive,
					},
				)
			}
		}
	}

	// If a qualifier is given and is not found, then return an empty array.
	// If a qualifier has been given and found, or not given, then return lines that matched the pattern.
	if len(ruleToScan.Qualifier) > 0 && !foundQualifier {
		var emptyOutput []rules.Occurrence
		return emptyOutput
	} else {
		return occurrences
	}
}

func RunIncludeExcludePatternsOnFile(fileName string, metaData rules.RuleMetadata) bool {

	if metaData.IncludePattern == "" && metaData.ExcludePattern == "" {
		return true
	}
	includePattern := regexp.MustCompile(metaData.IncludePattern)
	excludePattern := regexp.MustCompile(metaData.ExcludePattern)

	if (includePattern.MatchString("") || includePattern.MatchString(filepath.ToSlash(fileName))) && (excludePattern.MatchString("") || !excludePattern.MatchString(filepath.ToSlash(fileName))) {
		return true
	}
	return false
}

func replaceEncodedString(line string, tempStr string, pointer int) string {
	openingBracketBalance := 1
	length := len(line)
	for pointer < length {
		if line[pointer] == '(' {
			openingBracketBalance++
		}
		if line[pointer] == ')' {
			openingBracketBalance--
		}
		if openingBracketBalance == 0 {
			return tempStr
		}
		tempStr += "#"
		pointer++
	}
	return tempStr
}

func ReplaceEncodePartByHash(line string, encodeMethodsRegexp *regexp.Regexp) string {
	linePartReplacedByHash := ""
	pointer := 0

	for _, replaceRange := range encodeMethodsRegexp.FindAllStringIndex(line, -1) {
		/*If next encode occurrance is already replace by # then continue loop for that range
		EXP : {!JSENCODE(HTMLENCODE($Label.xyz) + HTMLENCODE($Label.xyz))}
		*/
		if replaceRange[0] < pointer {
			continue
		}

		linePartReplacedByHash += line[pointer:replaceRange[0]]
		pointer = replaceRange[0]
		for pointer < replaceRange[1] {
			linePartReplacedByHash += "#"
			pointer++
		}
		//Replace the value till the encode function bracket is closed
		linePartReplacedByHash = replaceEncodedString(line, linePartReplacedByHash, pointer)
		pointer = len(linePartReplacedByHash)
	}
	linePartReplacedByHash += line[pointer:]

	return linePartReplacedByHash
}
