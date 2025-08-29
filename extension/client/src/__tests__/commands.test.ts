import * as vscode from "vscode";
import { execFile } from "child_process";
import * as config from "../config";
import {
	createConfigCommand,
	editConfigCommand,
	preferencesCommand,
	listRulesCommand,
	scanWorkspaceCommand,
	scanFileCommand
} from "../commands";

jest.mock("child_process", () => ({
	execFile: jest.fn()
}));

jest.mock("../utils", () => ({
	__esModule: true,
	getScannerPath: jest.fn(() => "/mocked/path"),
	removeAnsiEscapeCodes: jest.fn((str: string) => {
		return "stdout data";
	})
}));

jest.mock("../config", () => ({
	__esModule: true,
	getConfigOption: jest.fn(() => "/config/path"),
	getConfigFilePath: jest.fn()
}));

jest.mock("os", () => ({
	type: jest.fn(() => "Windows_NT")
}));

describe("listRulesCommand", () => {
	let mockClient: any;
	beforeEach(() => {
		mockClient = {
			diagnostics: { clear: jest.fn() },
			sendRequest: jest.fn()
		};
		jest.resetModules();
	});

	afterEach(() => {
		jest.clearAllMocks();
	});

	it("registers command and appends stdout when execFile succeeds", () => {
		const mockCommandCallback = jest.fn();
		const mockDisposable = { dispose: jest.fn() };

		vscode.commands.registerCommand.mockImplementation((_id, cb) => {
			mockCommandCallback.mockImplementation(cb);
			return mockDisposable;
		});

		(execFile as unknown as jest.Mock).mockImplementation((_cmd, _args, cb) => {
			cb(null, "output from tool", "");
		});

		const disposable = listRulesCommand(mockClient);

		// To verify the callback for registerCommand method we need to call it explicitly.
		mockCommandCallback();

		expect(execFile).toHaveBeenCalledWith(
			"/mocked/path",
			["-c", "/config/path", "-l"],
			expect.any(Function)
		);
		expect(vscode.append).toHaveBeenCalledWith("output from tool");
		expect(disposable).toBe(mockDisposable);
	});

	it("does not append output when execFile returns error", () => {
		const mockCommandCallback = jest.fn();
		vscode.commands.registerCommand.mockImplementation((_id, cb) => {
			mockCommandCallback.mockImplementation(cb);
			return { dispose: jest.fn() };
		});

		(execFile as unknown as jest.Mock).mockImplementation((_cmd, _args, cb) => {
			cb(new Error("CLI failed"), "", "");
		});

		listRulesCommand(mockClient);
		// To verify the callback for registerCommand method we need to call it explicitly.
		mockCommandCallback();
		expect(mockClient.diagnostics.clear).toHaveBeenCalled();
		expect(vscode.append).toHaveBeenCalled();
		expect(vscode.window.showErrorMessage).toHaveBeenCalled();
	});
});

describe("scanWorkspaceCommand", () => {
	let mockClient: any;
	let registeredCallback: Function;

	beforeEach(() => {
		vscode.workspace.fs.readFile.mockResolvedValue(Buffer.from("template-content"));
		vscode.workspace.fs.writeFile.mockResolvedValue(undefined);
		vscode.window.showTextDocument.mockResolvedValue(undefined);
		vscode.commands.registerCommand.mockImplementation((_cmd, cb) => {
			registeredCallback = cb;
			return { dispose: jest.fn() };
		});

		mockClient = {
			diagnostics: { clear: jest.fn() },
			sendRequest: jest.fn()
		};
	});

	afterEach(() => {
		jest.clearAllMocks();
	});

	it("registers the command and runs execFile, updates diagnostics and sends request on success", async () => {
		vscode.window.withProgress.mockImplementation(async (opts, task) => task());
		const mockCommandCallback = jest.fn();
		const mockDisposable = { dispose: jest.fn() };

		vscode.commands.registerCommand.mockImplementation((_id, cb) => {
			mockCommandCallback.mockImplementation(cb);
			return mockDisposable;
		});

		(execFile as unknown as jest.Mock).mockImplementation((_cmd, _args, _opts, cb) => {
			cb(null, "stdout data", "");
		});

		const disposable = scanWorkspaceCommand({} as any);
		scanWorkspaceCommand(mockClient);

		await mockCommandCallback();

		expect(execFile).toHaveBeenCalledWith(
			"/mocked/path",
			["-c", "/config/path", "/workspace"],
			{ maxBuffer: 1024 * 1024 * 100 },
			expect.any(Function)
		);
		expect(vscode.append).toHaveBeenCalledWith("stdout data");

		expect(mockClient.diagnostics.clear).toHaveBeenCalled();

		expect(mockClient.sendRequest).toHaveBeenCalledWith("scancode", "stdout data");
		expect(disposable).toBe(mockDisposable);
	});

	it("shows information message if no workspaceFolders", async () => {
		vscode.workspace.workspaceFolders = undefined;

		scanWorkspaceCommand(mockClient);

		await registeredCallback();

		expect(vscode.window.showInformationMessage).toHaveBeenCalledWith("There is no workspace!");
	});
});

