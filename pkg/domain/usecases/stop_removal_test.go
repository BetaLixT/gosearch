package usecases

import "testing"

func TestCheckFiltered(t *testing.T) {
	type test struct {
		word     string
		expected bool
	}

	tests := []test{}
	tests = append(tests, test{"no", true})
	for idx := range stops {
		tests = append(tests, test{stops[idx], true})
	}
	tests = append(
		tests,
		test{"banana", false},
		test{"anddr", false},
		test{"zzzahan", false},
		test{"ab", false},
		test{"ana", false},
		test{"ano", false},
		test{"asz", false},
		test{"au", false},
		test{"bf", false},
		test{"buu", false},
		test{"bz", false},
		test{"fos", false},
		test{"ig", false},
		test{"noa", false},
		test{"nou", false},
		test{"og", false},
		test{"oo", false},
		test{"os", false},
		test{"suci", false},
		test{"thau", false},
		test{"thea", false},
		test{"theis", false},
		test{"theo", false},
		test{"therf", false},
		test{"thesf", false},
		test{"thez", false},
		test{"thit", false},
		test{"tp", false},
		test{"wat", false},
		test{"wilm", false},
		test{"witi", false},
	)

	for idx := range tests {
		if checkFiltered(tests[idx].word) != tests[idx].expected {
			t.Logf(
				"invalid result for %s expected [%v] got [%v]",
				tests[idx].word,
				tests[idx].expected,
				!tests[idx].expected,
			)
			t.Fail()
		}
	}
}
