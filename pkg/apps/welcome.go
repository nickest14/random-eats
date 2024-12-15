package apps

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowWelcome(w fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("歡迎使用晚餐選擇器"),
		widget.NewLabel("功能介紹:"),
		widget.NewLabel("1. 隨機選擇 - 從已保存的餐廳中隨機挑選"),
		widget.NewLabel("2. 餐廳管理 - 新增或刪除您喜歡的餐廳"),
	)
}
