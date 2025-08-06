package ignore

import (
	"bufio"
	"os"
	"os/user"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
	gogitignore "github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

const (
	commentPrefix   = "#"
	gitDir          = ".git"
	sfDir           = ".sf"
	sfdxDir         = ".sfdx"
	forceignoreFile = ".forceignore"
	gitignoreFile   = ".gitignore"
	infoExcludeFile = gitDir + "/info/exclude"
)

var defaultPatterns = []gogitignore.Pattern{
	gogitignore.ParsePattern(gitDir, nil),
	gogitignore.ParsePattern(sfdxDir, nil),
	gogitignore.ParsePattern(sfDir, nil),
}

type IgnoreOptions struct {
	RootPath        string
	DontGitIgnore   bool
	DontForceIgnore bool
}

func GetIgnoreFilesPatterns(ignoreoptions IgnoreOptions) (gogitignore.Matcher, error) {
	var patterns []gogitignore.Pattern
	var fileSystem billy.Filesystem
	var ignoreFiles []string

	if ignoreoptions.DontGitIgnore && ignoreoptions.DontForceIgnore {
		return nil, nil
	}

	fileInfo, err := os.Stat(ignoreoptions.RootPath)

	if err != nil {
		return nil, err
	}

	if !fileInfo.IsDir() {
		return nil, nil
	}

	fileSystem = osfs.New(ignoreoptions.RootPath)

	if !ignoreoptions.DontGitIgnore {
		ignoreFiles = append(ignoreFiles, gitignoreFile)
		ignoreFiles = append(ignoreFiles, infoExcludeFile)

	}
	if !ignoreoptions.DontForceIgnore {
		ignoreFiles = append(ignoreFiles, forceignoreFile)

	}

	patterns, err = readPatterns(fileSystem, nil, ignoreFiles)
	if err != nil {
		return nil, err
	}

	patterns = append(patterns, defaultPatterns...)

	return gogitignore.NewMatcher(patterns), nil

}

func readIgnoreFile(fs billy.Filesystem, path []string, ignoreFile string) (ps []gogitignore.Pattern, err error) {

	ignoreFile, err = replaceTildeWithHome(ignoreFile)

	if err != nil {
		return nil, err
	}

	f, err := fs.Open(fs.Join(append(path, ignoreFile)...))

	if err == nil {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			s := scanner.Text()
			if !strings.HasPrefix(s, commentPrefix) && len(strings.TrimSpace(s)) > 0 {
				ps = append(ps, gogitignore.ParsePattern(s, path))
			}
		}
	} else if os.IsNotExist(err) {
		// Set err to nil if the ignore file doesn't exist, as it's an ignorable case
		err = nil
	}

	return
}

func replaceTildeWithHome(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		firstSlash := strings.Index(path, "/")
		if firstSlash == 1 {
			home, err := os.UserHomeDir()
			if err != nil {
				return path, err
			}
			return strings.Replace(path, "~", home, 1), nil
		} else if firstSlash > 1 {
			username := path[1:firstSlash]
			userAccount, err := user.Lookup(username)
			if err != nil {
				return path, err
			}
			return strings.Replace(path, path[:firstSlash], userAccount.HomeDir, 1), nil
		}
	}

	return path, nil
}

func readPatterns(fs billy.Filesystem, path []string, ignoreFiles []string) (ps []gogitignore.Pattern, err error) {

	for _, ignoreFile := range ignoreFiles {
		subps, err := readIgnoreFile(fs, path, ignoreFile)
		if err != nil {
			return nil, err
		}
		ps = append(ps, subps...)
	}

	var fis []os.FileInfo
	fis, err = fs.ReadDir(fs.Join(path...))
	if err != nil {
		return nil, err
	}

	for _, fi := range fis {
		if fi.IsDir() {
			if gogitignore.NewMatcher(ps).Match(append(path, fi.Name()), true) || gogitignore.NewMatcher(defaultPatterns).Match(append(path, fi.Name()), true) {
				continue
			}

			var subps []gogitignore.Pattern
			subps, err = readPatterns(fs, append(path, fi.Name()), ignoreFiles)
			if err != nil {
				return nil, err
			}

			if len(subps) > 0 {
				ps = append(ps, subps...)
			}
		}
	}

	return ps, err
}
