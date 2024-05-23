package main

import (
	"testing"
)

func TestBuildDockerCompose(t *testing.T) {
	services := []Service{{Name: "test", Build: "./back"}}
	buildDockerCompose(services, "docker-compose.yaml")
}

func TestUpDockerCompose(t *testing.T) {
	services := []string{"test"}
	upDockerCompose(services, "docker-compose.yaml")
}
