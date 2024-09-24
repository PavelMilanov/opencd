package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/fatih/color"
)

// Производит git fetch, git merge, docker build, docker up исходя из изменений в коммитах.
// Собирает и запускает сервисы в указанном файле docker-compose, где обновились файлы.
func deploy(config Environments, stage string) {
	var STEPS = []string{"[1/4] Анализ изменений проекта", "[2/4] Сборка проекта", "[3/4] Обновление проекта", "[4/4] Очистка кеша docker"}
	info := fmt.Sprintf("Запуск обновления для окружения %s", config.Name)
	color.Green(info)
	switch stage {
	case "merge":
		color.Cyan(STEPS[0])
		if err := gitFetch(); err != nil {
			color.Red("Не найдена удаленная ветка")
			os.Exit(0)
		}
		changes, err := gitDiff(config.Local, config.Remote)
		if err != nil {
			color.Red("Ошибка при анализе изменений удаленной ветки")
			os.Exit(0)
		}
		services, err := parseDockerCompose(config.Docker)
		if err != nil {
			color.Red("Ошибка при чтении файла docker compose")
			os.Exit(0)
		}
		updateServices := analuzeChanges(services, changes)
		if len(updateServices) == 0 {
			color.Red("Изменений не обнаружено")
			os.Exit(0)
		}
		barListName := []string{}
		for _, service := range updateServices {
			barListName = append(barListName, service.Name)
		}
		color.Cyan("Обновление проекта")
		branch, err := createDeployBranch(config.Remote)
		if err != nil {
			return
		}
		err = gitMerge(config.Local, branch)
		if err != nil {
			fmt.Println(err)
		}
		err = deleteDeployBranch(branch)
		if err != nil {
			fmt.Println(err)
		}
		color.Cyan(STEPS[1])
		buildServices, err := buildDockerCompose(updateServices, config.Docker)
		if err != nil {
			fmt.Println(err)
		}
		color.Cyan(STEPS[2])
		upDescription := fmt.Sprintf("Обновление сервисов %s", strings.Join(barListName, " "))
		fmt.Println(upDescription)
		err = upDockerCompose(buildServices, config.Docker)
		if err != nil {
			fmt.Println(err)
		}
		color.Green("Обновление прошло успешно")
	case "docker":
		services, err := parseDockerCompose(config.Docker)
		if err != nil {
			fmt.Println(err)
		}
		barListName := []string{}
		for _, service := range services {
			barListName = append(barListName, service.Name)
		}
		color.Cyan(STEPS[1])
		buildServices, err := buildDockerCompose(services, config.Docker)
		if err != nil {
			fmt.Println(err)
		}
		color.Cyan(STEPS[2])
		upDescription := fmt.Sprintf("Обновление сервисов %s", strings.Join(barListName, " "))
		fmt.Println(upDescription)
		err = upDockerCompose(buildServices, config.Docker)
		if err != nil {
			fmt.Println(err)
		}
		color.Green("Обновление прошло успешно")
	default:
		fmt.Println("флаг не распознан")
	}
	color.Cyan(STEPS[3])
	dockerPrnune()
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
		cmd.Stderr = os.Stderr
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
