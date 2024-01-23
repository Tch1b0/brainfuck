package token

import "fmt"

func Tokenize(data string) ([]int, error) {
	tokens := []int{}

	for _, c := range data {
		if c == ' ' || c == '\n' || c == '\t' || c == '\r' {
			continue
		}

		t, err := getToken(c)
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, t)
	}

	return tokens, nil
}

func getToken(c rune) (int, error) {
	val := 0
	var err error = nil

	switch c {
	case '>':
		val = INC_DP
	case '<':
		val = DEC_DP
	case '+':
		val = INC_VAL
	case '-':
		val = DEC_VAL
	case '.':
		val = OUT
	case ',':
		val = IN
	case '[':
		val = BEGIN_LOOP
	case ']':
		val = END_LOOP
	default:
		err = fmt.Errorf("invalid token '%c'", c)
	}

	return val, err
}
