import { execFile } from "child_process";
import { TextDocument } from "vscode-languageserver-textdocument";
import {
	_Connection,
	Diagnostic,
	DiagnosticSeverity,
	WorkspaceFolder
} from "vscode-languageserver/node";
import { checkAndGetConfigFileOption, getDefaultScannerPath } from "./utils";
import { ASIST, JSON_OBJECT_RESULT_KEYS, OPTIONS } from "./constants";
import { fileURLToPath, pathToFileURL } from "url";

//Severity mapping from Go binary to Diagnostics feature
export const SEVERITY_DICTIONARY = new Map<string, DiagnosticSeverity>([
	["Critical", DiagnosticSeverity.Error],
	["High", DiagnosticSeverity.Error],
	["Medium", DiagnosticSeverity.Warning],
	["Low", DiagnosticSeverity.Information]
]);
/*
 * validateTextDocument - method used to scan the code
 */
export async function validateTextDocument(
	textDocument: TextDocument,
	workspace: WorkspaceFolder[] | null | undefined,
	connection: _Connection
): Promise<void> {
	const customBinarySettings = await connection.workspace.getConfiguration({
		section: "ASIST.customBinary"
	});

	//get the path to the Go binary, if there is a custom path we use it, otherwise we use the default binary
	let goScannerPath = "";
	if (customBinarySettings.enabled) {
		goScannerPath = customBinarySettings.path;
	} else {
		goScannerPath = getDefaultScannerPath();
	}

	//Check if there is a config file
	const configOption = await getConfigFileOption(workspace, connection);

	const fileToScan = fileURLToPath(textDocument.uri);
	const cmdOptions =
		configOption == null || configOption.trim() == ""
			? [fileToScan]
			: [OPTIONS.CONFIG, configOption, fileToScan];
	//execute the go scan
	execFile(goScannerPath, cmdOptions, function callback(error, stdout, stderr) {
		if (error == null) {
			// Clear out old diagnostics
			connection.sendDiagnostics({ uri: textDocument.uri, diagnostics: [] });
			//Create the diagnostics with the results
			createDiagnostics(stdout.toString(), connection);
		} else if (stderr != null) {
			connection.sendDiagnostics({ uri: textDocument.uri, diagnostics: [] });
		}
	});
}

/*
 * createDiagnostics - method used to create the diagnostics objects from the Go result and send them to the client
 */
export function createDiagnostics(json: string, connection: _Connection) {
	const jsonObject = JSON.parse(json);
	const count = jsonObject[JSON_OBJECT_RESULT_KEYS.COUNT];
	const mapDiagnosticsByUri = new Map<string, Diagnostic[]>();
	for (let i = 0; i < count; i++) {
		const diagnostic: Diagnostic = createDiagnosticFromJSONObject(
			jsonObject[JSON_OBJECT_RESULT_KEYS.RESULT][i]
		);

		const uri = diagnostic!.relatedInformation!.at(0)!.location.uri;

		if (mapDiagnosticsByUri.has(uri)) {
			mapDiagnosticsByUri.get(uri)!.push(diagnostic);
		} else {
			const diagnosticsArray: Diagnostic[] = [];
			diagnosticsArray.push(diagnostic);
			mapDiagnosticsByUri.set(uri, diagnosticsArray);
		}
	}

	mapDiagnosticsByUri.forEach(function (diagnostics, key, map) {
		connection.sendDiagnostics({ uri: key, diagnostics });
	});
}

/*
 * createDiagnosticFromJSONObject - method used to create a Diagnostic object from a JSON object
 */
export function createDiagnosticFromJSONObject(jsonObject: any): Diagnostic {
	const diagnostic: Diagnostic = {
		severity: SEVERITY_DICTIONARY.get(jsonObject[JSON_OBJECT_RESULT_KEYS.SEVERITY]),
		range: {
			start: {
				character:
					jsonObject[JSON_OBJECT_RESULT_KEYS.OCCURRENCE][JSON_OBJECT_RESULT_KEYS.COLUMN_RANGE][0],
				line:
					jsonObject[JSON_OBJECT_RESULT_KEYS.OCCURRENCE][JSON_OBJECT_RESULT_KEYS.LINE_NUMBER] - 1
			},
			end: {
				character:
					jsonObject[JSON_OBJECT_RESULT_KEYS.OCCURRENCE][JSON_OBJECT_RESULT_KEYS.COLUMN_RANGE][1],
				line:
					jsonObject[JSON_OBJECT_RESULT_KEYS.OCCURRENCE][JSON_OBJECT_RESULT_KEYS.LINE_NUMBER] - 1
			}
		},
		message: `[${jsonObject[JSON_OBJECT_RESULT_KEYS.ID]}] ${jsonObject[JSON_OBJECT_RESULT_KEYS.NAME]}`,
		source: ASIST
	};

	diagnostic.relatedInformation = [
		{
			location: {
				uri: pathToFileURL(
					jsonObject[JSON_OBJECT_RESULT_KEYS.OCCURRENCE][JSON_OBJECT_RESULT_KEYS.FILE]
				).href,
				range: Object.assign({}, diagnostic.range)
			},
			message: jsonObject[JSON_OBJECT_RESULT_KEYS.DESCRIPTION]
		}
	];

	return diagnostic;
}

/*
 * getConfigFileOption - method used to check if the configuration file exists and returns the path depending on the OS
 */
async function getConfigFileOption(
	workspace: WorkspaceFolder[] | null | undefined,
	connection: _Connection
): Promise<string> {
	if (!workspace) {
		return "";
	}

	const configFilePath: string = await connection.workspace
		.getConfiguration({
			section: ASIST
		})
		.then(config => config.configFilePath);

	return checkAndGetConfigFileOption(workspace[0].uri, configFilePath);
}
