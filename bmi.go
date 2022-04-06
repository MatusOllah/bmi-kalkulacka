package main

import "github.com/fatih/color"

func bmi(hmotnost, vyska float64) (float64, string) {
	var _type string

	bmi := hmotnost / (vyska * vyska)

	switch {
	case bmi <= 18.5:
		_type = color.New(color.BgCyan, color.FgWhite, color.Bold).Sprint("Podváha")
	case bmi > 18.5 && bmi < 24.9:
		_type = color.New(color.BgGreen, color.FgWhite, color.Bold).Sprint("Normálna hmotnosť")
	case bmi > 25 && bmi < 29.9:
		_type = color.New(color.BgYellow, color.FgWhite, color.Bold).Sprint("Nadváha")
	case bmi >= 30:
		_type = color.New(color.BgRed, color.FgWhite, color.Bold).Sprint("Obezita")
	}

	return bmi, _type
}
