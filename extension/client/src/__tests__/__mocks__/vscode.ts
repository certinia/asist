export const append = jest.fn();
export const outputChannel = { append };
const openTextDocUriPath = "/mock/workspace/mockfile.yaml";
const workspaceFoldersUriPath = "/workspace";
const windowsDocUriPath = "/mock/file.js";

export const workspace = {
	fs: {
		readFile: jest.fn().mockResolvedValue(Buffer.from("template content")),
		writeFile: jest.fn().mockResolvedValue(undefined)
	},
	getConfiguration: jest.fn().mockReturnValue({
		get: jest.fn().mockReturnValue("mocked-file.yaml")
	}),
	createFileSystemWatcher: jest.fn(() => ({
		onDidCreate: jest.fn(),
		onDidChange: jest.fn(),
		onDidDelete: jest.fn(),
		dispose: jest.fn()
	})),
	openTextDocument: jest.fn().mockResolvedValue({ uri: { path: openTextDocUriPath } }),
	workspaceFolders: [
		{
			uri: {
				fsPath: workspaceFoldersUriPath
			}
		}
	]
};
export const commands = {
	registerCommand: jest.fn(),
	executeCommand: jest.fn()
};

export const Disposable = class {
	dispose() {}
};

export const window = {
	createOutputChannel: jest.fn(() => outputChannel),
	showInformationMessage: jest.fn(),
	showErrorMessage: jest.fn(),
	showTextDocument: jest.fn().mockResolvedValue(undefined),
	openTextDocument: jest.fn().mockResolvedValue({}),
	withProgress: jest.fn(),
	activeTextEditor: {
		document: {
			uri: { fsPath: windowsDocUriPath },
			save: jest.fn()
		}
	}
};

export const Uri = {
	file: jest.fn(path => ({ fsPath: path })),
	parse: jest.fn((uri: string) => ({
		fsPath: uri,
		toString: () => uri
	}))
};

export const ProgressLocation = {
	Notification: 1
};
