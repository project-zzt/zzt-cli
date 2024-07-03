package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var headline = lipgloss.NewStyle().Bold(true).PaddingTop(1)
var aqua = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
var magenta = lipgloss.NewStyle().Foreground(lipgloss.Color("5"))

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := headline.Render("Welcome to the " + aqua.Render("zzt-cli") + " villain!\n")
	s += "\nWhat type of app do you want to build?\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		l := fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
		if checked == "x" {
			s += aqua.Render(l)
		} else {
			s += l
		}
	}

	s += magenta.Render("\nPress q to quit. \n")
	return s
}

func initialModel() model {
	return model{
		choices:  []string{"Web Application", "Static Site", "Poser (Blog, Forum)"},
		selected: make(map[int]struct{}),
	}
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Captain, we have an error: %v", err)
		os.Exit(1)
	}
}
