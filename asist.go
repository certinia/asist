package main

import (
	"time"

	"github.com/certinia/asist/errorhandler"
	"github.com/certinia/asist/output"
	"github.com/certinia/asist/scanner"
)

func main() {
	scanTime := output.ScanTime{
		StartedTime: time.Now().String(),
	}
	//Load required resources for scan
	paths, rules, err := scanner.LoadResources()
	if err != nil {
		errorhandler.ExitWithError(err)
	}
	//Run active rules on all files
	finalResult, err := scanner.RunRulesOnFiles(paths, rules)
	if err != nil {
		errorhandler.ExitWithError(err)
	}
	scanTime.EndingTime = time.Now().String()
	output.DisplayOutput(finalResult, &scanTime)
}
