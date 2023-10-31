package util
import "github.com/fatih/color"

var red = color.New(color.FgRed).PrintfFunc()
var green = color.New(color.FgGreen).PrintfFunc()

func PrintError(msg string) {
	red(msg)
}

func PrintSuccess(msg string) {
	green(msg)
}
