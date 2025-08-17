package cli

import "encoding/hex"

func isHex(hash string) bool {
	_, err := hex.DecodeString(hash)
	return hash != "" && err == nil
}

func validateMD5(hash string) bool {
	return len(hash) == 32 && isHex(hash)
}

func validateSHA1(hash string) bool {
	return len(hash) == 40 && isHex(hash)
}

func validateSHA256(hash string) bool {
	return len(hash) == 64 && isHex(hash)
}

func validateSHA384(hash string) bool {
	return len(hash) == 96 && isHex(hash)
}

func validateSHA512(hash string) bool {
	return len(hash) == 128 && isHex(hash)
}
