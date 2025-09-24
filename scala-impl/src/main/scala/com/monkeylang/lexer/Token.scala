package com.monkeylang.lexer

enum Token:
  case EOF
  case ILLEGAL
  case IDENT(value: String)
  case INT(value: Int)
  case ASSIGN
  case PLUS
  case COMMA
  case SEMICOLON
  case LPAREN
  case RPAREN
  case LBRACE
  case RBRACE
  case FUNCTION
  case LET
  case MINUS
  case BANG
  case ASTERISK
  case SLASH
  case LT
  case GT
  case IF
  case ELSE
  case RETURN
  case EQ
  case NOT_EQ
  case TRUE
  case FALSE

  override def toString(): String =
    this match
      case EOF           => "EOF"
      case ILLEGAL       => "ILLEGAL"
      case IDENT(value)  => s"IDENT($value)"
      case INT(value)    => s"INT($value)"
      case ASSIGN        => "="
      case PLUS          => "+"
      case COMMA         => ","
      case SEMICOLON     => ";"
      case LPAREN        => "("
      case RPAREN        => ")"
      case LBRACE        => "{"
      case RBRACE        => "}"
      case FUNCTION      => "FUNCTION"
      case LET           => "LET"
      case MINUS         => "-"
      case BANG          => "!"
      case ASTERISK      => "*"
      case SLASH         => "/"
      case LT            => "<"
      case GT            => ">"
      case IF            => "IF"
      case ELSE          => "ELSE"
      case RETURN        => "RETURN"
      case EQ            => "=="
      case NOT_EQ        => "!="
      case TRUE          => "TRUE"
      case FALSE         => "FALSE"

// A simple keyword lookup function that is outside of the Token enum
def lookup(ident: String): Token =
  ident match
    case s if s.isBlank() => Token.EOF
    case "fn"     => Token.FUNCTION
    case "let"    => Token.LET
    case "if"     => Token.IF
    case "else"   => Token.ELSE
    case "return" => Token.RETURN
    case "true"   => Token.TRUE
    case "false"  => Token.FALSE
    case s if s.forall(_.isLetter) => Token.IDENT(s)
    case i if i.forall(_.isDigit) => Token.INT(i.toInt)
    case _        => Token.ILLEGAL
  

