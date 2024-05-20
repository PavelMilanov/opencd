package main

import (
	"reflect"
	"testing"
)

func TestGitPull(t *testing.T) {
	gitDiff("origin/dev", "dev")
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

func TestSetTmpBranch(t *testing.T) {
	createDeployBranch("origin/test")
	// fmt.Println(data)
}

func TestDeploy(t *testing.T) {
	// deploy()
	// command := fmt.Sprintf("git log %s | head  -1", "origin/test")
	// exec, err := exec.Command("bash", "-c", command).Output()
	// if err != nil {
	// 	panic(err)
	// }
	// shortSha := strings.Split(string(exec), " ")[1][:7] // commit 11e00c3b19f88ec7602c4d115871113e49f63e07 => 11e00c3
	// fmt.Println("deploy" + "-" + shortSha)
}
