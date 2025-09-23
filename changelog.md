# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

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
