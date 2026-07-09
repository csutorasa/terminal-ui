package main

import (
	"fmt"
	"os"

	"github.com/csutorasa/terminal-ui/ansi"
	"github.com/csutorasa/terminal-ui/application"
	"github.com/csutorasa/terminal-ui/components"
	"github.com/csutorasa/terminal-ui/document"
	"github.com/csutorasa/terminal-ui/style"
)

func main() {
	//c := components.NewDefaultCreator()
	c := components.NewCreator(document.Default, style.DefaultTheme)
	text := c.NewText()
	text.SetText(ansi.NewFormattedText([]rune("hello\n"), new(ansi.Format).ForegroundColor(ansi.FormatColorGreen)).Concat(ansi.NewFormattedText([]rune("world\ntest"), new(ansi.Format))))
	text2 := c.NewText()
	text2.SetText(c.Theme().CreateText([]rune("hello2\nworld\ntest")))
	textInput := c.NewTextInput()
	textInput2 := c.NewTextInput()
	border2 := c.NewBorder()
	border2.SetBorder(ansi.NewFormattedRune('X', new(ansi.Format))).SetChild(text2)
	i := 0
	button := c.NewButton().SetText("OK").SetOnAction(func() {
		i++
		text.SetText(ansi.NewFormattedText([]rune(fmt.Sprintf("%d", i)), new(ansi.Format)))
	})
	//scroll := c.NewScroll(textInput)
	grid := c.NewGrid()
	grid.
		AddColumn(35).
		AddColumnOfRatio(1).
		AddRow(2).
		AddRowOfRatio(1).
		AddRowOfRatio(1).
		AddChildren(text, textInput, border2, textInput2, button, text)
	border := c.NewBorder().SetThickness(2)
	border.SetChild(grid)
	c.SetRoot(border)
	app := application.New(c.Document())
	app.Document().FocusNext()
	err := app.Run(application.NewTerminalDecoder(os.Stdin), application.NewTerminalRenderer(os.Stderr))
	if err != nil {
		panic(err)
	}
}
