package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

/************************************************
*  CONTAINER - Sync Configure - Construct full Form
***********************************************/
func buildForm_SyncConfigure(mSyncStatus SyncStatus, mSyncConfig SyncConfig, mSyncControlWidgets SyncControlWidgets, mAppColors AppColors) fyne.CanvasObject {
	var cont_selectSyncInterval *fyne.Container = mSyncControlWidgets.init_SelectSyncInterval(mSyncStatus, mSyncConfig)
	var cont_selectExportType *fyne.Container = mSyncControlWidgets.init_SelectSyncExportType(mSyncStatus, mSyncConfig)
	var cont_selectExportMethod *fyne.Container = mSyncControlWidgets.init_SelectExportMethod(mSyncStatus, mSyncConfig)
	var cont_selectFolderToWatch *fyne.Container = mSyncControlWidgets.init_SelectFolderToWatchButton(mSyncStatus, mSyncConfig)
	form := container.New(layout.NewVBoxLayout(), cont_selectFolderToWatch, cont_selectSyncInterval, cont_selectExportType, cont_selectExportMethod)
	return form
}

/************************************************
*  CONTAINER - Sync Control Start/Stop - Construct full Form
***********************************************/
func buildSection_controlButtons(mSyncStatus SyncStatus, mSyncControlWidgets SyncControlWidgets, mSyncConfig SyncConfig) fyne.CanvasObject {
	mSyncControlWidgets.init_StartButton(mSyncStatus, mSyncConfig)
	mSyncControlWidgets.init_StopButton(mSyncStatus, mSyncConfig)
	// return container.New(layout.NewVBoxLayout(), mSyncControlWidgets.BTN_START, mSyncControlWidgets.BTN_STOP)
	// return container.New(layout.NewGridLayout(4), layout.NewSpacer(), mSyncControlWidgets.BTN_START, mSyncControlWidgets.BTN_STOP, layout.NewSpacer())

	var innerContainer *fyne.Container = container.New(layout.NewStackLayout(), mSyncControlWidgets.BTN_START, mSyncControlWidgets.BTN_STOP)

	return container.New(layout.NewGridLayout(3), layout.NewSpacer(), innerContainer, layout.NewSpacer())
}

/************************************************
*  SyncConfig struct
***********************************************/
type SyncConfig struct {
	syncInterval   int16
	logFileToWatch string
	exportType     string
	exportTarget   string
}

func init_SyncConfig() SyncConfig {
	var mSyncConfig SyncConfig = SyncConfig{
		syncInterval:   5,
		logFileToWatch: "",
		exportType:     "SQL Table",
		exportTarget:   "Send via HTTP POST",
	}
	return mSyncConfig
}

// func (sc *SyncConfig) reset(scw SyncControlWidgets) {
func (sc *SyncConfig) reset() {
	sc.syncInterval = 5
	sc.logFileToWatch = ""
	sc.exportType = "SQL Table"
	sc.exportTarget = "Send via HTTP POST"
}

/************************************************
*  SyncControlWidgets struct
***********************************************/
type SyncControlWidgets struct {
	BTN_START            *widget.Button
	BTN_STOP             *widget.Button
	SELECT_SYNC_INTERVAL *widget.Select
	SELECT_EXPORT_TYPE   *widget.Select
	SELECT_EXPORT_METHOD *widget.Select
	BTN_FILE_SELECT      *widget.Button
}

func init_SyncControlWidgets() SyncControlWidgets {
	var mSyncControlWidgets SyncControlWidgets = SyncControlWidgets{
		BTN_START:            new(widget.Button),
		BTN_STOP:             new(widget.Button),
		SELECT_SYNC_INTERVAL: new(widget.Select),
		SELECT_EXPORT_TYPE:   new(widget.Select),
		SELECT_EXPORT_METHOD: new(widget.Select),
		BTN_FILE_SELECT:      new(widget.Button),
	}
	return mSyncControlWidgets
}

/************************************************
*  BTN - Stop Sync - Setup and Click Handling
***********************************************/
func (scw *SyncControlWidgets) init_SelectFolderToWatchButton(mSyncStatus SyncStatus, mSyncConfig SyncConfig) *fyne.Container {
	var label *widget.Label = widget.NewLabel("Select Log Folder:")
	label.Alignment = fyne.TextAlignTrailing
	scw.BTN_FILE_SELECT = widget.NewButtonWithIcon("Select", theme.FolderIcon(), func() {
		scw.onClick_SelectFileToWatch(mSyncStatus, mSyncConfig)
	})
	// scw.BTN_FILE_SELECT.Importance = widget.HighImportance //Colors the button Blue
	var full_container *fyne.Container = container.New(layout.NewGridLayout(2), label, scw.BTN_FILE_SELECT)
	return full_container
}

func (scw *SyncControlWidgets) onClick_SelectFileToWatch(mSyncStatus SyncStatus, mSyncConfig SyncConfig) {
	fmt.Println("File Select Clicked")
}

/************************************************
*  SELECT - Sync Export Type - Setup and Change Handling
***********************************************/
func (scw *SyncControlWidgets) init_SelectExportMethod(mSyncStatus SyncStatus, mSyncConfig SyncConfig) *fyne.Container {
	var label *widget.Label = widget.NewLabel("Export Method:")
	label.Alignment = fyne.TextAlignTrailing
	scw.SELECT_EXPORT_METHOD = widget.NewSelect([]string{"Send via HTTP POST", "Email"}, func(newValue string) {
		scw.onChange_SyncExportMethod(mSyncConfig, newValue)
	})
	scw.SELECT_EXPORT_METHOD.SetSelected("Send via HTTP POST")
	scw.SELECT_EXPORT_METHOD.Disable()
	var full_container *fyne.Container = container.New(layout.NewGridLayout(2), label, scw.SELECT_EXPORT_METHOD)
	return full_container
}

