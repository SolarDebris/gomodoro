package main

import (
    "fmt"
    "time"
    "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	startButton         *widget.Button
	startShortBreakButton *widget.Button
	startLongBreakButton  *widget.Button
	ticker              *time.Ticker
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Pomodoro Timer")

	timerLabel := widget.NewLabel("25:00")
	startButton := widget.NewButton("Start Pomodoro", func() {
		go pomodoroTimer(timerLabel, 25*time.Minute)
		startButton.Disable()
	})
	startShortBreakButton := widget.NewButton("Start Short Break", func() {
		go pomodoroTimer(timerLabel, 5*time.Minute)
		startShortBreakButton.Disable()
	})
	startLongBreakButton := widget.NewButton("Start Long Break", func() {
		go pomodoroTimer(timerLabel, 15*time.Minute)
		startLongBreakButton.Disable()
	})
	stopButton := widget.NewButton("Stop", func() {
		stopTimer()
		startButton.Enable()
		startShortBreakButton.Enable()
		startLongBreakButton.Enable()
	})

	buttonContainer := container.NewHBox(startButton, startShortBreakButton, startLongBreakButton, stopButton)
	content := container.NewVBox(timerLabel, buttonContainer)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}


func pomodoroTimer(timerLabel *widget.Label, duration time.Duration) {
	ticker = time.NewTicker(duration)

	for remainingTime := duration; remainingTime > 0; remainingTime -= time.Second {
		select {
		case <-ticker.C:
			timerLabel.SetText("00:00")
			// You can add a notification here, e.g., displaying a message box or playing a sound
			break
		default:
			timerLabel.SetText(fmt.Sprintf("%02d:%02d", remainingTime.Minutes(), int(remainingTime.Seconds())%60))
			time.Sleep(1 * time.Second)
		}
	}

	stopTimer()
}

func stopTimer() {
	if ticker != nil {
		ticker.Stop()
	}
}
