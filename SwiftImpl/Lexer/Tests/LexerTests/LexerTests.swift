import Testing

@testable import Lexer
@testable import Token

// Minimal language lexer support expected in this phase:
// - `let` bindings, identifiers, integer literals
// - function literal token `fn`
// - operators: + - * /, < >, ==, !=
// - punctuation: = , ; ( ) { }
// - keywords: if, else, true, false, return

@Test func lexerMinimalExample() async throws {
  let input = """
    let five = 5;
    let ten = 10;
    let mul = fn(x, y) { x * y };
    let add = fn(x, y) { x + y };
    let result = if(five < ten) { mul(five, ten) } else { add(five, ten) };
    """

  var l = Lexer(input)

  let expected: [Token] = [
    .Let,
    .Identifier("five"),
    .Assign,
    .Integer(5),
    .Semicolon,
    .Let,
    .Identifier("ten"),
    .Assign,
    .Integer(10),
    .Semicolon,
    .Let,
    .Identifier("mul"),
    .Assign,
    .Function,
    .LParen,
    .Identifier("x"),
    .Comma,
    .Identifier("y"),
    .RParen,
    .LBrace,
    .Identifier("x"),
    .Asterisk,
    .Identifier("y"),
    .RBrace,
    .Semicolon,
    .Let,
    .Identifier("add"),
    .Assign,
    .Function,
    .LParen,
    .Identifier("x"),
    .Comma,
    .Identifier("y"),
    .RParen,
    .LBrace,
    .Identifier("x"),
    .Plus,
    .Identifier("y"),
    .RBrace,
    .Semicolon,
    .Let,
    .Identifier("result"),
    .Assign,
    .If,
    .LParen,
    .Identifier("five"),
    .LT,
    .Identifier("ten"),
    .RParen,
    .LBrace,
    .Identifier("mul"),
    .LParen,
    .Identifier("five"),
    .Comma,
    .Identifier("ten"),
    .RParen,
    .RBrace,
    .Else,
    .LBrace,
    .Identifier("add"),
    .LParen,
    .Identifier("five"),
    .Comma,
    .Identifier("ten"),
    .RParen,
    .RBrace,
    .Semicolon,
    .EOF,
  ]

  var out: [Token] = []
  while true {
    let tok = l.nextToken()
    out.append(tok)
    if case .EOF = tok { break }
  }

  assert(out == expected)
}

@Test func lexerIsValueType() async throws {
  let input = "let a = 1; let b = 2;"
  var l1 = Lexer(input)
  var l2 = l1 // copy

  // advance l1 twice
  let firstFromL1 = l1.nextToken()
  let secondFromL1 = l1.nextToken()

  // l2 should still produce the original first token when advanced
  let firstFromL2 = l2.nextToken()

  assert(firstFromL1 == firstFromL2)
  assert(secondFromL1 != firstFromL2)
}
