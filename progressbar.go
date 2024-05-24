package main

import "github.com/schollz/progressbar/v3"

var PULL_UPDATE = progressbar.NewOptions(-1,
	progressbar.OptionEnableColorCodes(true),
	progressbar.OptionSetElapsedTime(false),
	progressbar.OptionSetWidth(20),
	progressbar.OptionSetDescription("[cyan][1/3][reset] Подготовка к обновлению компонентов...\n"),
	progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[green]=[reset]",
		SaucerHead:    "[green]>[reset]",
		SaucerPadding: " ",
		BarStart:      "[",
		BarEnd:        "]",
	}))

// var STAGE2 = progressbar.NewOptions(-1,
// 	progressbar.OptionEnableColorCodes(true),
// 	progressbar.OptionSetElapsedTime(false),
// 	progressbar.OptionSetWidth(20),
// 	progressbar.OptionSetDescription("[cyan][2/4][reset] Анализ конфигурации docker...\n"),
// 	progressbar.OptionSetTheme(progressbar.Theme{
// 		Saucer:        "[green]=[reset]",
// 		SaucerHead:    "[green]>[reset]",
// 		SaucerPadding: " ",
// 		BarStart:      "[",
// 		BarEnd:        "]",
// 	}))

var MERGE_UPDATE = progressbar.NewOptions(-1,
	progressbar.OptionEnableColorCodes(true),
	progressbar.OptionSetElapsedTime(false),
	progressbar.OptionSetWidth(20),
	progressbar.OptionSetDescription("[cyan][2/3][reset] Слияние веток...\n"),
	progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[green]=[reset]",
		SaucerHead:    "[green]>[reset]",
		SaucerPadding: " ",
		BarStart:      "[",
		BarEnd:        "]",
	}))

var BUILD_UPDATE = progressbar.NewOptions(-1,
	progressbar.OptionEnableColorCodes(true),
	progressbar.OptionSetElapsedTime(false),
	progressbar.OptionSetWidth(20),
	progressbar.OptionSetDescription("[cyan][3/3][reset] Обновление компонентов...\n"),
	progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[green]=[reset]",
		SaucerHead:    "[green]>[reset]",
		SaucerPadding: " ",
		BarStart:      "[",
		BarEnd:        "]",
	}))
