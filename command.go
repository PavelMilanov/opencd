package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// Производит git fetch, git merge, docker build, docker up исходя из изменений в коммитах.
// Собирает и запускает сервисы в указанном файле docker-compose, где обновились файлы.
func deploy(config Environments, stage string) {
	info := fmt.Sprintf("Запуск обновления для окружения %s", config.Name)
	fmt.Println(info)
	fmt.Println()
	switch stage {
	case "merge":
		PROGRESSBAR.Describe("[cyan][1/5][reset] Анализ изменений проекта...")
		gitFetch()
		changes, err := gitDiff(config.Local, config.Remote)
		if err != nil {
			errorbar(10)
		}
		services, err := parseDockerCompose(config.Docker)
		if err != nil {
			errorbar(10)
		}
		updateServices := analuzeChanges(services, changes)
		PROGRESSBAR.Add(10)
		if len(updateServices) == 0 {
			PROGRESSBAR.Describe("[cyan][5/5][reset] Изменений не обнаружено...")
			PROGRESSBAR.Finish()
			os.Exit(0)
		}
		barListName := []string{}
		for _, service := range updateServices {
			barListName = append(barListName, service.Name)
		}
		PROGRESSBAR.Describe("[cyan][2/5][reset] Обновление проекта")
		branch, err := createDeployBranch(config.Remote)
		if err != nil {
			return
		}
		err = gitMerge(config.Local, branch)
		if err != nil {
			errorbar(20)
		}
		err = deleteDeployBranch(branch)
		if err != nil {
			errorbar(20)
		}
		PROGRESSBAR.Add(10)
		buildDescription := fmt.Sprintf("[cyan][3/5][reset] Сборка новых образов docker %s", strings.Join(barListName, " "))
		PROGRESSBAR.Describe(buildDescription)
		buildServices, err := buildDockerCompose(updateServices, config.Docker)
		if err != nil {
			print(err, "Ошибка")
			errorbar(30)
		}
		PROGRESSBAR.Add(10)
		upDescription := fmt.Sprintf("[cyan][4/5][reset] Обновление сервисов %s", strings.Join(barListName, " "))
		PROGRESSBAR.Describe(upDescription)
		err = upDockerCompose(buildServices, config.Docker)
		if err != nil {
			errorbar(40)
		}
		PROGRESSBAR.Describe("[cyan][5/5][reset] Обновление прошло успешно")
		PROGRESSBAR.Finish()
	case "docker":
		services, err := parseDockerCompose(config.Docker)
		if err != nil {
			errorbar(10)
		}
		barListName := []string{}
		for _, service := range services {
			barListName = append(barListName, service.Name)
		}
		buildDescription := fmt.Sprintf("[cyan][1/3][reset] Сборка новых образов docker %s", strings.Join(barListName, " "))
		PROGRESSBAR.Describe(buildDescription)
		buildServices, err := buildDockerCompose(services, config.Docker)
		if err != nil {
			errorbar(10)
		}
		fmt.Println(buildServices)
		PROGRESSBAR.Add(50)
		upDescription := fmt.Sprintf("[cyan][2/3][reset] Обновление сервисов %s", strings.Join(barListName, " "))
		PROGRESSBAR.Describe(upDescription)
		err = upDockerCompose(buildServices, config.Docker)
		if err != nil {
			errorbar(50)
		}
		PROGRESSBAR.Describe("[cyan][3/3][reset] Обновление прошло успешно")
		PROGRESSBAR.Finish()
		dockerPrnune()
	default:
		fmt.Println("флаг не распознан")
	}
}

// Вывод информации о версии используемого ПО
func version() {
	fmt.Println("opencd version", VERSION)
	err := cmd("git version")
	if err != nil {
		fmt.Println(err)
	}
	err = cmd("docker compose version")
	if err != nil {
		fmt.Println(err)
	}
}

// Очищает кеш docker
func dockerPrnune() {
	docker := [3]string{"docker container  rm $(docker ps -a -f status=exited -q)", "docker image rm $(docker image ls -f dangling=true -q)", "docker volume rm $(docker volume ls -f dangling=true -q)"}

	var wg sync.WaitGroup
	wg.Add(3)

	worker := func(command string) {
		defer wg.Done()
		cmd := exec.Command("bash", "-c", command)
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			return
		}
	}

	for _, work := range docker {
		go worker(work)
	}
	wg.Wait()
}
