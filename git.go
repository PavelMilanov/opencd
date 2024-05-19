package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func displayCommits() {
	commits, err := exec.Command("bash", "-c", "git log --pretty=format:\"%h - %an, %ar : %s\"").Output()
	if err != nil {
		fmt.Println("error from display commits: ", err)
		os.Exit(1)
	}
	fmt.Println(string(commits))
}

// Проверяет изменения в ветках репозитория и возвращает директории, где были изменения
func gitPull() {
	diff, err := exec.Command("bash", "-c", "git diff origin/dev dev").Output()
	if err != nil {
		fmt.Println("error from git diff: ", err)
		os.Exit(1)
	}
	buffer := bytes.NewBuffer(diff)
	for {
		line, err := buffer.ReadString('\n')
		if err == io.EOF {
			break
		}
		// ишем строчки diff --git .....
		if strings.HasPrefix(line, "diff --git") {
			commit := strings.Split(line, " ")[2]          // diff --git a/test1/file1.txt b/test1/file1.txt => a/test1/file1.txt
			commitChange := strings.Split(commit, "a/")[1] // [ test1/file1.txt] => test1/file1.txt
			fmt.Println(commitChange)
		}
	}
}

func deploy() {
	gitPull()
}
