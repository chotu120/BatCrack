package cracker

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

func hasherMD5(s string) string {
	h := md5.New()
	io.WriteString(h, s)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func hasherSHA1(s string) string {
	h := sha1.New()
	io.WriteString(h, s)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func hasherSHA256(s string) string {
	h := sha256.New()
	io.WriteString(h, s)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func hasherSHA384(s string) string {
	h := sha512.New384()
	io.WriteString(h, s)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func hasherSha512(s string) string {
	h := sha512.New()
	io.WriteString(h, s)

	return fmt.Sprintf("%x", h.Sum(nil))
}
