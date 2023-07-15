package usecases

import "testing"

func TestTokenize(t *testing.T) {
	query := "this  is \"a test\" \\\"help \\\\wow"
	res := tokenize(query, SpaceBreakCheck)
	t.Log(res)
	if len(res) != 5 {
		t.Log("failed")
		t.FailNow()
	}

	if res[0] != "this" {
		t.Log("failed")
		t.FailNow()
	}

	if res[1] != "is" {
		t.Log("failed")
		t.FailNow()
	}

	if res[2] != "a test" {
		t.Log("failed")
		t.FailNow()
	}

	if res[3] != "\"help" {
		t.Log("failed")
		t.FailNow()
	}

	if res[4] != "\\wow" {
		t.Log("failed")
		t.FailNow()
	}
}
