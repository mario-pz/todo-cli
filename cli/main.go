package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	pages := [3]*tview.Form{
		MakeForm("Tasks", &colors[0]),
		MakeForm("About", &colors[1]),
		MakeForm("Login", &colors[2]),
	}

	flex := FlexMenu(pages)

	counter := 0
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlJ:
			counter = (counter + 1) % 3
		case tcell.KeyCtrlK:
			counter = (counter - 1 + 3) % 3
		}

		pages[counter].SetBorderColor(hoverColor)
		for i, page := range pages {
			if i != counter {
				page.SetBorderColor(colors[i])
			}
		}

		return event
	})

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
