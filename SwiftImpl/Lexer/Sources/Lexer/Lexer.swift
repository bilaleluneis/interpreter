import Token

public struct Lexer {
  let input: String
  var position: String.Index
  var readPosition: String.Index
  var ch: Character?

  public init(_ input: String) {
    self.input = input
    self.position = input.startIndex
    self.readPosition = input.startIndex
    self.ch = nil
    self.readChar()
  }

  mutating func readChar() {
    if readPosition >= input.endIndex {
      ch = nil
    } else {
      ch = input[readPosition]
    }
    position = readPosition
    if readPosition < input.endIndex {
      readPosition = input.index(after: readPosition)
    } else {
      readPosition = input.endIndex
    }
  }

  mutating func peekChar() -> Character? {
    if readPosition >= input.endIndex { return nil }
    return input[readPosition]
  }

  mutating public func nextToken() -> Token {
    skipWhitespace()

    guard let ch = ch else { return .EOF }

    // two-char tokens: ==, !=
    if ch == "=" && peekChar() == "=" {
      readChar()
      readChar()
      return .Equal
    }
    if ch == "!" && peekChar() == "=" {
      readChar()
      readChar()
      return .NotEqual
    }

    let tok: Token = {
      switch ch {
      case "=": return .Assign
      case ";": return .Semicolon
      case ",": return .Comma
      case "+": return .Plus
      case "-": return .Minus
      case "!": return .Bang
      case "/": return .Slash
      case "*": return .Asterisk
      case "(": return .LParen
      case ")": return .RParen
      case "{": return .LBrace
      case "}": return .RBrace
      case "<": return .LT
      case ">": return .GT
      default:
        if isLetter(ch) {
          return lookupIdent(readIdentifier())
        } else if isDigit(ch) {
          return .Integer(Int(readNumber()) ?? 0)
        } else {
          return .Illigal
        }
      }
    }()

    readChar()
    return tok
  }

  mutating func skipWhitespace() {
    while let c = ch, c.isWhitespace { readChar() }
  }

  mutating func readIdentifier() -> String {
    let start = position
    while let c = ch, isLetter(c) {
      readChar()
    }
    return String(input[start..<position])
  }

  mutating func readNumber() -> String {
    let start = position
    while let c = ch, isDigit(c) {
      readChar()
    }
    return String(input[start..<position])
  }

  func isLetter(_ ch: Character) -> Bool {
    return ch.isLetter || ch == "_"
  }

  func isDigit(_ ch: Character) -> Bool {
    return ch.isNumber
  }

  func lookupIdent(_ ident: String) -> Token {
    switch ident {
    case "let": return .Let
    case "fn": return .Function
    case "if": return .If
    case "else": return .Else
    case "return": return .Return
    case "true": return .True
    case "false": return .False
    default: return .Identifier(ident)
    }
  }

}
