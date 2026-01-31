use std::fmt::{Debug, Formatter, Result};

use crate::token::Token;

impl Debug for Token {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result {
        write!(f, "{}", self)
    }
}
