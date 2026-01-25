indirect enum Expression: CustomDebugStringConvertible {
  case literal(String)
  case variable(String)
  case binaryOperation(left: Expression, operator: String, right: Expression)
  case functionCall(name: String, arguments: [Expression])

  var debugDescription: String {
    switch self {
    case .literal(let value):
      return "literal(\(value))"
    case .variable(let name):
      return "variable(\"\(name)\")"
    case .binaryOperation(let left, let op, let right):
      let left = left.debugDescription
      let right = right.debugDescription
      return "binaryOperation(left: \(left), operator: \"\(op)\", right: \(right))"
    case .functionCall(let name, let arguments):
      let args = arguments.map { $0.debugDescription }.joined(separator: ", ")
      return "functionCall(name: \"\(name)\", arguments: [\(args)])"
    }
  }
}
