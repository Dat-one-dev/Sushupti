package main

import (
	"strings"

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
	if m.w < 20 || m.h < 5 {
		return ""
	}

	header := " SUSHUPTI "
	header = lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Render(header)

	s := "╔"
	s += strings.Repeat("═", (m.w-len(" SUSHUPTI ")-2)/2)
	s += header
	s += strings.Repeat("═", (m.w-len(" SUSHUPTI ")-2)/2)
	s += "╗\n"

	for i := 0; i < m.h-2; i++ {
		s += "║"
		s += strings.Repeat(" ", 20)
		s += "│"
		s += strings.Repeat(" ", m.w-23)
		s += "║\n"
	}

	s += "╚"
	s += strings.Repeat("═", m.w-2)
	s += "╝"

	return s
}
