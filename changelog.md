# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## \[1.2.0\] \- 2026-03-06

### Added

* Added property to `cicdmaxissues` in `ruleoverrides` config, allowing a maximum number of issues to be permitted per rule in CI/CD mode. This enables gradual remediation of existing issues by setting a threshold that can be reduced over time. Exceeding the threshold exits with a non-zero code and prints a violation summary to stderr.
* When running `--list-rules`, each rule now displays its `CicdMaxIssues` value. Rules without a `cicdmaxissues` override show a default value of `0`.
* Added `cicdmaxissues` support to `customregexrules`, allowing custom rules to define their own CI/CD issue threshold directly in their rule definition.

### Fixed

* Enhanced the regex to be case-insensitive and to support `.` in repository names, enabling correct extraction of the repository name from SSH URLs in baseline scan.

## \[1.1.1\] \- 2025-09-26

### Fixed

* Fixed `-V` flag not able to display version for go package.

## \[1.1.0\] \- 2025-09-23

### Added

* Added coloring for logs, errors, and warnings, like red `Error` text for errors and yellow `Warning` text for warnings for ASIST CLI.

### Changed

* Converted the InvalidRuleId error into a warning; now the execution will continue even if an invalid rule ID is encountered.

### Fixed

* Fixed custom rules not getting picked in CI/CD and specific rule options.

## \[1.0.0\] \- 2025-08-11

### Added

* **Core Functionality:**  
  * Introduced the primary feature: `SAST tool to scan a file or an entire directory for security vulnerabilities, during and after development.`  
  * Language support: `It is mainly built for Salesforce code, but compatible with any text file. It helps developers find potential security vulnerabilities in their code early in the development process.`  
  * Command-line interface (CLI) for easy execution.  
  * Configuration options via a YAML and JSON file.
  * Introduced the official VS Code extension for ASIST, enabling direct IDE integration for SAST analysis.
