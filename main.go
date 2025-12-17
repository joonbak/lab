package main

import (
	"fmt"
	"os"
	"time"
)

func createDir(dir string) {
	path := "./" + dir

	if err := os.Mkdir(path, 0755); os.IsExist(err) {
		fmt.Printf("Experiment '%s' already exists!\n", dir)
	} else {
		fmt.Printf("New experiment '%s' created!\n", dir)
	}
}

func main() {
	currentDate := time.Now().Format("2006-01-02")
	var name string

	fmt.Println("Enter name of experiment:")
	fmt.Printf("%s-", currentDate)
	fmt.Scanln(&name)

	dir := currentDate + "-" + name
	createDir(dir)
	fmt.Println("Copy below command to go to directory")
	fmt.Println("cd", dir)
}
