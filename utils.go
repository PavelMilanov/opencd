package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"gopkg.in/yaml.v3"
)

// Возвращает текущую директорию.
func getCurrentDirectory() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return pwd, nil
}

// Проверка на наличие необходимых файлов.
func checkComponents() error {
	pwd, err := getCurrentDirectory()
	if err != nil {
		return err
	}
	configFile := pwd + "/" + OPENCD_CONFIG
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return fmt.Errorf("не найден файл %s", configFile)
	}
	config, err := readOpencdFile()
	if err != nil {
		return err
	}
	for _, data := range config.Environments {
		file := pwd + "/" + data.Docker
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return fmt.Errorf("не найден файл %s", file)
		}
	}
	return nil
}

// Парсинг файла конфигурации opencd.
func readOpencdFile() (OpenCd, error) {
	var config OpenCd
	file, err := os.ReadFile(OPENCD_CONFIG)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}
	for _, env := range config.Environments {
		if env.Name == "" {
			return config, fmt.Errorf("не указано окружение в файле конфигурации opencd.yaml")
		} else if env.Local == "" {
			return config, fmt.Errorf("не указана локальная ветка в окружении %s opencd.yaml", env.Name)
		} else if env.Remote == "" {
			return config, fmt.Errorf("не указана удаленная ветка в окружении %s opencd.yaml", env.Name)
		} else if env.Docker == "" {
			return config, fmt.Errorf("не указан файл docker-compose в окружении %s opencd.yaml", env.Name)
		}
	}

	return config, nil
}

// Парсинг пути файла (относитеьный и абсолютный путь).
func parsePathFile(filepath string) (string, error) {
	// парсинг относительного пути файла
	relativePath, _ := regexp.Compile(`^.\/`)
	idx := relativePath.FindStringIndex(filepath)
	pwd, err := getCurrentDirectory()
	if err != nil {
		return "", err
	}
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
			return "", errors.New("файл не найден: " + filepath)
		}
	}
}

// Выполняет указанную Unix-команду
func cmd(command string) error {
	run := exec.Command("bash", "-c", command)
	run.Stdout = os.Stdout
	run.Stderr = os.Stderr
	err := run.Run()
	if err != nil {
		return err
	}
	return nil
}
