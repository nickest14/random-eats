package apps

import (
	"fmt"
	"net/url"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/nickest14/random-eats/pkg/db"
)

func ShowManage(w fyne.Window) fyne.CanvasObject {
	return listRestaurants()
}

type restaurantList struct {
	itemsPerPage   int
	currentPage    int
	restaurants    []db.Restaurant
	allRestaurants []db.Restaurant

	searchEntry   *widget.Entry
	prevButton    *widget.Button
	nextButton    *widget.Button
	pageLabel     *widget.Label
	listContent   *fyne.Container
	mainContainer *fyne.Container
}

func newRestaurantList() *restaurantList {
	rl := &restaurantList{
		itemsPerPage:   5,
		currentPage:    0,
		allRestaurants: db.GetAllRestaurants(),
	}
	rl.restaurants = rl.allRestaurants
	rl.initUI()
	return rl
}

func (rl *restaurantList) initUI() {
	rl.searchEntry = widget.NewEntry()
	rl.searchEntry.SetPlaceHolder("搜尋餐廳名稱、標籤或地址...")
	rl.searchEntry.OnSubmitted = rl.handleSearch

	rl.prevButton = widget.NewButton("上一頁", rl.handlePrevPage)
	rl.nextButton = widget.NewButton("下一頁", rl.handleNextPage)
	rl.pageLabel = widget.NewLabel("")

	rl.listContent = container.NewVBox()

	paginationControls := container.NewHBox(
		layout.NewSpacer(),
		container.NewHBox(rl.prevButton, rl.pageLabel, rl.nextButton),
		layout.NewSpacer(),
	)

	clearButton := widget.NewButton("清除", func() {
		rl.searchEntry.SetText("")
		rl.handleSearch("")
	})

	searchBox := container.NewBorder(
		nil, nil, nil, clearButton,
		rl.searchEntry,
	)

	rl.mainContainer = container.NewVBox(
		searchBox,
		layout.NewSpacer(),
		rl.listContent,
		layout.NewSpacer(),
		paginationControls,
	)
}

func (rl *restaurantList) handleSearch(searchText string) {
	rl.restaurants = filterRestaurants(rl.allRestaurants, searchText)
	rl.currentPage = 0
	rl.updateList()
}

func (rl *restaurantList) handlePrevPage() {
	if rl.currentPage > 0 {
		rl.currentPage--
		rl.updateList()
	}
}

func (rl *restaurantList) handleNextPage() {
	totalPages := rl.getTotalPages()
	if rl.currentPage < totalPages-1 {
		rl.currentPage++
		rl.updateList()
	}
}

func (rl *restaurantList) getTotalPages() int {
	return (len(rl.restaurants) + rl.itemsPerPage - 1) / rl.itemsPerPage
}

func (rl *restaurantList) createRestaurantItem(r db.Restaurant) *widget.AccordionItem {
	nameLabel := widget.NewLabel(r.Name)
	nameLabel.TextStyle = fyne.TextStyle{}

	link, err := url.Parse(r.Link)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	deleteBtn := widget.NewButton("刪除", func() {
		dialog.ShowConfirm("確認刪除", "確定要刪除這間餐廳嗎？", func(ok bool) {
			if ok {
				if err := db.DeleteRestaurant(r.ID); err != nil {
					dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
					return
				}
				rl.allRestaurants = db.GetAllRestaurants()
				rl.restaurants = filterRestaurants(rl.allRestaurants, rl.searchEntry.Text)
				rl.updateList()
			}
		}, fyne.CurrentApp().Driver().AllWindows()[0])
	})
	deleteBtn.Importance = widget.LowImportance

	details := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("餐廳地址"),
			widget.NewLabel(r.Address),
		),
		container.NewHBox(
			widget.NewLabel("連結"),
			widget.NewHyperlink("link", link),
		),
		container.NewHBox(
			widget.NewLabel("標籤"),
			widget.NewLabel(r.Tags),
		),
		container.NewHBox(
			widget.NewLabel("備註"),
			widget.NewLabel(r.Memo),
		),
	)

	contentContainer := container.NewBorder(
		nil,
		nil,
		details,
		container.NewCenter(deleteBtn),
	)

	locationText := r.City + r.District
	title := nameLabel.Text + "  -  " + locationText
	return widget.NewAccordionItem(title, contentContainer)
}

func (rl *restaurantList) updateList() {
	start := rl.currentPage * rl.itemsPerPage
	end := start + rl.itemsPerPage
	if end > len(rl.restaurants) {
		end = len(rl.restaurants)
	}

	items := make([]*widget.AccordionItem, end-start)
	for i, r := range rl.restaurants[start:end] {
		items[i] = rl.createRestaurantItem(r)
	}

	accordion := widget.NewAccordion(items...)
	accordion.MultiOpen = true

	totalPages := rl.getTotalPages()
	rl.pageLabel.SetText(fmt.Sprintf("%d / %d", rl.currentPage+1, totalPages))

	rl.prevButton.Enable()
	rl.nextButton.Enable()
	if rl.currentPage == 0 {
		rl.prevButton.Disable()
	}
	if rl.currentPage >= totalPages-1 {
		rl.nextButton.Disable()
	}

	rl.listContent.Objects = []fyne.CanvasObject{accordion}
	rl.listContent.Refresh()
}

func listRestaurants() fyne.CanvasObject {
	rl := newRestaurantList()
	rl.updateList()
	return rl.mainContainer
}

func filterRestaurants(restaurants []db.Restaurant, searchText string) []db.Restaurant {
	if searchText == "" {
		return restaurants
	}

	var filtered []db.Restaurant
	for _, r := range restaurants {
		if strings.Contains(strings.ToLower(r.Name), strings.ToLower(searchText)) ||
			strings.Contains(strings.ToLower(r.Tags), strings.ToLower(searchText)) ||
			strings.Contains(strings.ToLower(r.Memo), strings.ToLower(searchText)) {
			filtered = append(filtered, r)
		}
	}
	return filtered
}
