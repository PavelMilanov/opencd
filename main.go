package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

const OPENCD_CONFIG = "opencd.yaml"

var VERSION string
var LOGFILE string // "opencd.log"

var log = logrus.New()

func init() {
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006/01/02 15:04:00",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			_, filename := path.Split(f.File)
			filename = fmt.Sprintf("[ %s:%d]", filename, f.Line)
			return "", filename
		},
	})
	file, err := os.OpenFile(LOGFILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Не удалось открыть файл логов, используется стандартный stderr")
	}
}

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "deploy":
			err := checkComponents()
			if err != nil {
				errShutdown(err)
			}
			config, err := readOpencdFile()
			if err != nil {
				errShutdown(err)
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
