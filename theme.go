//go:generate fyne bundle -o bundled.go  res

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

//Theme set up: https://www.youtube.com/watch?v=tYHD7OU9xfM

type mcsTheme struct {
	fyne.Theme
}

func newMcsTheme() fyne.Theme {
	return &mcsTheme{Theme: theme.DefaultTheme()}
}
