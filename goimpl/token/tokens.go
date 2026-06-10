package token

import "iter"

// token is value type with internal mutation
// internally values share same slice of tokens
// but each value has its own start index and count of tokens
// read and write use mutex to prevent race conditions
type tokens struct {
	toks       []Token
	startIndex int
	count      int
}

func NewTokens(toks ...Token) tokens {
	return tokens{
		toks:       toks,
		startIndex: 0,
		count:      len(toks),
	}
}

func (t tokens) Len() int {
	return t.startIndex + t.count
}

// implement value iterators based on iter.Seq[Token]
func (t tokens) All() iter.Seq[Token] {
	return func(yield func(Token) bool) {
		for i := t.startIndex; i < t.startIndex+t.count; i++ {
			if !yield(t.toks[i]) {
				return
			}
		}
	}
}
