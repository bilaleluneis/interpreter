import Testing

@testable import Ast

@Test func example() async throws {
  let ast: Expression = .literal("42")
  #expect("\(ast)" == "literal(42)")
}
