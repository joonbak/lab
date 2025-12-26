package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var faint = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Faint(true)

func (m model) View() string {
	// The header
	s := "List of your Lab Experiments\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}
		// Render the row
		row := fmt.Sprintf(" [%s] %s", checked, choice)

		if m.cursor != i {
			row = faint.Render(row)
		}

		s += cursor + row + "\n"
	}

	// The footer
	s += "\nq" + faint.Render(" quit ") + "d" + faint.Render(" delete ") + "j/k" + faint.Render(" down/up ") + "enter" + faint.Render(" select\n")

	// Send the UI for rendering
	return s
}
