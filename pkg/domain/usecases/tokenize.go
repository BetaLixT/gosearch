package usecases

import "unicode"

// TODO: might remove this since it probs wont be inlined
type WordBreakCheck func(rune) bool

func tokenize(
	in string,
	brCheck WordBreakCheck,
	onlyLiteral bool,
) (out []string) {
	word := make([]rune, 0, 50)

	escaped := false
	literal := false
	for _, r := range in {
		if escaped || literal || !brCheck(r) {

			if !escaped {
				// escape handling
				if r == '\\' {
					escaped = true
					continue
				}

				// litral handling
				if r == '"' {
					if literal {
						literal = false
						out = append(out, string(word))
						word = word[:0]
					} else {
						literal = true
					}
					continue
				}
			} else {
				escaped = false
			}

			word = append(word, r)
			continue
		}

		if len(word) != 0 {
			if !checkFiltered(string(word)) {
				out = append(out, string(stem(word)))
			}
			word = word[:0]
		}

	}

	if len(word) != 0 {
		if !checkFiltered(string(word)) {
			out = append(out, string(stem(word)))
		}
		word = word[:0]
	}
	return
}

func SpaceBreakCheck(r rune) bool {
	return r == ' '
}

func SpecialCharacterBreakCheck(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsNumber(r)
}
