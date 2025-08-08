import { workspace } from "vscode";
import { existsSync } from "fs";
import {
	ASIST,
	ASIST_CONFIGURATION_OPTIONS,
	EMPTY_STRING,
	JSON_CONFIG_FILE,
	PATH_SEPARATOR,
	YAML_CONFIG_FILE
} from "./constants";

/*
 * getConfigOption - method used to get the config option
 */
export function getConfigOption(): string {
	const configFilePath: string = workspace
		.getConfiguration(ASIST)
		.get(ASIST_CONFIGURATION_OPTIONS.CUSTOM_CONFIG_FILE_PATH);

	if (workspace.workspaceFolders !== undefined) {
		return getConfigFileOption(workspace.workspaceFolders[0].uri.fsPath, configFilePath);
	} else {
		return EMPTY_STRING;
	}
}

/*
 * getConfigFilePath - method used to get the config file path
 */
export function getConfigFilePath(workspacePath: string, configFilePath: string = null): string {
	if (!!configFilePath && checkConfigFileExists(workspacePath, configFilePath)) {
		return workspacePath + PATH_SEPARATOR + configFilePath;
	}

	if (checkConfigFileExists(workspacePath, YAML_CONFIG_FILE)) {
		return workspacePath + PATH_SEPARATOR + YAML_CONFIG_FILE;
	}

	if (checkConfigFileExists(workspacePath, JSON_CONFIG_FILE)) {
		return workspacePath + PATH_SEPARATOR + JSON_CONFIG_FILE;
	}

	return null;
}

/*
 * getConfigFileOption - method used to get config file option with path depending on the OS
 */
export function getConfigFileOption(workspacePath: string, configFilePath: string): string {
	if (!!configFilePath && checkConfigFileExists(workspacePath, configFilePath)) {
		return workspacePath + PATH_SEPARATOR + configFilePath;
	}

	if (checkConfigFileExists(workspacePath, YAML_CONFIG_FILE)) {
		return workspacePath + PATH_SEPARATOR + YAML_CONFIG_FILE;
	}

	if (checkConfigFileExists(workspacePath, JSON_CONFIG_FILE)) {
		return workspacePath + PATH_SEPARATOR + JSON_CONFIG_FILE;
	}

	return EMPTY_STRING;
}

/*
 * checkConfigFileExists - method used to check config file exists or not
 */
function checkConfigFileExists(workspacePath: string, configFile: string) {
	return existsSync(workspacePath + PATH_SEPARATOR + configFile);
}
