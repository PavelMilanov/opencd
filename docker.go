package main

import (
	"fmt"
	"os"
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

func formatChankData(chank string) (string, string) {
	re := regexp.MustCompile(`context:[^\s]+`)
	build := re.FindString(chank)       // context:./test/test]
	str := strings.Split(build, ":")[1] // ./test/test]
	re2 := regexp.MustCompile(`^\w+`)   // имя сервиса всегда с новой строки во фрагменте файла
	name := re2.FindString(chank)
	return name, str[:len(str)-1] // test, ./test/test
}

func parseDockerCompose(filename string) []Service {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var conf DockerCompose
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		panic(err)
	}
	data := fmt.Sprintf("%v", conf.Services)

	re := regexp.MustCompile(`map`)                           // ищем в сыром тексте map и заменяем на пустой символ
	formatData := re.ReplaceAllString(data, "")               // замена "map" => ""
	reService := regexp.MustCompile(`[a-z]+:\[build`)         // для разделение текста на сервисы
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
	return services
}
