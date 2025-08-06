import { execFile } from "child_process";
import { Disposable, LanguageClient } from "vscode-languageclient/node";
import { type } from "os";
import { getScannerPath } from "./utils";
import { join } from "path";
import { getConfigFilePath, getConfigOption } from "./config";
import { window, commands, workspace, ProgressLocation, Uri } from "vscode";
import {
	ASIST,
	ASIST_CONFIGURATION_OPTIONS,
	COMMANDS,
	EMPTY_STRING,
	EXISTING_CONFIG_FILE,
	NO_CONFIG_FILE,
	NO_OPEN_FILE,
	NO_WORKSPACE,
	OPTIONS,
	SCANNING_WORKSPACE,
	WINDOWS_OS_TYPE,
	YAML_CONFIG_FILE
} from "./constants";

const outputChannel = window.createOutputChannel(ASIST);

/*
 * listRulesCommand - method used to execute go binary and display active rules metadata
 */
export function listRulesCommand(client: LanguageClient): Disposable {
	return commands.registerCommand(COMMANDS.LIST_RULES, () => {
		const configFile = getConfigOption();
		const cmdOptions =
			configFile == null || configFile.trim() == EMPTY_STRING
				? [OPTIONS.LIST_RULES]
				: [OPTIONS.CONFIG, configFile, OPTIONS.LIST_RULES];
		execFile(getScannerPath(), cmdOptions, function callback(error, stdout, stderr) {
			const errorMessage = stderr ? stderr : error?.message;
			if (errorMessage) {
				client.diagnostics.clear();
				outputChannel.append(errorMessage);
				window.showErrorMessage(errorMessage);
				return;
			}
			outputChannel.append(stdout);
		});
	});
}

/*
 * scanWorkspaceCommand - method used to execute go binary and scan whole workspace
 */
export function scanWorkspaceCommand(client: LanguageClient): Disposable {
	return commands.registerCommand(COMMANDS.SCAN_WORKSPACE, () => {
		if (workspace.workspaceFolders !== undefined) {
			const PROGRESS_OPTIONS = {
				location: ProgressLocation.Notification,
				title: SCANNING_WORKSPACE,
				cancellable: false
			};

			window.withProgress(PROGRESS_OPTIONS, () => {
				return initiateScanWorkspaceTask(client);
			});
		} else {
			window.showInformationMessage(NO_WORKSPACE);
		}
	});
}

/*
 * initiateScanWorkspaceTask - method used to initiate scan workspace task with ASIST GO binary and display output
 */
function initiateScanWorkspaceTask(client: LanguageClient): Promise<void> {
	const CHILD_PROCESS_OPTIONS = { maxBuffer: 1024 * 1024 * 100 };
	return new Promise<void>(resolve => {
		const configFile = getConfigOption();
		const cmdOptions =
			configFile == null || configFile.trim() == EMPTY_STRING
				? [workspace.workspaceFolders[0].uri.fsPath]
				: [OPTIONS.CONFIG, configFile, workspace.workspaceFolders[0].uri.fsPath];
		execFile(getScannerPath(), cmdOptions, CHILD_PROCESS_OPTIONS, (error, stdout, stderr) => {
			displayOutputAndClearRequest(error, stdout, stderr, client);
			resolve();
		});
	});
}

/*
 * displayOutputAndClearRequest - method used to display output on output channel and completes request
 */
function displayOutputAndClearRequest(
	error,
	stdout: string,
	stderr: string,
	client: LanguageClient
) {
	const errorMessage = stderr ? stderr : error?.message;
	if (errorMessage) {
		client.diagnostics.clear();
		outputChannel.append(errorMessage);
		window.showErrorMessage(errorMessage);
		return;
	}
	outputChannel.append(stdout);
	client.diagnostics.clear();
	client.sendRequest("scancode", stdout);
}

/*
 * scanFileCommand - method used to execute the go binary and scan a particular file
 */
export function scanFileCommand(client: LanguageClient): Disposable {
	return commands.registerCommand(COMMANDS.SCAN_FILE, () => {
		if (window.activeTextEditor !== undefined) {
			window.activeTextEditor.document.save();
			const configFile = getConfigOption();
			const cmdOptions =
				configFile == null || configFile.trim() == EMPTY_STRING
					? [window.activeTextEditor.document.uri.fsPath]
					: [OPTIONS.CONFIG, configFile, window.activeTextEditor.document.uri.fsPath];
			execFile(getScannerPath(), cmdOptions, (error, stdout, stderr) => {
				displayOutputAndClearRequest(error, stdout, stderr, client);
			});
		} else {
			window.showInformationMessage(NO_OPEN_FILE);
		}
	});
}

/*
 * preferencesCommand - method used to display user prefererences
 */
export function preferencesCommand(): Disposable {
	return commands.registerCommand(COMMANDS.PREFERENCES, () => {
		commands.executeCommand("workbench.action.openSettings", ASIST);
	});
}

/*
 * createConfigCommand - method used to create a sample config file in workspace
 */
export function createConfigCommand(): Disposable {
	return commands.registerCommand(COMMANDS.CREATE_CONFIG, () => {
		if (workspace.workspaceFolders !== undefined) {
			const configPath = getConfigFilePath(workspace.workspaceFolders[0].uri.fsPath);
			const PATH_SEPARATOR = "/";
			if (configPath == null) {
				const TEMPLATE_CONFIG_FILENAME = "configfiletemplate.yaml";
				let templatePath = join(__dirname, "..", TEMPLATE_CONFIG_FILENAME);
				if (type() == WINDOWS_OS_TYPE) {
					templatePath = PATH_SEPARATOR + templatePath.replaceAll("\\", "/");
				}
				const configFileTemplate = Uri.parse("file://" + templatePath);
				const newConfigFile = Uri.parse(
					workspace.workspaceFolders[0].uri.path + PATH_SEPARATOR + YAML_CONFIG_FILE
				);
				workspace.fs.readFile(configFileTemplate).then(value => {
					workspace.fs.writeFile(newConfigFile, value).then(() => {
						window.showTextDocument(newConfigFile, { preview: false });
					});
				});
			} else {
				window.showInformationMessage(EXISTING_CONFIG_FILE);
			}
		} else {
			window.showInformationMessage(NO_WORKSPACE);
		}
	});
}

/*
 * editConfigCommand - method used to open the ASIST config file for editing
 */
export function editConfigCommand(): Disposable {
	return commands.registerCommand(COMMANDS.EDIT_CONFIG, () => {
		if (!workspace.workspaceFolders) {
			window.showInformationMessage(NO_WORKSPACE);

			return;
		}

		const configFilePath: string = workspace
			.getConfiguration(ASIST)
			.get(ASIST_CONFIGURATION_OPTIONS.CUSTOM_CONFIG_FILE_PATH);

		const configPath = getConfigFilePath(workspace.workspaceFolders[0].uri.fsPath, configFilePath);

		if (!configPath) {
			window.showInformationMessage(NO_CONFIG_FILE);

			return;
		}

		const openPath = Uri.file(configPath);

		if (openPath != null) {
			workspace.openTextDocument(openPath).then(doc => {
				window.showTextDocument(doc);
			});
		}
	});
}
