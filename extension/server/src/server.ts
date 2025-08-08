import { TextDocument } from "vscode-languageserver-textdocument";
import {
	CodeAction,
	CodeActionKind,
	CodeActionParams,
	createConnection,
	DidChangeConfigurationNotification,
	InitializeParams,
	InitializeResult,
	ProposedFeatures,
	TextDocuments,
	TextDocumentSyncKind,
	WorkspaceFolder
} from "vscode-languageserver/node";
import { createDiagnostics, validateTextDocument } from "./scanner";
import { applyQuickFix } from "./quickfix";

//Workspace of the scanned project
let workspace: WorkspaceFolder[] | null | undefined = null;

// Create a connection for the server, using Node's IPC as a transport.
// Also include all preview / proposed LSP features.
const connection = createConnection(ProposedFeatures.all);

// Create a simple text document manager.
const documents: TextDocuments<TextDocument> = new TextDocuments(TextDocument);

connection.onInitialize((params: InitializeParams) => {
	workspace = params.workspaceFolders;
	const result: InitializeResult = {
		capabilities: {
			textDocumentSync: TextDocumentSyncKind.Incremental
		}
	};

	result.capabilities.workspace = {
		workspaceFolders: {
			supported: true
		}
	};

	result.capabilities.codeActionProvider = {
		codeActionKinds: [CodeActionKind.QuickFix]
	};

	return result;
});

connection.onInitialized(() => {
	// Register for all configuration changes.
	connection.client.register(DidChangeConfigurationNotification.type, undefined);

	connection.onRequest("scancode", param => createDiagnostics(param, connection));
});

connection.onCodeAction(provideCodeActions);

/*
 * provideCodeActions - Provides a list of code actions (quick fixes) for the specified document based on the provided diagnostics.
 */
async function provideCodeActions(parms: CodeActionParams): Promise<CodeAction[]> {
	if (!parms.context.diagnostics.length) {
		return [];
	}
	const document = documents.get(parms.textDocument.uri);
	if (document == null && document == undefined) {
		return [];
	}

	return applyQuickFix(document, parms);
}

documents.onDidOpen(open => {
	validateTextDocument(open.document, workspace, connection);
});

documents.onDidSave(save => {
	validateTextDocument(save.document, workspace, connection);
});

// Make the text document manager listen on the connection
// for open, change and close text document events
documents.listen(connection);

connection.listen();
