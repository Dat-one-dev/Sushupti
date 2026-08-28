package tui

import (
	"time"

	"github.com/Dat-one-dev/Sushupti/style"
)

func RenderDate() []string {
	now := time.Now()

	content := []string{
		style.Title.Render("DATE"),
		style.Divider,
		style.Value.Render(now.Format("Monday")),
		style.Value.Render(now.Format("January 02, 2006")),
	}

	return content
}
