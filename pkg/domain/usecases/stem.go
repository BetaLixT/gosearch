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

import "unicode"

// based of https://github.com/reiver/go-porterstemmer/tree/master
func stem(s []rune) []rune {

	if len(s) == 0 {
		return s
	}

	// Make all runes lowercase.
	for i := 0; i < len(s); i++ {
		s[i] = unicode.ToLower(s[i])
	}

	if len(s) < 3 {
		return s
	}

	// 1a
	if hasSuffix(s, []rune("sses")) ||
		hasSuffix(s, []rune("ies")) {
		s = s[:len(s)-2]
	} else if hasSuffix(s, []rune("ss")) {

	} else if hasSuffix(s, []rune("s")) {
		s = s[:len(s)-1]
	}

	// 1b
	if hasSuffix(s, []rune("eed")) {
		if measure(s[:len(s)-3]) > 0 {
			s = s[:len(s)-1]
		}
	} else if hasSuffix(s, []rune("ed")) {
		sub := s[:len(s)-2]
		if containsVowel(sub) {
			if hasSuffix(sub, []rune("at")) ||
				hasSuffix(sub, []rune("bl")) ||
				hasSuffix(sub, []rune("iz")) {
				s = s[:len(s)-1]
			} else if c := sub[len(sub)-1]; c != 'l' && c != 's' && c != 'z' && hasRepeatDoubleConsonantSuffix(sub) {
				s = sub[:len(sub)-1]
			} else if c := sub[len(sub)-1]; 1 == measure(sub) && hasConsonantVowelConsonantSuffix(sub) && c != 'w' && c != 'x' && c != 'y' {
				s = s[:len(s)-1]
				s[len(s)-1] = 'e'
			} else {
				s = sub
			}
		}
	} else if hasSuffix(s, []rune("ing")) {
		sub := s[:len(s)-3]
		if containsVowel(sub) {
			if hasSuffix(sub, []rune("at")) ||
				hasSuffix(sub, []rune("bl")) ||
				hasSuffix(sub, []rune("iz")) {
				s = s[:len(s)-2]
				s[len(s)-1] = 'e'
			} else if c := sub[len(sub)-1]; c != 'l' && c != 's' && c != 'z' && hasRepeatDoubleConsonantSuffix(sub) {
				s = sub[:len(sub)-1]
			} else if c := sub[len(sub)-1]; 1 == measure(sub) && hasConsonantVowelConsonantSuffix(sub) && 'w' != c && 'x' != c && 'y' != c {
				s = s[:len(s)-2]
				s[len(s)-1] = 'e'
			} else {
				s = sub
			}
		}
	}

	// 1c
	if len(s) < 2 {
		return s
	}
	if s[len(s)-1] == 'y' && containsVowel(s[:len(s)-1]) {
		s[len(s)-1] = 'i'
	}

	// 2
	if hasSuffix(s, []rune("ational")) {
		if measure(s[:len(s)-7]) > 0 {
			s[len(s)-5] = 'e'
			s = s[:len(s)-4]
		}
	} else if hasSuffix(s, []rune("tional")) {
		if measure(s[:len(s)-6]) > 0 {
			s = s[:len(s)-2]
		}
	} else if hasSuffix(s, []rune("enci")) {
		if measure(s[:len(s)-4]) > 0 {
			s[len(s)-1] = 'e'
		}
	} else if hasSuffix(s, []rune("anci")) {
		if measure(s[:len(s)-4]) > 0 {
			s[len(s)-1] = 'e'
		}
	} else if hasSuffix(s, []rune("izer")) {
		if measure(s[:len(s)-4]) > 0 {
			s = s[:len(s)-1]
		}
	} else if hasSuffix(s, []rune("bli")) { // --DEPARTURE--
		if measure(s[:len(s)-3]) > 0 {
			s[len(s)-1] = 'e'
		}
	} else if hasSuffix(s, []rune("alli")) {
		if measure(s[:len(s)-4]) > 0 {
			s = s[:len(s)-2]
		}
	} else if hasSuffix(s, []rune("entli")) {
		if measure(s[:len(s)-5]) > 0 {
			s = s[:len(s)-2]
		}
	} else if hasSuffix(s, []rune("eli")) {
		if measure(s[:len(s)-3]) > 0 {
			s = s[:len(s)-2]
		}
	} else if hasSuffix(s, []rune("ousli")) {
		if measure(s[:len(s)-5]) > 0 {
			s = s[:len(s)-2]
		}
	} else if hasSuffix(s, []rune("ization")) {
		if measure(s[:len(s)-7]) > 0 {
			s[len(s)-5] = 'e'
			s = s[:len(s)-4]
		}
	} else if hasSuffix(s, []rune("ation")) {
		if measure(s[:len(s)-5]) > 0 {
			s[len(s)-3] = 'e'
			s = s[:len(s)-2]
		}
	} else if hasSuffix(s, []rune("ator")) {
		if measure(s[:len(s)-4]) > 0 {
			s[len(s)-2] = 'e'
			s = s[:len(s)-1]
		}
	} else if hasSuffix(s, []rune("alism")) {
		if measure(s[:len(s)-5]) > 0 {
			s = s[:len(s)-3]
		}
	} else if hasSuffix(s, []rune("iveness")) {
		if measure(s[:len(s)-7]) > 0 {
			s = s[:len(s)-4]
		}
	} else if hasSuffix(s, []rune("fulness")) {
		if measure(s[:len(s)-7]) > 0 {
			s = s[:len(s)-4]
		}
	} else if hasSuffix(s, []rune("ousness")) {
		if measure(s[:len(s)-7]) > 0 {
			s = s[:len(s)-4]
		}
	} else if hasSuffix(s, []rune("aliti")) {
		if measure(s[:len(s)-5]) > 0 {
			s = s[:len(s)-3]
		}
	} else if hasSuffix(s, []rune("iviti")) {
		if measure(s[:len(s)-5]) > 0 {
			s[len(s)-3] = 'e'
			s = s[:len(s)-2]
		}
	} else if hasSuffix(s, []rune("biliti")) {
		if measure(s[:len(s)-6]) > 0 {
			s[len(s)-5] = 'l'
			s[len(s)-4] = 'e'
			s = s[:len(s)-3]
		}
	} else if hasSuffix(s, []rune("logi")) { // --DEPARTURE--
		if measure(s[:len(s)-4]) > 0 {
			s = s[:len(s)-1]
		}
	}

	// 3
	if hasSuffix(s, []rune("icate")) {
		if measure(s[:len(s)-5]) > 0 {
			s = s[:len(s)-3]
		}
	} else if hasSuffix(s, []rune("ative")) {
		sub := s[:len(s)-5]
		if measure(sub) > 0 {
			s = sub
		}
	} else if hasSuffix(s, []rune("alize")) {
		if measure(s[:len(s)-5]) > 0 {
			s = s[:len(s)-3]
		}
	} else if hasSuffix(s, []rune("iciti")) {
		if measure(s[:len(s)-5]) > 0 {
			s = s[:len(s)-3]
		}
	} else if hasSuffix(s, []rune("ical")) {
		if measure(s[:len(s)-4]) > 0 {
			s = s[:len(s)-2]
		}
	} else if hasSuffix(s, []rune("ful")) {
		sub := s[:len(s)-3]
		if measure(sub) > 0 {
			s = sub
		}
	} else if hasSuffix(s, []rune("ness")) {
		sub := s[:len(s)-4]
		if measure(sub) > 0 {
			s = sub
		}
	}

	// 4
	if hasSuffix(s, []rune("al")) {
		sub := s[:len(s)-2]

		if measure(sub) > 1 {
			s = s[:len(s)-2]
		}
	} else if hasSuffix(s, []rune("ance")) {
		sub := s[:len(s)-4]
		if measure(sub) > 1 {
			s = s[:len(s)-4]
		}
	} else if hasSuffix(s, []rune("ence")) {
		if measure(s[:len(s)-4]) > 1 {
			s = s[:len(s)-4]
		}
	} else if hasSuffix(s, []rune("er")) {
		sub := s[:len(s)-2]
		if measure(sub) > 1 {
			s = sub
		}
	} else if hasSuffix(s, []rune("ic")) {
		sub := s[:len(s)-2]
		if measure(sub) > 1 {
			s = sub
		}
	} else if hasSuffix(s, []rune("able")) {
		sub := s[:len(s)-4]

		if measure(sub) > 1 {
			s = sub
		}
	} else if hasSuffix(s, []rune("ible")) {
		sub := s[:len(s)-4]

		if measure(sub) > 1 {
			s = sub
		}
	} else if hasSuffix(s, []rune("ant")) {
		sub := s[:len(s)-3]

		if measure(sub) > 1 {
			s = sub
		}
	} else if hasSuffix(s, []rune("ement")) {
		sub := s[:len(s)-5]

		if measure(sub) > 1 {
			s = sub
		}
	} else if hasSuffix(s, []rune("ment")) {

		sub := s[:len(s)-4]

		if measure(sub) > 1 {
			s = sub
		}
	} else if hasSuffix(s, []rune("ent")) {

		sub := s[:len(s)-3]

		if measure(sub) > 1 {
			s = sub
		}
	} else if hasSuffix(s, []rune("ion")) {

		sub := s[:len(s)-3]

		c := sub[len(sub)-1]

		if measure(sub) > 1 && ('s' == c || 't' == c) {
			s = sub
		}
	} else if hasSuffix(s, []rune("ou")) {

		sub := s[:len(s)-2]

		if measure(sub) > 1 {
			s = sub
		}
	} else if hasSuffix(s, []rune("ism")) {

		sub := s[:len(s)-3]

		if measure(sub) > 1 {
			s = sub
		}
	} else if hasSuffix(s, []rune("ate")) {

		sub := s[:len(s)-3]

		if measure(sub) > 1 {
			s = sub
		}
	} else if hasSuffix(s, []rune("iti")) {

		sub := s[:len(s)-3]

		if measure(sub) > 1 {
			s = sub
		}
	} else if hasSuffix(s, []rune("ous")) {

		sub := s[:len(s)-3]

		if measure(sub) > 1 {
			s = sub
		}
	} else if hasSuffix(s, []rune("ive")) {

		sub := s[:len(s)-3]

		if measure(sub) > 1 {
			s = sub
		}
	} else if hasSuffix(s, []rune("ize")) {

		sub := s[:len(s)-3]

		if measure(sub) > 1 {
			s = sub
		}
	}

	if s[len(s)-1] == 'e' {
		sub := s[:len(s)-1]
		if len(sub) != 0 {
			m := measure(sub)
			if m > 1 {
				s = sub
			} else if m == 1 {
				if c := sub[len(sub)-1]; !(hasConsonantVowelConsonantSuffix(sub) && 'w' != c && 'x' != c && 'y' != c) {
					s = sub
				}
			}
		}
	}

	if len(s) > 2 && s[len(s)-2] == 'l' && s[len(s)-1] == 'l' {
		sub := s[:len(s)-1]
		if measure(sub) > 1 {
			s = sub
		}
	}

	// Return.
	return s
}

