package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	left := newPrimitive("Left")
	//main := newPrimitive("Main content")
	right := newPrimitive("Right")

	main := tview.NewTextView().
		SetTextColor(tcell.ColorYellow).
		SetScrollable(true)

	main.SetChangedFunc(func() {
		if main.HasFocus() {
			app.Draw()
		}
	})

	//main.SetBorder(true).SetTitle("TextView implements io.Writer")

	inputField := tview.NewInputField().
		SetLabel("Query: ").
		SetFieldWidth(200)

	inputField.SetDoneFunc(func(key tcell.Key) {
		//app.Stop()
		//fmt.Printf("finishing and entered: %s\n", inputField.GetText())
		fmt.Fprintf(main,"%s\n", inputField.GetText())
		inputField.SetText("")

	})

	inputField.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		//fmt.Printf("GOT %s\n", inputField.GetText())

		//main.SetText(inputField.GetText())
		//main.Set
		return event
	})

	//footer := newPrimitive("Footer)
	//frame := tview.NewFrame(inputField)

	grid := tview.NewGrid().
		SetRows(3, 0, 3).
		SetColumns(30, 0, 30).
		SetBorders(true).
		AddItem(newPrimitive("DataDog Logger"), 0, 0, 1, 3, 0, 0, false).
		AddItem(inputField, 2, 0, 1, 3, 0, 0, true)

	// Layout for screens wider than 100 cells.
	grid.AddItem(left, 1, 0, 1, 1, 0, 100, false).
		AddItem(main, 1, 1, 1, 1, 0, 100, false).
		AddItem(right, 1, 2, 1, 1, 0, 100, false)

	if err := app.SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}
