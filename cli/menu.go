package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func MakeMenu() menuModel {
	m := menuModel{
		usernameInput: textinput.NewModel(),
		passwordInput: textinput.NewModel(),
	}
	m.usernameInput.Prompt = "username"
	m.passwordInput.Prompt = "password"
	m.passwordInput.EchoMode = textinput.EchoPassword
	m.windowStyle = lipgloss.NewStyle().
		Padding(5, 12).
		Margin(0, 0).
		Border(lipgloss.NormalBorder(), true)
	return m
}

type menuModel struct {
	usernameInput textinput.Model
	passwordInput textinput.Model
	windowStyle   lipgloss.Style
	statusMessage string
}

func (m menuModel) Init() tea.Cmd {
	return nil
}

func (m menuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "tab":
			if m.usernameInput.Focused() {
				m.passwordInput.Prompt = ""
				m.usernameInput.Blur()
				m.passwordInput.Focus()
			} else {
				m.usernameInput.Prompt = ""
				m.passwordInput.Blur()
				m.usernameInput.Focus()
			}
		case "enter":
			if m.usernameInput.Value() != "" && m.passwordInput.Value() != "" {
				if checkRegex(m.usernameInput.Value()) {
					// Do something
					m.statusMessage = "Email is valid"
				} else {
					m.statusMessage = "Email is not valid"
				}
			}
			return m, nil
		case "ctrl+r":
			if m.usernameInput.Value() != "" && m.passwordInput.Value() != "" {
				if checkRegex(m.usernameInput.Value()) {
					// Do something
					m.statusMessage = "Registration successful"
				} else {
					m.statusMessage = "Email is not valid"
				}
			}
			return m, nil
		}
	}

	m.usernameInput, _ = m.usernameInput.Update(msg)
	m.passwordInput, _ = m.passwordInput.Update(msg)
	return m, nil
}

func (m menuModel) View() string {
	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#1E90FF")).
		Bold(true)

	descriptionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#bfff00")).
		Background(lipgloss.Color("#000000")).Bold(true)

	inputStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true)

	windowStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(lipgloss.Color("#FFFFFF")).
		Padding(1, 3)

	title := titleStyle.Render("Login Form")

	desc := descriptionStyle.Render("Ctrl R to Register, Tab to Navigate")

	usernameInput := inputStyle.Render(m.usernameInput.View())
	passwordInput := inputStyle.Render(m.passwordInput.View())

	form := fmt.Sprintf("%s\n%s\n",
		usernameInput, passwordInput,
	)

	statusStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F00")).
		Border(lipgloss.NormalBorder(), true)

	if m.statusMessage != "" {
		form += statusStyle.Render(fmt.Sprintf("%s\n", m.statusMessage))
	}

	window := windowStyle.Render(fmt.Sprintf("%s\n%s\n%s", title, form, desc))

	return lipgloss.JoinVertical(lipgloss.Top,
		window,
	)
}
