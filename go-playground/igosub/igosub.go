package igosub

import "github.com/fatih/color"

func PrintColor(str string) {
	if len(str) == 0 {
		return
	}
	color.Red(str)
}
