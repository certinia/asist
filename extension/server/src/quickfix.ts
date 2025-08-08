import { TextDocument } from "vscode-languageserver-textdocument";
import {
	CodeAction,
	CodeActionKind,
	CodeActionParams,
	Diagnostic,
	Range
} from "vscode-languageserver/node";
import { getFileExtension } from "./utils";
import { ASIST, MARK_FALSE_POSITIVE } from "./constants";

const HTML_COMMENTS = ["page", "html", "htm", "component", "xml"];
const BLOCK_COMMENT = ["css"];

/*
 * quickfix - method to provide quickfix feature
 */
export function applyQuickFix(textDocument: TextDocument, parms: CodeActionParams): CodeAction[] {
	const diagnostics = parms.context.diagnostics;
	const codeActions: CodeAction[] = [];

	if (diagnostics == null || diagnostics == undefined || diagnostics.length === 0) {
		return [];
	}

	diagnostics.forEach(diag => {
		if (diag.source == ASIST) {
			codeActions.push(createFalsePositiveQuickFix(diag, textDocument));
			return;
		}
	});

	return codeActions;
}

/*
 * createFalsePositiveQuickFix - method to mark occurrence as false positive using quickfix
 */
function createFalsePositiveQuickFix(diag: Diagnostic, textDocument: TextDocument) {
	const ruleID = diag.message.split("[")[1].split("]")[0];
	const IGNORE_BEGIN_COMMENT = "asist-ignore-begin";
	const IGNORE_END_COMMENT = "asist-ignore-end";
	const SAMPLE_DESCRIPTION = "TODO Write comment here";
	let previousLineRange = null;
	let modificationRange = null;
	let textChanged = "";

	if (diag.range.start.line > 0) {
		previousLineRange = Range.create(diag.range.start.line - 1, 0, diag.range.start.line, 0);
	}
	if (
		previousLineRange != null &&
		textDocument.getText(previousLineRange).includes(IGNORE_BEGIN_COMMENT)
	) {
		//A false positive exists already
		modificationRange = previousLineRange;
		const commentText = textDocument.getText(previousLineRange).split("]");
		const preBracketText = commentText.shift();
		const postBracketText = commentText.join("]");
		textChanged = `${preBracketText},${ruleID}]${postBracketText}`;
	} else {
		//Create new false positives comment
		modificationRange = Range.create(diag.range.start.line, 0, diag.range.start.line + 1, 0);

		const leadingSpaceCount = textDocument.getText(modificationRange).search(/\S|$/);

		const leadingSpace = textDocument
			.getText(modificationRange)
			.substring(0, leadingSpaceCount - 1);

		const currentCode = textDocument.getText(modificationRange);

		const fileExtension = getFileExtension(textDocument.uri);
		let openComment,
			closeComment = "";

		if (HTML_COMMENTS.includes(fileExtension)) {
			openComment = "<!-- ";
			closeComment = "-->\n";
		} else if (BLOCK_COMMENT.includes(fileExtension)) {
			openComment = "/*";
			closeComment = "*/\n";
		} else {
			openComment = "//";
			closeComment = "\n";
		}

		textChanged = `${leadingSpace}${openComment}${IGNORE_BEGIN_COMMENT}:[${ruleID}] ${SAMPLE_DESCRIPTION} ${closeComment}${currentCode}${leadingSpace}${openComment}${IGNORE_END_COMMENT}${closeComment}`;
	}

	return {
		title: MARK_FALSE_POSITIVE,
		kind: CodeActionKind.QuickFix,
		diagnostics: [diag],
		edit: {
			changes: {
				[textDocument.uri]: [
					{
						range: modificationRange,
						newText: textChanged
					}
				]
			}
		}
	};
}
