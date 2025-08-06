import * as fs from "fs";
import * as os from "os";
import * as path from "path";
import { pathToFileURL } from "url";
import { getFileExtension, checkAndGetConfigFileOption, getDefaultScannerPath } from "../utils";

jest.mock("os");

const mockedOs = os as jest.Mocked<typeof os>;
describe("server utils", () => {
	let existsSyncSpy: jest.SpyInstance;
	let osTypeSpy: jest.SpyInstance;

	beforeEach(() => {
		existsSyncSpy = jest.spyOn(fs, "existsSync");
		osTypeSpy = jest.spyOn(os, "type");

		osTypeSpy.mockReturnValue("Darwin");
	});

	afterEach(jest.clearAllMocks);

	describe("checkAndGetConfigFileOption", () => {
		const testPath = path.resolve("TestFolder");
		const workspaceUri = pathToFileURL(testPath).href;

		it("should return config file path if custom config file paths exists", () => {
			// Given file exists
			existsSyncSpy.mockReturnValue(true);

			// When
			const result = checkAndGetConfigFileOption(workspaceUri, ".asist.yaml");

			// Then
			expect(result).toBe(path.resolve("TestFolder", ".asist.yaml"));
			expect(existsSyncSpy).toHaveBeenCalled();
		});

		it("should return option with yaml config file if yaml config exists and custom config not exists", () => {
			// Given
			existsSyncSpy.mockReturnValueOnce(true);

			// When we call with the incorrect file path
			const result = checkAndGetConfigFileOption(workspaceUri, "");

			// Then we should go to the next valid file
			expect(result).toBe(path.resolve("TestFolder", ".asist.yaml"));
			expect(existsSyncSpy).toHaveBeenCalled();
		});

		it("should return option with json file if json config file exist and custom config not exist", () => {
			// Given the defined config file does not exist
			existsSyncSpy.mockReturnValueOnce(false);
			existsSyncSpy.mockReturnValueOnce(true);

			// When we call with the incorrect file path
			const result = checkAndGetConfigFileOption(workspaceUri, "");

			// Then we should go to the next valid file
			expect(result).toBe(path.resolve("TestFolder", ".asist.json"));
			expect(existsSyncSpy).toHaveBeenCalled();
		});

		it("should return empty string when no config file exist", () => {
			// Given the defined config file does not exist
			existsSyncSpy.mockReturnValueOnce(false);
			// But the workspace/.asist.yaml does
			existsSyncSpy.mockReturnValueOnce(false);

			// When we call with the incorrect file path
			const result = checkAndGetConfigFileOption(workspaceUri, "");

			// Then we should go to the next valid file
			expect(result).toBe(" ");
			expect(existsSyncSpy).toHaveBeenCalled();
		});
	});

	describe("getDefaultScannerPath", () => {
		const basePath = path.join(__dirname, "../../..");

		test.each([
			["win32", "x64", "asist.exe"],
			["win32", "arm64", "asist.exe"],
			["darwin", "x64", "asist_darwin_amd64"],
			["darwin", "arm64", "asist_darwin_arm64"],
			["linux", "arm64", "asist_linux_arm64"],
			["linux", "x64", "asist_linux_amd64"]
		] as const)("returns correct path for %s/%s", (platform, arch, expectedBinary) => {
			mockedOs.platform.mockReturnValue(platform);
			mockedOs.arch.mockReturnValue(arch);

			const expectedPath = path.join(basePath, `./${expectedBinary}`); // nosemgrep: path-join-resolve-traversal
			const actualPath = getDefaultScannerPath();

			expect(actualPath).toBe(expectedPath);
		});

		test("When os.platform() and os.arch() is unknown then return default", () => {
			mockedOs.platform.mockReturnValue(
				"unknown_platform" as unknown as ReturnType<typeof os.platform>
			);
			mockedOs.arch.mockReturnValue("unknown_arch");

			const expectedPath = path.join(basePath, "./asist_linux_amd64");
			expect(getDefaultScannerPath()).toBe(expectedPath);
		});
	});
});

describe("getFileExtension", () => {
	it('should return "js" for JavaScript file', () => {
		expect(getFileExtension("file.js")).toBe("js");
	});

	it('should return "ts" for TypeScript file', () => {
		expect(getFileExtension("example/path/file.ts")).toBe("ts");
	});

	it('should return "json" for JSON file', () => {
		expect(getFileExtension("/some/dir/data.json")).toBe("json");
	});
});
