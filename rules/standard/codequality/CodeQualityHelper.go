package codequality

import (
	"regexp"
	"slices"
	"strings"

	"github.com/certinia/asist/files"
	"github.com/certinia/asist/parser/options"
	"github.com/certinia/asist/utils"
)

// ******************************** FIND FUNCTION DEFINITION REGEX - START ********************************
var caseInsensitive = `(?i)`

// SELECT ,a-z,0-9,_,.,<,>,[,] AND AT LEAST ONE SPACE
var returnType = `\b[a-z0-9_.\[\]]+\s*(<[a-z0-9_.,\s<>]+>)?\s+`

// SELECT ,a-z,0-9 WORD AND SPACE IF PRESENT
var functionName = `\w+\s*`

/*
*

	\( 									 	- SELECT OPEN PARENTHESIS
		\s* 							 	- IF ANY SPACE PRESENT AFTER OPEN PARENTHESIS
		(
			(
				(final\s+)?					- USE TO SELECT FINAL KEYWORD
				[a-z0-9_.\[\]]+		 		- USE TO SELECT DATA_TYPE EXP: WHERE IN WORD CONTAINS LOWER(a-z), UPPER(), DIGIT(0-9),ARRAY(\[\])
				(\s*<[a-z0-9,_. <>]+>)* 	- USE TO MAINTAIN OPEN(<) CLOSE(>) ANGULAR BRACKET REASON - FALSE POSITIVE(else if(a > b){)
				\s+				 		 	- AFTER THE DATA_TYPE MUST REQUIRE ONE SPACE
				[a-z0-9_\[\]]+		 		- USE TO SELECT ARGUMENT NAME
			)?					 		 	- To make this group optional as there is a possibility that no arguments are present inside parenthesis.
			(					 		 	- IF MORE THEN ONE PARAMETER PRESENT THEN THIS GROUP EXECUTE
				\s*,\s*			 		 	- SELECT COMMA(,)
				(final\s+)?					- USE TO SELECT FINAL KEYWORD
				[a-z0-9_.\[\]]+ 		 	- USE TO SELECT DATA_TYPE EXP: WHERE IN WORD CONTAINS LOWER(a-z), UPPER(), DIGIT(0-9),ARRAY(\[\])
				(\s*<[a-z0-9,_. <>]+>)* 	- USE TO MAINTAIN OPEN(<) CLOSE(>) ANGULAR BRACKET REASON - FALSE POSITIVE(else if(a > b){)
				\s+						 	- AFTER THE DATA_TYPE MUST REQUIRE ONE SPACE
				[a-z0-9_\[\]]+		 		- USE TO SELECT ARGUMENT NAME
			)*							 	- IF THIS GROUP MATCH PRESENT NONE OR MULTIPLE TIME
		)
		\s*								 	- IF ANY SPACE PRESENT BEFORE CLOSE PARENTHESIS
	\)									 	- SELECT CLOSE PARENTHESIS
*/
var functionArguments = `\(\s*(((final\s+)?[a-z0-9_.\[\]]+(\s*<[a-z0-9,_.\s<>]+>)?\s+[a-z0-9_\[\]]+)?(\s*,\s*(final\s+)?[a-z0-9_.\[\]]+(\s*<[a-z0-9,_.\s<>]+>)?\s+[a-z0-9_\[\]]+)*)\s*\)`
var openCurlyBraces = `\s*\{`

//******************************** FIND FUNCTION DEFINITION REGEX - END ********************************
/**
Find the commented Line
Explaination -
	(\/\/.*) - SELECT THE SIGNLE LINE COMMENT
	|
	(\/\*+(.*?)\*+\/) - SELECT THE MULTILINE COMMENT
*/
var commentsRegexp = regexp.MustCompile(`(\/\/.*)|(\/\*+(.*?)\*+\/)`)
var classRegexp = regexp.MustCompile(`(?i)\bclass\s+\w+`)

type functionsDetail struct {
	lineIndex          int
	functionDefinition string
}

/**
 * getFunctionsList - method used to get the list of all the functions in a file
 */
