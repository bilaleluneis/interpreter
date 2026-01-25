extension Token: CustomDebugStringConvertible {
  var debugDescription: String {
    switch self {
    case .EOF:
      return "EOF"
    case .Illigal:
      return "Illigal"
    case .Identifier(let name):
      return "Identifier(\(name))"
    case .Integer(let value):
      return "Integer(\(value))"
    case .Let:
      return "Let"
    case .Return:
      return "Return"
    case .Assign:
      return "Assign"
    case .Plus:
      return "Plus"
    case .Comma:
      return "Comma"
    case .Semicolon:
      return "Semicolon"
    case .LParen:
      return "LParen"
    case .RParen:
      return "RParen"
    case .LBrace:
      return "LBrace"
    case .RBrace:
      return "RBrace"
    case .Function:
      return "Function"
    case .If:
      return "If"
    case .Else:
      return "Else"
    case .True:
      return "True"
    case .False:
      return "False"
    case .Bang:
      return "Bang"
    case .Minus:
      return "Minus"
    }
  }
}
