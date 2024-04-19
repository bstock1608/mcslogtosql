package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const SYNC_STATE_IDLE int = 1
const SYNC_STATE_SEARCHING_FOR_LOG_FILE int = 2
const SYNC_STATE_IN_SYNC int = 3
const SYNC_STATE_FINISHING_SYNC int = 4

type SyncStatus struct {
	syncStateID      binding.Int
	syncState        binding.String
	syncStartTime    binding.String
	lastSync         binding.String
	nextSync         binding.String
	piecesSuccess    binding.String
	piecesFailed     binding.String
	fileBeingWatched binding.String
}

func init_SyncStatus() SyncStatus {
	var mSyncStatus SyncStatus = SyncStatus{
		syncStateID:      binding.NewInt(),
		syncState:        binding.NewString(),
		syncStartTime:    binding.NewString(),
		lastSync:         binding.NewString(),
		nextSync:         binding.NewString(),
		piecesSuccess:    binding.NewString(),
		fileBeingWatched: binding.NewString(),
		piecesFailed:     binding.NewString()}

	return mSyncStatus
}

func (sySt SyncStatus) resetSyncStatus(mSyncControlWidgets SyncControlWidgets, mSyncConfig SyncConfig) {
	sySt.syncStateID.Set(SYNC_STATE_IDLE)
	sySt.syncState.Set("Not Running") //Sync Running: Running || Not Running
	sySt.syncStartTime.Set("-")       //Start Time: - || 04/19/2024 08:34 am
	sySt.lastSync.Set("-")            //Last Sync: - || 04/19/2024 08:34 am
	sySt.nextSync.Set("-")            //Next Sync: - || 54s ... 53s ... 52s ...
	sySt.piecesSuccess.Set("-")       //Successful Pieces: - || 3245
	sySt.piecesFailed.Set("-")        //Failed Pieces: - || 3245
	sySt.fileBeingWatched.Set("-")    //Log File Selected: - || ../../myLog.txt
}

/************************************************
*  CONTROL - Changing sync state is basically our UI state controller
***********************************************/
func (sySt SyncStatus) setSyncState(syncStateID int, mSyncControlWidgets *SyncControlWidgets, mSyncConfig SyncConfig) {

	switch syncStateID {

	case SYNC_STATE_SEARCHING_FOR_LOG_FILE:
		sySt.syncStateID.Set(SYNC_STATE_SEARCHING_FOR_LOG_FILE)
		sySt.syncState.Set("Waiting for Log File...")

		//Hide Start Button
		mSyncControlWidgets.BTN_STOP.Show()
		mSyncControlWidgets.BTN_START.Hide()

	case SYNC_STATE_IN_SYNC:
		sySt.syncStateID.Set(SYNC_STATE_IN_SYNC)
		sySt.syncState.Set("Sync In Progress...")

		//Hide Start Button
		mSyncControlWidgets.BTN_STOP.Show()
		mSyncControlWidgets.BTN_START.Hide()

	case SYNC_STATE_FINISHING_SYNC:
		sySt.syncStateID.Set(SYNC_STATE_FINISHING_SYNC)
		sySt.syncState.Set("Finishing Sync...")

		//Hide Start Button
		mSyncControlWidgets.BTN_STOP.Show()
		mSyncControlWidgets.BTN_START.Hide()

	// case SYNC_STATE_IDLE:
	default:
		sySt.syncStateID.Set(SYNC_STATE_IDLE)
		sySt.syncState.Set("Not Running")

		//Show Start Button
		mSyncControlWidgets.BTN_STOP.Hide()
		mSyncControlWidgets.BTN_START.Show()

		//Do stuff
	}
}

// func (sySt SyncStatus) setSyncStartTime(newState string) {
// 	sySt.syncStartTime.Set(newState)
// }
// func (sySt SyncStatus) setLastSyncTime(newState string) {
// 	sySt.lastSync.Set(newState)
// }
// func (sySt SyncStatus) setNextSyncTime(newState string) {
// 	sySt.nextSync.Set(newState)
// }
// func (sySt SyncStatus) setPiecesSuccess(newState string) {
// 	sySt.piecesSuccess.Set(newState)
// }
// func (sySt SyncStatus) setPiecesFailed(newState string) {
// 	sySt.piecesFailed.Set(newState)
// }

/************************************************
*  CONTAINER - Sync Status - Construct full details Form
***********************************************/
func buildField_SyncStatus(fieldLabel string, boundWidgetVal binding.String, mSyncControlWidgets SyncControlWidgets, valColor color.RGBA) *fyne.Container {
	label := canvas.NewText("   "+fieldLabel+":", valColor)
	widget_value := widget.NewLabelWithData(boundWidgetVal)
	// item3_widget.Resize(fyne.NewSize(100, item3_widget.MinSize().Height))
	return container.New(layout.NewHBoxLayout(), label, widget_value)
}

func buildForm_SyncStatus(mSyncStatus SyncStatus, mSyncControlWidgets SyncControlWidgets, mAppColors AppColors) fyne.CanvasObject {
	container_fileBeingWatched := buildField_SyncStatus("Log File Selected", mSyncStatus.fileBeingWatched, mSyncControlWidgets, mAppColors.GREY)
	container_syncStatus := buildField_SyncStatus("Sync Status", mSyncStatus.syncState, mSyncControlWidgets, mAppColors.GREY)
	container_syncStartTime := buildField_SyncStatus("Sync Start Time", mSyncStatus.syncStartTime, mSyncControlWidgets, mAppColors.GREY)
	container_lastSync := buildField_SyncStatus("Last Sync Time", mSyncStatus.lastSync, mSyncControlWidgets, mAppColors.GREY)
	container_nextSync := buildField_SyncStatus("Next Sync Time", mSyncStatus.nextSync, mSyncControlWidgets, mAppColors.GREY)
	container_piecesSuccess := buildField_SyncStatus("Successful Pieces", mSyncStatus.piecesSuccess, mSyncControlWidgets, mAppColors.GREY)
	container_piecesFailed := buildField_SyncStatus("Failed Pieces", mSyncStatus.piecesFailed, mSyncControlWidgets, mAppColors.GREY)

	form := container.New(layout.NewVBoxLayout(),
		container_fileBeingWatched,
		container_syncStatus,
		container_syncStartTime,
		container_lastSync,
		container_nextSync,
		container_piecesSuccess,
		container_piecesFailed,
	)

	return form
}