func getFunctionsList(currentRuleId string, fileToScan files.File) []functionsDetail {
	fileContentAsString, lengthOfEachLine := convertFileIntoSingleString(currentRuleId, fileToScan)
	constructor := getConstructorRegex(fileContentAsString)
	functionOrConstructor := ""

	if constructor == "" {
		return nil
	}

	constructor += `\s*`
	functionOrConstructor = caseInsensitive +
		`(` +
		constructor + `|` +
		`(` +
		returnType +
		functionName +
		`)` +
		`)` +
		functionArguments +
		openCurlyBraces
	functionRegexp := regexp.MustCompile(functionOrConstructor)

	functionsDetails := findAllFunctions(fileContentAsString, lengthOfEachLine, functionRegexp)
	return functionsDetails
}

/**
 * findAllFunctions - method used to find detail of all the functions in string of file content
 */
func findAllFunctions(fileContentAsString string, lengthOfEachLine []int, functionRegexp *regexp.Regexp) []functionsDetail {
	functionsDetails := []functionsDetail{}
	//FIND ALL THE FUNCTION PRESENT IN CONCATE FILE STRING AND ITERATE ONE BY ONE
	for _, functionIndex := range functionRegexp.FindAllStringIndex(fileContentAsString, -1) {
		/**
		SUBSTRING THE LINE OF FUNCTION
		EXP:
		concateFileString = `class abc { public static void xyz(){ }}`
		o/p =  public static void xyz(){
		*/
		functionDefinition := fileContentAsString[functionIndex[0]:functionIndex[1]]
		// functionNameIndex := (functionIndex[0]+parenthesisRegexp.FindStringIndex(functionDefinition)[0])-1;
		/**
		FIND THE LINE NUMBER IN FILE WITH FunctionNameIndex AND lengthOfEachLine
		EXP:
			lengthOfEachLine = [11, 31, 36, 41, 55, 65, 65, 75, 85, 95, 104]
			functionNameIndex = 45
			WITH BINARY SEARCH FIND FIRST INDEX OF GREATER THAN 45
			O/P = 5
		*/
		lineIndex, _ := slices.BinarySearch(lengthOfEachLine, functionIndex[0])
		/**
		IN functionsDetails APPEND LINE NUMBER AND VIRTUAL LINE CONTENT
		EXP:
			abc.cls
			1	class abc{
			2		void xyz(
			3			int a
			4		){
			5
			6		}
			7	}

			functionDefinition = "void xyz(int a){""
		*/
		functionsDetails = append(functionsDetails,
			functionsDetail{
				lineIndex:          lineIndex,
				functionDefinition: functionDefinition,
			},
		)
	}
	return functionsDetails
}

/**
 * getConstructorRegex - method used to get constructor regex of a string as file content
 */
func getConstructorRegex(fileContentAsString string) string {
	classNames := []string{}
	constructorRegex := ""

	for _, classAndName := range classRegexp.FindAllString(fileContentAsString, -1) {
		className := strings.Fields(classAndName)[1]
		classNames = append(classNames, className)
	}
	if len(classNames) != 0 {
		constructorRegex = `(` + utils.CreateRegex(classNames) + `)`
	}
	return constructorRegex
}

/**
 * convertFileIntoSingleString - method used to convert a file content into a single string
 */
func convertFileIntoSingleString(currentRuleId string, fileToScan files.File) (string, []int) {
	fileContentAsString := ""
	sumOfCurrentAndPreviousLineLength := 0
	/**
	TO STORE THE LENGTH OF EACH LINE
	EXP : "previos_line_length+current_line_length"
	*/
	lengthOfEachLine := []int{}

	for _, line := range fileToScan.Lines {
		isFalsePositive := fileToScan.IsLineMarkedFalsePositive(currentRuleId, line.LineNumber)
		if !line.IsCommentedLine && (!isFalsePositive || options.IsBaselineScan()) {
			commentIndexs := commentsRegexp.FindStringIndex(line.Text)
			lineLength := len(line.Text)
			if len(commentIndexs) != 0 {
				/*IGNORE COMMENT
				EXP :
				line = `void abc(int a,/*comment*\/int b)`
				o/p = void abc(int a,int b)
				*/
				createWithoutCommentedCodeString := line.Text[0:commentIndexs[0]] + line.Text[commentIndexs[1]:lineLength]
				sumOfCurrentAndPreviousLineLength += len(createWithoutCommentedCodeString)
				fileContentAsString += createWithoutCommentedCodeString
			} else {
				sumOfCurrentAndPreviousLineLength += lineLength
				fileContentAsString += line.Text
			}
		}
		//IF : Line is commented then append the same length of previous line
		lengthOfEachLine = append(lengthOfEachLine, sumOfCurrentAndPreviousLineLength)
	}
	return fileContentAsString, lengthOfEachLine
}
