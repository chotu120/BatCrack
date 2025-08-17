package cracker

import (
	"math"
)

func worker(id, step int, hash string, chars []string, ch chan string, hasher func(s string) string) {
	length := 1

	for {
		comb := int(math.Pow(float64(len(chars)), float64(length)))

		for i := id; i < comb; i += step {
			cand := getCandidate(i, length, chars)
			hashedCand := hasher(cand)

			if hash == hashedCand {
				ch <- cand
			}
		}

		length++
	}
}

func WorkerMD5(id, step int, hash string, ch chan string) {
	worker(id, step, hash, printable, ch, hasherMD5)
}

func WorkerSHA1(id, step int, hash string, ch chan string) {
	worker(id, step, hash, printable, ch, hasherSHA1)
}

func WorkerSHA256(id, step int, hash string, ch chan string) {
	worker(id, step, hash, printable, ch, hasherSHA256)
}

func WorkerSHA384(id, step int, hash string, ch chan string) {
	worker(id, step, hash, printable, ch, hasherSHA384)
}

func WorkerSHA512(id, step int, hash string, ch chan string) {
	worker(id, step, hash, printable, ch, hasherSha512)
}
