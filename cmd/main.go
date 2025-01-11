package main

import (
	"github.com/nickest14/random-eats/pkg/apps"
	"github.com/nickest14/random-eats/pkg/db"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	DB, err := db.Connect()
	if err != nil {
		panic("error connecting to database: " + err.Error())
	}

	defer db.Close(DB)

	myApp := app.New()
	myWindow := myApp.NewWindow("餐點選擇器")

	var tabs []*container.TabItem
	for _, appInfo := range apps.Apps {
		tabs = append(tabs, container.NewTabItemWithIcon(
			appInfo.Name,
			appInfo.Icon,
			container.NewPadded(appInfo.Show(myWindow)),
		))
	}

	appTabs := container.NewAppTabs(tabs...)
	appTabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(appTabs)
	myWindow.Resize(fyne.NewSize(600, 500))
	myWindow.ShowAndRun()
}
