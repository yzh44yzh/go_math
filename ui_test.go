package main

import (
	"math/rand"
	"strconv"
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
	test.Tap(u.primary)
	if u.practice == nil || len(u.practice.pairs) != 20 || u.window.Canvas().Focused() != u.primary {
		t.Fatal("practice did not start focused on answer")
	}
	test.Tap(u.primary)
	if u.errorLabel.Text != "Enter a whole number" || u.practice.state != questionPhase {
		t.Fatal("blank input did not preserve question")
	}
	test.Tap(u.digits[3])
	u.primary.TypedKey(&fyne.KeyEvent{Name: fyne.KeyReturn})
	if u.practice.state != feedbackPhase || len(u.practice.mistakes) != 1 || u.window.Canvas().Focused() != u.primary {
		t.Fatal("answer did not show focused feedback")
	}
	u.primary.TypedKey(&fyne.KeyEvent{Name: fyne.KeyReturn})
	for index := 1; index < 20; index++ {
		if u.practice.state != questionPhase || u.practice.index != index {
			t.Fatal("practice ended before the configured question count")
		}
		test.Type(u.primary, "4")
		test.Tap(u.primary)
		test.Tap(u.primary)
	}
	if u.practice.state != resultsPhase || u.primary.Text != "Retry mistakes" {
		t.Fatal("final feedback did not reach results")
	}
	test.Tap(u.primary)
	if u.practice.round != 1 || u.practice.state != questionPhase || len(u.practice.pairs) != 1 {
		t.Fatal("retry action failed")
	}
	test.Type(u.primary, "4")
	test.Tap(u.primary)
	test.Tap(u.primary)
	if u.primary.Text != "New Session" || u.practice.correct != 1 {
		t.Fatal("correct retry did not finish")
	}
	test.Tap(u.primary)
	if u.practice != nil || u.from.Value != 2 || u.count.Value != 20 {
		t.Fatal("new session did not retain settings")
	}
}

