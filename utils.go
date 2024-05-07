package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"gopkg.in/yaml.v3"
)

type OpenCd struct {
	Environments []Environments `yaml:"environments"`
}
type Environments struct {
	Name   string `yaml:"name"`
	Docker string `yaml:"docker"`
}

// Возвращает текущую директорию
func getCurrentDirectory() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd
}

// Проверка на наличие необходимых файлов
func checkComponents() {
	pwd := getCurrentDirectory()
	configFile := pwd + "/" + OPENCD_CONFIG
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Парсинг файла конфигурации opencd
func readOpencdFile() {
	var config OpenCd
	file, err := os.ReadFile(OPENCD_CONFIG)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}

// Парсинг пути файла (относитеьный и абсолютный путь)
func parsePathFile(filepath string) (string, error) {
	// парсинг относительного пути файла
	relativePath, _ := regexp.Compile(`^.\/`)
	idx := relativePath.FindStringIndex(filepath)
	pwd := getCurrentDirectory()
	if len(idx) > 0 {
		i := idx[1]
		return pwd + "/" + filepath[i:], nil
	} else {
		if filepath == OPENCD_CONFIG {
			return pwd + "/" + filepath, nil
		} else {
			// парсинг абсолютного пути файла
			absolutePath, _ := regexp.Compile(`^/`)
			match := absolutePath.MatchString(filepath)
			if match {
				return filepath, nil
			}
			// ошибка, если не найдет файл
			return "", errors.New("file not found: " + filepath)
		}
	}
}
