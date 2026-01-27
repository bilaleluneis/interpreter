use crate::node::{Expression, Node};
use std::fmt::{Display, Formatter};

pub struct Boolean {
    token: Token,
    value: bool,
}

impl Boolean {
    pub fn new(token: Token, value: bool) -> Boolean {
        Boolean { token, value }
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
    use scanner::token::Token;

    #[test]
    fn test_boolean() {
        let b = Boolean::new(Token::Boolean(true), true);
        assert_eq!(b.token_literal(), "Boolean(true)");
    }
}
