package main

import (
	"fmt"
	"os"

	"github.com/schollz/progressbar/v3"
)

var PROGRESSBAR = progressbar.NewOptions(100,
	progressbar.OptionEnableColorCodes(true),
	progressbar.OptionSetElapsedTime(true),
	progressbar.OptionShowIts(),
	progressbar.OptionSetWidth(20),
	progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[green]=[reset]",
		SaucerHead:    "[green]>[reset]",
		SaucerPadding: " ",
		BarStart:      "[",
		BarEnd:        "]",
	}))

var FAILE = progressbar.NewOptions(100,
	progressbar.OptionEnableColorCodes(true),
	progressbar.OptionSetElapsedTime(true),
	progressbar.OptionShowIts(),
	progressbar.OptionSetWidth(20),
	progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[red]=[red][reset]",
		SaucerHead:    "[red]>[red][reset]",
		SaucerPadding: " ",
		BarStart:      "[",
		BarEnd:        "]",
	}))

// изменяет тему прогрессбара для вывода ошибок.
// Значение state должно быть кратйно 10.
func errorbar(state int) {
	idx := state / 10
	text := fmt.Sprintf("[cyan][%d/5][reset] Ошибка", idx)
	PROGRESSBAR = FAILE
	PROGRESSBAR.Describe(text)
	PROGRESSBAR.Add(state)
	PROGRESSBAR.Exit()
	os.Exit(1)
}
