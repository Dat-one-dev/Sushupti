package style

import "github.com/charmbracelet/lipgloss"

var (
	Title     = lipgloss.NewStyle().Bold(true).Foreground(Primary)
	Value     = lipgloss.NewStyle().Bold(true).Foreground(Secondary)
	Highlight = lipgloss.NewStyle().Bold(true).Foreground(Accent)
	Muted     = lipgloss.NewStyle().Foreground(MutedC)

	ProjectColors = []lipgloss.Style{
		lipgloss.NewStyle().Bold(true).Foreground(Project1),
		lipgloss.NewStyle().Bold(true).Foreground(Project2),
		lipgloss.NewStyle().Bold(true).Foreground(Project3),
		lipgloss.NewStyle().Bold(true).Foreground(Project4),
		lipgloss.NewStyle().Bold(true).Foreground(Project5),
	}
)

const Divider = "──────────────────────────────────────────"
