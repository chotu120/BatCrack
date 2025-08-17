package cracker

import (
	"strings"
)

func getCandidate(num int, length int, chars []string) string {
	base := len(chars)
	cand := make([]string, length)

	for i := length - 1; i >= 0; i-- {
		index := num % base
		cand[i] = chars[index]
		num /= base
	}

	return strings.Join(cand, "")
}
