package main

import (
	"strings"
	"time"

	"github.com/Dat-one-dev/Sushupti/data"
	"github.com/Dat-one-dev/Sushupti/tui"
	"github.com/Dat-one-dev/Sushupti/utils"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	w         int
	h         int
	daily     []data.DailyStat
	anim      int
	symb      string
	symbColor string
	symbframe int
	catPos    int
	catDir    int
}

type tickMsg struct{}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		tick(),
		refresh(),
	)
}

func tick() tea.Cmd {
	return tea.Tick(150*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

func refresh() tea.Cmd {
	return tea.Tick(1*time.Minute, func(t time.Time) tea.Msg {
		start, end, err := utils.DateRange()
		if err != nil {
			return nil
		}

		config, err := utils.LoadConfig(start, end)
		if err != nil {
			return nil
		}

		daily, err := utils.FetchDaily(config)
		if err != nil {
			return nil
		}

		return refreshMsg(daily)
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		m.tick()
		return m, tick()
	}

	if msg, ok := msg.(refreshMsg); ok {
		m.daily = msg
		return m, refresh()
	}

	return m, nil
}

func (m model) View() string {
	if m.w < 40 || m.h < 10 {
		return ""
	}

	sidebarWidth := 24
	contentWidth := m.w - sidebarWidth - 3
	boxWidth := (contentWidth - 2) / 2

	sidebar := utils.Box(
		tui.RenderSidebar(m.daily, m.symb),
		sidebarWidth,
	)

	dashboard := m.dashboard(boxWidth)

	var s strings.Builder

	s.WriteString(m.header())
	s.WriteString("\n")
	s.WriteString(m.body(sidebar, dashboard, sidebarWidth, contentWidth))
	s.WriteString(m.footer())

	return s.String()
}
