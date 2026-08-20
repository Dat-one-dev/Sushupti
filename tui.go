package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	w        int
	h        int
	selected int
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
		if key.String() == "down" || key.String() == "j" {
			m.selected++
		}

		if key.String() == "up" || key.String() == "k" {
			m.selected--
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
	menu := lipgloss.NewStyle().Foreground(lipgloss.Color("1")).Render("MENU")

	s := "╔"
	s += strings.Repeat("═", (m.w-len(" SUSHUPTI ")-2)/2)
	s += header
	s += strings.Repeat("═", (m.w-len(" SUSHUPTI ")-2)/2)
	s += "╗\n"

	for i := 0; i < m.h-2; i++ {
		s += "║"

		if i == 0 {
			s += "┌──────────────────┐"
		} else if i == m.h-3 {
			s += "└──────────────────┘"
		} else if i == 1 {
			s += "│       " + menu + "       │"
		} else if i == 2 {
			if m.selected == 0 {
				s += "│   > Overview     │"
			} else {
				s += "│     Overview     │"
			}
		} else if i == 3 {
			if m.selected == 1 {
				s += "│   > Projects     │"
			} else {
				s += "│     Projects     │"
			}
		} else if i == 4 {
			if m.selected == 2 {
				s += "│   > Activity     │"
			} else {
				s += "│     Activity     │"
			}
		} else {
			s += "│                  │"
		}

		s += "│"
		s += strings.Repeat(" ", m.w-23)
		s += "║\n"
	}

	s += "╚"
	s += strings.Repeat("═", m.w-2)
	s += "╝"

	return s
}
