package cli

import "testing"

type detectionTest struct {
	input    string
	expected AlgoType
}

func TestDetectHashAlgo(t *testing.T) {
	tests := []detectionTest{
		// MD5
		{"d41d8cd98f00b204e9800998ecf8427e", AlgoMD5},
		{"5d41402abc4b2a76b9719d911017c592", AlgoMD5},
		{"d41d8cd98f00b204e9800998ecf8427", AlgoUnsupported},
		{"z41d8cd98f00b204e9800998ecf8427e", AlgoUnsupported},

		// SHA-1
		{"2aae6c35c94fcfb415dbe95f408b9ce91ee846ed", AlgoSHA1},
		{"2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED", AlgoSHA1},
		{"2aae6c35c94fcfb415dbe95f408b9ce91ee846e", AlgoUnsupported},
		{"g2aae6c35c94fcfb415dbe95f408b9ce91ee846ed", AlgoUnsupported},

		// SHA-256
		{"9d5ed678fe57bcca610140957afab571ca8c6c3a8f8a9f4b6a9a7a8c4c1d4ca8", AlgoSHA256},
		{"9D5ED678FE57BCCA610140957AFAB571CA8C6C3A8F8A9F4B6A9A7A8C4C1D4CA8", AlgoSHA256},
		{"9d5ed678fe57bcca610140957afab571ca8c6c3a8f8a9f4b6a9a7a8c4c1d4ca", AlgoUnsupported},
		{"z9d5ed678fe57bcca610140957afab571ca8c6c3a8f8a9f4b6a9a7a8c4c1d4ca8c", AlgoUnsupported},

		// SHA-384
		{"3391fdddfc8dc7393707a65b1b4709397cf8b1d162af05abfe8f450de5f36bc6d166db4bb6459f9ed48bf0cbbd0a6990", AlgoSHA384},
		{"3391FDDDFC8DC7393707A65B1B4709397CF8B1D162AF05ABFE8F450DE5F36BC6D166DB4BB6459F9ED48BF0CBBD0A6990", AlgoSHA384},
		{"3391fdddfc8dc7393707a65b1b4709397cf8b1d162af05abfe8f450de5f36bc6d166db4bb6459f9ed48bf0cbbd0a699", AlgoUnsupported},
		{"z391fdddfc8dc7393707a65b1b4709397cf8b1d162af05abfe8f450de5f36bc6d166db4bb6459f9ed48bf0cbbd0a6990f", AlgoUnsupported},

		// SHA-512
		{"cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", AlgoSHA512},
		{"CF83E1357EEFB8BDF1542850D66D8007D620E4050B5715DC83F4A921D36CE9CE47D0D13C5D85F2B0FF8318D2877EEC2F63B931BD47417A81A538327AF927DA3E", AlgoSHA512},
		{"cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3", AlgoUnsupported},
		{"zf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3ef", AlgoUnsupported},
	}

	for i, test := range tests {
		if ok := detectHashAlgo(test.input); ok != test.expected {
			t.Errorf("Test %d failed: expected %s, got %s for input '%s'", i+1, test.expected.String(), ok.String(), test.input)
		}
	}
}
