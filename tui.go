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

func renderOverview(daily []DailyStat) []string {
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

	content := []string{
		overviewTitle,
		"──────────────────────────────────────────",
		"Total       " + totalStyle.Render(fmt.Sprintf("%dh %02dm", total/3600, (total%3600)/60)),
		"Active Days " + totalStyle.Render(fmt.Sprintf("%d", activeDays)),
		"Best Day    " + totalStyle.Render(bestDay[5:]) + "  " +
			totalStyle.Render(fmt.Sprintf("%dh %02dm", bestTime/3600, (bestTime%3600)/60)),
		"Daily Avg   " + totalStyle.Render(fmt.Sprintf("%dh %02dm", dailyAverage/3600, (dailyAverage%3600)/60)),
		"",
	}

	for _, day := range daily {
		bar := 0

		if maxSec > 0 {
			bar = day.TotalSeconds * 20 / maxSec
		}

		line := day.Date[5:] + "  "

		if bar == 0 {
			line += "─"
		} else {
			for j := 0; j < bar; j++ {
				line += charBar.Render("█")
			}
		}

		hours := day.TotalSeconds / 3600
		minutes := (day.TotalSeconds % 3600) / 60

		line += fmt.Sprintf("  %dh %02dm", hours, minutes)

		content = append(content, line)
	}

	return content
}
func renderProjects(daily []DailyStat) []string {
	most := mostUsedProject(daily)
	projectTime := projectTotalTime(daily, most.ProjectName)
	projectShare := projectShareCalc(daily, projectTime)

	projectsTitle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("4")).
		Render("Projects")

	mostProjectName := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("1")).
		Render(most.ProjectName)

	shareTxt := fmt.Sprintf("%.1f%%", projectShare)

	content := []string{
		projectsTitle,
		"──────────────────────────────────────────",
		"Most Hours  " + mostProjectName,
		"Total Time  " + fmt.Sprintf("%dh %02dm",
			projectTime/3600,
			(projectTime%3600)/60),
		"Share       " + shareTxt,
	}

	return content
}
func renderLeaderboard(daily []DailyStat) []string {
	projects := ProjectLeaderboard(daily)

	leaderboardTitle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("4")).
		Render("Projects Leaderboard")

	content := []string{
		leaderboardTitle,
		"──────────────────────────────────────────",
	}

	for i, project := range projects {
		hours := project.TotalSeconds / 3600
		minutes := (project.TotalSeconds % 3600) / 60

		content = append(content, fmt.Sprintf(
			"%d. %-20s %dh %02dm",
			i+1,
			project.ProjectName,
			hours,
			minutes,
		))
	}

	return content
}
func (m model) View() string {
	if m.w < 40 || m.h < 10 {
		return ""
	}

	// HEADER
	header := lipgloss.NewStyle().
		Foreground(lipgloss.Color("4")).
		Render(" SUSHUPTI ")

	headerLine := "╭"
	headerLine += strings.Repeat("─", (m.w-lipgloss.Width(header)-2)/2)
	headerLine += header
	headerLine += strings.Repeat("─", (m.w-lipgloss.Width(header)-2)/2)
	headerLine += "╮"

	// WIDTHS
	menuWidth := 20
	gap := 2

	contentWidth := m.w - menuWidth - gap - 3

	if contentWidth < 20 {
		contentWidth = 20
	}

	boxWidth := contentWidth / 2

	// MENU
	menuStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("1"))

	selectedStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("6"))

	menuContent := []string{
		menuStyle.Render("MENU"),
		"",
		"Overview",
		"Projects",
		"Activity",
	}

	if m.selected == 0 {
		menuContent[2] = selectedStyle.Render("▸ Overview")
	}

	if m.selected == 1 {
		menuContent[3] = selectedStyle.Render("▸ Projects")
	}

	if m.selected == 2 {
		menuContent[4] = selectedStyle.Render("▸ Activity")
	}

	menuBox := box(menuContent, menuWidth)

	// DASHBOARD BOXES
	overviewBox := box(renderOverview(m.daily), boxWidth)
	leaderboardBox := box(renderLeaderboard(m.daily), boxWidth)

	top := joinBoxes(overviewBox, leaderboardBox)

	projectsBox := box(
		renderProjects(m.daily),
		boxWidth,
	)

	bottom := projectsBox

	// JOIN MENU + DASHBOARD
	dashboard := top + "\n" + bottom

	menuLines := strings.Split(menuBox, "\n")
	dashboardLines := strings.Split(dashboard, "\n")

	height := m.h - 2

	var s strings.Builder

	s.WriteString(headerLine)
	s.WriteString("\n")

	for i := 0; i < height; i++ {
		s.WriteString("│")

		// MENU
		if i < len(menuLines) {
			s.WriteString(menuLines[i])
		} else {
			s.WriteString(strings.Repeat(" ", menuWidth))
		}

		s.WriteString("│")

		// GAP
		s.WriteString(" ")

		// DASHBOARD
		if i < len(dashboardLines) {
			s.WriteString(dashboardLines[i])
		} else {
			s.WriteString(strings.Repeat(" ", len(top)))
		}

		s.WriteString("\n")
	}

	s.WriteString("╰")
	s.WriteString(strings.Repeat("─", m.w-2))
	s.WriteString("╯")

	return s.String()
}
