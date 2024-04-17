package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const windowDefaultWidth = 500
const windowDefaultHeight = 500

func buildWindow(mApp fyne.App, mAppName string) fyne.Window {
	// Create a new window
	mWindow := mApp.NewWindow(mAppName)

	mWindow.SetContent(buildGUI())

	//Resize window
	// mWindow.Resize(fyne.NewSize(500, 500))
	mWindow.Resize(fyne.NewSize(windowDefaultWidth, windowDefaultHeight))

	return mWindow
}

func buildGUI() fyne.CanvasObject {

	top := makeTopBanner() //Toolbar

	left := widget.NewLabel("Left")
	right := widget.NewLabel("Right")

	content := widget.NewLabel("Content")
	content.Alignment = fyne.TextAlignCenter

	//OLD
	// // Create a label widget
	// mLabel := widget.NewLabel("MCS Log to SQL Watcher")

	// // Create a button widget
	// mButton := widget.NewButton("Click me", func() {
	// 	// mLabel.SetText("Button Clicked!")
	// 	updateLabel(mLabel)
	// })

	// // Create a container to hold the widgets
	// mContentContainer := container.NewVBox(
	// 	mLabel,
	// 	mButton,
	// )
	// // Set the content of the window
	// mWindow.SetContent(mContentContainer)

	// container.New

	return container.NewBorder(top, nil, left, right, content)

	//Custom Layout
	// objs := []fyne.CanvasObject{content, top, left, right}
	// return container.New(newMcsLayout(top, left, right, content), objs...)
}

// func updateLabel(label *widget.Label) {
// 	label.SetText("Button Clicked!")
// }

func makeTopBanner() fyne.CanvasObject {

	top := widget.NewToolbar(
		// widget.NewToolbarAction(theme.HomeIcon(), func() {}),
		widget.NewToolbarAction(nil, func() {}),
	)

	mLogo := canvas.NewImageFromResource(resourceMcsLogo2019Jpg)
	mLogo.FillMode = canvas.ImageFillContain

	return container.NewStack(top, mLogo)
	// return container.NewStack(top, container.NewPadded(mLogo))
}
