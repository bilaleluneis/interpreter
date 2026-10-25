use std::fmt::{Display, Formatter};

use crate::expression::Expression;

/// Statement nodes: constructs executed for effect.
pub enum Statement {
    /// `let <name> = <value>;`
    Let {
        /// The bound identifier.
        name: String,
        /// The initializer expression.
        value: Expression,
    },
    /// An expression used in statement position.
    Expression(Expression),
    /// `{ ... }`
    Block(Vec<Self>),
}

impl Display for Statement {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::Let { name, value } => write!(f, "let {name} = {value};"),
            Self::Expression(e) => write!(f, "{e}"),
            Self::Block(statements) if statements.is_empty() => write!(f, "{{}}"),
            Self::Block(statements) => {
                writeln!(f, "{{")?;
                statements.iter().try_for_each(|s| writeln!(f, "{s}"))?;
                write!(f, "}}")
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    fn let_five() -> Statement {
        Statement::Let {
            name: "five".to_string(),
            value: Expression::Integer(5),
        }
    }

    #[test]
    fn let_displays_name_and_value() {
        assert_eq!(let_five().to_string(), "let five = 5;");
    }

    #[test]
    fn expression_statement_displays_its_expression() {
        let s = Statement::Expression(Expression::Identifier("five".to_string()));
        assert_eq!(s.to_string(), "five");
    }

    #[test]
    fn block_displays_each_statement_on_its_own_line() {
        let block = Statement::Block(vec![
            let_five(),
            Statement::Expression(Expression::Boolean(true)),
        ]);
        assert_eq!(block.to_string(), "{\nlet five = 5;\ntrue\n}");
    }

    #[test]
    fn empty_block_displays_empty_braces() {
        assert_eq!(Statement::Block(vec![]).to_string(), "{}");
    }

    #[test]
    fn variants_are_distinguished_by_matching() {
        let statements = [
            let_five(),
            Statement::Expression(Expression::Integer(1)),
            Statement::Block(vec![]),
        ];

        let names: Vec<&str> = statements
            .iter()
            .map(|s| match s {
                Statement::Let { .. } => "let",
                Statement::Expression(_) => "expression",
                Statement::Block(_) => "block",
            })
            .collect();

        assert_eq!(names, ["let", "expression", "block"]);
    }
}
