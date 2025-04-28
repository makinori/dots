package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func radioListDialog(options []string) int {
	a := app.New()

	w := a.NewWindow("Hello")
	w.Resize(fyne.Size{Width: 300, Height: 350})

	var selected int

	w.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		switch event.Name {
		case "Escape":
			selected = -1
			w.Close()
		case "Return":
			w.Close()
		}
	})

	list := widget.NewList(
		func() int { return len(options) },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(options[i])
		},
	)

	list.OnSelected = func(i widget.ListItemID) {
		selected = i
	}

	listScroll := container.NewVScroll(list)

	cancelButton := widget.NewButtonWithIcon(
		"Cancel", theme.CancelIcon(),
		func() {
			selected = -1
			w.Close()
		},
	)

	okButton := widget.NewButtonWithIcon(
		"Ok", theme.ConfirmIcon(),
		func() {
			w.Close()
		},
	)

	label := container.NewCenter(
		widget.NewLabel("Select an album to upload to"),
	)

	buttons := container.NewBorder(nil, nil, nil, container.NewHBox(
		cancelButton,
		okButton,
	))

	box := container.NewBorder(
		label,
		buttons,
		nil,
		nil,
		listScroll,
	)

	w.SetContent(box)

	w.ShowAndRun()

	return selected
}
