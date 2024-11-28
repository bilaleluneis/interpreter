use std::fmt;

pub trait Node : fmt::Display {
    // Dump returns a string representation of the node object
    fn dump(&self, ident: usize) -> String;
    fn token_literal(&self) -> String;
}

pub trait Statement: Node {}

pub trait Expression: Node {}
