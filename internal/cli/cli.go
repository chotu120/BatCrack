package cli

import (
	"batcrack/internal/cracker"
	"runtime"
)

func Run() {
	printLogo()
	flags := parseFlags()

	runtime.GOMAXPROCS(flags.Threads)
	runtime.LockOSThread()

	ch := make(chan string)

	for i := range flags.Threads {
		switch flags.Algo {
		case AlgoMD5:
			go cracker.WorkerMD5(i, flags.Threads, flags.Hash, ch)
		case AlgoSHA1:
			go cracker.WorkerSHA1(i, flags.Threads, flags.Hash, ch)
		case AlgoSHA256:
			go cracker.WorkerSHA256(i, flags.Threads, flags.Hash, ch)
		case AlgoSHA384:
			go cracker.WorkerSHA384(i, flags.Threads, flags.Hash, ch)
		case AlgoSHA512:
			go cracker.WorkerSHA512(i, flags.Threads, flags.Hash, ch)
		}
	}

	logInfo("Algorithm:", flags.Algo.String())
	logInfo("Cracking...")

	result := <-ch

	logSuccess("Cracked:", result)
}
