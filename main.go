package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

func main() {
	hmotnost := *flag.Float64("h", 0, "Hmotnosť (v kg)")
	vyska := *flag.Float64("v", 0, "Výška (v cm)")

	flag.Parse()

	fmt.Fprintf(color.Output, "%s %s\n", color.New(color.FgWhite, color.Bold).Sprint("== bmi-kalkulacka verzia "), color.New(color.FgBlue, color.Bold).Sprint("1.0"))
	fmt.Fprintf(color.Output, "%s %s\n", color.New(color.FgWhite, color.Bold).Sprint("== Go verzia "), color.New(color.FgBlue, color.Bold).Sprint(runtime.Version()))
	fmt.Println()

	if flag.NFlag() != 0 || flag.NFlag() != 2 {
		logError("použitie: bmi-kalkulacka -h [hmotnosť v kg] -v [výška v cm]")
	}

	if flag.NFlag() == 2 {
		bmi, _type := bmi(hmotnost, vyska)

		logInfo(fmt.Sprintf("BMI: %v %s", bmi, _type))
	}

	if flag.NFlag() == 0 {
		var _hmotnost float64
		var _vyska float64
		survey.AskOne(&survey.Input{
			Message: "Hmotnosť",
		}, &_hmotnost)
		survey.AskOne(&survey.Input{
			Message: "Výška",
		}, &_vyska)

		bmi, _type := bmi(hmotnost, vyska)

		logInfo(fmt.Sprintf("BMI: %v %s", bmi, _type))
	}
}
