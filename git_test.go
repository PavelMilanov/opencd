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
	createDeployBranch("origin/testing")
	// fmt.Println(data)
}

func TestDeploy(t *testing.T) {
	deploy()
	// exec, err := exec.Command("bash", "-c", "git checkout -b test2 && git merge origin/testing").Output()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(exec))
	// shortSha := strings.Split(string(exec), " ")[1][:7] // commit 11e00c3b19f88ec7602c4d115871113e49f63e07 => 11e00c3
	// fmt.Println("deploy" + "-" + shortSha)
}
