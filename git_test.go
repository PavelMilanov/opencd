package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGitDiff(t *testing.T) {
	data := gitDiff("origin/testing", "dev2")
	fmt.Println(data)
}

func TestAnaluzeChanges(t *testing.T) {
	commits1 := []string{"x/text.txt", "z/x/y.txt"}
	commits2 := []string{"test/text.txt"}
	services := []Service{{Name: "test", Build: "test"}, {Name: "test2", Build: "test2/x"}}
	result1 := []Service{}
	result2 := []Service{{Name: "test", Build: "test"}}
	data1 := analuzeChanges(services, commits1)
	if !reflect.DeepEqual(data1, result1) {
		t.Errorf("%s not equal to result %s", data1, result1)
	}
	data2 := analuzeChanges(services, commits2)
	if !reflect.DeepEqual(data2, result2) {
		t.Errorf("%s not equal to result %s", data2, result2)
	}
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