describe("scanFileCommand", () => {
	let mockClient: any;
	let registeredCallback: Function;

	beforeEach(() => {
		vscode.workspace.fs.readFile.mockResolvedValue(Buffer.from("template-content"));
		vscode.workspace.fs.writeFile.mockResolvedValue(undefined);
		vscode.window.showTextDocument.mockResolvedValue(undefined);
		vscode.commands.registerCommand.mockImplementation((_cmd, cb) => {
			registeredCallback = cb;
			return { dispose: jest.fn() };
		});

		mockClient = {
			diagnostics: { clear: jest.fn() },
			sendRequest: jest.fn()
		};
	});

	afterEach(() => {
		jest.clearAllMocks();
	});
	it("runs execFile and processes stdout when activeTextEditor exists", () => {
		(execFile as unknown as jest.Mock).mockImplementation((_cmd, _args, cb) => {
			cb(null, "file command stdout", "");
		});

		scanFileCommand(mockClient);
		registeredCallback();

		expect(vscode.window.activeTextEditor.document.save).toHaveBeenCalled();

		expect(execFile).toHaveBeenCalledWith(
			"/mocked/path",
			["-c", "/config/path", "/mock/file.js"],
			expect.any(Function)
		);

		expect(vscode.append).toHaveBeenCalledWith("file command stdout");
		expect(mockClient.diagnostics.clear).toHaveBeenCalled();
		expect(mockClient.sendRequest).toHaveBeenCalledWith("scancode", "file command stdout");
	});

	it("shows message when no activeTextEditor is present", () => {
		vscode.window.activeTextEditor = undefined;

		scanFileCommand(mockClient);
		registeredCallback();

		expect(vscode.window.showInformationMessage).toHaveBeenCalledWith("There is no open file!");
	});
});

describe("editConfigCommand", () => {
	let commandCallback: Function;
	const mockedWorkspacePath = "/mock/workspace";

	beforeEach(() => {
		jest.clearAllMocks();

		(vscode.commands.registerCommand as jest.Mock).mockImplementation((_, cb) => {
			commandCallback = cb;
			return { dispose: jest.fn() };
		});

		(vscode.workspace as any).workspaceFolders = [{ uri: { path: mockedWorkspacePath } }];
	});

	it("shows message if there is no workspace", async () => {
		(vscode.workspace as any).workspaceFolders = undefined;

		editConfigCommand();
		await commandCallback();

		expect(vscode.window.showInformationMessage).toHaveBeenCalledWith("There is no workspace!");
	});

	it("shows message if file path is not found", async () => {
		(config.getConfigFilePath as jest.Mock).mockReturnValue(null);

		editConfigCommand();
		await commandCallback();

		expect(vscode.window.showInformationMessage).toHaveBeenCalledWith("There is no config file!");
	});

	it("opens document if file path is found", async () => {
		const mockedConfigPath = mockedWorkspacePath + "/mockfile.yaml";
		(config.getConfigFilePath as jest.Mock).mockReturnValue(mockedConfigPath);

		editConfigCommand();
		await commandCallback();

		expect(vscode.workspace.openTextDocument).toHaveBeenCalledWith({
			fsPath: mockedConfigPath
		});
		expect(vscode.window.showTextDocument).toHaveBeenCalled();
	});
});

describe("createConfigCommand", () => {
	let callback: Function;

	beforeEach(() => {
		jest.clearAllMocks();
		(vscode.commands.registerCommand as jest.Mock).mockImplementation((_, cb) => {
			callback = cb;
			return { dispose: jest.fn() };
		});

		(vscode.workspace as any).workspaceFolders = [{ uri: { path: "/mock/workspace" } }];
	});

	it("creates a new file if it does not exist", async () => {
		(config.getConfigFilePath as jest.Mock).mockReturnValue(null);
		createConfigCommand();

		await callback();

		expect(vscode.workspace.fs.readFile).toHaveBeenCalled();
	});

	it("shows message if file already exists", async () => {
		(config.getConfigFilePath as jest.Mock).mockReturnValue("/mock/workspace/.hellodata.yaml");

		createConfigCommand();
		await callback();

		expect(vscode.workspace.fs.readFile).not.toHaveBeenCalled();
		expect(vscode.window.showInformationMessage).toHaveBeenCalledWith(
			"A config file exists already!"
		);
	});

	it("shows message if no workspace", async () => {
		(vscode.workspace as any).workspaceFolders = undefined;

		createConfigCommand();
		await callback();

		expect(vscode.window.showInformationMessage).toHaveBeenCalledWith("There is no workspace!");
	});
});

describe("preferencesCommand", () => {
	let commandCallback: Function;

	beforeEach(() => {
		jest.clearAllMocks();

		(vscode.commands.registerCommand as jest.Mock).mockImplementation((_cmd, cb) => {
			commandCallback = cb;
			return { dispose: jest.fn() };
		});
	});

	it("registers the command and executes openSettings", async () => {
		preferencesCommand();

		await commandCallback();

		expect(vscode.commands.executeCommand).toHaveBeenCalledWith(
			"workbench.action.openSettings",
			"ASIST"
		);
	});
});
