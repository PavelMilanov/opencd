package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

type DockerCompose struct {
	Version  string                 `yaml:"version"`
	Services map[string]interface{} `yaml:"services"`
}

type Service struct {
	Name  string
	Build string
}

// Форматированный вывод строк.
func formatChankData(chank string) (string, string) {
	re := regexp.MustCompile(`context:[^\s]+`)
	re2 := regexp.MustCompile(`^\w+(-\w+)*`) // имя сервиса всегда с новой строки во фрагменте файла
	build := re.FindString(chank)            // context:./test/test]
	str := strings.Split(build, ":")[1]      // ./test/test]
	name := re2.FindString(chank)

	re3 := regexp.MustCompile(`^.\/[a-zA-Z]+`) // ищет ./build_dir
	dir := re3.FindString(str)
	if dir == "" { // если нет совпадений -> ищем build: .
		re4 := regexp.MustCompile(`^.`) // ищет ./build_dir
		dir2 := re4.FindString(str)
		if dir2 != "." {
			fmt.Printf("%s:\n  build:\n    context: <./app_dir> или <.>\n", name)
			os.Exit(0)
		}
		return name, dir2
	}
	return name, dir
}

// Выполняет парсинг файла docker-compose. Возврашает список сервисов, у которых собираемый локально образ.
func parseDockerCompose(filename string) ([]Service, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var conf DockerCompose
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		return nil, err
	}
	data := fmt.Sprintf("%v", conf.Services)

	re := regexp.MustCompile(`map`)                           // ищем в сыром тексте map и заменяем на пустой символ
	formatData := re.ReplaceAllString(data, "")               // замена "map" => ""
	reService := regexp.MustCompile(`\w+(-\w+)*:\[build`)     // для разделение текста на сервисы
	parseData := reService.FindAllStringIndex(formatData, -1) // [[1 15] [338 355] [675 690] [963 979]]

	services := []Service{}
	for i := 0; i < len(parseData); i++ {
		if i == len(parseData)-1 {
			chank := formatData[parseData[i][0] : len(formatData)-1]
			name, build := formatChankData(chank)
			services = append(services, Service{Name: name, Build: build})
		} else {
			chank := formatData[parseData[i][0]:parseData[i+1][0]] // последний кусок из списка
			name, build := formatChankData(chank)
			services = append(services, Service{Name: name, Build: build})
		}
	}
	return services, nil
}

// Запускает сборку сервисов из docker-compose файла, у которых изменился код.
// Возвращает список названий сервисов для сборки.
func buildDockerCompose(services []Service, composeFile string) ([]string, error) {
	serviceNameList := []string{}
	for _, s := range services {
		serviceNameList = append(serviceNameList, s.Name)
	}
	command := fmt.Sprintf("docker compose -f %s build %s", composeFile, strings.Join(serviceNameList, " "))
	run := exec.Command("bash", "-c", command)
	run.Stdout = os.Stdout
	run.Stderr = os.Stderr
	err := run.Run()
	if err != nil {
		return serviceNameList, err
	}
	return serviceNameList, nil
}

// Запускает сервисы docker-compose, переданные в параметры функции.
func upDockerCompose(services []string, composeFile string) error {
	command := fmt.Sprintf("docker compose -f %s up -d %s", composeFile, strings.Join(services, " "))
	run := exec.Command("bash", "-c", command)
	run.Stdout = os.Stdout
	run.Stderr = os.Stderr
	err := run.Run()
	if err != nil {
		return err
	}
	return nil
}
