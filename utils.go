package main

import (
	"fmt"
	"os"
)

// Возвращает текущую директорию
func getCurrentDirectory() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

func checkComponents() {
	pwd := getCurrentDirectory()
	configFile := pwd + "/" + OPENCD_CONFIG
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(1)
	}
}
