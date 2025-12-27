package main

import (
	"log"
	"os"
	"sort"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor      int
	choices     []string
	selected    map[int]struct{}
	selectedDir string
}

func initialModel() model {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	path := home + "/lab/experiments"

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	type dirWithTime struct {
		name string
		mod  time.Time
	}
	var dirs []dirWithTime
	for _, entry := range entries {
		if entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				continue
			}
			dirs = append(dirs, dirWithTime{name: entry.Name(), mod: info.ModTime()})
		}
	}

	// Sort by modification time descending (latest first)
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].mod.After(dirs[j].mod)
	})

	choices := []string{}
	for _, d := range dirs {
		choices = append(choices, d.name)
	}
	return model{
		choices:     choices,
		selected:    make(map[int]struct{}),
		selectedDir: "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
