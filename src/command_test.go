package main

import "testing"

func TestDeploy(t *testing.T) {
	config, err := readOpencdFile()
	if err != nil {
		panic(err)
	}
	deploy(config.Environments[0], config.Settings, "merge")
}

func TestDockerPrune(t *testing.T) {
	dockerPrnune()
}
