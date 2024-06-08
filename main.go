package main

import (
	"flag"
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
			deployCommand := flag.NewFlagSet("deploy", flag.ExitOnError)
			stage := deployCommand.String("s", "merge", "запускает обновление проекта;\nдопустимые флаги [merge, docker];\nmerge - полный цикл сборки;\ndocker - сборка и запуск контейнеров в текущем состоянии.\n")
			deployCommand.Parse(os.Args[2:])
			deploy(config.Environments[0], *stage)
		case "version":
			version()
		case "help":
			fmt.Print(MENU_TEXT)
		default:
			fmt.Println("неверная команда")
		}
	} else {
		fmt.Print("неверная команда\n", MENU_TEXT)
	}
}
