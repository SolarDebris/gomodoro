package main



import (
    "fmt"
    "os"
    "time"
    "github.com/SolarDebris/pomodoro/pomodoro"
    "github.com/charmbracelet/bubbles/help"
    "github.com/charmbracelet/bubbles/key"
    "github.com/charmbracelet/bubbles/timer"
    tea "github.com/charmbracelet/bubbletea"
)

func main(){

    const timeout = time.Second * 5

    m := pomodoro.StopWatchModel{
		timer: timer.NewWithInterval(timeout, time.Millisecond),
		keymap: pomodoro.StopWatchKeymap{
			start: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "start"),
			),
			stop: key.NewBinding(
				key.WithKeys("s"),
				key.WithHelp("s", "stop"),
			),
			reset: key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "reset"),
			),
			quit: key.NewBinding(
				key.WithKeys("q", "ctrl+c"),
				key.WithHelp("q", "quit"),
			),
		},
		help: help.New(),
	}

	m.keymap.start.SetEnabled(false)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Uh oh, we encountered an error:", err)
		os.Exit(1)
	}
}
