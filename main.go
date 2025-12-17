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
		fmt.Fprintln(os.Stderr, "Failed to get gome directory:", err)
		return
	}

	basePath := filepath.Join(home, "lab", "experiments")
	if err := os.MkdirAll(basePath, 0755); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create base directory:", err)
		return
	}

	path := filepath.Join(basePath, dir)
	pathString := "~/lab/experiments/" + dir

	if err := os.Mkdir(path, 0755); os.IsExist(err) {
		fmt.Fprintf(os.Stderr, "Experiment '%s' already exists!\n", dir)
		return
	} else if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create directory:", err)
	} else {
		fmt.Fprintf(os.Stderr, "New experiment '%s' created!\n", dir)
	}

	fmt.Println("cd", pathString)
}

func main() {
	currentDate := time.Now().Format("2006-01-02")
	var name string

	fmt.Fprintln(os.Stderr, "Enter name of experiment:")
	fmt.Fprintf(os.Stderr, "%s-", currentDate)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		name = scanner.Text()
	}

	name = strings.ReplaceAll(name, " ", "-")

	if name == "" {
		fmt.Fprintln(os.Stderr, "Name connot be empty!")
		return
	}

	dir := currentDate + "-" + name
	createDir(dir)
}
