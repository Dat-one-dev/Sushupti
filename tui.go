package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	w        int
	h        int
	selected int
	daily    []DailyStat
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
			if m.selected < 2 {
				m.selected++
			}
		}

		if key.String() == "up" || key.String() == "k" {
			if m.selected > 0 {
				m.selected--
			}
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
	menu := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("4")).
		Render("MENU")
	selected := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("6"))

	s := "╭"
	s += strings.Repeat("─", (m.w-len(" SUSHUPTI ")-2)/2)
	s += header
	s += strings.Repeat("─", (m.w-len(" SUSHUPTI ")-2)/2)
	s += "╮\n"

	//doing ts again for 6-7th time instead of making a better way to do it
	maxSec := 0

	for _, day := range m.daily {
		if day.TotalSeconds > maxSec {
			maxSec = day.TotalSeconds
		}
	}

	//right
	total := 0
	for _, day := range m.daily {
		total += day.TotalSeconds
	}
	right := "  DAILY ACTIVITY\n"
	right += fmt.Sprintf("  Total: %dh %02dm\n\n", total/3600, (total%3600)/60)

	for _, day := range m.daily {
		bar := 0

		if maxSec > 0 {
			bar = day.TotalSeconds * 30 / maxSec
		}

		right += "  " + day.Date + " "

		for j := 0; j < bar; j++ {
			right += "█"
		}

		hours := day.TotalSeconds / 3600
		minutes := (day.TotalSeconds % 3600) / 60

		right += " " + fmt.Sprintf("%dh %02dm", hours, minutes) + "\n"
	}
	rightLines := strings.Split(right, "\n")

	for i := 0; i < m.h-2; i++ {
		s += "│"

		if i == 0 {
			s += "┌──────────────────┐"
		} else if i == m.h-3 {
			s += "└──────────────────┘"
		} else if i == 1 {
			s += "│       " + menu + "       │"
		} else if i == 2 {
			if m.selected == 0 {
				s += "│   " + selected.Render("▸  Overview") + "    │"
			} else {
				s += "│     Overview     │"
			}
		} else if i == 3 {
			if m.selected == 1 {
				s += "│   " + selected.Render("▸  Projects") + "    │"
			} else {
				s += "│     Projects     │"
			}
		} else if i == 4 {
			if m.selected == 2 {
				s += "│   " + selected.Render("▸  Activity") + "    │"
			} else {
				s += "│     Activity     │"
			}
		} else {
			s += "│                  │"
		}

		s += "│"

		if i < len(rightLines) {
			s += rightLines[i]

			space := m.w - 23 - lipgloss.Width(rightLines[i])

			if space > 0 {
				s += strings.Repeat(" ", space)
			}
		} else {
			s += strings.Repeat(" ", m.w-23)
		}

		s += "│\n"
	}

	s += "╰"
	s += strings.Repeat("─", m.w-2)
	s += "╯"

	return s
}
