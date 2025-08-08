import * as vscode from "vscode";
import * as path from "path";
import * as extension from "../../src/extension";
import { LanguageClient } from "vscode-languageclient/node";
import { listRulesCommand } from "../commands";

jest.mock("vscode-languageclient/node", () => {
	return {
		LanguageClient: jest.fn().mockImplementation(() => ({
			start: jest.fn(),
			stop: jest.fn(() => Promise.resolve())
		})),
		TransportKind: {
			ipc: 0
		}
	};
});

jest.mock("../../src/commands", () => {
	return {
		Commands: jest.fn(),
		listRulesCommand: jest.fn(() => {}),
		scanWorkspaceCommand: jest.fn(() => {}),
		scanFileCommand: jest.fn(() => {}),
		editConfigCommand: jest.fn(() => {}),
		createConfigCommand: jest.fn(() => {}),
		preferencesCommand: jest.fn(() => {})
	};
});

describe("Extension Activate and deactivate Function check", () => {
	const mockContext: vscode.ExtensionContext = {
		extensionPath: "/mock/path",
		asAbsolutePath: jest.fn(relPath => path.join("/mock/path", relPath)), // nosemgrep: path-join-resolve-traversal
		subscriptions: [],
		globalState: {} as any,
		workspaceState: {} as any,
		extensionUri: {} as any,
		environmentVariableCollection: {} as any,
		storageUri: {} as any,
		globalStorageUri: {} as any,
		logUri: {} as any,
		extensionMode: 1,
		extension: {} as any,
		secrets: {} as any
	};

	afterEach(() => {
		jest.resetModules();
		jest.clearAllMocks();
	});

	it("should start the client and register Commands", () => {
		extension.activate(mockContext);

		expect(LanguageClient).toHaveBeenCalledWith(
			"languageServer",
			"Language Server",
			expect.any(Object),
			expect.any(Object)
		);

		const clientInstance = (LanguageClient as jest.Mock).mock.results[0].value;
		expect(clientInstance.start).toHaveBeenCalled();
		expect(listRulesCommand).toHaveBeenCalled();
	});

	it("Should return undefined when client is not set or undefined", () => {
		const extension = require("../../src/extension");
		const result = extension.deactivate();
		expect(result).toBeUndefined();
	});

	it("when extension is deactivate then calls client.stop() and returns its result ", async () => {
		extension.activate(mockContext);
		const clientInstance = (LanguageClient as jest.Mock).mock.results[0].value;
		const result = extension.deactivate();
		expect(clientInstance.stop).toHaveBeenCalled();
		await expect(result).resolves.toBeUndefined();
	});
});
