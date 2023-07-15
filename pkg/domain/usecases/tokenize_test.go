package usecases

import "testing"

func TestTokenize(t *testing.T) {
	query := "this  is a test"
	res := tokenize(query)
	t.Log(res)
	if len(res) != 4 {
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

	if res[2] != "a" {
		t.Log("failed")
		t.FailNow()
	}

	if res[3] != "test" {
		t.Log("failed")
		t.FailNow()
	}
}
