package main

import (
	// "fmt"

	"fyne.io/fyne/v2/app"
)

const mAppName = "Perfect Track Log Interpreter"

func main() {
	// fmt.Println("Launching " + mAppName + "...")

	// Create a new application instance
	mApp := app.New()
	mApp.Settings().SetTheme(newMcsTheme())

	//Build Window
	mWindow := buildWindow(mApp, mAppName)

	// Show the window
	mWindow.ShowAndRun()

	// fmt.Println("Terminating " + mAppName + "...")
}
