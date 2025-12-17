package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func createDir(dir string) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Failed to get gome directory:", err)
		return
	}

	basePath := filepath.Join(home, "lab", "experiments")
	if err := os.MkdirAll(basePath, 0755); err != nil {
		fmt.Println("Failed to create base directory:", err)
		return
	}

	path := filepath.Join(basePath, dir)
	pathString := "~/lab/experiments/" + dir

	if err := os.Mkdir(path, 0755); os.IsExist(err) {
		fmt.Printf("Experiment '%s' already exists!\n", dir)
		return
	} else if err != nil {
		fmt.Println("Failed to create directory:", err)
	} else {
		fmt.Printf("New experiment '%s' created!\n", dir)
	}

	fmt.Println("Copy below command to go to directory:")
	fmt.Println("cd", pathString)
}

func main() {
	currentDate := time.Now().Format("2006-01-02")
	var name string

	fmt.Println("Enter name of experiment:")
	fmt.Printf("%s-", currentDate)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		name = scanner.Text()
	}

	name = strings.ReplaceAll(name, " ", "-")

	if name == "" {
		fmt.Println("Name connot be empty!")
		return
	}

	dir := currentDate + "-" + name
	createDir(dir)
}
