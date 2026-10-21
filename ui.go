package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Keep text readable on both desktop and mobile without shrinking for rotation.
type practiceTheme struct{ fyne.Theme }

func (t practiceTheme) Size(name fyne.ThemeSizeName) float32 {
	switch name {
	case theme.SizeNameText, theme.SizeNameCaptionText:
		return 20
	case theme.SizeNameHeadingText:
		return 36
	case theme.SizeNameSubHeadingText:
		return 24
	}
	return t.Theme.Size(name)
}

type progressTheme struct{ fyne.Theme }

func (t progressTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNamePrimary {
		if variant == theme.VariantLight {
			return color.NRGBA{R: 237, G: 139, B: 35, A: 255}
		}
		return color.NRGBA{R: 168, G: 82, B: 10, A: 255}
	}
	if name == theme.ColorNameForegroundOnPrimary {
		if variant == theme.VariantLight {
			return color.NRGBA{R: 51, G: 32, B: 10, A: 255}
		}
		return color.White
	}
	return t.Theme.Color(name, variant)
}

// Reserve two lines for validation or feedback so the action never jumps.
type messageSlotLayout struct{}

func (messageSlotLayout) MinSize([]fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(0, 64)
}

func (messageSlotLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	objects[0].Move(fyne.NewPos(0, 0))
	objects[0].Resize(size)
}

// Fyne buttons activate with Space by default. Also support Enter for practice.
type actionButton struct {
	widget.Button
	onRune func(rune)
	onKey  func(*fyne.KeyEvent) bool
}

func newActionButton(text string, tapped func()) *actionButton {
	b := &actionButton{}
	b.Text, b.OnTapped = text, tapped
	b.ExtendBaseWidget(b)
	return b
}

func (b *actionButton) MinSize() fyne.Size {
	size := b.Button.MinSize()
	size.Height = max(size.Height, 56)
	return size
}
func (b *actionButton) TypedKey(event *fyne.KeyEvent) {
	if b.onKey != nil && b.onKey(event) {
		return
	}
	if event.Name == fyne.KeyReturn || event.Name == fyne.KeyEnter {
		b.Tapped(nil)
		return
	}
	b.Button.TypedKey(event)
}

func (b *actionButton) TypedRune(char rune) {
	if b.onRune != nil {
		b.onRune(char)
	}
}

// Keep a bounded, centered column while allowing it to shrink with the window.
type focusColumnLayout struct{}

func (focusColumnLayout) MinSize(objects []fyne.CanvasObject) fyne.Size { return objects[0].MinSize() }
func (focusColumnLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	width := min(float32(420), size.Width)
	objects[0].Move(fyne.NewPos((size.Width-width)/2, 0))
	objects[0].Resize(fyne.NewSize(width, max(size.Height, objects[0].MinSize().Height)))
}

type practiceUI struct {
	window                         fyne.Window
	scroll                         *container.Scroll
	rng                            *rand.Rand
	practice                       *practice
	from, to, count                *widget.Slider
	answer                         *widget.Label
	answerText                     string
	digits                         [10]*actionButton
	clear, erase                   *actionButton
	fromLabel, toLabel, countLabel *widget.Label
	errorLabel                     *widget.Label
	primary, back                  *actionButton
	confirmation                   *dialog.CustomDialog
	discard                        *actionButton
}

func newPracticeUI(window fyne.Window, rng *rand.Rand) *practiceUI {
	fyne.CurrentApp().Settings().SetTheme(practiceTheme{theme.DefaultTheme()})
	u := &practiceUI{window: window, rng: rng}
	window.Canvas().SetOnTypedRune(u.appendDigit)
	window.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) { u.answerKey(event) })
	u.from, u.to, u.count = widget.NewSlider(1, 20), widget.NewSlider(1, 20), widget.NewSlider(1, 50)
	u.fromLabel, u.toLabel, u.countLabel = wrappedLabel(""), wrappedLabel(""), wrappedLabel("")
	u.from.SetValue(float64(defaultSettings.from))
	u.to.SetValue(float64(defaultSettings.to))
	u.count.SetValue(float64(defaultSettings.count))
	u.from.OnChanged = func(value float64) {
		if value > u.to.Value {
			u.to.SetValue(value)
		}
		u.updateSettings()
	}
	u.to.OnChanged = func(value float64) {
		if value < u.from.Value {
			u.from.SetValue(value)
		}
		u.updateSettings()
	}
	u.count.OnChanged = func(float64) { u.updateSettings() }
	u.updateSettings()
	u.showSettings()
	return u
}

func wrappedLabel(text string) *widget.Label {
	label := widget.NewLabel(text)
	label.Wrapping = fyne.TextWrapWord
	return label
}