func (scw *SyncControlWidgets) onChange_SyncExportMethod(mSyncConfig SyncConfig, newValue string) {
	switch {
	case newValue == "Email":
		mSyncConfig.exportTarget = "Email"
	default:
		mSyncConfig.exportTarget = "Send via HTTP POST"
	}
	// fmt.Printf("Export Method set to %v\n", mSyncConfig.exportType)
}

/************************************************
*  SELECT - Sync Export Type - Setup and Change Handling
***********************************************/
func (scw *SyncControlWidgets) init_SelectSyncExportType(mSyncStatus SyncStatus, mSyncConfig SyncConfig) *fyne.Container {
	var label *widget.Label = widget.NewLabel("Export Format:")
	label.Alignment = fyne.TextAlignTrailing
	scw.SELECT_EXPORT_TYPE = widget.NewSelect([]string{"SQL Table", "JSON"}, func(newValue string) {
		scw.onChange_SyncExportType(mSyncConfig, newValue)
	})
	scw.SELECT_EXPORT_TYPE.SetSelected("SQL Table")
	scw.SELECT_EXPORT_TYPE.Disable()
	var full_container *fyne.Container = container.New(layout.NewGridLayout(2), label, scw.SELECT_EXPORT_TYPE)
	return full_container
}

func (scw *SyncControlWidgets) onChange_SyncExportType(mSyncConfig SyncConfig, newValue string) {
	switch {
	case newValue == "JSON":
		mSyncConfig.exportType = "JSON"
	default:
		mSyncConfig.exportType = "SQL Table"
	}
	// fmt.Printf("Export Type set to %v\n", mSyncConfig.exportType)
}

/************************************************
*  SELECT - Sync Interval - Setup and Change Handling
***********************************************/
func (scw *SyncControlWidgets) init_SelectSyncInterval(mSyncStatus SyncStatus, mSyncConfig SyncConfig) *fyne.Container {
	var label *widget.Label = widget.NewLabel("Sync Interval:")
	label.Alignment = fyne.TextAlignTrailing
	scw.SELECT_SYNC_INTERVAL = widget.NewSelect([]string{"1 min", "5 min", "10 min"}, func(newValue string) {
		scw.onChange_SyncInterval(mSyncConfig, newValue)
	})
	scw.SELECT_SYNC_INTERVAL.SetSelected("5 min")
	var full_container *fyne.Container = container.New(layout.NewGridLayout(2), label, scw.SELECT_SYNC_INTERVAL)
	return full_container
}

func (scw *SyncControlWidgets) onChange_SyncInterval(mSyncConfig SyncConfig, newValue string) {
	switch {
	case newValue == "1 min":
		mSyncConfig.syncInterval = 1
	case newValue == "5 min":
		mSyncConfig.syncInterval = 5
	// case newValue == "10 min":
	default:
		mSyncConfig.syncInterval = 10
	}

	fmt.Printf("Select set to %v minutes\n", mSyncConfig.syncInterval)
}

/************************************************
*  BTN Start Sync Setup and Click Handling
***********************************************/
func (scw *SyncControlWidgets) init_StartButton(mSyncStatus SyncStatus, mSyncConfig SyncConfig) {
	scw.BTN_START = widget.NewButtonWithIcon("Start", theme.MediaPlayIcon(), func() {

		scw.onClick_StartButton(mSyncStatus, mSyncConfig)
	})
	scw.BTN_START.Importance = widget.HighImportance //Colors the button Blue

	// btnWithBkg := container.NewStack(canvas.NewRectangle(COLOR_PURPLE), btn_startStop)

	// blue := color.NRGBA{R: 52, G: 152, B: 219, A: 255}
	// white := color.White

	// buttonStyle := &fyne.buttonStyle{
	// 	BackgroundColor: &blue,
	// 	TextColor:       &white,
	// }
	// btn_startStop.Renderer = buttonStyle.Renderer(btn_startStop)
}

func (scw *SyncControlWidgets) onClick_StartButton(mSyncStatus SyncStatus, mSyncConfig SyncConfig) {
	// Switch to theme.MediaStopIcon()
	fmt.Println("Start Sync ...")

	//TODO: Verify if folder has been selected
	var folderSelected bool = false

	if folderSelected {
		startSync(mSyncStatus, scw, mSyncConfig)
	} else {
		showDialog_startSyncError("Please select the folder which will contain the Perfect Track log file.")
	}
}

/************************************************
*  BTN - Stop Sync - Setup and Click Handling
***********************************************/
func (scw *SyncControlWidgets) init_StopButton(mSyncStatus SyncStatus, mSyncConfig SyncConfig) {
	scw.BTN_STOP = widget.NewButtonWithIcon("Stop", theme.MediaStopIcon(), func() {
		scw.onClick_StopButton(mSyncStatus, mSyncConfig)
	})
	scw.BTN_STOP.Importance = widget.DangerImportance //Colors the button Blue
	scw.BTN_STOP.Hide()                               //Initial state has Stop button hidden
}

func (scw *SyncControlWidgets) onClick_StopButton(mSyncStatus SyncStatus, mSyncConfig SyncConfig) {
	fmt.Println("Stop Sync ...")
	stopSync(mSyncStatus, scw, mSyncConfig)
}
