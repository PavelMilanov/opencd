package main

import (
	"testing"
)

func TestBuildDockerCompose(t *testing.T) {
	services := []Service{{Name: "nginx", Build: "./nginx"}, {Name: "postgres", Build: "./postgres"}}
	buildDockerCompose(services, "docker-compose.yaml")
}

func TestUpDockerCompose(t *testing.T) {
	services := []string{"nginx", "postgres"}
	upDockerCompose(services, "docker-compose.yaml")
}
