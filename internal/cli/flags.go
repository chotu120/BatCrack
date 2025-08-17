package cli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Flags struct {
	Hash    string
	Algo    AlgoType
	Threads int
}

func printUsage() {
	fmt.Printf("%s %s %s\n\n",
		pink("Usage:"),
		white(filepath.Base(os.Args[0])),
		white("[options]"),
	)
	fmt.Println(pink("Options:"))
	fmt.Println(white("  -h, --help                Show this help message and exit"))
	fmt.Printf("  %s %s         %s%s%s\n",
		white("-H, --hash"),
		pink("string"),
		white("Hash value to crack ("),
		pink("required"),
		white(")"),
	)
	fmt.Printf("  %s %s    %s%s%s\n",
		white("-a, --algorithm"),
		pink("string"),
		white("Hash algorithm to use ("),
		pink("leave empty for auto-detect"),
		white(")"),
	)
	fmt.Printf("  %s %s         %s%s%s\n\n",
		white("-t, --threads"),
		pink("int"),
		white("Number of CPU threads to use ("),
		pink("default: 1"),
		white(")"),
	)
	fmt.Println(pink("Supported algorithms:"))
	fmt.Println(white(`  md5
  sha1
  sha256
  sha384
  sha512`))
}

func parseFlags() Flags {
	var flags Flags
	var algo string

	flag.StringVar(&flags.Hash, "hash", "", "Hash value to crack (required)")
	flag.StringVar(&flags.Hash, "H", "", "Hash value to crack (required)")
	flag.StringVar(&algo, "algorithm", "", "Hash algorithm to use (leave empty for auto-detect)")
	flag.StringVar(&algo, "a", "", "Hash algorithm to use (leave empty for auto-detect)")
	flag.IntVar(&flags.Threads, "threads", 1, "Number of CPU threads to use (default: 1)")
	flag.IntVar(&flags.Threads, "t", 1, "Number of CPU threads to use (default: 1)")

	flag.Usage = printUsage
	flag.Parse()

	if flags.Hash == "" {
		logError("Hash value is required.")
		os.Exit(1)
	}

	switch strings.ToLower(algo) {
	case "":
		flags.Algo = detectHashAlgo(flags.Hash)
	case "md5":
		flags.Algo = AlgoMD5
	case "sha1", "sha-1":
		flags.Algo = AlgoSHA1
	case "sha256", "sha-256":
		flags.Algo = AlgoSHA256
	case "sha384", "sha-384":
		flags.Algo = AlgoSHA384
	case "sha512", "sha-512":
		flags.Algo = AlgoSHA512
	default:
		flags.Algo = AlgoUnsupported
	}

	if flags.Algo == AlgoUnsupported {
		logError("Unsupported algorithm.")
		os.Exit(1)
	}

	return flags
}
