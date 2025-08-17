package cli

import (
	"testing"
)

type validationTest struct {
	input    string
	expected bool
}

func TestIsHex(t *testing.T) {
	tests := []validationTest{
		{"abcdef123456", true},
		{"ABCDEF123456", true},
		{"1234567890", true},
		{"deadbeef", true},
		{"00ff00", true},
		{"", false},
		{"g12345", false},
		{"xyz", false},
		{"12345z", false},
		{"abc def", false},
		{"0x123abc", false},
		{"123\nabc", false},
		{"12-34", false},
		{"1234_5678", false},
	}

	for i, test := range tests {
		if ok := isHex(test.input); ok != test.expected {
			t.Errorf("Test %d failed: expected %t, got %t for input '%s'", i+1, test.expected, ok, test.input)
		}
	}
}

func TestValidateMD5(t *testing.T) {
	tests := []validationTest{
		{"d41d8cd98f00b204e9800998ecf8427e", true},
		{"5d41402abc4b2a76b9719d911017c592", true},
		{"D41D8CD98F00B204E9800998ECF8427E", true},
		{"5d41402abc4b2a76b9719d911017c59", false},
		{"5d41402abc4b2a76b9719d911017c592a", false},
		{"5d41402abc4b2a76b9719d911017c59g", false},
		{"", false},
		{"zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz", false},
	}

	for i, test := range tests {
		if ok := validateMD5(test.input); ok != test.expected {
			t.Errorf("Test %d failed: expected %t, got %t for input '%s'", i+1, test.expected, ok, test.input)
		}
	}
}

func TestValidateSHA1(t *testing.T) {
	tests := []validationTest{
		{"2aae6c35c94fcfb415dbe95f408b9ce91ee846ed", true},
		{"2AAE6C35C94FCFB415DBE95F408B9CE91EE846ED", true},
		{"2aae6c35c94fcfb415dbe95f408b9ce91ee846e", false},
		{"2aae6c35c94fcfb415dbe95f408b9ce91ee846edd", false},
		{"2aae6c35c94fcfb415dbe95f408b9ce91ee846eg", false},
		{"", false},
		{"1234567890123456789012345678901234567890", true},
		{"xyzxyzxyzxyzxyzxyzxyzxyzxyzxyzxyzxyzxyzx", false},
	}

	for i, test := range tests {
		if ok := validateSHA1(test.input); ok != test.expected {
			t.Errorf("Test %d failed: expected %t, got %t for input '%s'", i+1, test.expected, ok, test.input)
		}
	}
}

func TestValidateSHA256(t *testing.T) {
	tests := []validationTest{
		{"9d5ed678fe57bcca610140957afab571ca8c6c3a8f8a9f4b6a9a7a8c4c1d4ca8", true},
		{"9D5ED678FE57BCCA610140957AFAB571CA8C6C3A8F8A9F4B6A9A7A8C4C1D4CA8", true},
		{"9d5ed678fe57bcca610140957afab571ca8c6c3a8f8a9f4b6a9a7a8c4c1d4ca", false},
		{"9d5ed678fe57bcca610140957afab571ca8c6c3a8f8a9f4b6a9a7a8c4c1d4ca8c1", false},
		{"9d5ed678fe57bcca610140957afab571ca8c6c3a8f8a9f4b6a9a7a8c4c1d4ca8g", false},
		{"", false},
		{"abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890", true},
		{"zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz", false},
	}

	for i, test := range tests {
		if ok := validateSHA256(test.input); ok != test.expected {
			t.Errorf("Test %d failed: expected %t, got %t for input '%s'", i+1, test.expected, ok, test.input)
		}
	}
}

func TestValidateSHA384(t *testing.T) {
	tests := []validationTest{
		{"3391fdddfc8dc7393707a65b1b4709397cf8b1d162af05abfe8f450de5f36bc6d166db4bb6459f9ed48bf0cbbd0a6990", true},
		{"3391FDDDFC8DC7393707A65B1B4709397CF8B1D162AF05ABFE8F450DE5F36BC6D166DB4BB6459F9ED48BF0CBBD0A6990", true},
		{"3391fdddfc8dc7393707a65b1b4709397cf8b1d162af05abfe8f450de5f36bc6d166db4bb6459f9ed48bf0cbbd0a699", false},
		{"3391fdddfc8dc7393707a65b1b4709397cf8b1d162af05abfe8f450de5f36bc6d166db4bb6459f9ed48bf0cbbd0a6990g", false},
		{"", false},
		{"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", true},
		{"0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef", true},
		{"0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdeg0123456789abcdef0123456789abcdef", false},
	}

	for i, test := range tests {
		if ok := validateSHA384(test.input); ok != test.expected {
			t.Errorf("Test %d failed: expected %t, got %t for input '%s'", i+1, test.expected, ok, test.input)
		}
	}
}

func TestValidateSHA512(t *testing.T) {
	tests := []validationTest{
		{"cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e", true},
		{"CF83E1357EEFB8BDF1542850D66D8007D620E4050B5715DC83F4A921D36CE9CE47D0D13C5D85F2B0FF8318D2877EEC2F63B931BD47417A81A538327AF927DA3E", true},
		{"cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3", false},
		{"cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3eg", false},
		{"", false},
		{"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", true},
		{"0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef", true},
		{"0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdeg", false},
	}

	for i, test := range tests {
		if ok := validateSHA512(test.input); ok != test.expected {
			t.Errorf("Test %d failed: expected %t, got %t for input '%s'", i+1, test.expected, ok, test.input)
		}
	}
}