func TestPracticeActionPositionAndDisabledFeedbackKeypad(t *testing.T) {
	u := setupUI(t)
	u.count.SetValue(2)
	test.Tap(u.primary)
	position, size := u.primary.Position(), u.primary.Size()
	u.scroll.ScrollToBottom()
	offset := u.scroll.Offset
	assertPosition := func() {
		t.Helper()
		if u.primary.Position() != position || u.primary.Size() != size {
			t.Fatalf("action moved: position %v, size %v; want %v, %v", u.primary.Position(), u.primary.Size(), position, size)
		}
		if u.scroll.Offset != offset {
			t.Fatalf("practice scroll moved: got %v, want %v", u.scroll.Offset, offset)
		}
	}
	test.Tap(u.primary)
	assertPosition()
	test.Tap(u.digits[0])
	test.Tap(u.primary)
	assertPosition()
	for _, digit := range u.digits {
		if !digit.Visible() || !digit.Disabled() {
			t.Fatal("feedback keypad must stay visible and disabled")
		}
		test.Tap(digit)
	}
	if u.answerText != "0" || !u.clear.Disabled() || !u.erase.Disabled() {
		t.Fatal("feedback allowed answer editing")
	}
	test.Tap(u.primary)
	assertPosition()
	if u.digits[0].Disabled() {
		t.Fatal("next question did not enable keypad")
	}
	test.Type(u.primary, strconv.Itoa(u.practice.current().product()))
	test.Tap(u.primary)
	assertPosition()
	if u.primary.Text != "See results" || u.back.Text != "New Session" {
		t.Fatal("unexpected final feedback actions")
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
	u.count.SetValue(50)
	test.Tap(u.primary)
	if len(u.practice.pairs) != 50 {
		t.Fatal("practice did not honor the 50-question setting")
	}
	test.Type(u.primary, "12")
	before := u.practice
	test.Tap(u.back)
	if u.confirmation == nil {
		t.Fatal("discard confirmation not shown")
	}
	u.window.Canvas().Focused().TypedKey(&fyne.KeyEvent{Name: fyne.KeyReturn})
	if u.practice != before || u.answerText != "12" || u.window.Canvas().Focused() != u.primary {
		t.Fatal("cancelled navigation changed progress/focus")
	}
	test.Tap(u.back)
	u.discard.TypedKey(&fyne.KeyEvent{Name: fyne.KeyReturn})
	if u.practice != nil || u.count.Value != 50 {
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
	u.count.SetValue(51)
	if u.count.Value != 50 {
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
	test.Type(u.primary, strconv.Itoa(q.product()))
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
	if u.practice.index != 1 || u.practice.state != questionPhase || u.window.Canvas().Focused() != u.primary {
		t.Fatal("next question did not restore answer focus")
	}
	u.window.Canvas().Focus(u.digits[1])
	u.window.Canvas().FocusNext()
	if u.window.Canvas().Focused() != u.digits[2] {
		t.Fatal("Tab traversal did not reach the next digit")
	}
	u.window.Canvas().FocusPrevious()
	if u.window.Canvas().Focused() != u.digits[1] {
		t.Fatal("reverse Tab traversal did not restore the first digit")
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

func TestAnswerKeypadEditing(t *testing.T) {
	u := setupUI(t)
	u.from.SetValue(7)
	u.to.SetValue(8)
	u.count.SetValue(1)
	test.Tap(u.primary)
	test.Tap(u.digits[0])
	test.Tap(u.digits[0])
	if u.answerText != "0" {
		t.Fatal("leading zero was duplicated")
	}
	test.Tap(u.digits[5])
	test.Tap(u.digits[6])
	test.Tap(u.digits[7])
	test.Tap(u.digits[8])
	if u.answerText != "567" || u.answer.Text != "Your answer: 567" || u.practice.state != questionPhase {
		t.Fatal("keypad did not replace the initial zero, limit length, or wait for submission")
	}
	test.Tap(u.erase)
	if u.answerText != "56" {
		t.Fatal("backspace did not remove the final digit")
	}
	test.Tap(u.clear)
	if u.answerText != "" || u.answer.Text != "Your answer: —" {
		t.Fatal("Clear did not reset the display")
	}
	test.Tap(u.erase)
	if u.answerText != "" {
		t.Fatal("backspace on an empty answer changed it")
	}
}

func TestAnswerKeyboardRouting(t *testing.T) {
	u := setupUI(t)
	u.count.SetValue(2)
	test.Tap(u.primary)
	test.Tap(u.digits[1])
	test.Type(u.digits[1], "2x-34")
	if u.answerText != "123" {
		t.Fatal("focused keypad button did not route digits, ignore non-digits, or limit length")
	}
	u.digits[1].TypedKey(&fyne.KeyEvent{Name: fyne.KeyBackspace})
	if u.answerText != "12" {
		t.Fatal("keyboard Backspace failed")
	}
	u.digits[1].TypedKey(&fyne.KeyEvent{Name: fyne.KeyDelete})
	if u.answerText != "" {
		t.Fatal("keyboard Delete failed")
	}
	u.window.Canvas().Unfocus()
	u.window.Canvas().OnTypedRune()('9')
	if u.answerText != "9" {
		t.Fatal("unfocused canvas did not route numeric input")
	}
	u.digits[9].TypedKey(&fyne.KeyEvent{Name: fyne.KeyReturn})
	if u.practice.state != feedbackPhase || u.practice.answer != 9 {
		t.Fatal("Enter on a digit button must submit rather than append that digit")
	}
	u.window.Canvas().OnTypedRune()('7')
	if u.answerText != "9" {
		t.Fatal("numeric keys changed the answer during feedback")
	}
	test.Tap(u.primary)
	if u.answerText != "" || u.answer.Text != "Your answer: —" {
		t.Fatal("next question did not reset answer entry")
	}
	test.Type(u.primary, "4")
	test.Tap(u.back)
	u.window.Canvas().OnTypedRune()('5')
	u.window.Canvas().OnTypedKey()(&fyne.KeyEvent{Name: fyne.KeyReturn})
	if u.answerText != "4" || u.practice.state != questionPhase {
		t.Fatal("answer entry remained active behind the confirmation dialog")
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

func TestUITwelveTimesTwelve(t *testing.T) {
	u := setupUI(t)
	u.from.SetValue(13)
	if u.from.Value != 12 || u.to.Value != 12 {
		t.Fatal("factor sliders must allow 12 and clamp larger values")
	}
	u.count.SetValue(1)
	test.Tap(u.primary)
	for _, digit := range []int{1, 4, 4} {
		test.Tap(u.digits[digit])
	}
	test.Tap(u.primary)
	if u.practice.current() != (question{12, 12}) || u.practice.answer != 144 || u.practice.correct != 1 {
		t.Fatal("12 × 12 could not be answered correctly using the keypad")
	}
}
