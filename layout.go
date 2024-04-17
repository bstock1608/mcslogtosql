package main

import "fyne.io/fyne/v2"

const minWindowWidth = 300
const minWindowHeight = 300

type mcsLayout struct {
	top, left, right, content fyne.CanvasObject
}

func newMcsLayout(top, left, right, content fyne.CanvasObject) fyne.Layout {
	return &mcsLayout{top: top, left: left, right: right, content: content}
}

func (l *mcsLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	topHeight := l.top.MinSize().Height
	l.top.Resize(fyne.NewSize(size.Width, topHeight))

	l.left.Move(fyne.NewPos(0, topHeight))
	l.left.Resize(fyne.NewSize(0, topHeight))

	//Continue arranging layout ...
}

func (l *mcsLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(minWindowWidth, minWindowHeight)
}