func (u *practiceUI) setScreen(body fyne.CanvasObject, focus fyne.Focusable) {
	header := widget.NewLabelWithStyle("Go Math", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	u.back = nil
	var right fyne.CanvasObject
	if u.practice != nil {
		u.back = newActionButton("New Session", u.backToSettings)
		if u.practice.state == questionPhase {
			u.back.onRune = u.appendDigit
			u.back.onKey = func(event *fyne.KeyEvent) bool {
				if event.Name == fyne.KeyReturn || event.Name == fyne.KeyEnter {
					return false
				}
				return u.answerKey(event)
			}
		}
		right = u.back
	}
	top := container.NewBorder(nil, nil, header, right)
	column := container.New(focusColumnLayout{}, body)
	u.scroll = container.NewVScroll(container.NewPadded(column))
	content := container.NewBorder(container.NewVBox(top, widget.NewSeparator()), nil, nil, nil, u.scroll)
	u.window.SetContent(container.NewPadded(content))
	u.focus(focus)
}

// focus moves keyboard focus for desktop hardware entry. On mobile it would
// only pop the soft keyboard, which this app never needs: all entry is touch.
func (u *practiceUI) focus(f fyne.Focusable) {
	if fyne.CurrentDevice().IsMobile() {
		return
	}
	u.window.Canvas().Focus(f)
}

func (u *practiceUI) showSettings() {
	u.practice = nil
	u.errorLabel = wrappedLabel("")
	u.errorLabel.Importance = widget.DangerImportance
	u.errorLabel.Hide()
	u.primary = newActionButton("Start practice", u.start)
	u.primary.Importance = widget.HighImportance
	fields := container.NewVBox(u.fromLabel, u.from, u.toLabel, u.to, u.countLabel, u.count)
	u.setScreen(container.NewVBox(widget.NewLabelWithStyle("Choose your practice", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}), fields, wrappedLabel("Factors 1–20 · Questions 1–50"), u.errorLabel, u.primary), u.from)
}

func (u *practiceUI) selectedSettings() settings {
	return settings{from: int(u.from.Value), to: int(u.to.Value), count: int(u.count.Value)}
}

func (u *practiceUI) updateSettings() {
	u.fromLabel.SetText(fmt.Sprintf("Minimum factor: %d", int(u.from.Value)))
	u.toLabel.SetText(fmt.Sprintf("Maximum factor: %d", int(u.to.Value)))
	u.countLabel.SetText(fmt.Sprintf("Questions: %d", int(u.count.Value)))
}

func (u *practiceUI) start() {
	if u.practice != nil {
		return
	}
	p, err := newPractice(u.selectedSettings(), u.rng)
	if err != nil {
		u.errorLabel.SetText(err.Error())
		u.errorLabel.Show()
		return
	}
	u.practice = p
	u.showPractice()
}

