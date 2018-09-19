package console

import (
	"fmt"

	"github.com/fatih/color"
)

/* console.Error("!@#$%^&*()") */

func Dev(info string) {
	color.Set(color.FgBlue)
	fmt.Printf("[DEV]]] %s\n", info)
	color.Unset()
}

func Info(info string) {
	fmt.Printf("[INFO]] %s\n", info)
}

func Warn(info string) {
	color.Set(color.FgYellow)
	fmt.Printf("[WARN]] %s\n", info)
	color.Unset()
}

func Error(info string) {
	color.Set(color.FgRed)
	fmt.Printf("[ERROR] %s\n", info)
	color.Unset()
}
