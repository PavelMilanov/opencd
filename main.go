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
				fmt.Println(err)
				os.Exit(1)
			}
			deployCommand := flag.NewFlagSet("deploy", flag.ExitOnError)
			stage := deployCommand.String("s", "merge", "запускает обновление проекта;\nдопустимые флаги [merge, docker];\nmerge - полный цикл сборки;\ndocker - сборка и запуск контейнеров в текущем состоянии.\n")
			env := deployCommand.String("e", "", "название окружения для обновления проекта;\nпараметр opencd.name.")
			deployCommand.Parse(os.Args[2:])
			for _, item := range config.Environments {
				if item.Name == *env {
					deploy(item, config.Settings, *stage)
					os.Exit(0)
				}
			}
			fmt.Println("не указано окружение для обновления проекта. Подробнее - opencd help")
			os.Exit(1)
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
