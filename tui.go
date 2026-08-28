package main

import (
	"strings"
	"time"

	"github.com/Dat-one-dev/Sushupti/data"
	"github.com/Dat-one-dev/Sushupti/tui"
	"github.com/Dat-one-dev/Sushupti/utils"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	w         int
	h         int
	daily     []data.DailyStat
	anim      int
	symb      string
	symbColor string
	symbframe int
}

type tickMsg struct{}

func (m model) Init() tea.Cmd {
	return tick()
}

func tick() tea.Cmd {
	return tea.Tick(150*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
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

	if _, ok := msg.(tickMsg); ok {
		m.anim++

		m.symbframe++

		switch m.symbframe % 4 {
		case 0:
			m.symb = "│"
			m.symbColor = "1"
		case 1:
			m.symb = "╲"
			m.symbColor = "1"
		case 2:
			m.symb = "─"
			m.symbColor = "1"
		case 3:
			m.symb = "╱"
			m.symbColor = "1"
		}

		return m, tick()
	}

	return m, nil
}
func (m model) View() string {
	if m.w < 40 || m.h < 10 {
		return ""
	}

	// HEADER
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

	headerLine := "╭"
	headerLine += strings.Repeat("─", left)
	headerLine += header
	headerLine += strings.Repeat("─", right)
	headerLine += "╮"

	// WIDTHS
	sidebarWidth := 24
	gap := 2

	contentWidth := m.w - sidebarWidth - gap - 1

	if contentWidth < 20 {
		contentWidth = 20
	}

	boxWidth := (contentWidth - 2) / 2

	// DASHBOARD BOXES
	overviewBox := utils.Box(tui.RenderOverview(m.daily, m.anim), boxWidth)
	leaderboardBox := utils.Box(tui.RenderLeaderboard(m.daily), boxWidth)
	projectsBox := utils.Box(tui.RenderProjects(m.daily), boxWidth)
	languagesBox := utils.Box(tui.ProjectBar(m.daily, boxWidth, m.anim), boxWidth)
	sidebarBox := utils.Box(tui.RenderSidebar(m.daily, m.symb), sidebarWidth)

	rightC := leaderboardBox + "\n" +
		projectsBox + "\n" +
		languagesBox

	cat := tui.RenderCat(m.symbframe)
	// SIDEBAR + DASHBOARD
	dashboard := utils.JoinBoxes(overviewBox, rightC)
	sidebarLines := strings.Split(sidebarBox, "\n")
	dashboardLines := strings.Split(dashboard, "\n")
	catlines := cat
	height := m.h - 2

	var s strings.Builder

	s.WriteString(headerLine)
	s.WriteString("\n")

	for i := 0; i < height; i++ {
		s.WriteString("│")

		// SIDEBAR
		if i < len(sidebarLines) {
			s.WriteString(sidebarLines[i])
		} else {
			// CAT
			// CAT
			catStart := height - len(catlines)
			catIndex := i - catStart

			if catIndex >= 0 && catIndex < len(catlines) {
				catLine := catlines[catIndex]

				catWidth := lipgloss.Width(catLine)
				leftPadding := (sidebarWidth - catWidth) / 2

				s.WriteString(strings.Repeat(" ", leftPadding))
				s.WriteString(catLine)

				remaining := sidebarWidth - leftPadding - catWidth
				if remaining > 0 {
					s.WriteString(strings.Repeat(" ", remaining))
				}
			} else {
				s.WriteString(strings.Repeat(" ", sidebarWidth))
			}
		}

		s.WriteString(" ")

		// DASHBOARD
		if i < len(dashboardLines) {
			line := dashboardLines[i]
			padding := contentWidth - lipgloss.Width(line)

			s.WriteString(line)

			if padding > 0 {
				s.WriteString(strings.Repeat(" ", padding))
			}
		} else {
			s.WriteString(strings.Repeat(" ", contentWidth))
		}

		s.WriteString("│\n")
	}

	quitText := lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render(" q/ctrl+c to quit")
	brdw := m.w - lipgloss.Width(quitText) - 2 //sometime my variable name are so I lwk start regeretting my life choices....lets see if anyone in future finds this repo then does he know what bdrw means....ig if i see this code after 7-8 months i will forget it
	s.WriteString("╰")
	s.WriteString(strings.Repeat("─", brdw/2))
	s.WriteString(quitText)
	s.WriteString(strings.Repeat("─", brdw-brdw/2))
	s.WriteString("╯")

	return s.String()
}
