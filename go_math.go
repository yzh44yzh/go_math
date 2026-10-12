package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"math/rand"
	"time"
)

func main() {
	a := app.NewWithID("io.github.yzh44yzh.go_math")
	w := a.NewWindow("Go Math")
	w.Resize(fyne.NewSize(560, 600))
	newPracticeUI(w, rand.New(rand.NewSource(time.Now().UnixNano())))
	w.ShowAndRun()
}
