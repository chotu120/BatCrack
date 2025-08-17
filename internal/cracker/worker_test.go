package cracker

import (
	"testing"
)

func TestWorker(t *testing.T) {
	ch := make(chan string, 1)
	go worker(0, 2, "e2fc714c4727ee9395f324cd2e7f331f", testChars, ch, hasherMD5)
	go worker(1, 2, "e2fc714c4727ee9395f324cd2e7f331f", testChars, ch, hasherMD5)

	result := <-ch

	if result != "abcd" {
		t.Errorf("Test failed: expected abcd, got %s for input 'e2fc714c4727ee9395f324cd2e7f331f'", result)
	}
}
