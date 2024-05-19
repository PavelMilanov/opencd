package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestGitPull(t *testing.T) {
	gitPull("origin/dev", "dev")
}

func TestAnaluzeChanges(t *testing.T) {
	commits1 := []string{"x/text.txt", "z/x/y.txt"}
	commits2 := []string{"test/text.txt"}
	services := []Service{{Name: "test", Build: "test"}, {Name: "test2", Build: "test2/x"}}
	result1 := []Service{}
	result2 := []Service{{Name: "test", Build: "test"}}
	data1 := analuzeChanges(services, commits1)
	if reflect.DeepEqual(data1, result1) {
		t.Errorf("%s not equal to result %s", data1, result1)
	}
	data2 := analuzeChanges(services, commits2)
	if reflect.DeepEqual(data2, result2) {
		t.Errorf("%s not equal to result %s", data2, result2)
	}
}

func TestDeploy(t *testing.T) {
	// changes := gitPull("origin/dev", "dev")
	testChanges := []string{"backend/test1", "test2.txt", "web/test3"}
	services := parseDockerCompose("docker-compose.yaml")
	changeService := []Service{}
	for _, service := range services {
		match := strings.Split(service.Build, "/")[0] // название папки
		for _, change := range testChanges {          //ищем совпадения папок в docker-compose и изменениях коммита
			if strings.Contains(change, match) {
				changeService = append(changeService, Service{Name: service.Name, Build: service.Build})
			}
		}
	}
	fmt.Println(changeService, services)
}
