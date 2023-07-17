package usecases

import "testing"

func TestTokenize(t *testing.T) {
	query := "this  is \"a test\" \\\"help \\\\wow"
	res := tokenize(query, SpaceBreakCheck)
	t.Log(res)
	if len(res) != 5 {
		t.Log("failed len")
		t.FailNow()
	}

	if res[0] != "this" {
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
}
