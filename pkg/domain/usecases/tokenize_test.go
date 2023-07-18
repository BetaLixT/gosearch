package usecases

import "testing"

func TestTokenize(t *testing.T) {
	query := "this  is \"a test\" \\\"help \\\\wow helping the dogs that were washed eeg"
	res := tokenize(query, SpaceBreakCheck)
	t.Log(res)
	if len(res) != 8 {
		t.Log("failed len")
		t.FailNow()
	}

	if res[0] != "a test" {
		t.Log("failed 0")
		t.FailNow()
	}

	if res[1] != "\"help" {
		t.Log("failed 1")
		t.FailNow()
	}

	if res[2] != "\\wow" {
		t.Log("failed 2")
		t.FailNow()
	}

	if res[3] != "help" {
		t.Log("failed 3")
		t.FailNow()
	}

	if res[4] != "dog" {
		t.Log("failed 4")
		t.FailNow()
	}

	if res[5] != "were" {
		t.Log("failed 6")
		t.FailNow()
	}

	if res[6] != "wash" {
		t.Log("failed 7")
		t.FailNow()
	}
}
