package cli

func detectHashAlgo(hash string) AlgoType {
	if validateMD5(hash) {
		return AlgoMD5
	}

	if validateSHA1(hash) {
		return AlgoSHA1
	}

	if validateSHA256(hash) {
		return AlgoSHA256
	}

	if validateSHA384(hash) {
		return AlgoSHA384
	}

	if validateSHA512(hash) {
		return AlgoSHA512
	}

	return AlgoUnsupported
}
