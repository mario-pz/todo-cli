package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	program := tea.NewProgram(initialModel)
	if err := program.Start(); err != nil {
		fmt.Println("Error running program:", err)
	}
}

func initialModel() tea.Model {
	return "Hello, World!"
}

func update(msg tea.Msg, model tea.Model) (tea.Model, tea.Cmd) {
	return model, nil
}

func view(model tea.Model) string {
	return fmt.Sprintf("  %v\n\n", model)
}
