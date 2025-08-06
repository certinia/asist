import * as config from "../config";
import * as fs from "fs";
import { join } from "path";
import * as vscode from "vscode";

const mockedWorkspace = vscode.workspace as unknown as {
	getConfiguration: jest.Mock;
	workspaceFolders: any;
};

jest.mock("fs", () => ({
	existsSync: jest.fn()
}));

describe("Check Config functionality", () => {
	const mockExistsSync = fs.existsSync as jest.Mock;
	const mockedWorkspacePath = "/workspace";
	const mockedConfigYamlName = "custom.yaml";

	afterEach(() => {
		jest.clearAllMocks();
	});

	describe("getConfigFileOption", () => {
		it("returns path for custom file if it exists", () => {
			const customPath = join(mockedWorkspacePath, mockedConfigYamlName);
			mockExistsSync.mockReturnValueOnce(true);

			const result = config.getConfigFileOption(mockedWorkspacePath, mockedConfigYamlName);
			expect(result).toBe(`${customPath}`);
		});

		it("falls back to .asst.yaml if custom file does not exist", () => {
			mockExistsSync.mockReturnValueOnce(false).mockReturnValueOnce(true);

			const result = config.getConfigFileOption(mockedWorkspacePath, mockedConfigYamlName);
			expect(result).toBe(join(mockedWorkspacePath, ".asist.yaml"));
		});

		it("falls back to .asist.json if yaml is missing", () => {
			mockExistsSync
				.mockReturnValueOnce(false)
				.mockReturnValueOnce(false)
				.mockReturnValueOnce(true);

			const result = config.getConfigFileOption(mockedWorkspacePath, mockedConfigYamlName);
			expect(result).toBe(join(mockedWorkspacePath, ".asist.json"));
		});

		it("returns empty string if no file exists", () => {
			mockExistsSync.mockReturnValue(false);

			const result = config.getConfigFileOption(mockedWorkspacePath, "none.yaml");
			expect(result).toBe("");
		});
	});

	describe("getConfigOption", () => {
		it("returns formatted path when file exists", () => {
			const mockGet = jest.fn().mockReturnValue(mockedConfigYamlName);
			mockedWorkspace.workspaceFolders = [{ uri: { fsPath: mockedWorkspacePath } }];
			mockedWorkspace.getConfiguration.mockReturnValue({ get: mockGet });

			const fullPath = join(mockedWorkspacePath, mockedConfigYamlName);
			(fs.existsSync as jest.Mock).mockReturnValue(true);

			const result = config.getConfigOption();

			expect(result).toBe(fullPath);
		});

		it("returns empty string when no workspace folders", () => {
			mockedWorkspace.workspaceFolders = undefined;

			(vscode.workspace.getConfiguration as jest.Mock).mockReturnValue({
				get: jest.fn().mockReturnValue("custom.yaml")
			});

			const result = config.getConfigOption();
			expect(result).toBe("");
		});
	});

	describe("getConfigFilePath", () => {
		const mockedConfigYamlPath1 = join(mockedWorkspacePath, "custom.yaml");
		const mockedConfigYamlPath2 = join(mockedWorkspacePath, ".asist.yaml");
		const mockedConfigJsonPath = join(mockedWorkspacePath, ".asist.json");

		it("returns path for custom file if it exists", () => {
			const customPath = mockedConfigYamlPath1;
			mockExistsSync.mockReturnValueOnce(true);

			const result = config.getConfigFilePath(mockedWorkspacePath, mockedConfigYamlName);
			expect(result).toBe(`${customPath}`);
		});

		it("falls back to .asist.yaml if custom file does not exist", () => {
			mockExistsSync.mockReturnValueOnce(false).mockReturnValueOnce(true);

			const result = config.getConfigFilePath(mockedWorkspacePath, mockedConfigYamlName);
			expect(result).toBe(mockedConfigYamlPath2);
		});

		it("falls back to .asist.json if yaml and custom config file is missing", () => {
			mockExistsSync
				.mockReturnValueOnce(false)
				.mockReturnValueOnce(false)
				.mockReturnValueOnce(true);

			const result = config.getConfigFilePath(mockedWorkspacePath, mockedConfigYamlName);
			expect(result).toBe(mockedConfigJsonPath);
		});

		it("returns empty string if no file exists", () => {
			mockExistsSync.mockReturnValue(false);

			const result = config.getConfigFilePath(mockedWorkspacePath, "none.yaml");
			expect(result).toBe(null);
		});
	});
});
