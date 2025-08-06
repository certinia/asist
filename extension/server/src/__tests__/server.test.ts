import { CodeActionKind, DidChangeConfigurationNotification } from "vscode-languageserver/node";
import * as scanner from "../../src/scanner";
import * as quickfix from "../../src/quickfix";
import { TextDocument } from "vscode-languageserver-textdocument";

const mockRegister = jest.fn();
const mockOnRequest = jest.fn();
const mockOnInitialize = jest.fn();
const mockOnInitialized = jest.fn();
const mockOnCodeAction = jest.fn();
const mockListen = jest.fn();
const mockOnDidOpen = jest.fn();
const mockOnDidSave = jest.fn();

const mockConnection = {
	onInitialize: mockOnInitialize,
	onInitialized: mockOnInitialized,
	onRequest: mockOnRequest,
	client: {
		register: mockRegister
	},
	onCodeAction: mockOnCodeAction,
	listen: mockListen
};

const mockDocuments = {
	onDidOpen: mockOnDidOpen,
	onDidSave: mockOnDidSave,
	listen: jest.fn(),
	get: jest.fn()
};
jest.mock("../../src/scanner");
jest.mock("../../src/quickfix");
jest.mock("vscode-languageserver/node", () => {
	const actual = jest.requireActual("vscode-languageserver/node");
	return {
		...actual,
		createConnection: jest.fn(() => mockConnection),
		TextDocuments: jest.fn(() => mockDocuments)
	};
});

// imported after Jest mock to load the (mocked) version of the vscode-languageserver/node.
import "../../src/server";

describe("Language Server", () => {
	it("registers onInitialize and returns correct capabilities", () => {
		const onInitCallback = mockOnInitialize.mock.calls[0][0];
		const capabilities = onInitCallback({
			workspaceFolders: [{ name: "my-folder", uri: "file:///my-folder" }]
		});

		expect(capabilities).toEqual({
			capabilities: {
				textDocumentSync: 2,
				workspace: {
					workspaceFolders: {
						supported: true
					}
				},
				codeActionProvider: {
					codeActionKinds: [CodeActionKind.QuickFix]
				}
			}
		});
	});

	it("registers onInitialized and sets up config and scan request", () => {
		const onInitCallback = mockOnInitialized.mock.calls[0][0];

		onInitCallback();

		expect(mockRegister).toHaveBeenCalledWith(DidChangeConfigurationNotification.type, undefined);
		expect(mockOnRequest).toHaveBeenCalledWith("scancode", expect.any(Function));
	});

	it("calls quickfix.applyQuickFix in onCodeAction", async () => {
		const fakeUri = "file:///sample.ts";
		const mockDoc = { uri: fakeUri } as TextDocument;

		mockDocuments.get.mockReturnValue(mockDoc);

		const diagnostics = [{ message: "sample error" }];
		const params = {
			textDocument: { uri: fakeUri },
			context: { diagnostics }
		};

		const quickfixMock = jest.fn().mockResolvedValue(["fix"]);
		(quickfix.applyQuickFix as jest.Mock).mockImplementation(quickfixMock);

		const callback = mockOnCodeAction.mock.calls[0][0];
		const result = await callback(params);

		expect(mockDocuments.get).toHaveBeenCalledWith(fakeUri);
		expect(quickfix.applyQuickFix).toHaveBeenCalledWith(mockDoc, params);
		expect(result).toEqual(["fix"]);
	});

	it("calls scanner.validateTextDocument on open and save", () => {
		const mockValidate = scanner.validateTextDocument as jest.Mock;

		expect(mockOnDidOpen).toHaveBeenCalledWith(expect.any(Function));
		expect(mockOnDidSave).toHaveBeenCalledWith(expect.any(Function));

		const fakeDoc = { uri: "file:///doc.ts" };
		const docEvent = { document: fakeDoc };

		const openCb = mockOnDidOpen.mock.calls[0][0];
		openCb(docEvent);
		expect(mockValidate).toHaveBeenCalledWith(fakeDoc, expect.anything(), mockConnection);

		const saveCb = mockOnDidSave.mock.calls[0][0];
		saveCb(docEvent);
		expect(mockValidate).toHaveBeenCalledWith(fakeDoc, expect.anything(), mockConnection);
	});
});
