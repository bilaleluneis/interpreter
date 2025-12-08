use scanner::lexer::Lexer;

#[derive(Clone, Debug)]
pub struct CombinatorParser<L: Lexer> {
    lexer: L,
}

impl<L: Lexer> CombinatorParser<L> {
    pub fn new(lexer: L) -> Self {
        CombinatorParser { lexer }
    }

    // Example: a method to parse something
    pub fn parse(&mut self) {
        // Implement your parsing logic here
    }
}

// Example of a combinator function
pub fn map<L, F>(mut parser: CombinatorParser<L>, f: F) -> CombinatorParser<L>
where
    L: Lexer,
    F: Fn(&mut CombinatorParser<L>),
{
    f(&mut parser);
    parser
}

#[cfg(test)]
mod tests {
    use crate::combinator::{CombinatorParser, map};
    use scanner::lazy::LazyLexer;

    #[test]
    fn test_combinator_parser() {
        let lexer = LazyLexer::new("=+,;(){}");
        let mut parser = CombinatorParser::new(lexer);
        parser.parse();
        let _ = map(parser, |p| p.parse());
    }
}
