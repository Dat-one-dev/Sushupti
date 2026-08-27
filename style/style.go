package style

import "github.com/charmbracelet/lipgloss"

var (
	Title     = lipgloss.NewStyle().Bold(true).Foreground(Primary)
	Value     = lipgloss.NewStyle().Bold(true).Foreground(Secondary)
	Highlight = lipgloss.NewStyle().Bold(true).Foreground(Accent)
	Muted     = lipgloss.NewStyle().Foreground(MutedC)
)
