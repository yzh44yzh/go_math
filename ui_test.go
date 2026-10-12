package main

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
)

func setupUI(t *testing.T) *practiceUI {
	t.Helper()
	a := test.NewApp()
	w := a.NewWindow("Go Math")
	w.Resize(fyne.NewSize(560, 600))
	t.Cleanup(func() { w.Close(); a.Quit() })
	return newPracticeUI(w, rand.New(rand.NewSource(1)))
}

func TestUIPracticeAndRetry(t *testing.T) {
	u := setupUI(t)
	u.from.SetValue(2)
	u.to.SetValue(2)
	u.count.SetValue(20)
	if !strings.Contains(u.capLabel.Text, "1 question.") {
		t.Fatalf("missing cap: %q", u.capLabel.Text)
	}
	test.Tap(u.primary)
	if u.practice == nil || len(u.practice.pairs) != 1 || u.window.Canvas().Focused() != u.answer {
		t.Fatal("practice did not start focused on answer")
	}
	test.Tap(u.primary)
	if u.errorLabel.Text != "Enter a whole number" || u.practice.state != questionPhase {
		t.Fatal("blank input did not preserve question")
	}
	u.answer.SetText("3")
	u.answer.TypedKey(&fyne.KeyEvent{Name: fyne.KeyReturn})
	if u.practice.state != feedbackPhase || len(u.practice.mistakes) != 1 || u.window.Canvas().Focused() != u.primary {
		t.Fatal("answer did not show focused feedback")
	}
	u.primary.TypedKey(&fyne.KeyEvent{Name: fyne.KeyReturn})
	if u.practice.state != resultsPhase || u.primary.Text != "Retry mistakes" {
		t.Fatal("final feedback did not reach results")
	}
	test.Tap(u.primary)
	if u.practice.round != 1 || u.practice.state != questionPhase {
		t.Fatal("retry action failed")
	}
	test.Type(u.answer, "4")
	test.Tap(u.primary)
	test.Tap(u.primary)
	if u.primary.Text != "New session" || u.practice.correct != 1 {
		t.Fatal("correct retry did not finish")
	}
	test.Tap(u.primary)
	if u.practice != nil || u.from.Value != 2 || u.count.Value != 20 {
		t.Fatal("new session did not retain settings")
	}
}

func TestUISettingsAndDiscard(t *testing.T) {
	u := setupUI(t)
	u.from.SetValue(9)
	u.to.SetValue(1)
	if u.from.Value != 1 || u.to.Value != 1 {
		t.Fatal("crossing factor sliders did not preserve a valid range")
	}
	u.from.SetValue(1)
	u.to.SetValue(9)
	u.count.SetValue(100)
	if !strings.Contains(u.capLabel.Text, "81") {
		t.Fatal("full range cap not explained")
	}
	test.Tap(u.primary)
	u.answer.SetText("123")
	before := u.practice
	test.Tap(u.back)
	if u.confirmation == nil {
		t.Fatal("discard confirmation not shown")
	}
	u.window.Canvas().Focused().TypedKey(&fyne.KeyEvent{Name: fyne.KeyReturn})
	if u.practice != before || u.answer.Text != "123" || u.window.Canvas().Focused() != u.answer {
		t.Fatal("cancelled navigation changed progress/focus")
	}
	test.Tap(u.back)
	u.discard.TypedKey(&fyne.KeyEvent{Name: fyne.KeyReturn})
	if u.practice != nil || u.count.Value != 100 {
		t.Fatal("confirmed navigation did not retain settings")
	}
}

