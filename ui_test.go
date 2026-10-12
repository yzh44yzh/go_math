package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"strconv"
	"strings"
	"testing"
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
	u.from.SetText("2")
	u.to.SetText("2")
	u.count.SetText("20")
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
	if u.practice != nil || u.from.Text != "2" || u.count.Text != "20" {
		t.Fatal("new session did not retain settings")
	}
}

func TestUISettingsAndDiscard(t *testing.T) {
	u := setupUI(t)
	u.from.SetText("9")
	u.to.SetText("1")
	test.Tap(u.primary)
	if u.practice != nil || u.errorLabel.Text == "" {
		t.Fatal("invalid settings started practice")
	}
	u.from.SetText("1")
	u.to.SetText("9")
	u.count.SetText("100")
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
	if u.practice != nil || u.count.Text != "100" {
		t.Fatal("confirmed navigation did not retain settings")
	}
}

func TestUIManualAdvancementAndFocus(t *testing.T) {
	u := setupUI(t)
	u.count.SetText("2")
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
