use std::fmt;

use crate::token::Token;

impl fmt::Debug for Token {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        // Use Display implementation for Debug output
        write!(f, "DUH{}", self)
    }
}
