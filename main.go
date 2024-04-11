package main

import (
    "os"
    "fmt"
    "time"
    "github.com/SolarDebris/gomodoro/pomodoro"
    "github.com/charmbracelet/bubbles/help"
    "github.com/charmbracelet/bubbles/key"
    "github.com/charmbracelet/bubbles/timer"
    tea "github.com/charmbracelet/bubbletea"
)

func main(){

    const timeout = time.Second * 5
    fmt.Printf("%T", timeout)

    m := pomodoro.StopWatchModel{
		Timer: timer.NewWithInterval(timeout, time.Millisecond),
		Keymap: pomodoro.StopWatchKeymap{
			Start: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "start"),
			),
			Stop: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "stop"),
			),
			Reset: key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "reset"),
			),
			Quit: key.NewBinding(
				key.WithKeys("q", "ctrl+c"),
				key.WithHelp("q", "quit"),
			),
		},
		Help: help.New(),
	}

	m.Keymap.Start.SetEnabled(false)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Uh oh, we encountered an error:", err)
		os.Exit(1)
	}
}
