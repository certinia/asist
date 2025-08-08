import * as scanner from "../../src/scanner";
import { _Connection, Diagnostic } from "vscode-languageserver";
import { TextDocument } from "vscode-languageserver-textdocument";
import * as utils from "../../src/utils";
import * as url from "url";
import { execFile } from "child_process";

jest.mock("vscode-languageserver", () => ({
	_Connection: jest.fn(),
	DiagnosticSeverity: {
		Error: 1,
		Warning: 2,
		Information: 3,
		Hint: 4
	},
	Diagnostic: jest.fn()
}));

jest.mock("child_process", () => ({
	execFile: jest.fn()
}));

jest.mock("../../src/utils", () => ({
	getDefaultScannerPath: jest.fn(),
	checkAndGetConfigFileOption: jest.fn()
}));

const testFilePath = "/test/file1.ts";

describe("Test functionality for createDiagnostic and createDiagnosticFromJSONObject", () => {
	const mockConnection: _Connection = {
		sendDiagnostics: jest.fn(),
		workspace: {} as any
	};

	const mockUri = "file:///test/file1.ts";

	beforeEach(() => {
		jest.spyOn(url, "fileURLToPath").mockReturnValue(testFilePath);
	});

	afterEach(() => {
		jest.restoreAllMocks();
	});

	it("should correctly create diagnostics and group them by URI", () => {
		const jsonInput = JSON.stringify({
			Count: 2,
			Result: [
				{
					ID: "1",
					Name: "Diagnostic 1",
					Severity: "Critical",
					Occurrence: {
						File: testFilePath,
						LineNumber: 1,
						ColumnRange: [0, 10]
					},
					Description: "First diagnostic"
				},
				{
					ID: "2",
					Name: "Diagnostic 2",
					Severity: "Medium",
					Occurrence: {
						File: testFilePath,
						LineNumber: 2,
						ColumnRange: [0, 5]
					},
					Description: "Second diagnostic"
				}
			]
		});

		jest.spyOn(url, "pathToFileURL").mockReturnValue(new URL(mockUri));

		scanner.createDiagnostics(jsonInput, mockConnection);
		expect(mockConnection.sendDiagnostics).toHaveBeenCalledWith({
			uri: mockUri,
			diagnostics: [
				expect.objectContaining({ message: "[1] Diagnostic 1", severity: 1 }),
				expect.objectContaining({ message: "[2] Diagnostic 2", severity: 2 })
			]
		});
	});

	it("should correctly create a Diagnostic from a JSON object", () => {
		const jsonObject = {
			ID: "1",
			Name: "Test Diagnostic",
			Severity: "High",
			Occurrence: {
				File: testFilePath,
				LineNumber: 1,
				ColumnRange: [5, 15]
			},
			Description: "Test description"
		};
		jest.spyOn(url, "pathToFileURL").mockReturnValue(new URL(mockUri));

		const diagnostic: Diagnostic = scanner.createDiagnosticFromJSONObject(jsonObject);

		expect(diagnostic).toEqual(
			expect.objectContaining({
				severity: 1,
				message: "[1] Test Diagnostic",
				range: {
					start: { line: 0, character: 5 },
					end: { line: 0, character: 15 }
				}
			})
		);

		expect(diagnostic.relatedInformation).toHaveLength(1);
		expect(diagnostic.relatedInformation[0].location.uri).toBe(mockUri);
	});

	it("should handle empty Result array gracefully", () => {
		const emptyJsonInput = JSON.stringify({
			Count: 0,
			Result: []
		});
		scanner.createDiagnostics(emptyJsonInput, mockConnection);

		expect(mockConnection.sendDiagnostics).toHaveBeenCalled();
	});
});

describe("Test functionality for validateTextDocument", () => {
	const mockUri = "file://mockFile.ts";

	const mockTextDocument = {
		uri: mockUri
	} as TextDocument;

	const mockWorkspaceFolders = [{ uri: "file://mockWorkspace", name: "mockWorkspace" }];

	const mockSendDiagnostics = jest.fn();

	const mockConnection = {
		workspace: {
			getConfiguration: jest.fn()
		},
		sendDiagnostics: mockSendDiagnostics
	};

	const mockStdout = {
		toString: () =>
			JSON.stringify([
				{
					message: "Simulated diagnostic",
					range: { start: { line: 0, character: 0 }, end: { line: 0, character: 5 } },
					severity: 1
				}
			])
	};

	let createDiagnosticsSpy: jest.SpyInstance;

	beforeEach(() => {
		jest.clearAllMocks();

		createDiagnosticsSpy = jest
			.spyOn(scanner as any, "createDiagnostics")
			.mockImplementation(() => {});
	});

	it("should call execFile and handle diagnostics correctly", async () => {
		(mockConnection.workspace.getConfiguration as jest.Mock)
			.mockResolvedValueOnce({ enabled: true, path: "/bin/tool" })
			.mockResolvedValueOnce({ FilePath: "some/file/path" });

		jest.spyOn(url, "fileURLToPath").mockReturnValue(testFilePath);

		(utils.checkAndGetConfigFileOption as jest.Mock).mockResolvedValueOnce(" ");

		(execFile as unknown as jest.Mock).mockImplementation((_cmd, _args, cb) => {
			cb(null, mockStdout, "");
		});

		await scanner.validateTextDocument(
			mockTextDocument,
			mockWorkspaceFolders,
			mockConnection as any
		);

		expect(execFile).toHaveBeenCalledWith("/bin/tool", [testFilePath], expect.any(Function));
		expect(mockConnection.sendDiagnostics).toHaveBeenCalledWith({
			uri: mockUri,
			diagnostics: []
		});
	});
});
