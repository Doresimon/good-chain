package console

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

/* console.Error("!@#$%^&*()") */
func Dev(info string) {
	color.Set(color.FgCyan)
	log.Printf("[DEV]   %s\n", info)
	color.Unset()
}

func Info(info string) {
	log.Printf("[INFO]  %s\n", info)
}

func Bingo(info string) {
	color.Set(color.FgGreen)
	log.Printf("[BINGO] %s\n", info)
	color.Unset()
}

func Warn(info string) {
	color.Set(color.FgYellow)
	log.Printf("[WARN]  %s\n", info)
	color.Unset()
}

func Error(info string) {
	color.Set(color.FgHiRed)
	log.Printf("[ERROR] %s\n", info)
	color.Unset()
}

func Fatal(info string) {
	color.Set(color.FgRed)
	log.Printf("[FATAL] %s\n", info)
	color.Unset()
}

func Title(title string) {
	fmt.Printf("******** %s ********\n", title)
}

// ShowColors prints all color types of console package
func ShowColors() {
	Title("Colors of Console")
	Dev("Dev")
	Info("Info")
	Bingo("Bingo")
	Warn("Warn")
	Error("Error")
	Fatal("Fatal\n")
}
