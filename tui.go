package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	w         int
	h         int
	daily     []DailyStat
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
		if m.anim < 20 {
			m.anim++
		}

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

func renderOverview(daily []DailyStat, anim int) []string {
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

			if bar > anim {
				bar = anim
			}
		}

		line := day.Date[5:] + "  "

		if bar == 0 {
			line += "─"
		} else {
			for j := 0; j < bar; j++ {
				line += charBar.Render("#")
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
func renderSidebar(daily []DailyStat, symb string) []string {
	total := 0
	today := 0

	for _, day := range daily {
		total += day.TotalSeconds
	}
	if len(daily) > 0 {
		today = daily[len(daily)-1].TotalSeconds
	}

	totalHr := total / 3600
	totalMin := (total % 3600) / 60
	todayHr := today / 3600
	todayMin := (today % 3600) / 60

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("4"))

	valueStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("6"))

	liveStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("1"))

	content := []string{
		titleStyle.Render("SUSHUPTI"),
		"",
		liveStyle.Render(symb + " LIVE"),
		"",
		titleStyle.Render("TODAY"),
		valueStyle.Render(fmt.Sprintf("%02dh %02dm", todayHr, todayMin)),
		"",
		titleStyle.Render("DATA"),
		fmt.Sprintf("%d days", len(daily)),
		valueStyle.Render(fmt.Sprintf("%02dh %02dm", totalHr, totalMin)),
		"",
		titleStyle.Render("STATUS"),
		liveStyle.Render("● ACTIVE"),
	}
	return content
}
func renderCat(frame int) []string {
	cats := [][]string{
		{
			"    HACK CAT     ",
			"                 ",
			"  ^~^  ,         ",
			" (`Y`) )         ",
			" /   \\/          ",
			"(\\|||/)          ",
		},
		{
			"    HACK CAT     ",
			"                 ",
			"  ^~^   ,        ",
			" ('Y') )         ",
			" /   \\/          ",
			"(\\|||/)          ",
		},
		{
			"    HACK CAT     ",
			"                 ",
			"  ^~^    ,       ",
			" ('Y') )         ",
			" /   \\/          ",
			"(\\|||/)          ",
		},
		{
			"    HACK CAT     ",
			"                 ",

			"  ^~^   ,        ",
			" (`Y`) )         ",
			" /   \\/          ",
			"(\\|||/)          ",
		},
		{
			"    HACK CAT     ",
			"                 ",

			"  ^~^  ,         ",
			" ('Y') )         ",
			" /   \\/          ",
			"(\\|||/)          ",
		},
		{
			"    HACK CAT     ",
			"                 ",

			"  ^~^ ,          ",
			" ('Y') )         ",
			" /   \\/          ",
			"(\\|||/)          ",
		},
		{
			"    HACK CAT     ",
			"                 ",

			"  ^~^,           ",
			" ('Y') )         ",
			" /   \\/          ",
			"(\\|||/)          ",
		},
		{
			"    HACK CAT     ",
			"                 ",

			"  ^~^ ,          ",
			" ('Y') )         ",
			" /   \\/          ",
			"(\\|||/)          ",
		},
		{
			"    HACK CAT     ",
			"                 ",

			"  ^~^  ,         ",
			" ('Y') )         ",
			" /   \\/          ",
			"(\\|||/)          ",
		},
		{
			"    HACK CAT     ",
			"                 ",

			"  ^~^   ,        ",
			" ('Y') )         ",
			" /   \\/          ",
			"(\\|||/)          ",
		},
	}

	return cats[frame%len(cats)]
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
	overviewBox := box(renderOverview(m.daily, m.anim), boxWidth)
	leaderboardBox := box(renderLeaderboard(m.daily), boxWidth)

	top := joinBoxes(overviewBox, leaderboardBox)

	projectsBox := box(
		renderProjects(m.daily),
		boxWidth,
	)

	bottom := projectsBox
	sidebarBox := box(
		renderSidebar(m.daily, m.symb),
		sidebarWidth,
	)
	cat := renderCat(m.symbframe)
	// SIDEBAR + DASHBOARD
	dashboard := top + "\n" + bottom
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
