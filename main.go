package main

import (
	"fmt"
	"os"
)

const OPENCD_CONFIG = "opencd.yaml"

var VERSION string

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "deploy":
			err := checkComponents()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			config, err := readOpencdFile()
			if err != nil {
				panic(err)
			}
			deploy(config.Environments[0])
		case "rollback":
			err := checkComponents()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			return
		case "version":
			fmt.Println("opencd version:", VERSION)
		case "help":
			fmt.Print(MENU_TEXT)
		default:
			fmt.Println("неверная команда")
		}
	} else {
		fmt.Print("неверная команда\n", MENU_TEXT)
	}
}
