package main

import (
	"academie/model"
	"fmt"
	"os"

	_ "student"

	_ "github.com/01-edu/go-tests/lib/challenge"
	_ "github.com/01-edu/z01"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/inancgumus/screen"
)

func main() {
	bar := model.Progress()

	if _, err := tea.NewProgram(bar).Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
	screen.Clear()

}
