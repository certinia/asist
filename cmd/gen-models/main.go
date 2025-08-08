// The following code look for all rules under rules/standard directory and generates ruleset.go file with ruleId -> rule mapping

package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/certinia/asist/files"
	"github.com/dave/jennifer/jen"
)

type RuleData map[string][]string

const RULES_PKG = "github.com/certinia/asist/rules"
const STANDARD_PKG = "github.com/certinia/asist/rules/standard/"

func main() {
	ruleMap := make(RuleData)
	getStandardRules(ruleMap)

	ruleSetFile := jen.NewFile("ruleset")
	ruleSetFile.HeaderComment("// This file is dynamically generated using code generator and contains rule mapping")

	// Generate ruleId -> rule mapping in ruleset file
	addRuleMappingToFile(ruleSetFile, ruleMap)

	if err := ruleSetFile.Save("../../ruleset/ruleset.go"); err != nil {
		log.Fatalf("Error in saving ruleset file: %v", err)
	}
}

func extractRuleName(file string) string {
	base := filepath.Base(file)
	return strings.TrimSuffix(base, ".go")
}

func isRuleFile(file string) bool {
	return strings.HasSuffix(file, ".go") &&
		!strings.HasSuffix(file, "_test.go") &&
		!strings.Contains(file, "Helper")
}

func containingDir(filePath string) string {
	dir := filepath.Dir(filePath)
	return filepath.Base(dir)
}

func getStandardRules(ruleMap RuleData) {
	currDir, _ := os.Getwd()
	pathSeparator := string(os.PathSeparator)
	standardRulePath := strings.Replace(currDir, "cmd"+pathSeparator+"gen-models", "rules"+pathSeparator+"standard", 1)

	// Get all standard rule files
	fileOptions := files.FileOptions{
		RootPath:        standardRulePath,
		DontForceIgnore: true,
		DontGitIgnore:   true,
	}
	files, err := files.GetAllFilePaths(fileOptions)
	if err != nil {
		log.Fatalf("Not able to list files at location rules/standard: %v", err)
		return
	}

	// Create Package_name -> Ruleset mapping
	for _, file := range files {
		if !isRuleFile(file) {
			continue
		}
		ruleName := extractRuleName(file)
		packageName := containingDir(file)
		ruleMap[packageName] = append(ruleMap[packageName], ruleName)
	}
}

func addRuleMappingToFile(file *jen.File, ruleMap RuleData) {
	file.Var().Id("ruleMapping").Op("=").Map(
		jen.Qual(RULES_PKG, "RuleID"),
	).Qual(RULES_PKG, "Rule").BlockFunc(func(g *jen.Group) {
		for pkg, ruleList := range ruleMap {
			for _, rule := range ruleList {
				g.Qual(STANDARD_PKG+pkg, rule+"RuleID").
					Op(":").
					Qual(STANDARD_PKG+pkg, "New"+rule+"Rule").
					Call().Op(",")
			}
		}
	})
}
