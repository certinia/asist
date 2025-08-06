import { arch, platform, type } from "os";
import { join } from "path";
import { ASIST_BINARY_OS_TYPE, ASIST_CONFIGURATION_OPTIONS, OS_TYPE } from "./constants";
import { workspace } from "vscode";

/*
 * getDefaultScannerPath - method returns the default path of go binary depending on the OS
 */
export function getDefaultScannerPath(): string {
	const platarch = `${platform()}/${arch()}`;
	const EXTENSION_ROOT_PATH = "../..";

	switch (platarch) {
		case OS_TYPE.WINDOWS:
			return join(__dirname, EXTENSION_ROOT_PATH, ASIST_BINARY_OS_TYPE.WINDOWS);
		case OS_TYPE.WINDOWS_ARM64:
			return join(__dirname, EXTENSION_ROOT_PATH, ASIST_BINARY_OS_TYPE.WINDOWS);
		case OS_TYPE.DARWIN_AMD64:
			return join(__dirname, EXTENSION_ROOT_PATH, ASIST_BINARY_OS_TYPE.DARWIN_AMD64);
		case OS_TYPE.DARWIN_ARM64:
			return join(__dirname, EXTENSION_ROOT_PATH, ASIST_BINARY_OS_TYPE.DARWIN_ARM64);
		case OS_TYPE.LINUX_ARM64:
			return join(__dirname, EXTENSION_ROOT_PATH, ASIST_BINARY_OS_TYPE.LINUX_ARM64);
		default: // "linux/amd64" or other OS
			return join(__dirname, EXTENSION_ROOT_PATH, ASIST_BINARY_OS_TYPE.LINUX_AMD64);
	}
}

/*
 * getScannerPath - method used to get the go binary path for scan
 */
export function getScannerPath(): string {
	if (workspace.getConfiguration(ASIST_CONFIGURATION_OPTIONS.CUSTOM_BINARY).enabled) {
		return workspace.getConfiguration(ASIST_CONFIGURATION_OPTIONS.CUSTOM_BINARY).path;
	} else {
		return getDefaultScannerPath();
	}
}
