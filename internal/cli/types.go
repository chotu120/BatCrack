package cli

type AlgoType int

const (
	AlgoUnsupported AlgoType = iota
	AlgoMD5
	AlgoSHA1
	AlgoSHA256
	AlgoSHA384
	AlgoSHA512
)

func (a AlgoType) String() string {
	switch a {
	case AlgoMD5:
		return "MD5"
	case AlgoSHA1:
		return "SHA-1"
	case AlgoSHA256:
		return "SHA-256"
	case AlgoSHA384:
		return "SHA-384"
	case AlgoSHA512:
		return "SHA-512"
	default:
		return "Unsupported"
	}
}
