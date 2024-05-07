package main

import (
	"testing"
)

func TestReadOpencdFile(t *testing.T) {
	readOpencdFile()
}

func TestParsePathFile(t *testing.T) {
	test1, err := parsePathFile("opencd.yaml")
	if err != nil {
		t.Errorf("%s is not found", test1)
	}
	test2, err := parsePathFile("./opencd.yaml")
	if err != nil {
		t.Errorf("%s is not found", test2)
	}
	test3, err := parsePathFile("/home/pavel/projects/opencd/opencd.yaml")
	if err != nil {
		t.Errorf("%s is not found", test3)
	}
}
