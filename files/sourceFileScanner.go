package files

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

/*
Regex to check commented line is start
EXP :
 1. <!-- SOME COMMENTED CODE OR WORDS
 2. /* or /** or /*(n number of star) SOME COMMENTED CODE OR WORDS
*/
var startOfCommentedLineRegexp = regexp.MustCompile(`^\s*(/\*+|<!--)`)

// Regex to check commented line is end
//
//	EXP :
//		1) SOME COMMENTED CODE OR WORDS -->
//		2) SOME COMMENTED CODE OR WORDS */ or **/ or (n number of star)*/
var endOfCommentedLineRegexp = regexp.MustCompile(`\*+/\s*$|-->\s*$`)

// Regex to check commented line  is end  between the code
//
//	EXP :
//		1) SOME COMMENTED CODE OR WORDS --> Executable code
//		2) SOME COMMENTED CODE OR WORDS */ Executable code
var endOfcommentedLineBetweenRegexp = regexp.MustCompile(`\*+/|-->`)

// Regex to find line is single line commented
var singleLineCommentedRegexp = regexp.MustCompile(`^\s*//`)

/*
Regex to find occurrance is false positive - BEGIN
EXP: asist-ignore-Begin:[RuleID1,RuleID2]
*/
var multiLineFalsePositiveBeginPatternRegexp = regexp.MustCompile(`asist\s*-\s*ignore\s*-\s*begin:\s*\[.*\]`)

/*
Regex to find occurrance is false positive - END
EXP: asist-ignore-End:[RuleID1,RuleID2]
*/
var multiLineFalsePositiveEndPatternRegexp = regexp.MustCompile(`asist\s*-\s*ignore\s*-\s*end`)

/*
Regex to find list of ruleIDs
EXP: [RuleID1, RuleID2]
*/
var selectBracketRegexp = regexp.MustCompile(`(\[.*\])`)

/**
 * Read - method used to read the file content and store in file struct
 */
func Read(filename string) (*File, error) {
	var fileLines []Line
	var ignoreSelectedLines []IgnoreSelected

	readFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	lineNumber := 0
	//Use as a flag to check line is commented or not
	isCommentedLine := false
	foundIgnoreMismatch := false
	lineContent := ""

	for fileScanner.Scan() {
		lineNumber += 1
		lineContent = fileScanner.Text()
		isEndOfCommentedLine := false
		//TRUE : IF isCommentedLine value is false
		if !isCommentedLine {
			isCommentedLine = startOfCommentedLineRegexp.MatchString(lineContent)
		}
		//TRUE : IF isCommentedLine value is true or we have found comment is started
		if isCommentedLine {
			//Use as a flag to check commented line is end and use also in Line.IsCommentedLine
			isEndOfCommentedLine = endOfCommentedLineRegexp.MatchString(lineContent)
			//TRUE : if In the line comment end tag is present end of the line
			//  SOME COMMENTED CODE OR WORDS -->
			//  SOME COMMENTED CODE OR WORDS */ or **/ or  (n number of star)*/
			if isEndOfCommentedLine {
				isCommentedLine = false
				//TRUE : if isCommentedLine flag value is true and In the line comment end tag is present between of the line
				//  SOME COMMENTED CODE OR WORDS --> Executable Code
				//  SOME COMMENTED CODE OR WORDS */ or **/ or  (n number of star)*/ Executable Code
			} else if endOfcommentedLineBetweenRegexp.MatchString(lineContent) {
				isCommentedLine = false
			}
		}

		isCommentedLine := (isCommentedLine || isEndOfCommentedLine || singleLineCommentedRegexp.MatchString(lineContent))

		if isCommentedLine && !foundIgnoreMismatch {
			//IF : "asist-ignore-end:[RuleIDs]" Present in lineContent
			if multiLineFalsePositiveEndPatternRegexp.MatchString(lineContent) {
				foundIgnoreMismatch = (len(ignoreSelectedLines) > 0 && ignoreSelectedLines[len(ignoreSelectedLines)-1].EndLine != -1) || len(ignoreSelectedLines) == 0
				if !foundIgnoreMismatch {
					ignoreSelectedLines[len(ignoreSelectedLines)-1].EndLine = lineNumber
				}
			} else {
				falsePositiveRange := multiLineFalsePositiveBeginPatternRegexp.FindStringIndex(lineContent)

				if len(falsePositiveRange) > 0 {
					foundIgnoreMismatch = len(ignoreSelectedLines) > 0 && ignoreSelectedLines[len(ignoreSelectedLines)-1].EndLine == -1

					if !foundIgnoreMismatch {
						ignoreSelectedLines = append(ignoreSelectedLines, getIgnoreSelectedLine(falsePositiveRange, lineContent, lineNumber))
					}
				}
			}
		}
		fileLines = append(fileLines, Line{
			LineNumber:      lineNumber,
			Text:            lineContent,
			IsCommentedLine: isCommentedLine,
		})
	}
	masterFile := File{Lines: fileLines, FileName: filename, IgnoresSelected: ignoreSelectedLines}
	return &masterFile, nil
}

/**
 * getIgnoreSelectedLine - method to get the lines numbers of code with RuleIds which are marked as false positive
 */
func getIgnoreSelectedLine(falsePositiveRange []int, lineContent string, lineNumber int) IgnoreSelected {
	falsePositiveComment := lineContent[falsePositiveRange[0]:falsePositiveRange[1]]
	//find the range of the RuleIDs bracket
	ruleIDsRange := selectBracketRegexp.FindStringIndex(falsePositiveComment)
	//MAP to add RuleIDs
	ruleIDsMAP := map[string]bool{}
	/*
		Substring of all ruleIDS present between bracket
		EXP : "[RuleID1,RuleId2...]" : string
		ruleIDs = "RuleID1,RuleId2..." : string
	*/
	ruleIDs := falsePositiveComment[ruleIDsRange[0]+1 : ruleIDsRange[1]-1]
	//Add all split RuleIDs to the map
	for _, ruleId := range strings.Split(ruleIDs, ",") {
		ruleIDsMAP[ruleId] = true
	}
	return IgnoreSelected{
		BeginLine: lineNumber,
		EndLine:   -1,
		RuleIDs:   ruleIDsMAP,
	}
}
