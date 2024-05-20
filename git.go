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
func gitDiff(localBranch string, remoteBranch string) []string {
	command := fmt.Sprintf("git diff %s %s", remoteBranch, localBranch)
	diff, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		fmt.Println("error from git diff: ", err)
		os.Exit(1)
	}
	changes := []string{}
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
			changes = append(changes, commitChange)
		}
	}
	return changes
}

func analuzeChanges(services []Service, commits []string) []Service {
	changeService := []Service{}
	for _, service := range services {
		match := strings.Split(service.Build, "/")[0] // название папки
		for _, change := range commits {              //ищем совпадения папок в docker-compose и изменениях коммита
			if strings.Contains(change, match) {
				changeService = append(changeService, Service{Name: service.Name, Build: service.Build})
			}
		}
	}
	return changeService
}

func deploy() {
	changes := gitDiff("origin/test", "dev")
	services := parseDockerCompose("docker-compose.yaml")
	updateServices := analuzeChanges(services, changes)
	fmt.Println(updateServices)
}
