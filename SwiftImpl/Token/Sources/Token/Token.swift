enum Token {
  case EOF
  case Illigal
  case Identifier(String)
  case Integer(Int)
  case Let
  case Return
  case Assign
  case Plus
  case Comma
  case Semicolon
  case LParen
  case RParen
  case LBrace
  case RBrace
  case Function
  case If
  case Else
  case True
  case False
  case Bang
  case Minus

  static func collection(_ tokens: [Token]) -> [Token] {
    return tokens
  }

}
