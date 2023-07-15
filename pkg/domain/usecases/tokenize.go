package usecases

func tokenize(in string) (out []string) {
	word := make([]rune, 0, 50)
	for _, r := range in {
		if r != ' ' {
			word = append(word, r)
		} else {
			if len(word) != 0 {
				out = append(out, string(word))
				word = word[:0]
			}
		}
	}

	if len(word) != 0 {
		out = append(out, string(word))
	}
	return
}
