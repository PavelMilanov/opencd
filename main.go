package main

import (
	"fmt"
	"os"
)

const OPENCD_CONFIG = "opencd.yaml"

var VERSION string

func main() {
	err := checkComponents()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "deploy":
			config, err := readOpencdFile()
			if err != nil {
				panic(err)
			}
			deploy(config.Environments[0])
		case "rollback":
			fmt.Println("в разработке")
		case "version":
			fmt.Println("opencd version:", VERSION)
		default:
			fmt.Print("неверная команда\n", MENU_TEXT)
		}
	} else {
		fmt.Print("неверная команда\n", MENU_TEXT)
	}
}
