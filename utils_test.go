package main

import (
	"os"
	"testing"
)

func TestReadOpencdFile(t *testing.T) {
	readOpencdFile()
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
