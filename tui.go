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

func renderOverview(daily []DailyStat) string {
	//ALL VARIABLES ARE HERE
	maxSec := 0
	total := 0
	bestDay := ""
	bestTime := 0
	activeDays := 0
	dailyAverage := 0

	//ALL LIPGLLOSS STUFF IS HERE
	overviewTitle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("4")).
		Render("Overview - Last 10 Days")

	totalStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("6"))

	charBar := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("6"))

	//doing ts again for 6-7th time instead of making a better way to do it

	// max seconds loop
	for _, day := range daily {
		if day.TotalSeconds > maxSec {
			maxSec = day.TotalSeconds
		}
	}
	//total seconds loop
	for _, day := range daily {
		total += day.TotalSeconds
	}

	//activeDays loop
	for _, day := range daily {
		if day.TotalSeconds > 0 {
			activeDays++
		}
	}

	//bestDay bestTime loop
	for _, day := range daily {
		if day.TotalSeconds > bestTime {
			bestTime = day.TotalSeconds
			bestDay = day.Date
		}
	}

	//Dailyavg
	if activeDays > 0 {
		dailyAverage = total / activeDays
	}

	//RIGHT HEADER
	right := "  ┌────────────────────────────────────────────┐\n"
	right += "  │  " + overviewTitle + "                   │\n"
	right += "  │  ──────────────────────────────────────────│\n"
	right += "  │  Total       " + totalStyle.Render(fmt.Sprintf("%dh %02dm", total/3600, (total%3600)/60)) + "                       │\n"
	right += "  │  Active Days " + totalStyle.Render(fmt.Sprintf("%d", activeDays)) + "                             │\n"
	right += "  │  Best Day    " + totalStyle.Render(bestDay[5:]) + "  " + totalStyle.Render(fmt.Sprintf("%dh %02dm", bestTime/3600, (bestTime%3600)/60)) + "                 │\n"
	right += "  │  Daily Avg   " + totalStyle.Render(fmt.Sprintf("%dh %02dm", dailyAverage/3600, (dailyAverage%3600)/60)) + "                        │\n"
	right += "  │                                            │\n"

	for _, day := range daily {
		bar := 0

		if maxSec > 0 {
			bar = day.TotalSeconds * 20 / maxSec
		}

		date := day.Date[5:]

		line := "  │  " + date + "  "

		if bar == 0 {
			line += "─"
		} else {
			for j := 0; j < bar; j++ {
				line += charBar.Render("█")
			}
		}

		hours := day.TotalSeconds / 3600
		minutes := (day.TotalSeconds % 3600) / 60

		line += "  " + fmt.Sprintf("%dh %02dm", hours, minutes)

		lineWidth := lipgloss.Width(line)
		padding := 47 - lineWidth

		if padding > 0 {
			line += strings.Repeat(" ", padding)
		}

		line += "│"

		right += line + "\n"
	}
	right += "  └────────────────────────────────────────────┘\n"

	return right
}

func renderProjects(daily []DailyStat) string {
	most := mostUsedProject(daily)
	editor := mostUsedEditor(daily)
	lang := mostUsedlang(daily)
	projectsTitle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("4")).
		Render("Projects")

	mostprojName := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("1")).
		Render(most.ProjectName)

	editorName := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("1")).
		Render(editor.Name)

	langName := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("1")).
		Render(lang.Name)

	time := fmt.Sprintf("%dh %02dm", most.TotalSeconds/3600, (most.TotalSeconds%3600)/60)

	right := "  ┌────────────────────────────────────────────┐\n"
	right += "  │  " + projectsTitle + "                                  │\n"
	right += "  │  ──────────────────────────────────────────│\n"
	right += "  │  Most Hours  " + mostprojName + "  " + time + "                │\n"
	right += "  │  Best Editor  " + editorName + "  " + "                           │\n"
	right += "  │  Best Lang  " + langName + "  " + "                             │\n"
	right += "  └────────────────────────────────────────────┘\n"

	return right
}

func (m model) View() string {
	if m.w < 20 || m.h < 5 {
		return ""
	}
	header := " SUSHUPTI "
	header = lipgloss.NewStyle().Foreground(lipgloss.Color("4")).Render(header)
	menu := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("1")).
		Render("MENU")
	selected := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("6"))

	s := "╭"
	s += strings.Repeat("─", (m.w-len(" SUSHUPTI ")-2)/2)
	s += header
	s += strings.Repeat("─", (m.w-len(" SUSHUPTI ")-2)/2)
	s += "╮\n"

	right := renderOverview(m.daily)
	right += renderProjects(m.daily)

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
