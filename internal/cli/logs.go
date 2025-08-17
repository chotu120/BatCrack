package cli

import (
	"fmt"
	"time"
)

func logInfo(msg string, feat ...string) {
	now := time.Now()

	if len(feat) > 0 {
		fmt.Printf("%s %s %s %s\n",
			white(now.Format("15:04:05")),
			black(pinkBg(" INFO ")),
			white(msg),
			pink(feat[0]),
		)

		return
	}

	fmt.Printf("%s %s %s\n",
		white(now.Format("15:04:05")),
		black(pinkBg(" INFO ")),
		white(msg),
	)
}

func logSuccess(msg string, feat ...string) {
	now := time.Now()

	if len(feat) > 0 {
		fmt.Printf("%s %s %s %s\n",
			white(now.Format("15:04:05")),
			black(greenBg(" SUCC ")),
			white(msg),
			pink(feat[0]),
		)

		return
	}

	fmt.Printf("%s %s %s\n",
		white(now.Format("15:04:05")),
		black(greenBg(" SUCC ")),
		white(msg),
	)
}

func logError(msg string) {
	now := time.Now()

	fmt.Printf("%s %s %s\n",
		white(now.Format("15:04:05")),
		black(redBg(" ERRO ")),
		white(msg),
	)
}
