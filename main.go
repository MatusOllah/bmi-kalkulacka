package main

import (
	"fmt"
	"runtime"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

func main() {
	fmt.Fprintf(color.Output, "%s %s\n", color.New(color.FgWhite, color.Bold).Sprint("== bmi-kalkulacka verzia "), color.New(color.FgBlue, color.Bold).Sprint("1.0"))
	fmt.Fprintf(color.Output, "%s %s\n", color.New(color.FgWhite, color.Bold).Sprint("== Go verzia "), color.New(color.FgBlue, color.Bold).Sprint(runtime.Version()))
	fmt.Println()

	var hmotnost float64

	survey.AskOne(&survey.Input{
		Message: "Hmotnosť (kg):",
	}, &hmotnost)

	var vyska float64

	survey.AskOne(&survey.Input{
		Message: "Výška (cm):",
	}, &vyska)

	bmi, _type := bmi(hmotnost, vyska)

	logInfo(fmt.Sprintf("BMI: %v %s", bmi, _type))

	fmt.Scanln()
}
