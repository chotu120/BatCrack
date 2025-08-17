package cli

import (
	"os"

	"github.com/fatih/color"
)

var (
	white, black, pink, pinkBg, greenBg, redBg func(a ...any) string
)

func init() {
	if os.Getenv("COLORTERM") == "truecolor" {
		white = color.RGB(255, 255, 255).SprintFunc()
		black = color.RGB(0, 0, 0).SprintFunc()
		pink = color.RGB(255, 102, 222).SprintFunc()

		pinkBg = color.BgRGB(255, 102, 222).SprintFunc()
		greenBg = color.BgRGB(77, 255, 77).SprintFunc()
		redBg = color.BgRGB(255, 79, 77).SprintFunc()
	} else {
		white = color.New(color.FgHiWhite).SprintFunc()
		black = color.New(color.FgBlack).SprintFunc()
		pink = color.New(color.FgHiMagenta).SprintFunc()

		pinkBg = color.New(color.BgHiMagenta).SprintFunc()
		greenBg = color.New(color.BgHiGreen).SprintFunc()
		redBg = color.New(color.BgHiRed).SprintFunc()
	}
}
