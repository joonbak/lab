package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
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
	args := os.Args[1:]
	if len(args) == 0 {
		currentDate := time.Now().Format("2006-01-02")
		var name string

		fmt.Fprintln(os.Stderr, "Enter name of experiment:")
		fmt.Fprintf(os.Stderr, "%s-", currentDate)

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			name = scanner.Text()
		}

		name = strings.ReplaceAll(strings.TrimSpace(name), " ", "-")

		if name == "" {
			fmt.Fprintln(os.Stderr, "Name connot be empty!")
			return
		}

		dir := currentDate + "-" + name
		createDir(dir)
	} else if args[0] == "list" {
		p := tea.NewProgram(initialModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Please enter a valid command!")
	}
}
