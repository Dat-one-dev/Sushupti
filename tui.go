package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	w int
	h int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// QUIT
	if key, ok := msg.(tea.KeyMsg); ok {
		if key.String() == "q" || key.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	if size, ok := msg.(tea.WindowSizeMsg); ok {
		m.w = size.Width
		m.h = size.Height
	}

	return m, nil
}

func (m model) View() string {
	header := " Sushupti "
	top := lipgloss.NewStyle().Bold(true).Render(header)
	box := lipgloss.NewStyle().Width(m.w-2).Height(m.h-2).Border(lipgloss.DoubleBorder()).Padding(1, 2).Render(top)

	return box
}
