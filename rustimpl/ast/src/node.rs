use std::fmt;

/// Common behavior shared by all AST nodes.
pub trait Node: fmt::Display {
    // Dump returns a string representation of the node object
    /// Renders a formatted representation with the given indentation level.
    fn dump(&self, ident: usize) -> String;
    /// Returns the literal form of the token that produced this node.
    fn token_literal(&self) -> String;
}

/// Marker trait for statement nodes.
pub trait Statement: Node {}

/// Marker trait for expression nodes.
pub trait Expression: Node {}
