package main

import "github.com/schollz/progressbar/v3"

var SUCCESSBAR = progressbar.NewOptions(50,
	progressbar.OptionEnableColorCodes(true),
	progressbar.OptionSetElapsedTime(true),
	progressbar.OptionShowIts(),
	progressbar.OptionSetWidth(40),
	progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[green]=[reset]",
		SaucerHead:    "[green]>[reset]",
		SaucerPadding: " ",
		BarStart:      "[",
		BarEnd:        "]",
	}))

var FAILEDBAR = progressbar.NewOptions(50,
	progressbar.OptionEnableColorCodes(true),
	progressbar.OptionSetElapsedTime(true),
	progressbar.OptionShowIts(),
	progressbar.OptionSetWidth(40),
	progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[red]=[red][reset]",
		SaucerHead:    "[red]=[red][reset]",
		SaucerPadding: " ",
		BarStart:      "[",
		BarEnd:        "]",
	}))
