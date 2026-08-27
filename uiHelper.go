package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func joinBoxes(left, right string) string {
	leftLines := strings.Split(left, "\n")
	rightLines := strings.Split(right, "\n")
	leftWidth := 0
	for _, line := range leftLines {
		if w := lipgloss.Width(line); w > leftWidth {
			leftWidth = w
		}
	}
	maxLines := len(leftLines)
	if len(rightLines) > maxLines {
		maxLines = len(rightLines)
	}
	var result strings.Builder
	for i := 0; i < maxLines; i++ {
		leftLine := ""
		rightLine := ""

		if i < len(leftLines) {
			leftLine = leftLines[i]
		}

		if i < len(rightLines) {
			rightLine = rightLines[i]
		}
		padding := leftWidth - lipgloss.Width(leftLine)
		if padding > 0 {
			leftLine += strings.Repeat(" ", padding)
		}

		result.WriteString(leftLine)
		result.WriteString("  ")
		result.WriteString(rightLine)

		if i < maxLines-1 {
			result.WriteString("\n")
		}
	}

	return result.String()
}

func box(content []string, width int) string {
	if width < 4 {
		width = 4
	}
	var s strings.Builder
	s.WriteString("┌")
	s.WriteString(strings.Repeat("─", width-2))
	s.WriteString("┐\n")
	for _, line := range content {
		linew := lipgloss.Width(line)

		if linew > width-4 {
			line = lipgloss.NewStyle().
				MaxWidth(width - 4).
				Render(line)

			linew = lipgloss.Width(line)
		}

		padding := width - 4 - linew

		s.WriteString("│ ")
		s.WriteString(line)

		if padding > 0 {
			s.WriteString(strings.Repeat(" ", padding))
		}

		s.WriteString(" │\n")
	}
	s.WriteString("└")
	s.WriteString(strings.Repeat("─", width-2))
	s.WriteString("┘")
	return s.String()
}
