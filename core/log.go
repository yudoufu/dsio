package core

import (
	"fmt"

	"github.com/fatih/color"
)

func Conform(message interface{}) {
	fmt.Printf("%v %v", color.CyanString("[CONFIRM]"), message)
}

func Conformf(format string, value ...interface{}) {
	fmt.Printf("%v ", color.CyanString("[CONFIRM]"))
	fmt.Printf(format, value...)
}
func Info(message interface{}) {
	fmt.Printf("%v %v\n", color.GreenString("[INFO]"), message)
}

func Infof(format string, value ...interface{}) {
	fmt.Printf("%v ", color.GreenString("[INFO]"))
	fmt.Printf(format, value...)
}

func Debug(message interface{}) {
	if ctx.Verbose {
		fmt.Printf("%v %v\n", color.CyanString("[DEBUG]"), message)
	}
}

func Debugf(format string, value ...interface{}) {
	if ctx.Verbose {
		fmt.Printf("%v ", color.CyanString("[DEBUG]"))
		fmt.Printf(format, value...)
	}
}

func Error(message interface{}) {
	color.Red("[ERROR] %v", message)
	fmt.Println("")
	fmt.Errorf("%v", message)
}

func Errorf(format string, value ...interface{}) {
	color.Red("[ERROR] "+format, value...)
	fmt.Println("")
	fmt.Errorf(format, value...)
}

func Panicf(format string, value ...interface{}) string {
	return color.RedString("[PANIC] "+format, value...)
}
