package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func MakeForm(title string, backgroundColor *tcell.Color) *tview.Form {
	page := tview.NewForm()
	page.SetBorder(true).SetTitle(title)
	page.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor).
		SetBorderColor(*backgroundColor)
	return page
}

func FlexMenu(pages [3]*tview.Form) *tview.Flex {
	return tview.NewFlex().
		AddItem(pages[0], 0, 2, true).
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(pages[1], 0, 1, false).
			AddItem(pages[2], 0, 1, false), 0, 1, false)
}
