import { existsSync } from "fs";
import { arch, platform } from "os";
import { fileURLToPath } from "url";
import { join } from "path";
import { ASIST_BINARY_OS_TYPE, OS_TYPE } from "./constants";

//Name of the configuration files. They should be included in the root of the project's workspace
const YAML_CONFIG_FILE = ".asist.yaml";
const JSON_CONFIG_FILE = ".asist.json";
const PATH_SEPARATOR = "/";

/*
 * checkAndGetConfigFileOption - method used to check if the configuration file exists and returns the path depending on the OS
 */
export function checkAndGetConfigFileOption(workspaceUri: string, configFilePath: string): string {
	let filePath: string | null = null;

	if (!!configFilePath && checkConfigFileExists(workspaceUri, configFilePath)) {
		filePath = fileURLToPath(workspaceUri + PATH_SEPARATOR + configFilePath);
	} else if (checkConfigFileExists(workspaceUri, YAML_CONFIG_FILE)) {
		filePath = fileURLToPath(workspaceUri + PATH_SEPARATOR + YAML_CONFIG_FILE);
	} else if (checkConfigFileExists(workspaceUri, JSON_CONFIG_FILE)) {
		filePath = fileURLToPath(workspaceUri + PATH_SEPARATOR + JSON_CONFIG_FILE);
	}

	if (filePath !== null) {
		return filePath;
	}

	return " ";
}

/*
 * checkConfigFileExists - method used to check if the configuration file  path exists or not
 */
function checkConfigFileExists(workspaceUri: string, configFile: string) {
	return existsSync(fileURLToPath(workspaceUri + PATH_SEPARATOR + configFile));
}

/*
 * getDefaultScannerPath - method returns the default path of go binary depending on the OS
 */
export function getDefaultScannerPath(): string {
	const platarch = `${platform()}/${arch()}`;
	const EXTENSION_ROOT_PATH = "../..";

	switch (platarch) {
		case OS_TYPE.WINDOWS:
			return join(__dirname, EXTENSION_ROOT_PATH, ASIST_BINARY_OS_TYPE.WINDOWS);
		case OS_TYPE.WINDOWS_ARM64:
			return join(__dirname, EXTENSION_ROOT_PATH, ASIST_BINARY_OS_TYPE.WINDOWS);
		case OS_TYPE.DARWIN_AMD64:
			return join(__dirname, EXTENSION_ROOT_PATH, ASIST_BINARY_OS_TYPE.DARWIN_AMD64);
		case OS_TYPE.DARWIN_ARM64:
			return join(__dirname, EXTENSION_ROOT_PATH, ASIST_BINARY_OS_TYPE.DARWIN_ARM64);
		case OS_TYPE.LINUX_ARM64:
			return join(__dirname, EXTENSION_ROOT_PATH, ASIST_BINARY_OS_TYPE.LINUX_ARM64);
		default: // "linux/amd64" or other OS
			return join(__dirname, EXTENSION_ROOT_PATH, ASIST_BINARY_OS_TYPE.LINUX_AMD64);
	}
}

/*
 * getFileExtension - method returns the extension of file
 */
export function getFileExtension(uri: string): string {
	const uriSplitted = uri.split(".");
	return uriSplitted[uriSplitted.length - 1];
}
