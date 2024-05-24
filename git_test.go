package main

import (
	"fmt"
	"testing"
)

func TestGitDiff(t *testing.T) {
	data := gitDiff("origin/dev", "dev2")
	fmt.Println(data)
}

func TestAnaluzeChanges(t *testing.T) {
	config, err := readOpencdFile()
	if err != nil {
		panic(err)
	}
	changes := gitDiff(config.Environments[0].Local, config.Environments[0].Remote)
	services := parseDockerCompose(config.Environments[0].Docker)
	data := analuzeChanges(services, changes)
	fmt.Println(data)
}

func TestCreateDeployBranch(t *testing.T) {
	createDeployBranch("origin/dev")
}

func TestGitMerge(t *testing.T) {
	gitMerge("dev3", "deploy-702996a")
}

func TestDeploy(t *testing.T) {
	config, err := readOpencdFile()
	if err != nil {
		panic(err)
	}
	deploy(config.Environments[0])
}
