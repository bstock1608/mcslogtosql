package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

const windowDefaultWidth = 750
const windowDefaultHeight = 500

func buildWindow(mApp fyne.App, mAppName string, mSyncStatus SyncStatus, mSyncConfig SyncConfig, mAppColors AppColors, mSyncControlWidgets SyncControlWidgets) fyne.Window {
	// Create a new window
	mWindow := mApp.NewWindow(mAppName)

	mWindow.SetContent(buildGUI_2(mSyncStatus, mSyncConfig, mAppColors, mSyncControlWidgets))

	//Resize window
	// mWindow.Resize(fyne.NewSize(500, 500))
	mWindow.Resize(fyne.NewSize(windowDefaultWidth, windowDefaultHeight))

	return mWindow
}

func buildGUI_2(mSyncStatus SyncStatus, mSyncConfig SyncConfig, mAppColors AppColors, mSyncControlWidgets SyncControlWidgets) fyne.CanvasObject {
	//Initialize Bound Vars

	// top := widget.NewToolbar(
	// 	// widget.NewToolbarAction(theme.HomeIcon(), func() {}),
	// 	widget.NewToolbarAction(nil, func() {}),
	// )

	section_topLogo := buildSection_topLogo()

	section_syncControl := buildSection_syncControlOuterContainer(mSyncStatus, mSyncConfig, mSyncControlWidgets, mAppColors)

	section_startStopButton := buildSection_controlButtons(mSyncStatus, mSyncControlWidgets, mSyncConfig)

	mSpacer_aboveBtn := canvas.NewImageFromResource(resource24pxspacerPng)
	mSpacer_aboveBtn.FillMode = canvas.ImageFillOriginal

	mSpacer_bottom := canvas.NewImageFromResource(resource50pxspacerPng)
	mSpacer_bottom.FillMode = canvas.ImageFillOriginal

	return container.New(
		layout.NewVBoxLayout(),
		section_topLogo,
		section_syncControl,
		mSpacer_aboveBtn,
		section_startStopButton,
		mSpacer_bottom)

	//Custom Layout
	// objs := []fyne.CanvasObject{content, top, left, right}
	// return container.New(newMcsLayout(top, left, right, content), objs...)
}

func buildSection_syncControlOuterContainer(mSyncStatus SyncStatus, mSyncConfig SyncConfig, mSyncControlWidgets SyncControlWidgets, mAppColors AppColors) fyne.CanvasObject {
	form_logConfigure := buildForm_SyncConfigure(mSyncStatus, mSyncConfig, mSyncControlWidgets, mAppColors)
	form_syncOutput := buildForm_SyncStatus(mSyncStatus, mSyncControlWidgets, mAppColors)

	horizLayoutContainer := container.New(layout.NewGridLayout(2), form_logConfigure, form_syncOutput)
	fullCenterSection := container.New(layout.NewGridLayout(1), horizLayoutContainer)

	return fullCenterSection
}

func buildSection_topLogo() fyne.CanvasObject {
	mSpacer_1 := canvas.NewImageFromResource(resource50pxspacerPng)
	mSpacer_1.FillMode = canvas.ImageFillOriginal

	mLogo := canvas.NewImageFromResource(resourceMcsLogo2019Jpg)
	mLogo.FillMode = canvas.ImageFillOriginal

	logoSubtext := canvas.NewText("Perfect Track Log Interpreter", color.Black)
	logoSubtext.Alignment = fyne.TextAlignCenter

	mSpacer_2 := canvas.NewImageFromResource(resource50pxspacerPng)
	mSpacer_2.FillMode = canvas.ImageFillOriginal

	logoSpacedContainer := container.New(layout.NewVBoxLayout(), mSpacer_1, mLogo, layout.NewSpacer(), logoSubtext, mSpacer_2)

	return logoSpacedContainer
}
