package console

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

/* console.Error("!@#$%^&*()") */
func Dev(text string) {
	color.Set(color.FgCyan)
	log.Printf("[DEV]   %s\n", text)
	color.Unset()
}

func Devf(text string, arg ...interface{}) {
	color.Set(color.FgCyan)
	log.Printf("[DEV]   "+text+"\n", arg...)
	color.Unset()
}

func Info(text string) {
	log.Printf("[INFO]  %s\n", text)
}

func Infof(text string, arg ...interface{}) {
	log.Printf("[INFO]  "+text+"\n", arg...)
}

func Bingo(text string) {
	color.Set(color.FgGreen)
	log.Printf("[BINGO] %s\n", text)
	color.Unset()
}
func Bingof(text string, arg ...interface{}) {
	color.Set(color.FgGreen)
	log.Printf("[BINGO] "+text+"\n", arg...)
	color.Unset()
}

func Warn(text string) {
	color.Set(color.FgYellow)
	log.Printf("[WARN]  %s\n", text)
	color.Unset()
}

func Warnf(text string, arg ...interface{}) {
	color.Set(color.FgYellow)
	log.Printf("[WARN]  "+text+"\n", arg...)
	color.Unset()
}

func Error(text string) {
	color.Set(color.FgHiRed)
	log.Printf("[ERROR] %s\n", text)
	color.Unset()
}
func Errorf(text string, arg ...interface{}) {
	color.Set(color.FgHiRed)
	log.Printf("[ERROR] "+text+"\n", arg...)
	color.Unset()
}

func Fatal(text string) {
	color.Set(color.FgRed)
	log.Printf("[FATAL] %s\n", text)
	color.Unset()
}
func Fatalf(text string, arg ...interface{}) {
	color.Set(color.FgRed)
	log.Printf("[FATAL] "+text+"\n", arg...)
	color.Unset()
}

func Title(title string) {
	fmt.Printf("-------- %s --------\n", title)
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