func (u *practiceUI) showPractice() {
	p := u.practice
	if p.state == resultsPhase {
		u.showResults()
		return
	}
	q := p.current()
	progressText := fmt.Sprintf("Question %d of %d", p.index+1, len(p.pairs))
	if p.round > 0 {
		progressText = fmt.Sprintf("Retry %d · %s", p.round, progressText)
	}
	progress := widget.NewProgressBar()
	completed := p.index
	if p.state == feedbackPhase {
		completed++
	}
	progress.SetValue(float64(completed) / float64(len(p.pairs)))
	problem := widget.NewRichText(&widget.TextSegment{Text: fmt.Sprintf("%d × %d = ?", q.x, q.y), Style: widget.RichTextStyle{Alignment: fyne.TextAlignCenter, SizeName: theme.SizeNameHeadingText, TextStyle: fyne.TextStyle{Bold: true}}})
	body := container.NewVBox(widget.NewLabelWithStyle(progressText, fyne.TextAlignCenter, fyne.TextStyle{}),
		container.NewThemeOverride(progress, progressTheme{fyne.CurrentApp().Settings().Theme()}), problem)
	u.answer = widget.NewLabelWithStyle("Your answer: —", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	u.errorLabel = wrappedLabel("")
	u.errorLabel.Importance = widget.DangerImportance
	body.Add(u.answer)
	body.Add(container.New(messageSlotLayout{}, u.errorLabel))
	body.Add(u.answerKeypad())
	var focus fyne.Focusable
	if p.state == questionPhase {
		u.answerText = ""
		u.primary = newActionButton("Check answer", u.submit)
		u.primary.onRune, u.primary.onKey = u.appendDigit, u.answerKey
		focus = u.primary
	} else {
		feedback := "Correct"
		if p.answer != q.product() {
			feedback = fmt.Sprintf("Correct answer: %d × %d = %d", q.x, q.y, q.product())
		}
		u.answer.SetText(fmt.Sprintf("Your answer: %d", p.answer))
		u.errorLabel.Alignment = fyne.TextAlignCenter
		if p.answer == q.product() {
			u.errorLabel.Importance = widget.HighImportance
		}
		u.errorLabel.SetText(feedback)
		for _, digit := range u.digits {
			digit.Disable()
		}
		u.clear.Disable()
		u.erase.Disable()
		nextText := "Next"
		if p.index == len(p.pairs)-1 {
			nextText = "See results"
		}
		u.primary = newActionButton(nextText, func() {
			offset := u.scroll.Offset
			if p.next() {
				u.showPractice()
				if p.state == questionPhase {
					u.scroll.ScrollToOffset(offset)
				}
			}
		})
		focus = u.primary
	}
	u.primary.Importance = widget.HighImportance
	body.Add(u.primary)
	u.setScreen(body, focus)
}

func (u *practiceUI) answerKeypad() fyne.CanvasObject {
	keys := container.NewGridWithColumns(3)
	for digit := 1; digit <= 9; digit++ {
		char := rune('0' + digit)
		u.digits[digit] = u.answerButton(string(char), func() { u.appendDigit(char) })
		keys.Add(u.digits[digit])
	}
	u.clear = u.answerButton("Clear", func() { u.editAnswer(false) })
	u.digits[0] = u.answerButton("0", func() { u.appendDigit('0') })
	u.erase = u.answerButton("⌫", func() { u.editAnswer(true) })
	keys.Add(u.clear)
	keys.Add(u.digits[0])
	keys.Add(u.erase)
	return keys
}

func (u *practiceUI) answerButton(text string, tapped func()) *actionButton {
	b := newActionButton(text, tapped)
	b.onRune, b.onKey = u.appendDigit, u.answerKey
	return b
}

func (u *practiceUI) enteringAnswer() bool {
	return u.practice != nil && u.practice.state == questionPhase && u.confirmation == nil
}

func (u *practiceUI) appendDigit(char rune) {
	if !u.enteringAnswer() || char < '0' || char > '9' {
		return
	}
	if u.answerText == "0" {
		u.answerText = string(char)
	} else if len(u.answerText) < 3 {
		u.answerText += string(char)
	}
	u.updateAnswer()
}

func (u *practiceUI) editAnswer(backspace bool) {
	if !u.enteringAnswer() {
		return
	}
	if backspace && len(u.answerText) > 0 {
		u.answerText = u.answerText[:len(u.answerText)-1]
	} else if !backspace {
		u.answerText = ""
	}
	u.updateAnswer()
}

func (u *practiceUI) updateAnswer() {
	text := u.answerText
	if text == "" {
		text = "—"
	}
	u.answer.SetText("Your answer: " + text)
	u.errorLabel.SetText("")
}

func (u *practiceUI) answerKey(event *fyne.KeyEvent) bool {
	if !u.enteringAnswer() {
		return false
	}
	switch event.Name {
	case fyne.KeyBackspace:
		u.editAnswer(true)
	case fyne.KeyDelete:
		u.editAnswer(false)
	case fyne.KeyReturn, fyne.KeyEnter:
		u.submit()
	default:
		return false
	}
	return true
}

func (u *practiceUI) submit() {
	if !u.enteringAnswer() {
		return
	}
	if err := u.practice.submit(u.answerText); err != nil {
		u.errorLabel.SetText(err.Error())
		u.errorLabel.Show()
		u.focus(u.primary)
		return
	}
	offset := u.scroll.Offset
	u.showPractice()
	u.scroll.ScrollToOffset(offset)
}

func (u *practiceUI) showResults() {
	p := u.practice
	title := "Round complete"
	if len(p.mistakes) == 0 {
		title = "All correct!"
	}
	newSession := newActionButton("New Session", u.showSettings)
	u.primary = newSession
	actions := container.NewVBox()
	if len(p.mistakes) > 0 {
		u.primary = newActionButton("Retry mistakes", func() {
			if p.retry() {
				u.showPractice()
			}
		})
		actions.Add(u.primary)
	}
	u.primary.Importance = widget.HighImportance
	actions.Add(newSession)
	summary := fmt.Sprintf("Correct: %d\nIncorrect: %d", p.correct, len(p.mistakes))
	body := container.NewVBox(widget.NewLabelWithStyle(title, fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))
	if p.round > 0 {
		body.Add(widget.NewLabel(fmt.Sprintf("Retry round %d", p.round)))
	}
	body.Add(wrappedLabel(summary))
	body.Add(actions)
	u.setScreen(body, u.primary)
}

func (u *practiceUI) backToSettings() {
	if u.practice == nil {
		return
	}
	if u.practice.state == resultsPhase {
		u.showSettings()
		return
	}
	var focus fyne.Focusable = u.primary
	var confirmation *dialog.CustomDialog
	keep := newActionButton("Keep practicing", func() { confirmation.Hide() })
	u.discard = newActionButton("New Session", func() {
		confirmation.Hide()
		u.showSettings()
	})
	keep.Importance = widget.HighImportance
	confirmation = dialog.NewCustomWithoutButtons("Leave practice?", container.NewVBox(
		wrappedLabel("Your current progress will be discarded."), keep, u.discard,
	), u.window)
	confirmation.SetOnClosed(func() {
		u.confirmation = nil
		u.focus(focus)
	})
	u.confirmation = confirmation
	confirmation.Show()
	u.focus(keep)
}