func TestSettingsSlidersAndKeyboard(t *testing.T) {
	u := setupUI(t)
	if u.selectedSettings() != defaultSettings {
		t.Fatal("slider defaults differ from practice defaults")
	}
	u.to.SetValue(3)
	u.from.SetValue(7)
	if u.to.Value != 7 || u.toLabel.Text != "Maximum factor: 7" {
		t.Fatal("raising minimum did not raise maximum and its label")
	}
	u.to.SetValue(2)
	if u.from.Value != 2 || u.fromLabel.Text != "Minimum factor: 2" {
		t.Fatal("lowering maximum did not lower minimum and its label")
	}
	u.from.TypedKey(&fyne.KeyEvent{Name: fyne.KeyRight})
	if u.from.Value != 3 || u.to.Value != 3 {
		t.Fatal("arrow adjustment must move by one and preserve the range")
	}
	u.count.SetValue(27.7)
	if u.count.Value != 28 || u.countLabel.Text != "Questions: 28" {
		t.Fatal("question slider did not round to whole numbers and update its label")
	}
	u.count.SetValue(101)
	if u.count.Value != 100 {
		t.Fatal("question slider exceeded its upper bound")
	}
	u.count.SetValue(0)
	if u.count.Value != 1 {
		t.Fatal("question slider exceeded its lower bound")
	}
	for _, slider := range []*widget.Slider{u.from, u.to, u.count} {
		u.window.Canvas().Focus(slider)
		slider.TypedKey(&fyne.KeyEvent{Name: fyne.KeyReturn})
		if u.practice != nil {
			t.Fatal("Enter on a slider started practice")
		}
	}
	u.window.Canvas().Focus(u.from)
	for _, expected := range []fyne.Focusable{u.to, u.count, u.primary} {
		u.window.Canvas().FocusNext()
		if u.window.Canvas().Focused() != expected {
			t.Fatal("unexpected settings Tab order")
		}
	}
	u.primary.TypedKey(&fyne.KeyEvent{Name: fyne.KeyReturn})
	if u.practice == nil {
		t.Fatal("Enter on Start practice did not start practice")
	}
}

func TestUIManualAdvancementAndFocus(t *testing.T) {
	u := setupUI(t)
	u.count.SetValue(2)
	test.Tap(u.primary)
	q := u.practice.current()
	u.answer.SetText(strconv.Itoa(q.product()))
	check := u.primary
	test.Tap(check)
	test.Tap(check)
	if u.practice.index != 0 || u.practice.correct != 1 {
		t.Fatal("duplicate event advanced or double-counted")
	}
	if u.primary.Text != "Next" {
		t.Fatal("wrong feedback action")
	}
	test.Tap(u.primary)
	if u.practice.index != 1 || u.practice.state != questionPhase || u.window.Canvas().Focused() != u.answer {
		t.Fatal("next question did not restore answer focus")
	}
	u.window.Canvas().FocusNext()
	if u.window.Canvas().Focused() != u.primary {
		t.Fatal("Tab traversal did not reach Check answer")
	}
	u.window.Canvas().FocusPrevious()
	if u.window.Canvas().Focused() != u.answer {
		t.Fatal("reverse Tab traversal did not restore answer entry")
	}
}

func TestActionButtonKeyboard(t *testing.T) {
	setupUI(t)
	for _, key := range []fyne.KeyName{fyne.KeyReturn, fyne.KeyEnter, fyne.KeySpace} {
		calls := 0
		b := newActionButton("Next", func() { calls++ })
		b.TypedKey(&fyne.KeyEvent{Name: key})
		if calls != 1 {
			t.Fatalf("key %s invoked action %d times", key, calls)
		}
		b.Disable()
		b.TypedKey(&fyne.KeyEvent{Name: key})
		if calls != 1 {
			t.Fatal("disabled button activated")
		}
	}
}

func TestFocusLayoutFitsNarrowWindow(t *testing.T) {
	setupUI(t)
	content := widget.NewLabel("Question")
	column := fyne.NewContainerWithLayout(focusColumnLayout{}, content)
	for _, width := range []float32{320, 560} {
		column.Resize(fyne.NewSize(width, 200))
		if content.Position().X < 0 || content.Position().X+content.Size().Width > width {
			t.Fatal("column extends beyond window")
		}
	}
}
