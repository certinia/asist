package options

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/certinia/asist/debugger"
	"github.com/certinia/asist/errorhandler"
	"github.com/certinia/asist/message"
	"github.com/certinia/asist/rules"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	RepoURL      string `short:"u" long:"repo-url" required:"false" description:"URL of the repo. Used for baseline scan output"`
	ConfigFile   string `short:"c" long:"config" required:"false" description:"JSON or YAML config file to read from"`
	Rules        string `short:"r" long:"rules" required:"false" description:"Rules comma separated to run (ignore rules enabled/disabled in config)"`
	ListRules    bool   `short:"l" long:"list-rules" required:"false" description:"List rules which would be run"`
	BaselineScan bool   `short:"b" long:"baseline-scan" required:"false" description:"For getting output of ASIST baseline scan as count of occurrences and false positive occurrences, number of custom rules occurrences, type of record and this data is used for creating metrics."`
	CICDScan     bool   `short:"j" long:"cicd-rules" required:"false" description:"For use in CI/CD pipelines. Tells ASIST to only run the CICD rules defined in config file. If there are no CI/CD rules defined, no rule will run. If there are any occurrences, returns a non-zero exit code which will make the pipeline step fail"`
	Debug        bool   `short:"v" long:"verbose" required:"false" description:"Print out debug messages with time elapsed since last message"`
	Version      bool   `short:"V" long:"version" required:"false" description:"Display the current version of ASIST binary"`

	Args struct {
		Path string `description:"Path to the file or folder to scan"`
	} `positional-args:"yes"`
}

var opts Options

func GetRepoURL() string {
	return opts.RepoURL
}
func IsCICDScan() bool {
	return opts.CICDScan
}

func Initilize() *Options {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(int(errorhandler.ExitCodeUserError))
	}
	validation()
	setup()
	return &opts
}

func GetPathToScan() string {
	if opts.ListRules || opts.Version {
		return ""
	}
	validation()
	path, err := filepath.Abs(opts.Args.Path)

	if err != nil {
		errorhandler.ExitWithCode(message.GetPathFetchingError(err), errorhandler.ExitCodeUserError)
	}
	return path
}

func IsBaselineScan() bool {
	return opts.BaselineScan
}

func IsListRules() bool {
	return opts.ListRules
}

func IsVersion() bool {
	return opts.Version
}

func (o *Options) SpecificRuleIds() []rules.RuleID {
	ruleIds := []rules.RuleID{}

	for _, ruleId := range strings.Split(o.Rules, ",") {
		ruleIds = append(ruleIds, rules.RuleID(ruleId))
	}
	return ruleIds
}

func setup() {
	if opts.Debug {
		debugger.EnableDebugMode()
	}
}

func validation() {
	if len(opts.Args.Path) == 0 && !opts.ListRules && !opts.Version {
		errorhandler.ExitWithCode(message.GetMissingFileOrFolderError(), errorhandler.ExitCodeUserError)
	}
}
