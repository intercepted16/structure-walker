package main

import "github.com/charmbracelet/huh"

var (
	seeCode string
)

func promptForCode(fileNames ...string) {
	err := huh.NewSelect[string]().
		Title("Do you want to see the code?").
		Options(
			huh.NewOption("Yes", "yes"),
			huh.NewOption("No", "no"),
		).
		Value(&seeCode).Run()
	if err != nil {
		return
	}
	if seeCode == "yes" {
		for _, fileName := range fileNames {
			err = readAndPrintCode(fileName)
			if err != nil {
				return
			}
		}
	}

}
