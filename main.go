package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
)

const mAppName = "Perfect Track Log Interpreter"

func main() {
	// Create a new application instance
	mApp := app.New()
	var mAppColors AppColors = init_AppColors()
	mApp.Settings().SetTheme(newMcsTheme(mAppColors))

	var mSyncConfig SyncConfig = init_SyncConfig()
	var mSyncControlWidgets SyncControlWidgets = init_SyncControlWidgets()
	var mSyncStatus SyncStatus = init_SyncStatus()

	//Build Window
	mWindow := buildWindow(mApp, mAppName, mSyncStatus, mSyncConfig, mAppColors, mSyncControlWidgets)

	//Initialize Ui Data with default statuses
	mSyncStatus.resetSyncStatus(mSyncControlWidgets, mSyncConfig)

	// Show the window
	mWindow.ShowAndRun()
}

func startSync(mSyncStatus SyncStatus, mSyncControlWidgets *SyncControlWidgets, mSyncConfig SyncConfig) {
	//Change Sync State
	mSyncStatus.setSyncState(SYNC_STATE_SEARCHING_FOR_LOG_FILE, mSyncControlWidgets, mSyncConfig)

	//Begin sync routine
	// ...
}

func stopSync(mSyncStatus SyncStatus, mSyncControlWidgets *SyncControlWidgets, mSyncConfig SyncConfig) {
	// var intVal = mSyncStatus.syncStateID.Get()
	syncStateID, err := mSyncStatus.syncStateID.Get()
	if err != nil {
		fmt.Println("Sync ID Status Get() Err: ", err)
	}

	switch syncStateID {
	case SYNC_STATE_SEARCHING_FOR_LOG_FILE:
		//TODO: End log file search / sync routine
		//Don't worry about sending anything since we never even started

		//Update sync state to IDLE
		mSyncStatus.setSyncState(SYNC_STATE_IDLE, mSyncControlWidgets, mSyncConfig)

	case SYNC_STATE_IN_SYNC:
		//Update sync state to finishing state
		mSyncStatus.setSyncState(SYNC_STATE_FINISHING_SYNC, mSyncControlWidgets, mSyncConfig)

		//TODO: End log file sync routine
		//TODO: Send last transmission

		//Update sync state to idle state
		mSyncStatus.setSyncState(SYNC_STATE_IDLE, mSyncControlWidgets, mSyncConfig)

	case SYNC_STATE_FINISHING_SYNC:
		//Do nothing, we're already in the finishing state

	// case SYNC_STATE_IDLE:
	default:
		//Do nothing, we're not running
		//I guess we should already be in an idle state, so let's put the app there
		mSyncStatus.setSyncState(SYNC_STATE_IDLE, mSyncControlWidgets, mSyncConfig)
	}
}

func showDialog_startSyncError(errorMsg string) {
	mWarningDialog := dialog.NewInformation("Error", errorMsg, fyne.CurrentApp().Driver().AllWindows()[0])
	mWarningDialog.Show()
}
