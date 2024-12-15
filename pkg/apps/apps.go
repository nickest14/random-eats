package apps

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type AppInfo struct {
	Name string
	Icon fyne.Resource
	Show func(w fyne.Window) fyne.CanvasObject
}

var Apps = []AppInfo{
	{
		Name: "主頁",
		Icon: theme.HomeIcon(),
		Show: ShowWelcome,
	},
	{
		Name: "隨機選擇",
		Icon: theme.ViewRefreshIcon(),
		Show: ShowRandom,
	},
	{
		Name: "餐廳管理",
		Icon: theme.ListIcon(),
		Show: ShowManage,
	},
}
