## What's Changed

## \[1.2.0\] \- 2026-03-06

### Added

* Added property to `cicdmaxissues` in `ruleoverrides` config, allowing a maximum number of issues to be permitted per rule in CI/CD mode. This enables gradual remediation of existing issues by setting a threshold that can be reduced over time. Exceeding the threshold exits with a non-zero code and prints a violation summary to stderr.
* When running `--list-rules`, each rule now displays its `CicdMaxIssues` value. Rules without a `cicdmaxissues` override show a default value of `0`.
* Added `cicdmaxissues` support to `customregexrules`, allowing custom rules to define their own CI/CD issue threshold directly in their rule definition.

### Fixed

* Enhanced the regex to be case-insensitive and to support `.` in repository names, enabling correct extraction of the repository name from SSH URLs in baseline scan.

### [Full Changelog](https://github.com/certinia/asist/blob/main/changelog.md)
