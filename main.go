package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Device Name")

	name := widget.NewLabel("(tap lookup)")
	w.SetContent(container.NewVBox(name,
		widget.NewButton("Lookup", func() {
			name.SetText(lookupName())
		})))

	w.ShowAndRun()
}
