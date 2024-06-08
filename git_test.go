package main

import (
	"fmt"
	"testing"
)

func TestGitDiff(t *testing.T) {
	data, err := gitDiff("origin/dev", "dev2")
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Println(data)
}
func TestAnaluzeChanges(t *testing.T) {
	config, err := readOpencdFile()
	if err != nil {
		t.Errorf("%s", err)
	}
	changes, err := gitDiff(config.Environments[0].Local, config.Environments[0].Remote)
	if err != nil {
		t.Errorf("%s", err)
	}
	services, err := parseDockerCompose(config.Environments[0].Docker)
	if err != nil {
		t.Errorf("%s", err)
	}
	data := analuzeChanges(services, changes)
	fmt.Println(data)
}

func TestCreateDeployBranch(t *testing.T) {
	createDeployBranch("origin/dev")
}

func TestGitMerge(t *testing.T) {
	gitMerge("dev3", "deploy-702996a")
}
