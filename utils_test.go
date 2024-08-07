package main

import (
	"fmt"
	"os"
	"testing"
)

func TestReadOpencdFile(t *testing.T) {
	data, err := readOpencdFile()
	if err != nil {
		t.Errorf("%s", err)
	}
	for i, envs := range data.Environments {
		fmt.Println(i, envs.Name)
	}
}

func TestParsePathFile(t *testing.T) {
	pwd, _ := os.Getwd()
	test1, err := parsePathFile("opencd.yaml")
	if err != nil {
		t.Errorf("%s is not found", test1)
	}
	test2, err := parsePathFile("./opencd.yaml")
	if err != nil {
		t.Errorf("%s is not found", test2)
	}
	test3, err := parsePathFile(pwd + "/" + "opencd.yaml")
	if err != nil {
		t.Errorf("%s is not found", test3)
	}
}

func TestCheckComponents(t *testing.T) {
	err := checkComponents()
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestCmd(t *testing.T) {
	err := cmd("ls")
	if err != nil {
		t.Errorf("%s", err)
	}
}
