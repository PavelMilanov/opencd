package main

import (
	"fmt"
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
	const NotUpdate string = "Already up to date.\n"

	pull, err := exec.Command("bash", "-c", "git pull").Output()
	if err != nil {
		fmt.Println("error from git pull: ", err)
		os.Exit(1)
	}
	if string(pull) == NotUpdate {
		fmt.Println("Нет обновлений в текущей ветке")
		return
	}
	fmt.Print(string(pull))
}

func deploy() {
	gitPull()
}
