use std::fmt::{Display, Formatter};

use crate::statement::Statement;

/// Root of the AST: an ordered list of top-level statements.
pub struct Program(Box<[Statement]>);

impl Program {
    /// Creates a program from its top-level statements.
    #[must_use]
    pub fn new(statements: Vec<Statement>) -> Self {
        Self(statements.into_boxed_slice())
    }

    /// Borrows the top-level statements in source order.
    #[must_use]
    pub fn statements(&self) -> &[Statement] {
        &self.0
    }
}

impl Display for Program {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        self.0
            .iter()
            .enumerate()
            .try_for_each(|(index, statement)| {
                if index > 0 {
                    writeln!(f)?;
                }
                write!(f, "{statement}")
            })
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::expression::Expression;

    /// Minimal phase-one constructs: `let`, expression statement, and block.
    fn program() -> Program {
        Program::new(vec![
            Statement::Let {
                name: "five".to_string(),
                value: Expression::Integer(5),
            },
            Statement::Let {
                name: "ten".to_string(),
                value: Expression::Integer(10),
            },
            Statement::Expression(Expression::Identifier("five".to_string())),
            Statement::Block(vec![Statement::Expression(Expression::Boolean(true))]),
            Statement::Block(vec![]),
        ])
    }

    #[test]
    fn empty_program_displays_nothing() {
        assert_eq!(Program::new(vec![]).to_string(), "");
    }

    #[test]
    fn program_displays_each_statement_on_its_own_line() {
        assert_eq!(
            program().to_string(),
            "let five = 5;\nlet ten = 10;\nfive\n{\ntrue\n}\n{}"
        );
    }

    #[test]
    fn program_statements_are_matched_by_variant() {
        let summary: Vec<String> = program()
            .statements()
            .iter()
            .map(|statement| match statement {
                Statement::Let { name, value } => format!("let {name} = {value}"),
                Statement::Expression(expression) => format!("expression {expression}"),
                Statement::Block(statements) => format!("block of {}", statements.len()),
            })
            .collect();

        assert_eq!(
            summary,
            [
                "let five = 5",
                "let ten = 10",
                "expression five",
                "block of 1",
                "block of 0",
            ]
        );
    }

    #[test]
    fn nested_block_contents_are_matched_too() {
        let program = program();

        let block_expressions: Vec<String> = program
            .statements()
            .iter()
            .filter_map(|statement| match statement {
                Statement::Block(statements) => Some(statements),
                Statement::Let { .. } | Statement::Expression(_) => None,
            })
            .flatten()
            .filter_map(|inner| match inner {
                Statement::Expression(expression) => Some(expression.to_string()),
                Statement::Let { .. } | Statement::Block(_) => None,
            })
            .collect();

        assert_eq!(block_expressions, ["true"]);
    }

    #[test]
    fn program_counts_by_variant() {
        let program = program();

        let lets = program
            .statements()
            .iter()
            .filter(|s| matches!(s, Statement::Let { .. }))
            .count();
        let blocks = program
            .statements()
            .iter()
            .filter(|s| matches!(s, Statement::Block(_)))
            .count();

        assert_eq!((lets, blocks), (2, 2));
    }
}
