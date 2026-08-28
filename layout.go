package main

import (
	"fmt"
	"strings"

	"github.com/Dat-one-dev/Sushupti/tui"
	"github.com/Dat-one-dev/Sushupti/utils"
	"github.com/charmbracelet/lipgloss"
)

func logError(err error) bool {
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	return true
}

func (m model) header() string {
	symbol := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(m.symbColor)).
		Render(m.symb)

	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("4"))

	header := headerStyle.Render(" SUSHUPTI [") +
		symbol +
		headerStyle.Render("]")

	headerWidth := lipgloss.Width(header)
	remaining := m.w - headerWidth - 2

	left := remaining / 2
	right := remaining - left

	return "╭" +
		strings.Repeat("─", left) +
		header +
		strings.Repeat("─", right) +
		"╮"
}

func (m model) dashboard(boxW int) string {
	// DASHBOARD BOXES
	overviewBox := utils.Box(tui.RenderOverview(m.daily, m.anim), boxW)
	leaderboardBox := utils.Box(tui.RenderLeaderboard(m.daily), boxW)
	projectsBox := utils.Box(tui.RenderProjects(m.daily), boxW)
	languagesBox := utils.Box(tui.ProjectBar(m.daily, boxW, m.anim*2), boxW)
	dateBox := utils.Box(tui.RenderDate(), boxW)
	clockBox := utils.Box(tui.RenderClock(), boxW)

	rightC := leaderboardBox + "\n" +
		projectsBox + "\n" +
		languagesBox + "\n" +
		dateBox

	leftC := overviewBox + "\n" + clockBox

	return utils.JoinBoxes(leftC, rightC)
}

func (m model) catLine(cat []string, i, height, sidebarWidth, catWidth int) string {
	catStart := height - len(cat)
	catIndex := i - catStart

	if catIndex < 0 || catIndex >= len(cat) {
		return strings.Repeat(" ", sidebarWidth)
	}

	line := cat[catIndex]

	left := m.catPos
	maxLeft := sidebarWidth - catWidth

	if left > maxLeft {
		left = maxLeft
	}
	if left < 0 {
		left = 0
	}

	result := strings.Repeat(" ", left) + line

	if remaining := sidebarWidth - left - catWidth; remaining > 0 {
		result += strings.Repeat(" ", remaining)
	}

	return result
}

func (m model) body(sidebar, dashboard string, sidebarWidth, contentWidth int) string {
	sidebarLines := strings.Split(sidebar, "\n")
	dashboardLines := strings.Split(dashboard, "\n")
	cat := tui.RenderCat(m.symbframe)

	height := m.h - 2
	catWidth := tui.CatWidth()

	var s strings.Builder

	for i := 0; i < height; i++ {
		s.WriteString("│")

		if i < len(sidebarLines) {
			s.WriteString(sidebarLines[i])
		} else {
			s.WriteString(m.catLine(cat, i, height, sidebarWidth, catWidth))
		}

		s.WriteString(" ")

		if i < len(dashboardLines) {
			line := dashboardLines[i]
			s.WriteString(line)

			if padding := contentWidth - lipgloss.Width(line); padding > 0 {
				s.WriteString(strings.Repeat(" ", padding))
			}
		} else {
			s.WriteString(strings.Repeat(" ", contentWidth))
		}

		s.WriteString("│\n")
	}

	return s.String()
}

func (m model) footer() string {
	quitText := lipgloss.NewStyle().
		Foreground(lipgloss.Color("8")).
		Render(" q/ctrl+c to quit")

	width := m.w - lipgloss.Width(quitText) - 2

	return "╰" +
		strings.Repeat("─", width/2) +
		quitText +
		strings.Repeat("─", width-width/2) +
		"╯"
}

func (m *model) tick() {
	m.anim++
	m.symbframe++

	switch m.symbframe % 4 {
	case 0:
		m.symb = "│"
	case 1:
		m.symb = "╲"
	case 2:
		m.symb = "─"
	case 3:
		m.symb = "╱"
	}

	m.moveCat()
}

func (m *model) moveCat() {
	catWidth := tui.CatWidth()
	maxPos := 24 - catWidth

	if maxPos < 0 {
		maxPos = 0
	}

	if m.catDir == 0 {
		m.catDir = 1
	}

	m.catPos += m.catDir

	if m.catPos >= maxPos {
		m.catPos = maxPos
		m.catDir = -1
	}

	if m.catPos <= 0 {
		m.catPos = 0
		m.catDir = 1
	}
}
