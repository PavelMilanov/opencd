package main

import (
	"fmt"
	"os"
)

const OPENCD_CONFIG = "opencd.yaml"
const DOCKER_COMPOSE = "docker-compose.*"

func main() {
	err := checkComponents()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "deploy":
			deploy()
		case "rollback":
			fmt.Println("rollback")
		case "commits":
			displayCommits()
		default:
			fmt.Println("unknown command")
		}
	} else {
		fmt.Println("bad command")
	}
}
