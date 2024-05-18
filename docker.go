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

func ParseDockerCompose(filename string) {
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
			re := regexp.MustCompile(`context:[^\s]+`)
			build := re.FindString(chank)
			str := strings.Split(build, ":")[1]
			// fmt.Println(str[:len(str)-1])
			re2 := regexp.MustCompile(`^\w+`)
			name := re2.FindString(chank)
			// fmt.Println(name)
			services = append(services, Service{Name: name, Build: str})
		} else {
			chank := formatData[parseData[i][0]:parseData[i+1][0]] // последний кусок из списка
			re := regexp.MustCompile(`context:[^\s]+`)
			build := re.FindString(chank) // context:./backend/backend]
			str := strings.Split(build, ":")[1]
			// fmt.Println(str[:len(str)-1]) // ./backend/backend
			re2 := regexp.MustCompile(`^\w+`)
			name := re2.FindString(chank)
			// fmt.Println(name) // backend
			services = append(services, Service{Name: name, Build: str})
		}
	}
	fmt.Println(services)
}
