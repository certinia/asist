import * as os from "os";
import * as path from "path";
import * as vscode from "vscode";
import { getScannerPath, getDefaultScannerPath } from "../utils";

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
});
