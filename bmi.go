package main

import (
	"math"

	"github.com/fatih/color"
)

func bmi(hmotnost, vyska float64) (float64, string) {
	var _type string

	vyska /= 100

	bmi := hmotnost / math.Pow(vyska, 2)

	switch {
	//case bmi < 18.5:
	//	_type = color.New(color.BgCyan, color.FgWhite, color.Bold).Sprint("Podváha")
	case bmi < 16:
		_type = color.New(color.BgCyan, color.FgWhite, color.Bold).Sprint("Ťažká Podváha")
	case bmi > 16 && bmi < 16.99:
		_type = color.New(color.BgCyan, color.FgWhite, color.Bold).Sprint("Stredná Podváha")
	case bmi > 17 && bmi < 18.49:
		_type = color.New(color.BgCyan, color.FgWhite, color.Bold).Sprint("Mierna Podváha")

	case bmi > 18.5 && bmi < 24.99:
		_type = color.New(color.BgGreen, color.FgWhite, color.Bold).Sprint("Normálna hmotnosť")

	//case bmi >= 25:
	//	_type = color.New(color.BgYellow, color.FgWhite, color.Bold).Sprint("Nadváha")
	case bmi > 25 && bmi < 29.99:
		_type = color.New(color.BgYellow, color.FgWhite, color.Bold).Sprint("Mierna Nadváha")

	//case bmi >= 30:
	//	_type = color.New(color.BgRed, color.FgWhite, color.Bold).Sprint("Obezita")
	case bmi > 30 && bmi < 34.99:
		_type = color.New(color.BgRed, color.FgWhite, color.Bold).Sprint("Obezita 1. stupňa")
	case bmi > 35 && bmi < 39.99:
		_type = color.New(color.BgRed, color.FgWhite, color.Bold).Sprint("Obezita 2. stupňa")
	case bmi >= 40:
		_type = color.New(color.BgRed, color.FgWhite, color.Bold).Sprint("Obezita 3. stupňa")
	}

	return bmi, _type
}
