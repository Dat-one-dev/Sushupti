package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/Dat-one-dev/Sushupti/data"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

func DrawText(img *image.RGBA, text string, x, y, size float64, fontData *opentype.Font, textColor color.Color) {

	face, err := opentype.NewFace(fontData, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return
	}
	defer face.Close()

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(textColor),
		Face: face,
		Dot:  fixed.P(int(x), int(y)),
	}

	d.DrawString(text)
}

func Graph(dailyStats []data.DailyStat) {
	maxSec := 0
	for _, stat := range dailyStats {
		if stat.TotalSeconds > maxSec {
			maxSec = stat.TotalSeconds
		}
	}

	for _, stat := range dailyStats {
		barLength := 0
		if maxSec > 0 {
			barLength = stat.TotalSeconds * 30 / maxSec
		}
		bar := ""
		for i := 0; i < barLength; i++ {
			bar += "█"
		}
		hours := stat.TotalSeconds / 3600
		min := (stat.TotalSeconds % 3600) / 60

		fmt.Printf("%s | %-30s %dh %02dm\n", stat.Date, bar, hours, min)
	}
}

func ExportGraph(dailyStat []data.DailyStat, filename string, darkMode bool) error {
	//Dark Mode
	bgColor := color.White
	textColor := color.Black

	if darkMode {
		bgColor = color.Black
		textColor = color.White
	}

	//VARIABLES
	width := 1000
	barX := 202
	barWidth := 700
	barHeight := 20
	gap := 10
	left := 200
	right := 900
	top := 120
	bottom := top + len(dailyStat)*(barHeight+gap)
	maxSec := 0
	height := bottom + 70

	//font
	fontBytes, err := os.ReadFile("assets/font.ttf")
	if !logError(err) {
		return err
	}

	fontData, err := opentype.Parse(fontBytes)
	if !logError(err) {
		return err
	}

	fontBytes2, err := os.ReadFile("assets/font3.otf")
	if !logError(err) {
		return err
	}

	fontData2, err := opentype.Parse(fontBytes2)
	if !logError(err) {
		return err
	}

	//Image
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{
		C: bgColor,
	}, image.Point{}, draw.Src)
	DrawText(img, "SushupTi", 500-70, 48, 48, fontData2, textColor)
	DrawText(img, "Daily Coding Activity", 400, 80, 24, fontData, textColor)

	for _, stat := range dailyStat {
		if stat.TotalSeconds > maxSec {
			maxSec = stat.TotalSeconds
		}
	}

	draw.Draw(
		img,
		image.Rect(left, top, left+2, bottom),
		&image.Uniform{C: textColor},
		image.Point{},
		draw.Src,
	)

	draw.Draw(
		img,
		image.Rect(left, bottom-2, right, bottom),
		&image.Uniform{C: textColor},
		image.Point{},
		draw.Src,
	)
	maxHours := (maxSec + 3599) / 3600

	if maxHours < 1 {
		maxHours = 1
	}

	for hour := 0; hour <= maxHours; hour++ {
		x := left + hour*(right-left)/maxHours

		DrawText(
			img,
			fmt.Sprintf("%d Hour", hour),
			float64(x-7),
			float64(bottom+20),
			12,
			fontData,
			textColor,
		)
	}
	for i, stat := range dailyStat {
		if maxSec == 0 {
			continue
		}

		currentWidth := stat.TotalSeconds * barWidth / maxSec
		y := top + i*(barHeight+gap)
		DrawText(
			img,
			stat.Date,
			120,
			float64(y+15),
			12,
			fontData,
			textColor,
		)
		draw.Draw(
			img,
			image.Rect(
				barX-1,
				y-1,
				barX+currentWidth+1,
				y+barHeight+1,
			),
			&image.Uniform{
				C: textColor,
			},
			image.Point{},
			draw.Src,
		)

		draw.Draw(
			img,
			image.Rect(
				barX,
				y,
				barX+currentWidth,
				y+barHeight,
			),
			&image.Uniform{
				C: color.RGBA{R: 15, G: 191, B: 62, A: 255},
			},
			image.Point{},
			draw.Src,
		)
	}
	file, err := os.Create(filename)
	if !logError(err) {
		return nil
	}
	defer file.Close()
	return png.Encode(file, img)
}
