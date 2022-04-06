package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func logInfo(a interface{}) {
	fmt.Fprintf(color.Output, "%s: %v\n", color.New(color.FgCyan).Sprint("info"), a)
}

func logError(a interface{}) {
	fmt.Fprintf(color.Output, "%s: %v\n", color.New(color.FgRed).Sprint("chyba"), a)

	os.Exit(1)
}

func checkErr(a interface{}) {
	if a != nil {
		logError(a)
	}
}