func hasSuffix(s, suffix []rune) bool {

	lenSMinusOne := len(s) - 1
	criesfixlenMinusOne := len(suffix) - 1

	if lenSMinusOne <= criesfixlenMinusOne {
		return false
	}
	if s[lenSMinusOne] != suffix[criesfixlenMinusOne] {
		return false
	}

	for i := 0; i < criesfixlenMinusOne; i++ {
		if suffix[i] != s[lenSMinusOne-criesfixlenMinusOne+i] {
			return false
		}
	}
	return true
}

func measure(s []rune) uint {

	result := uint(0)
	i := 0

	// Short Circuit.
	if 0 == len(s) {
		/////////// RETURN
		return result
	}

	// Ignore (potential) consonant sequence at the beginning of word.
	for isConsonant(s, i) {
		i++
		if i >= len(s) {
			return result
		}
	}

	// For each pair of a vowel sequence followed by a consonant sequence, increment result.
Outer:
	for i < len(s) {
		for !isConsonant(s, i) {
			i++
			if i >= len(s) {
				break Outer
			}
		}
		for isConsonant(s, i) {

			i++
			if i >= len(s) {
				result++
				break Outer
			}
		}
		result++
	}

	// Return
	return result
}

func isConsonant(s []rune, i int) bool {
	result := true

	switch s[i] {
	case 'a', 'e', 'i', 'o', 'u':
		result = false
	case 'y':
		if 0 == i {
			result = true
		} else {
			result = !isConsonant(s, i-1)
		}
	default:
		result = true
	}

	return result
}

func containsVowel(s []rune) bool {

	for i := 0; i < len(s); i++ {
		if !isConsonant(s, i) {
			return true
		}
	}
	return false
}

func hasConsonantVowelConsonantSuffix(s []rune) bool {

	result := false

	// Do it!
	if len(s) < 3 {
		result = false
	} else if isConsonant(s, len(s)-3) && !isConsonant(s, len(s)-2) && isConsonant(s, len(s)-1) {
		result = true
	} else {
		result = false
	}

	// Return
	return result
}

func hasRepeatDoubleConsonantSuffix(s []rune) bool {

	result := false

	// Do it!
	if 2 > len(s) {
		result = false
	} else if s[len(s)-1] == s[len(s)-2] && isConsonant(s, len(s)-1) { // Will using isConsonant() cause a problem with "YY"?
		result = true
	} else {
		result = false
	}

	// Return,
	return result
}
