import XCTest
import SwiftTreeSitter
import TreeSitterJsonnet

final class TreeSitterJsonnetTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_jsonnet())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Jsonnet grammar")
    }
}
