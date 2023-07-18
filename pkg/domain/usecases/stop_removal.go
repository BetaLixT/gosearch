package usecases

import "strings"

func filterStop(in []string) []string {

	lastValid := len(in) - 1
	for idx := 0; idx < len(in) && idx >= lastValid; {
		if !checkFiltered(in[idx]) {
			idx++
		} else {
			in[idx] = in[lastValid]
			lastValid--
		}
	}
	return in[:lastValid]
}

func checkFiltered(word string) bool {

	first := StopsFirst
	last := StopsLast
	mid := StopsMiddle // len(stops) / 2
	for (mid >= first && mid <= last) && (mid != first || mid != last) {
		comp := strings.Compare(word, stops[mid])
		if comp == 0 {
			return true
		}
		if comp == -1 {
			last = mid - 1
		} else {
			first = mid + 1
		}
		mid = ((last - first) / 2) + first
	}

	if mid >= StopsFirst && mid <= StopsLast {
		return strings.Compare(word, stops[mid]) == 0
	}

	return false
}

const (
	StopsMiddle = 16
	StopsFirst  = 0
	StopsLast   = 32
)

var stops = []string{
	"a",
	"an",
	"and",
	"are",
	"as",
	"at",
	"be",
	"but",
	"by",
	"for",
	"if",
	"in",
	"into",
	"is",
	"it",
	"no",
	"not",
	"of",
	"on",
	"or",
	"such",
	"that",
	"the",
	"their",
	"then",
	"there",
	"these",
	"they",
	"this",
	"to",
	"was",
	"will",
	"with",
}
