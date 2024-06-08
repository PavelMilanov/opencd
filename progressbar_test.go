package main

import (
	"testing"
	"time"
)

func TestBar(t *testing.T) {
	// DEPLOY_STATUS.RenderBlank()
	SUCCESSBAR.Describe("[cyan][1/3][reset] Подготовка к обновлению компонентов...")
	// fmt.Println(DEPLOY_STATUS.String())
	SUCCESSBAR.Add(10)
	// fmt.Println(DEPLOY_STATUS.String())
	time.Sleep(1 * time.Second)
	SUCCESSBAR.Describe("[cyan][5/5][reset] Обновление прошло успешно")
	// SUCCESSBAR.Add(10)
	time.Sleep(1 * time.Second)
	SUCCESSBAR.Finish()
	// DEPLOY_STATUS.Add(10)
	// time.Sleep(1 * time.Second)
	// DEPLOY_STATUS.Describe("[cyan][3/3][reset] Обновление компонентов...")
	// DEPLOY_STATUS.Add(10)
}
