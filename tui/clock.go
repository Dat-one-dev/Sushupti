package tui

import (
	"strings"
	"time"

	"github.com/Dat-one-dev/Sushupti/style"
	"github.com/common-nighthawk/go-figure"
)

func RenderClock() []string {
	time := time.Now().Format("15:04")

	text := figure.NewFigure(time, "small", true)

	content := []string{
		style.Title.Render("TIME"),
		style.Divider,
	}

	for _, line := range strings.Split(text.String(), "\n") {
		content = append(content, style.Value.Render(line))
	}

	return content
}
