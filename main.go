package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	// create a new application
	app := tview.NewApplication()
	// var w, h int

	// modalA := func(p tview.Primitive, width, height int) tview.Primitive {
	// 	return tview.NewGrid().
	// 		SetColumns(0, width, 0).
	// 		SetRows(0, height, 0).
	// 		AddItem(p, 1, 1, 1, 1, 0, 0, true)
	// }

	modalB := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}

	background := tview.NewTextView().
		SetTextColor(tcell.ColorBlue).
		SetText(strings.Repeat("background ", 1000))
	// create a new menu
	menu := tview.NewList().
		AddItem("Option 1", "", '1', nil).
		AddItem("Option 2", "", '2', nil).
		AddItem("Option 3", "", '3', nil).
		AddItem("Quit", "", 'q', func() {
			app.Stop()
		})

	// add page title for menu
	menu.SetTitle("Menu").SetTitleAlign(tview.AlignCenter).SetBorder(true)

	pages := tview.NewPages().
		AddPage("background", background, true, true).
		AddPage("modal", modalB(menu, 50, 10), true, true)
	// app.SetAfterDrawFunc(func(screen tcell.Screen) {
	// 	w, h = screen.Size()
	// 	// log.Printf("%v %v", w, h)
	// 	// oriX, oriY, oriW, oriH := menu.GetRect()
	// 	// log.Printf("%v %v %v %v ", oriX, oriY, oriW, oriH)
	// 	newW := (w / 3) * 2
	// 	newH := (h / 3) * 2
	// 	newX := (w - newW) / 2
	// 	newY := (h - newH) / 2
	// 	menu.SetRect(newX, newY, 150, 150)
	// 	// log.Printf("%v %v %v %v ", newX, newY, newW, newH)
	// })
	// add frame to the menu
	menuFrame := tview.NewFrame(pages).SetBorders(0, 0, 0, 0, 0, 0)

	// add title to header in frame
	menuFrame.AddText("Application Title", true, tview.AlignLeft, tcell.ColorRed)

	// set the menu as the root of the application
	if err := app.SetRoot(menuFrame, true).Run(); err != nil {
		panic(err)
	}

}
