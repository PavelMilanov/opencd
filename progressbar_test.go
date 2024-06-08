package main

import (
	"testing"
	"time"
)

func TestBar(t *testing.T) {
	PROGRESSBAR.Describe("[cyan][1/3][reset] Подготовка к обновлению компонентов...")
	PROGRESSBAR.Add(10)
	time.Sleep(1 * time.Second)
	PROGRESSBAR.Describe("[cyan][2/3][reset] Обновление прошло успешно")
	PROGRESSBAR.Add(10)
	PROGRESSBAR.Describe("[cyan][3/3][reset] Обновление прошло успешно 1")
	PROGRESSBAR.Add(20)
	time.Sleep(1 * time.Second)
	errorbar(40)
}
