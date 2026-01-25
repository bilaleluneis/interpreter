import Testing

@testable import Token

@Test func initTokens() async throws {
  let tokens = Token.collection([
    .Let,
    .Identifier("five"),
    .Assign,
    .Integer(5),
    .Semicolon,
    .EOF,
  ])
  #expect(tokens.count == 6)
  #expect(tokens.filter({ $0 == .EOF }).count == 1)
}
