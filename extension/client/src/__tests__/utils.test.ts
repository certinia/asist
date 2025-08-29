import * as os from "os";
import * as path from "path";
import * as vscode from "vscode";
import { getScannerPath, getDefaultScannerPath, removeAnsiEscapeCodes } from "../utils";

jest.mock("os");

const mockedOs = os as jest.Mocked<typeof os>;
describe("client utils", () => {
	afterEach(jest.clearAllMocks);

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
			mockedOs.platform.mockReturnValue("unknown" as unknown as ReturnType<typeof os.platform>);
			mockedOs.arch.mockReturnValue("unknown");

			const expectedPath = path.join(basePath, "./asist_linux_amd64");
			expect(getDefaultScannerPath()).toBe(expectedPath);
		});
	});

	describe("getScannerPath", () => {
		const basePath = path.join(__dirname, "../../..");
		const mockedScannerPath = "/custom/path/to/scanner";
		afterEach(() => {
			jest.clearAllMocks();
		});

		it("returns custom binary path when enabled", () => {
			(vscode.workspace.getConfiguration as jest.Mock).mockReturnValue({
				enabled: true,
				path: mockedScannerPath
			});

			const result = getScannerPath();
			expect(result).toBe(mockedScannerPath);
		});

		it("returns default path when custom binary is disabled", () => {
			(vscode.workspace.getConfiguration as jest.Mock).mockReturnValue({
				enabled: false,
				path: "/custombinarypath"
			});

			const result = getScannerPath();
			const expectedPath = path.join(basePath, `/asist_linux_amd64`);
			expect(result).toBe(expectedPath);
		});
	});

	describe("removeAnsiEscapeCodes", () => {
		it("should return an empty string when input is empty", () => {
			expect(removeAnsiEscapeCodes("")).toBe("");
		});

		it("should return the same string if there are no ANSI escape codes", () => {
			const input = "Hello, world!";
			expect(removeAnsiEscapeCodes(input)).toBe(input);
		});

		it("should remove simple ANSI codes", () => {
			const input = "\u001b[31mHello\u001b[0m";
			const expected = "Hello";
			expect(removeAnsiEscapeCodes(input)).toBe(expected);
		});

		it("should handle multiple ANSI codes in the same string", () => {
			const input = "\x1b[31mRed\x1b[0m and \x1b[32mGreen\x1b[0m";
			const expected = "Red and Green";
			expect(removeAnsiEscapeCodes(input)).toBe(expected);
		});

		it("should not throw on undefined or null", () => {
			expect(removeAnsiEscapeCodes(undefined)).toBe("");
			expect(removeAnsiEscapeCodes(null)).toBe("");
		});
	});
});
