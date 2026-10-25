use std::fmt::{Display, Formatter};

/// Expression nodes: anything that produces a value.
pub enum Expression {
    /// A bare name such as `five`.
    Identifier(String),
    /// An integer literal such as `5`.
    Integer(u64),
    /// `true` or `false`.
    Boolean(bool),
}

impl Display for Expression {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::Identifier(name) => write!(f, "{name}"),
            Self::Integer(value) => write!(f, "{value}"),
            Self::Boolean(b) => write!(f, "{b}"),
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn integer_displays_its_value() {
        let e = Expression::Integer(5);
        assert!(matches!(e, Expression::Integer(5)));
        assert_eq!(e.to_string(), "5");
    }

    #[test]
    fn identifier_displays_its_name() {
        assert_eq!(
            Expression::Identifier("five".to_string()).to_string(),
            "five"
        );
    }

    #[test]
    fn boolean_expression_renders_its_value() {
        let e = Expression::Boolean(true);
        assert_eq!(e.to_string(), "true");
    }
}
