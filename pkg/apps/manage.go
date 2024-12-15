package apps

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var restaurants []string

func ShowManage(w fyne.Window) fyne.CanvasObject {
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("輸入餐廳名稱")

	restaurantList := widget.NewList(
		func() int {
			return len(restaurants)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("餐廳")
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(restaurants[id])
		},
	)

	return container.NewBorder(
		container.NewVBox(
			nameEntry,
			widget.NewButton("新增餐廳", func() {
				if name := nameEntry.Text; name != "" {
					restaurants = append(restaurants, name)
					nameEntry.SetText("")
					restaurantList.Refresh()
				}
			}),
		),
		widget.NewButton("刪除", func() {
			// TODO: Implement delete functionality
		}),
		nil, nil,
		restaurantList,
	)
}
