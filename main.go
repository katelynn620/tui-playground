package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var idx int = 0

func setBtnFocus(menu *tview.Flex, app *tview.Application) {
	for i := 0; i < menu.GetItemCount(); i++ {
		btn := menu.GetItem(i).(*tview.Button)
		// btn.SetStyle(tcell.StyleDefault.Background(tcell.ColorRed).Foreground(tview.Styles.PrimaryTextColor))
		if idx == i {
			// btn.SetStyle(tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tview.Styles.PrimaryTextColor))
			app.SetFocus(btn)
		}
	}
}
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
	subMenu := tview.NewList().
		AddItem("Option 1", "", '1', nil).
		AddItem("Option 2", "", '2', nil).
		AddItem("Option 3", "", '3', nil)

	// add page title for menu

	subMenu.ShowSecondaryText(false).SetTitle("Menu").SetTitleAlign(tview.AlignCenter).SetBorder(true).
		SetDrawFunc(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x + 2, y + 2, width, height
		})

	pages := tview.NewPages()
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

	subMenu.AddItem("Test", "", 't', func() {
		pages.SwitchToPage("background")
	})
	subMenu.AddItem("Test 2", "", 's', func() {
		pages.SwitchToPage("modal")
	})
	// add frame to the menu
	menuFrame := tview.NewFrame(pages).SetBorders(0, 0, 0, 0, 0, 0)

	// add title to header in frame
	menuFrame.AddText("Application Title", true, tview.AlignLeft, tcell.ColorRed)

	mainMenu := tview.NewFlex().SetDirection(tview.FlexRow)
	btnA := tview.NewButton("A").SetSelectedFunc(func() {
		// menuFrame.Clear().AddText("A", true, tview.AlignLeft, tcell.ColorRed)
		subMenu.SetTitle(fmt.Sprintf(" %v ", "A"))
		pages.SwitchToPage("modal")
		app.SetFocus(subMenu.SetCurrentItem(0))
	}).SetStyle(tcell.StyleDefault.Background(tcell.ColorRed).Foreground(tview.Styles.PrimaryTextColor))
	btnB := tview.NewButton("B").SetSelectedFunc(func() {
		// menuFrame.Clear().AddText("B", true, tview.AlignLeft, tcell.ColorRed)
		subMenu.SetTitle(fmt.Sprintf(" %v ", "B"))
		pages.SwitchToPage("modal")
		app.SetFocus(subMenu.SetCurrentItem(0))
	}).SetStyle(tcell.StyleDefault.Background(tcell.ColorRed).Foreground(tview.Styles.PrimaryTextColor))
	btnC := tview.NewButton("C").SetSelectedFunc(func() {
		// menuFrame.Clear().AddText("C", true, tview.AlignLeft, tcell.ColorRed)
		subMenu.SetTitle(fmt.Sprintf(" %v ", "C"))
		pages.SwitchToPage("modal")
		app.SetFocus(subMenu.SetCurrentItem(0))
	}).SetStyle(tcell.StyleDefault.Background(tcell.ColorRed).Foreground(tview.Styles.PrimaryTextColor))
	btnQ := tview.NewButton("Q").SetSelectedFunc(func() {
		app.Stop()
	}).SetStyle(tcell.StyleDefault.Background(tcell.ColorRed).Foreground(tview.Styles.PrimaryTextColor))

	mainMenu.AddItem(btnA, 0, 1, true)
	mainMenu.AddItem(btnB, 0, 1, false)
	mainMenu.AddItem(btnC, 0, 1, false)
	mainMenu.AddItem(btnQ, 0, 1, false)

	app.SetAfterDrawFunc(func(screen tcell.Screen) {
		menuFrame.Clear().AddText(fmt.Sprintf("%v item(s)", mainMenu.GetItemCount()), true, tview.AlignLeft, tcell.ColorRed)
		// setBtnFocus(mainMenu)
	})

	mainMenu.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyDown {
			if idx < mainMenu.GetItemCount()-1 {
				idx++
			}
		} else if event.Key() == tcell.KeyUp {
			if idx > 0 {
				idx--
			}
		} else if event.Key() == tcell.KeyHome {
			idx = 0
		} else if event.Key() == tcell.KeyEnd {
			idx = mainMenu.GetItemCount() - 1
		} else if event.Key() == tcell.KeyEsc {
			app.Stop()
		}
		setBtnFocus(mainMenu, app)
		menuFrame.Clear().AddText(fmt.Sprintf("%v / %v", event.Key(), idx), true, tview.AlignLeft, tcell.ColorRed)
		return event
	})

	pages.AddPage("background", background, true, true).
		AddPage("modal", modalB(subMenu, 50, 10), true, false).
		AddPage("main", mainMenu, true, true)

	subMenu.AddItem("Back", "", 'b', func() {
		pages.SwitchToPage("main")
		setBtnFocus(mainMenu, app)
	})
	background.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// menuFrame.Clear().AddText(fmt.Sprintf("%v", event.Key()), true, tview.AlignLeft, tcell.ColorRed)
		if event.Key() == tcell.KeyEscape {
			pages.SwitchToPage("modal")
		}
		return nil
	})
	// set the menu as the root of the application
	if err := app.SetRoot(menuFrame, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}

}
