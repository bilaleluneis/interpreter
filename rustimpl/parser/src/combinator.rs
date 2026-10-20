use scanner::lexer::Lexer;

#[derive(Clone, Debug)]
/// Minimal parser wrapper used to compose parser transformations.
pub struct CombinatorParser<L: Lexer> {
    lexer: L,
}

impl<L: Lexer> CombinatorParser<L> {
    pub(crate) const fn new(lexer: L) -> Self {
        Self { lexer }
    }

    // Example: a method to parse something
    /// Runs the parser entry point.
    pub const fn parse(&self) {
        let Self { lexer: _ } = self;
        // Implement your parsing logic here
    }
}

// Example of a combinator function
/// Applies a parser transformation and returns the updated parser.
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
    // use crate::combinator::{self, CombinatorParser, map};
    // use scanner::lazy::LazyLexer;

    // #[test]
    // fn test_combinator_parser() {
    //     let lexer = LazyLexer::new("=+,;(){}");
    //     let mut parser = CombinatorParser::new(lexer);
    //     parser.parse();
    //     let _ = map(
    //         parser,
    //         |arg0: &mut combinator::CombinatorParser<scanner::lazy::LazyLexer<'_>>| {
    //             super::CombinatorParser::parse(&*arg0)
    //         },
    //     );
    // }
}
