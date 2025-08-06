import { join } from "path";
import {
	LanguageClient,
	LanguageClientOptions,
	ServerOptions,
	TransportKind
} from "vscode-languageclient/node";
import {
	createConfigCommand,
	editConfigCommand,
	listRulesCommand,
	preferencesCommand,
	scanFileCommand,
	scanWorkspaceCommand
} from "./commands";
import { ExtensionContext, workspace } from "vscode";

let client: LanguageClient;

/*
 * activate - method triggered when the VScode extension gets activated
 */
export function activate(context: ExtensionContext) {
	// The server is implemented in node
	const serverModule = context.asAbsolutePath(join("server", "out", "server.js"));
	// The debug options for the server
	// --inspect=6009: runs the server in Node's Inspector mode so VS Code can attach to the server for debugging
	const debugOptions = { execArgv: ["--nolazy", "--inspect=6009"] };

	// If the extension is launched in debug mode then the debug server options are used
	// Otherwise the run options are used
	const serverOptions: ServerOptions = {
		run: { module: serverModule, transport: TransportKind.ipc },
		debug: {
			module: serverModule,
			transport: TransportKind.ipc,
			options: debugOptions
		}
	};

	// Options to control the language client
	const clientOptions: LanguageClientOptions = {
		// Register the server for plain text documents
		documentSelector: [{ scheme: "file" }],
		synchronize: {
			// Notify the server about file changes to '.clientrc files contained in the workspace
			fileEvents: workspace.createFileSystemWatcher("**/.clientrc")
		}
	};

	// Create the language client and start the client. This will also launch the server
	client = new LanguageClient("languageServer", "Language Server", serverOptions, clientOptions);
	client.start();

	registerCommands(context);
}

/*
 * deactivate - method used to clean up resources when extension is deactivated
 */
export function deactivate(): Thenable<void> | undefined {
	if (!client) {
		return undefined;
	}
	return client.stop();
}

/*
 * registerCommands - method used to register all the ASIST extension commands
 */
function registerCommands(context: ExtensionContext) {
	context.subscriptions.push(listRulesCommand(client));
	context.subscriptions.push(scanWorkspaceCommand(client));
	context.subscriptions.push(scanFileCommand(client));
	context.subscriptions.push(editConfigCommand());
	context.subscriptions.push(createConfigCommand());
	context.subscriptions.push(preferencesCommand());
}
