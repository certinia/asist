export const ASIST = "ASIST";
export const EXISTING_CONFIG_FILE = "A config file exists already!";
export const NO_WORKSPACE = "There is no workspace!";
export const NO_OPEN_FILE = "There is no open file!";
export const NO_CONFIG_FILE = "There is no config file!";
export const SCANNING_WORKSPACE = "Scanning workspace...";
export const WINDOWS_OS_TYPE = "Windows_NT";
export const YAML_CONFIG_FILE = ".asist.yaml";
export const JSON_CONFIG_FILE = ".asist.json";
export const EMPTY_STRING = "";
export const PATH_SEPARATOR = "/";

export const OPTIONS = {
	LIST_RULES: "-l",
	CONFIG: "-c"
};

export const COMMANDS = {
	LIST_RULES: "extension.listRules",
	SCAN_WORKSPACE: "extension.scanWorkspace",
	SCAN_FILE: "extension.scanFile",
	PREFERENCES: "extension.preferences",
	CREATE_CONFIG: "extension.createConfig",
	EDIT_CONFIG: "extension.editConfig"
};

export const ASIST_CONFIGURATION_OPTIONS = {
	CUSTOM_BINARY: "ASIST.customBinary",
	CUSTOM_CONFIG_FILE_PATH: "configFilePath"
};

export const ASIST_BINARY_OS_TYPE = {
	WINDOWS: "./asist.exe",
	DARWIN_AMD64: "./asist_darwin_amd64",
	DARWIN_ARM64: "./asist_darwin_arm64",
	LINUX_ARM64: "./asist_linux_arm64",
	LINUX_AMD64: "./asist_linux_amd64"
};

export const OS_TYPE = {
	WINDOWS: "win32/x64",
	WINDOWS_ARM64: "win32/arm64",
	DARWIN_AMD64: "darwin/x64",
	DARWIN_ARM64: "darwin/arm64",
	LINUX_ARM64: "linux/arm64"
};
