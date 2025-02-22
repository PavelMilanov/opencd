package main

import (
	"fmt"
	"testing"
)

func TestBuildDockerCompose(t *testing.T) {
	services := []Service{{Name: "nginx", Build: "./nginx"}, {Name: "postgres", Build: "./postgres"}}
	data, err := buildDockerCompose(services, "docker-compose.yaml")
	if err != nil {
		t.Errorf("Ошибка %s", err)
	}
	fmt.Println(data)
}

func TestUpDockerCompose(t *testing.T) {
	services := []string{"nginx", "postgres"}
	err := upDockerCompose(services, "docker-compose.yaml")
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestParseDockerCompose(t *testing.T) {
	data, err := parseDockerCompose("docker-compose.yaml")
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Println(data)
}
