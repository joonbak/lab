package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

		case "g":
			m.selectedDir = m.choices[m.cursor]
			return m, tea.Quit

		case "d":
			// Build a new slice without selected items

			home, err := os.UserHomeDir()
			if err != nil {
				log.Fatal(err)
			}

			path := home + "/lab/experiments/"
			dirs := []string{}
			newChoices := make([]string, 0, len(m.choices))
			newSelected := make(map[int]struct{})

			for i, choice := range m.choices {
				if _, selected := m.selected[i]; !selected {
					newChoices = append(newChoices, choice)
				} else {
					dirs = append(dirs, choice)
				}
			}
			for _, dir := range dirs {
				if err := os.RemoveAll(path + dir); err != nil {
					log.Printf("failed to remove %s: %v", dir, err)
				}
			}

			m.choices = newChoices
			m.selected = newSelected

			// Fix cursor position
			if m.cursor >= len(m.choices) {
				m.cursor = len(m.choices) - 1
			}
			if m.cursor < 0 {
				m.cursor = 0
			}

		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
