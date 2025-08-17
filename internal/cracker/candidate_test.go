package cracker

import (
	"testing"
)

type candidateTest struct {
	number   int
	length   int
	expected string
}

var testChars = []string{
	"a", "b", "c", "d",
}

func TestGetCandidate(t *testing.T) {
	tests := []candidateTest{
		{0, 1, "a"},
		{1, 1, "b"},
		{3, 1, "d"},
		{2, 1, "c"},

		{0, 2, "aa"},
		{1, 2, "ab"},
		{7, 2, "bd"},
		{15, 2, "dd"},

		{0, 3, "aaa"},
		{1, 3, "aab"},
		{31, 3, "bdd"},
		{63, 3, "ddd"},

		{0, 4, "aaaa"},
		{1, 4, "aaab"},
		{127, 4, "bddd"},
		{255, 4, "dddd"},

		{0, 5, "aaaaa"},
		{1, 5, "aaaab"},
		{511, 5, "bdddd"},
		{1023, 5, "ddddd"},

		{0, 6, "aaaaaa"},
		{1, 6, "aaaaab"},
		{2047, 6, "bddddd"},
		{4095, 6, "dddddd"},

		{0, 7, "aaaaaaa"},
		{1, 7, "aaaaaab"},
		{8191, 7, "bdddddd"},
		{16383, 7, "ddddddd"},

		{0, 8, "aaaaaaaa"},
		{1, 8, "aaaaaaab"},
		{32767, 8, "bddddddd"},
		{65535, 8, "dddddddd"},
	}

	for i, test := range tests {
		candidate := getCandidate(test.number, test.length, testChars)

		if candidate != test.expected {
			t.Errorf("Test %d failed: expected %s, got %s for input '%d'", i+1, test.expected, candidate, test.number)
		}
	}
}
