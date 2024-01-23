package token

func IsGrammarValid(tokens []int) bool {
	scopes := 0

	for _, tok := range tokens {
		if tok == BEGIN_LOOP {
			scopes++
		} else if tok == END_LOOP {
			scopes--
		}

		if scopes < 0 {
			return false
		}
	}

	return scopes == 0
}
