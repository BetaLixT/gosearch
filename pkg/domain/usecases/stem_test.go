/*
Copyright (c) 2013 Charles Iliya Krempeaux <charles@reptile.ca> :: http://changelog.ca/

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package usecases

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestStep1a(t *testing.T) {

	i := 0

	tests := make([]struct {
		S        []rune
		Expected []rune
	}, 12)

	tests[i].S = []rune("caresses")
	tests[i].Expected = []rune("caress")
	i++

	tests[i].S = []rune("ponies")
	tests[i].Expected = []rune("poni")
	i++

	tests[i].S = []rune("ties")
	tests[i].Expected = []rune("ti")
	i++

	tests[i].S = []rune("caress")
	tests[i].Expected = []rune("caress")
	i++

	tests[i].S = []rune("cats")
	tests[i].Expected = []rune("cat")
	i++

	for _, datum := range tests {
		for i = 0; i < len(datum.S); i++ {

			actual := make([]rune, len(datum.S))
			copy(actual, datum.S)

			actual = stem(actual)

			lenActual := len(actual)
			lenExpected := len(datum.Expected)

			equal := true
			if 0 == lenActual && 0 == lenExpected {
				equal = true
			} else if lenActual != lenExpected {
				equal = false
			} else if actual[0] != datum.Expected[0] {
				equal = false
			} else if actual[lenActual-1] != datum.Expected[lenExpected-1] {
				equal = false
			} else {
				for j := 0; j < lenActual; j++ {

					if actual[j] != datum.Expected[j] {
						equal = false
					}
				}
			}

			if !equal {
				t.Errorf(
					"Did NOT get what was expected for calling step1a() on [%s]. Expect [%s] but got [%s]",
					string(datum.S),
					string(datum.Expected),
					string(actual),
				)
			}
		} // for
	}
}

func TestStep1b(t *testing.T) {

	i := 0

	tests := make([]struct {
		S        []rune
		Expected []rune
	}, 17)

	tests[i].S = []rune("feed")
	tests[i].Expected = []rune("feed")
	i++

	tests[i].S = []rune("agreed")
	tests[i].Expected = []rune("agre")
	i++

	tests[i].S = []rune("plastered")
	tests[i].Expected = []rune("plaster")
	i++

	tests[i].S = []rune("bled")
	tests[i].Expected = []rune("bled")
	i++

	tests[i].S = []rune("motoring")
	tests[i].Expected = []rune("motor")
	i++

	tests[i].S = []rune("sing")
	tests[i].Expected = []rune("sing")
	i++

	tests[i].S = []rune("conflated")
	tests[i].Expected = []rune("conflat")
	i++

	tests[i].S = []rune("troubled")
	tests[i].Expected = []rune("troubl")
	i++

	tests[i].S = []rune("sized")
	tests[i].Expected = []rune("size")
	i++

	tests[i].S = []rune("hopping")
	tests[i].Expected = []rune("hop")
	i++

	tests[i].S = []rune("tanned")
	tests[i].Expected = []rune("tan")
	i++

	tests[i].S = []rune("falling")
	tests[i].Expected = []rune("fall")
	i++

	tests[i].S = []rune("hissing")
	tests[i].Expected = []rune("hiss")
	i++

	tests[i].S = []rune("fizzed")
	tests[i].Expected = []rune("fizz")
	i++

	tests[i].S = []rune("failing")
	tests[i].Expected = []rune("fail")
	i++

	tests[i].S = []rune("filing")
	tests[i].Expected = []rune("file")
	i++

	for _, datum := range tests {

		actual := make([]rune, len(datum.S))
		copy(actual, datum.S)

		actual = stem(actual)

		lenActual := len(actual)
		lenExpected := len(datum.Expected)

		equal := true
		if 0 == lenActual && 0 == lenExpected {
			equal = true
		} else if lenActual != lenExpected {
			equal = false
		} else if actual[0] != datum.Expected[0] {
			equal = false
		} else if actual[lenActual-1] != datum.Expected[lenExpected-1] {
			equal = false
		} else {
			for j := 0; j < lenActual; j++ {

				if actual[j] != datum.Expected[j] {
					equal = false
				}
			}
		}

		if !equal {
			t.Errorf(
				"Did NOT get what was expected for calling step1b() on [%s]. Expect [%s] but got [%s]",
				string(datum.S),
				string(datum.Expected),
				string(actual),
			)
		}
	}
}

// Test for issue listed here:
// https://github.com/reiver/go-porterstemmer/issues/1
//
// StemString("ion") was causing runtime exception
func TestStemStringIon(t *testing.T) {

	expected := "ion"

	s := "ion"
	actual := stem([]rune(s))
	if expected != string(actual) {
		t.Errorf("Input: [%s] -> Actual: [%s]. Expected: [%s]", s, string(actual), expected)
	}
}

// Test for issue listed here:
// https://github.com/reiver/go-porterstemmer/pull/10
//
// StemString("eeg") was causing runtime exception
func TestStemStringEeg(t *testing.T) {

	expected := "eeg"

	s := "eeg"
	actual := stem([]rune(s))
	if expected != string(actual) {
		t.Errorf("Input: [%s] -> Actual: [%s]. Expected: [%s]", s, string(actual), expected)
	}
}

func TestIsConsontant(t *testing.T) {

	i := 0

	tests := make([]struct {
		S        []rune
		Expected []bool
	}, 12)

	tests[i].S = []rune("apple")
	tests[i].Expected = []bool{false, true, true, true, false}
	i++

	tests[i].S = []rune("cyan")
	tests[i].Expected = []bool{true, false, false, true}
	i++

	tests[i].S = []rune("connects")
	tests[i].Expected = []bool{true, false, true, true, false, true, true, true}
	i++

	tests[i].S = []rune("yellow")
	tests[i].Expected = []bool{true, false, true, true, false, true}
	i++

	tests[i].S = []rune("excellent")
	tests[i].Expected = []bool{false, true, true, false, true, true, false, true, true}
	i++

	tests[i].S = []rune("yuk")
	tests[i].Expected = []bool{true, false, true}
	i++

	tests[i].S = []rune("syzygy")
	tests[i].Expected = []bool{true, false, true, false, true, false}
	i++

	tests[i].S = []rune("school")
	tests[i].Expected = []bool{true, true, true, false, false, true}
	i++

	tests[i].S = []rune("pay")
	tests[i].Expected = []bool{true, false, true}
	i++

	tests[i].S = []rune("golang")
	tests[i].Expected = []bool{true, false, true, false, true, true}
	i++

	// NOTE: The Porter Stemmer technical should make a mistake on the second "y".
	//       Really, both the 1st and 2nd "y" are consontants. But
	tests[i].S = []rune("sayyid")
	tests[i].Expected = []bool{true, false, true, false, false, true}
	i++

	tests[i].S = []rune("ya")
	tests[i].Expected = []bool{true, false}
	i++

	for _, datum := range tests {
		for i = 0; i < len(datum.S); i++ {

			if actual := isConsonant(datum.S, i); actual != datum.Expected[i] {
				t.Errorf(
					"Did NOT get what was expected for calling isConsonant() on [%s] at [%d] (i.e., [%s]). Expect [%t] but got [%t]",
					string(datum.S),
					i,
					string(datum.S[i]),
					datum.Expected[i],
					actual,
				)
			}
		} // for
	}
}

func TestStep1c(t *testing.T) {

	i := 0

	tests := make([]struct {
		S        []rune
		Expected []rune
	}, 17)

	tests[i].S = []rune("happy")
	tests[i].Expected = []rune("happi")
	i++

	tests[i].S = []rune("sky")
	tests[i].Expected = []rune("sky")
	i++

	tests[i].S = []rune("apology")
	tests[i].Expected = []rune("apolog")
	i++

	for _, datum := range tests {

		actual := make([]rune, len(datum.S))
		copy(actual, datum.S)

		actual = stem(actual)

		lenActual := len(actual)
		lenExpected := len(datum.Expected)

		equal := true
		if 0 == lenActual && 0 == lenExpected {
			equal = true
		} else if lenActual != lenExpected {
			equal = false
		} else if actual[0] != datum.Expected[0] {
			equal = false
		} else if actual[lenActual-1] != datum.Expected[lenExpected-1] {
			equal = false
		} else {
			for j := 0; j < lenActual; j++ {

				if actual[j] != datum.Expected[j] {
					equal = false
				}
			}
		}

		if !equal {
			t.Errorf(
				"Did NOT get what was expected for calling step1c() on [%s]. Expect [%s] but got [%s]",
				string(datum.S),
				string(datum.Expected),
				string(actual),
			)
		}
	}
}

func TestStep2(t *testing.T) {

	i := 0

	tests := make([]struct {
		S        []rune
		Expected []rune
	}, 22)

	tests[i].S = []rune("relational")
	tests[i].Expected = []rune("relat")
	i++

	tests[i].S = []rune("conditional")
	tests[i].Expected = []rune("condit")
	i++

	tests[i].S = []rune("rational")
	tests[i].Expected = []rune("ration")
	i++

	tests[i].S = []rune("valenci")
	tests[i].Expected = []rune("valenc")
	i++

	tests[i].S = []rune("hesitanci")
	tests[i].Expected = []rune("hesit")
	i++

	tests[i].S = []rune("digitizer")
	tests[i].Expected = []rune("digit")
	i++

	tests[i].S = []rune("conformabli")
	tests[i].Expected = []rune("conform")
	i++

	tests[i].S = []rune("radicalli")
	tests[i].Expected = []rune("radic")
	i++

	tests[i].S = []rune("differentli")
	tests[i].Expected = []rune("differ")
	i++

	tests[i].S = []rune("vileli")
	tests[i].Expected = []rune("vile")
	i++

	tests[i].S = []rune("analogousli")
	tests[i].Expected = []rune("analog")
	i++

	tests[i].S = []rune("vietnamization")
	tests[i].Expected = []rune("vietnam")
	i++

	tests[i].S = []rune("predication")
	tests[i].Expected = []rune("predic")
	i++

	tests[i].S = []rune("operator")
	tests[i].Expected = []rune("oper")
	i++

	tests[i].S = []rune("feudalism")
	tests[i].Expected = []rune("feudal")
	i++

	tests[i].S = []rune("decisiveness")
	tests[i].Expected = []rune("decis")
	i++

	tests[i].S = []rune("hopefulness")
	tests[i].Expected = []rune("hope")
	i++

	tests[i].S = []rune("callousness")
	tests[i].Expected = []rune("callous")
	i++

	tests[i].S = []rune("formaliti")
	tests[i].Expected = []rune("formal")
	i++

	tests[i].S = []rune("sensitiviti")
	tests[i].Expected = []rune("sensit")
	i++

	tests[i].S = []rune("sensibiliti")
	tests[i].Expected = []rune("sensibl")
	i++

	for _, datum := range tests {

		actual := make([]rune, len(datum.S))
		copy(actual, datum.S)

		actual = stem(actual)

		lenActual := len(actual)
		lenExpected := len(datum.Expected)

		equal := true
		if 0 == lenActual && 0 == lenExpected {
			equal = true
		} else if lenActual != lenExpected {
			equal = false
		} else if actual[0] != datum.Expected[0] {
			equal = false
		} else if actual[lenActual-1] != datum.Expected[lenExpected-1] {
			equal = false
		} else {
			for j := 0; j < lenActual; j++ {

				if actual[j] != datum.Expected[j] {
					equal = false
				}
			}
		}

		if !equal {
			t.Errorf(
				"Did NOT get what was expected for calling step2() on [%s]. Expect [%s] but got [%s]",
				string(datum.S),
				string(datum.Expected),
				string(actual),
			)
		}
	}
}

func TestStep3(t *testing.T) {

	i := 0

	tests := make([]struct {
		S        []rune
		Expected []rune
	}, 22)

	tests[i].S = []rune("triplicate")
	tests[i].Expected = []rune("triplic")
	i++

	tests[i].S = []rune("formative")
	tests[i].Expected = []rune("form")
	i++

	tests[i].S = []rune("formalize")
	tests[i].Expected = []rune("formal")
	i++

	tests[i].S = []rune("electriciti")
	tests[i].Expected = []rune("electr")
	i++

	tests[i].S = []rune("electrical")
	tests[i].Expected = []rune("electr")
	i++

	tests[i].S = []rune("hopeful")
	tests[i].Expected = []rune("hope")
	i++

	tests[i].S = []rune("goodness")
	tests[i].Expected = []rune("good")
	i++

	for _, datum := range tests {

		actual := make([]rune, len(datum.S))
		copy(actual, datum.S)

		actual = stem(actual)

		lenActual := len(actual)
		lenExpected := len(datum.Expected)

		equal := true
		if 0 == lenActual && 0 == lenExpected {
			equal = true
		} else if lenActual != lenExpected {
			equal = false
		} else if actual[0] != datum.Expected[0] {
			equal = false
		} else if actual[lenActual-1] != datum.Expected[lenExpected-1] {
			equal = false
		} else {
			for j := 0; j < lenActual; j++ {

				if actual[j] != datum.Expected[j] {
					equal = false
				}
			}
		}

		if !equal {
			t.Errorf(
				"Did NOT get what was expected for calling step3() on [%s]. Expect [%s] but got [%s]",
				string(datum.S),
				string(datum.Expected),
				string(actual),
			)
		}
	}
}

func TestStep4(t *testing.T) {

	i := 0

	tests := make([]struct {
		S        []rune
		Expected []rune
	}, 20)

	tests[i].S = []rune("revival")
	tests[i].Expected = []rune("reviv")
	i++

	tests[i].S = []rune("allowance")
	tests[i].Expected = []rune("allow")
	i++

	tests[i].S = []rune("inference")
	tests[i].Expected = []rune("infer")
	i++

	tests[i].S = []rune("airliner")
	tests[i].Expected = []rune("airlin")
	i++

	tests[i].S = []rune("gyroscopic")
	tests[i].Expected = []rune("gyroscop")
	i++

	tests[i].S = []rune("adjustable")
	tests[i].Expected = []rune("adjust")
	i++

	tests[i].S = []rune("defensible")
	tests[i].Expected = []rune("defens")
	i++

	tests[i].S = []rune("irritant")
	tests[i].Expected = []rune("irrit")
	i++

	tests[i].S = []rune("replacement")
	tests[i].Expected = []rune("replac")
	i++

	tests[i].S = []rune("adjustment")
	tests[i].Expected = []rune("adjust")
	i++

	tests[i].S = []rune("dependent")
	tests[i].Expected = []rune("depend")
	i++

	tests[i].S = []rune("adoption")
	tests[i].Expected = []rune("adopt")
	i++

	tests[i].S = []rune("homologou")
	tests[i].Expected = []rune("homolog")
	i++

	tests[i].S = []rune("communism")
	tests[i].Expected = []rune("commun")
	i++

	tests[i].S = []rune("activate")
	tests[i].Expected = []rune("activ")
	i++

	tests[i].S = []rune("angulariti")
	tests[i].Expected = []rune("angular")
	i++

	tests[i].S = []rune("homologous")
	tests[i].Expected = []rune("homolog")
	i++

	tests[i].S = []rune("effective")
	tests[i].Expected = []rune("effect")
	i++

	tests[i].S = []rune("bowdlerize")
	tests[i].Expected = []rune("bowdler")
	i++

	for _, datum := range tests {

		actual := make([]rune, len(datum.S))
		copy(actual, datum.S)

		actual = stem(actual)

		lenActual := len(actual)
		lenExpected := len(datum.Expected)

		equal := true
		if 0 == lenActual && 0 == lenExpected {
			equal = true
		} else if lenActual != lenExpected {
			equal = false
		} else if actual[0] != datum.Expected[0] {
			equal = false
		} else if actual[lenActual-1] != datum.Expected[lenExpected-1] {
			equal = false
		} else {
			for j := 0; j < lenActual; j++ {

				if actual[j] != datum.Expected[j] {
					equal = false
				}
			}
		}

		if !equal {
			t.Errorf(
				"Did NOT get what was expected for calling step4() on [%s]. Expect [%s] but got [%s]",
				string(datum.S),
				string(datum.Expected),
				string(actual),
			)
		}
	}
}

func TestStep5a(t *testing.T) {

	i := 0

	tests := make([]struct {
		S        []rune
		Expected []rune
	}, 3)

	tests[i].S = []rune("probate")
	tests[i].Expected = []rune("probat")
	i++

	tests[i].S = []rune("rate")
	tests[i].Expected = []rune("rate")
	i++

	tests[i].S = []rune("cease")
	tests[i].Expected = []rune("ceas")
	i++

	for _, datum := range tests {

		actual := make([]rune, len(datum.S))
		copy(actual, datum.S)

		actual = stem(actual)

		lenActual := len(actual)
		lenExpected := len(datum.Expected)

		equal := true
		if 0 == lenActual && 0 == lenExpected {
			equal = true
		} else if lenActual != lenExpected {
			equal = false
		} else if actual[0] != datum.Expected[0] {
			equal = false
		} else if actual[lenActual-1] != datum.Expected[lenExpected-1] {
			equal = false
		} else {
			for j := 0; j < lenActual; j++ {

				if actual[j] != datum.Expected[j] {
					equal = false
				}
			}
		}

		if !equal {
			t.Errorf(
				"Did NOT get what was expected for calling step5a() on [%s]. Expect [%s] but got [%s]",
				string(datum.S),
				string(datum.Expected),
				string(actual),
			)
		}
	}
}

func TestStep5b(t *testing.T) {

	i := 0

	tests := make([]struct {
		S        []rune
		Expected []rune
	}, 3)

	tests[i].S = []rune("controll")
	tests[i].Expected = []rune("control")
	i++

	tests[i].S = []rune("roll")
	tests[i].Expected = []rune("roll")
	i++

	for _, datum := range tests {

		actual := make([]rune, len(datum.S))
		copy(actual, datum.S)

		actual = stem(actual)

		lenActual := len(actual)
		lenExpected := len(datum.Expected)

		equal := true
		if 0 == lenActual && 0 == lenExpected {
			equal = true
		} else if lenActual != lenExpected {
			equal = false
		} else if actual[0] != datum.Expected[0] {
			equal = false
		} else if actual[lenActual-1] != datum.Expected[lenExpected-1] {
			equal = false
		} else {
			for j := 0; j < lenActual; j++ {

				if actual[j] != datum.Expected[j] {
					equal = false
				}
			}
		}

		if !equal {
			t.Errorf(
				"Did NOT get what was expected for calling step5b() on [%s]. Expect [%s] but got [%s]",
				string(datum.S),
				string(datum.Expected),
				string(actual),
			)
		}
	}
}

func TestStemString(t *testing.T) {

	testDataDirName := "testdata"

	_, err := os.Stat(testDataDirName)
	if nil != err {
		_ = os.Mkdir(testDataDirName, 0755)
	}
	_, err = os.Stat(testDataDirName)
	if nil != err {
		t.Errorf(
			"The test data folder ([%s]) does not exists (and could not create it). Received error: [%v]",
			testDataDirName,
			err,
		)
		/////// RETURN
		return
	}

	vocFileName := testDataDirName + "/voc.txt"
	_, err = os.Stat(vocFileName)
	if nil != err {

		vocHref := "http://tartarus.org/martin/PorterStemmer/voc.txt"

		resp, err := http.Get(vocHref)
		if nil != err {
			t.Errorf(
				"Could not download test file (from web) from URL: [%s]. Received error: [%v]",
				vocHref,
				err,
			)
			/////////// RETURN
			return
		}

		respBody, err := ioutil.ReadAll(resp.Body)
		if nil != err {
			t.Errorf(
				"Error loading the contents of from URL: [%s]. Received error: [%v].",
				vocHref,
				err,
			)
			/////////// RETURN
			return
		}

		_ = ioutil.WriteFile(vocFileName, respBody, 0644)

	}
	vocFd, err := os.Open(vocFileName)
	if nil != err {
		t.Errorf("Could NOT open testdata file: [%s]. Received error: [%v]", vocFileName, err)
		/////// RETURN
		return
	}
	defer vocFd.Close()

	voc := bufio.NewReaderSize(vocFd, 1024)

	outFileName := testDataDirName + "/output.txt"
	_, err = os.Stat(outFileName)
	if nil != err {

		outHref := "http://tartarus.org/martin/PorterStemmer/output.txt"

		resp, err := http.Get(outHref)
		if nil != err {
			t.Errorf(
				"Could not download test file (from web) from URL: [%s]. Received error: [%v]",
				outHref,
				err,
			)
			/////////// RETURN
			return
		}

		respBody, err := ioutil.ReadAll(resp.Body)
		if nil != err {
			t.Errorf(
				"Error loading the contents of from URL: [%s]. Received error: [%v].",
				outHref,
				err,
			)
			/////////// RETURN
			return
		}

		_ = ioutil.WriteFile(outFileName, respBody, 0644)

	}
	outFd, err := os.Open(outFileName)
	if nil != err {
		t.Errorf("Could NOT open testdata file: [%s]. Received error: [%v]", outFileName, err)
		/////// RETURN
		return
	}
	defer outFd.Close()

	out := bufio.NewReaderSize(outFd, 1024)

	for {

		vocS, err := voc.ReadString('\n')
		if nil != err {
			/////// BREAK
			break
		}

		vocS = strings.Trim(vocS, "\n\r\t ")

		expected, err := out.ReadString('\n')
		if nil != err {
			t.Errorf(
				"Received unexpected error when trying to read a line from [%s]. Received error: [%v]",
				outFileName,
				err,
			)
			/////// BREAK
			break

		}

		expected = strings.Trim(expected, "\n\r\t ")

		actual := stem([]rune(vocS))
		if expected != string(actual) {
			t.Errorf("Input: [%s] -> Actual: [%s]. Expected: [%s]", vocS, string(actual), expected)
		}
	}
}

const maxFuzzLen = 6

func TestStemFuzz(t *testing.T) {

	input := []byte{'a'}
	for len(input) < maxFuzzLen {
		// test input

		panicked := false
		func() {
			defer func() { panicked = recover() != nil }()
			stem([]rune(string(input)))
		}()
		if panicked {
			t.Errorf("StemString panicked for input '%s'", input)
		}

		// if all z's extend
		if allZs(input) {
			input = bytes.Repeat([]byte{'a'}, len(input)+1)
		} else {
			// increment
			input = incrementBytes(input)
		}
	}
}

func incrementBytes(in []byte) []byte {
	rv := make([]byte, len(in))
	copy(rv, in)
	for i := len(rv) - 1; i >= 0; i-- {
		if rv[i]+1 == '{' {
			rv[i] = 'a'
			continue
		}
		rv[i] = rv[i] + 1
		break

	}
	return rv
}

func allZs(in []byte) bool {
	for _, b := range in {
		if b != 'z' {
			return false
		}
	}
	return true
}
