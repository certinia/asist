jest.mock("../utils", () => ({
	getFileExtension: jest.fn()
}));

import { applyQuickFix } from "../quickfix";
import * as utils from "../utils";
import { Range, Diagnostic } from "vscode-languageserver-types";

const mockTextDoc = {
	uri: "file:///mock/file.html",
	getText: jest.fn()
};
describe("applyQuickFix", () => {
	beforeEach(() => {
		jest.clearAllMocks();
	});

	it("returns empty array when no diagnostics", () => {
		const result = applyQuickFix(mockTextDoc as any, { context: { diagnostics: [] } } as any);
		expect(result).toEqual([]);
	});

	it("generates comment for ASIST diagnostic with html extension", () => {
		const mockDiagnostic: Diagnostic = {
			message: "Issue [123]",
			range: Range.create(3, 0, 3, 10),
			source: "ASIST"
		};

		(utils.getFileExtension as jest.Mock).mockReturnValue("html");

		mockTextDoc.getText = jest.fn().mockImplementation(() => "    const a = 1;\n");

		const result = applyQuickFix(
			mockTextDoc as any,
			{
				context: { diagnostics: [mockDiagnostic] }
			} as any
		);

		expect(result).toHaveLength(1);
		expect(result[0].edit!.changes![mockTextDoc.uri][0].newText).toContain(
			"<!-- asist-ignore-begin:[123]"
		);
	});

	it('wraps existing comment when "asist-ignore-begin" exists in text above', () => {
		const mockDiagnostic: Diagnostic = {
			message: "Error [456]",
			range: Range.create(2, 0, 2, 10),
			source: "ASIST"
		};

		const codeParams = {
			context: { diagnostics: [mockDiagnostic] }
		};

		(utils.getFileExtension as jest.Mock).mockReturnValue("html");

		mockTextDoc.getText
			.mockReturnValueOnce("<!-- asist-ignore-begin:[001] here -->")
			.mockReturnValueOnce("other text");

		const result = applyQuickFix(mockTextDoc as any, codeParams as any);

		expect(result[0].edit!.changes![mockTextDoc.uri][0].newText).toContain("other text,456]");
	});
});
