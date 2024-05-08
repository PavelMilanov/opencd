package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func displayCommits() {
	commits, err := exec.Command("bash", "-c", "git log --pretty=format:\"%h - %an, %ar : %s\"").Output()
	if err != nil {
		fmt.Println("error from display commits: ", err)
		os.Exit(1)
	}
	fmt.Println(string(commits))
}

func gitPull() {
	diff, err := exec.Command("bash", "-c", "git diff origin/dev dev").Output()
	if err != nil {
		fmt.Println("error from git diff: ", err)
		os.Exit(1)
	}
	// fmt.Print(string(diff))
	buffer := bytes.NewBuffer(diff)
	for {
		n, err := buffer.ReadString('\n')
		fmt.Print(n)
		// fmt.Println(n, err, buf[:n])
		if err == io.EOF {
			break
		}
	}
	// fmt.Println(reader)
	// r, _ := regexp.Compile(`a/.+$`)
	// fmt.Println(r.FindAllString(string(diff), -1))
}

func deploy() {
	gitPull()
}
