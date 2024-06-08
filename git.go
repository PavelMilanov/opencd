package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// Выполняет git fetch для удаленного репозитория.
func gitFetch() error {
	run := exec.Command("bash", "-c", "git remote | git fetch")
	run.Stdout = os.Stdout
	run.Stderr = os.Stderr
	err := run.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Проверяет изменения в ветках репозитория и возвращает директории, где были изменения.
func gitDiff(localBranch string, remoteBranch string) ([]string, error) {
	command := fmt.Sprintf("git diff %s %s", remoteBranch, localBranch)
	diff, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		fmt.Println(err)
		return nil, err
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
	return changes, nil
}

// Возвращает список сервисов docker-compose, для которых былм изменения в полученных коммитах.
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

// Создает отдельную ветку git для применения изменений, возвращает название этой ветки.
func createDeployBranch(remoteBranch string) (string, error) {
	command := fmt.Sprintf("git log %s | head  -1", remoteBranch)
	commitsha, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	shortSha := strings.Split(string(commitsha), " ")[1][:7] // commit 11e00c3b19f88ec7602c4d115871113e49f63e07 => 11e00c3
	deployBranch := "deploy" + "-" + shortSha
	command2 := fmt.Sprintf("git checkout -b %s && git merge %s", deployBranch, remoteBranch) // переключение на ветку деплоя и применение изменений
	run := exec.Command("bash", "-c", command2)
	run.Stdout = os.Stdout
	run.Stderr = os.Stderr
	run.Run()
	return deployBranch, nil
}

// Удаляет ранее созданную временную ветку для деплоя.
func deleteDeployBranch(branch string) error {
	command := fmt.Sprintf("git branch -D %s", branch)
	run := exec.Command("bash", "-c", command)
	run.Stdout = os.Stdout
	run.Stderr = os.Stderr
	err := run.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Производит git merge для рабочей ветки из ветки, созданной в <createDeployBranch>. Если нет ошибок, временная ветка будет удалена.
func gitMerge(localBranch, deployBranch string) error {
	command := fmt.Sprintf("git checkout %s && git merge %s", localBranch, deployBranch)
	run := exec.Command("bash", "-c", command)
	run.Stdout = os.Stdout
	run.Stderr = os.Stderr
	err := run.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
