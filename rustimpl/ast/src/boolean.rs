use token::token::Token;

use crate::node::{Expression, Node};
use std::fmt::{Display, Formatter};

/// AST node representing a boolean literal expression.
pub struct Boolean {
    token: Token,
    value: bool,
}

impl Boolean {
    /// Creates a boolean literal node from its token and value.
    #[must_use]
    pub const fn new(token: Token, value: bool) -> Self {
        Self { token, value }
    }
}

impl Node for Boolean {
    fn dump(&self, ident: usize) -> String {
        format!("{}{}", "\t".repeat(ident), self.value)
    }

    fn token_literal(&self) -> String {
        self.token.to_string()
    }
}

impl Expression for Boolean {} // Boolean is an expression, marker trait

impl Display for Boolean {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", self.value)
    }
}

#[cfg(test)]
mod tests {
    use crate::boolean::Boolean;
    use crate::node::Node;
    use token::token::Token;

    #[test]
    fn test_boolean() {
        let b = Boolean::new(Token::Boolean(true), true);
        assert_eq!(b.token_literal(), "Boolean(true)");
    }
}
