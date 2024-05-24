package main

import "github.com/schollz/progressbar/v3"

var STAGE1 = progressbar.NewOptions(-1,
	progressbar.OptionEnableColorCodes(true),
	progressbar.OptionSetElapsedTime(false),
	progressbar.OptionSetWidth(20),
	progressbar.OptionSetDescription("[cyan][1/4][reset] Анализ удаленного репозитория...\n"),
	progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[green]=[reset]",
		SaucerHead:    "[green]>[reset]",
		SaucerPadding: " ",
		BarStart:      "[",
		BarEnd:        "]",
	}))

var STAGE2 = progressbar.NewOptions(-1,
	progressbar.OptionEnableColorCodes(true),
	progressbar.OptionSetElapsedTime(false),
	progressbar.OptionSetWidth(20),
	progressbar.OptionSetDescription("[cyan][2/4][reset] Анализ конфигурации docker...\n"),
	progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[green]=[reset]",
		SaucerHead:    "[green]>[reset]",
		SaucerPadding: " ",
		BarStart:      "[",
		BarEnd:        "]",
	}))

var STAGE3 = progressbar.NewOptions(-1,
	progressbar.OptionEnableColorCodes(true),
	progressbar.OptionSetElapsedTime(false),
	progressbar.OptionSetWidth(20),
	progressbar.OptionSetDescription("[cyan][3/4][reset] Слияние веток...\n"),
	progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[green]=[reset]",
		SaucerHead:    "[green]>[reset]",
		SaucerPadding: " ",
		BarStart:      "[",
		BarEnd:        "]",
	}))

var STAGE4 = progressbar.NewOptions(-1,
	progressbar.OptionEnableColorCodes(true),
	progressbar.OptionSetElapsedTime(false),
	progressbar.OptionSetWidth(20),
	progressbar.OptionSetDescription("[cyan][4/4][reset] Сборка и запуск новой версии...\n"),
	progressbar.OptionSetTheme(progressbar.Theme{
		Saucer:        "[green]=[reset]",
		SaucerHead:    "[green]>[reset]",
		SaucerPadding: " ",
		BarStart:      "[",
		BarEnd:        "]",
	}))
