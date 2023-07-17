package usecases

import "testing"

func TestTokenize(t *testing.T) {
	query := "this  is \"a test\" \\\"help \\\\wow helping the dogs that were washed eeg"
	res := tokenize(query, SpaceBreakCheck)
	t.Log(res)
	if len(res) != 12 {
		t.Log("failed len")
		t.FailNow()
	}

	if res[0] != "thi" {
		t.Log("failed 0")
		t.FailNow()
	}

	if res[1] != "is" {
		t.Log("failed 1")
		t.FailNow()
	}

	if res[2] != "a test" {
		t.Log("failed 2")
		t.FailNow()
	}

	if res[3] != "\"help" {
		t.Log("failed 3")
		t.FailNow()
	}

	if res[4] != "\\wow" {
		t.Log("failed 4")
		t.FailNow()
	}

	if res[5] != "help" {
		t.Log("failed 5")
		t.FailNow()
	}

	if res[6] != "the" {
		t.Log("failed 6")
		t.FailNow()
	}

	if res[7] != "dog" {
		t.Log("failed 7")
		t.FailNow()
	}

	if res[8] != "that" {
		t.Log("failed 8")
		t.FailNow()
	}

	if res[9] != "were" {
		t.Log("failed 9")
		t.FailNow()
	}

	if res[10] != "wash" {
		t.Log("failed 10")
		t.FailNow()
	}
}
