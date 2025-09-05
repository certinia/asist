## What's Changed

## \[1.1.0\] \- 2025-09-05

### Added

* Added coloring for logs, errors, and warnings, like red `Error` text for errors and yellow `Warning` text for warnings.
* Made compatibility changes in the extension to handle ANSI escape sequences coming from colored text outputs from the ASIST CLI.

### Changed

* Converted the InvalidRuleId error into a warning; now the execution will continue even if an invalid rule ID is encountered.

### Fixed

* Fixed custom rules not getting picked in CI/CD and specific rule options.

### [Full Changelog](https://github.com/certinia/asist/blob/main/changelog.md)