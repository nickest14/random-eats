package apps

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowRandom(w fyne.Window) fyne.CanvasObject {
	resultLabel := widget.NewLabel("等待選擇...")

	return container.NewVBox(
		widget.NewButton("開始隨機", func() {
			// TODO: Implement random selection logic
			resultLabel.SetText("隨機選擇的餐廳是: XXX")
		}),
		resultLabel,
	)
}
