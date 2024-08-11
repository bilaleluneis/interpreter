use scanner::lexer::Lexer;

#[derive(Clone, Copy, Debug)]
pub struct CombinatorParser<L: Lexer> {
    lexer: L,
}

impl<L: Lexer> FnOnce<()> for CombinatorParser<L> {
    type Output = CombinatorParser<L>;

    extern "rust-call" fn call_once(self, _: ()) -> Self::Output {
        self
    }
}

impl<L: Lexer> FnMut<()> for CombinatorParser<L> {
    extern "rust-call" fn call_mut(&mut self, _: ()) -> Self::Output {
        *self
    }
}

impl<L: Lexer> Fn<()> for CombinatorParser<L> {
    extern "rust-call" fn call(&self, _: ()) -> Self::Output {
        *self
    }
}

pub type CombinatorParserFn<L> = fn(L) -> CombinatorParser<L>;

#[cfg(test)]
mod tests {
    use scanner::lazy::LazyLexer;

    use crate::combinator::CombinatorParser;

    #[test]
    fn test_combinator_parser() {
        let lexer = LazyLexer::new("=+,;(){}");
        let parser = CombinatorParser { lexer };
        let _ = parser();
    }
}

