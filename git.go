package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// Выводит список всех коммитов
func displayCommits() {
	commits, err := exec.Command("bash", "-c", "git log --pretty=format:\"%h - %an, %ar : %s\"").Output()
	if err != nil {
		fmt.Println("error from display commits: ", err)
		os.Exit(1)
	}
	fmt.Println(string(commits))
}

// Выполняет git fetch для удаленного репозитория
func gitFetch() {
	// command := fmt.Sprintf("git remote | git fetch")
	fetch, err := exec.Command("bash", "-c", "git remote | git fetch").Output()
	if err != nil {
		fmt.Println("error from git diff: ", err)
		panic(err)
	}
	fmt.Println(string(fetch))
}

// Проверяет изменения в ветках репозитория и возвращает директории, где были изменения
func gitDiff(localBranch string, remoteBranch string) []string {
	command := fmt.Sprintf("git diff %s %s", remoteBranch, localBranch)
	diff, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		fmt.Println("error from git diff: ", err)
		panic(err)
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

// Возвращает список сервисов docker-compose, для которых былм изменения в полученных коммитах
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

// Создает отдельную ветку git для применения изменений, возвращает название этой ветки
func createDeployBranch(remoteBranch string) string {
	command := fmt.Sprintf("git log %s | head  -1", remoteBranch)
	run1, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		fmt.Println("Ошибка при работе с удаленной веткой")
		panic(err)
	}
	shortSha := strings.Split(string(run1), " ")[1][:7] // commit 11e00c3b19f88ec7602c4d115871113e49f63e07 => 11e00c3
	deployBranch := "deploy" + "-" + shortSha
	command2 := fmt.Sprintf("git checkout -b %s && git merge %s", deployBranch, remoteBranch) // переключение на ветку деплоя и применение изменений
	run2, err := exec.Command("bash", "-c", command2).Output()
	if err != nil {
		fmt.Println("Ошибка при создании ветки для деплоя")
		panic(err)
	}
	fmt.Println(string(run2)) // тут нужно переключение на рабочую ветку
	return string(run2)
}

func deleteDeployBranch(branch string) {
	command := fmt.Sprintf("git branch -D %s", branch)
	run, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		fmt.Println("Ошибка при создании ветки для деплоя")
		panic(err)
	}
	fmt.Println(string(run))
}

// Производит git merge для рабочей ветки из ветки, созданной в <createDeployBranch>. Если нет ошибок, временная ветка будет удалена
func gitMerge(localBranch, deployBranch string) {
	command := fmt.Sprintf("git checkout %s && git merge %s", localBranch, deployBranch)
	run, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		fmt.Println("Ошибка при слиянии в рабочую ветку")
		panic(err)
	}
	fmt.Println(string(run))
	// удалить временную ветку
}

func deploy() {
	gitFetch()
	changes := gitDiff("origin/dev", "dev2")
	services := parseDockerCompose("docker-compose.yaml")
	updateServices := analuzeChanges(services, changes)
	if len(updateServices) == 0 {
		fmt.Println("Изменений не обнаружено")
		// os.Exit(0)
	}
	fmt.Println("Обнаружены изменения в:")
	for _, service := range updateServices {
		text := fmt.Sprintf(" - %s ", service.Name)
		fmt.Println(text)
	}
	branch := createDeployBranch("origin/dev")
	gitMerge("dev2", branch)
	deleteDeployBranch(branch)
	// тут стартует докер

}
