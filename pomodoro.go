package main
import (
    "fmt"
    "time"
    "github.com/fatih/color"
)

func main() {
    var modeChoice string
    var timeChoice int
    printLogo()
    fmt.Print("Welcome to my Pomodoro timer.\nWould you like the timer to be automatic or manual (a or m): ")


    _, err := fmt.Scan(&modeChoice)
    if err != nil {
        return 
    }

    if modeChoice == "a" {
        for {
            startPomodoro()
            shortBreak()
        }
    } else if modeChoice == "m" {
        for {
            fmt.Println("1. Pomodoro (25 min)\n2. Short Break (5 min)\n3. Long Break (15 min)")
            fmt.Println("Enter your choice:")

            _, err := fmt.Scan(&timeChoice)
            if err != nil {
                return
            }

            if timeChoice == 1 {
                startPomodoro()
            } else if timeChoice == 2 {
                shortBreak()
            } else if timeChoice == 3 {
                longBreak()
            }
        }
    }
}


func printLogo(){
    logo := `
██▓███   ▒█████   ███▄ ▄███▓ ▒█████  ▓█████▄  ▒█████   ██▀███   ▒█████    
▓██░  ██▒▒██▒  ██▒▓██▒▀█▀ ██▒▒██▒  ██▒▒██▀ ██▌▒██▒  ██▒▓██ ▒ ██▒▒██▒  ██▒ 
▓██░ ██▓▒▒██░  ██▒▓██    ▓██░▒██░  ██▒░██   █▌▒██░  ██▒▓██ ░▄█ ▒▒██░  ██▒ 
▒██▄█▓▒ ▒▒██   ██░▒██    ▒██ ▒██   ██░░▓█▄   ▌▒██   ██░▒██▀▀█▄  ▒██   ██░ 
▒██▒ ░  ░░ ████▓▒░▒██▒   ░██▒░ ████▓▒░░▒████▓ ░ ████▓▒░░██▓ ▒██▒░ ████▓▒░ 
▒▓▒░ ░  ░░ ▒░▒░▒░ ░ ▒░   ░  ░░ ▒░▒░▒░  ▒▒▓  ▒ ░ ▒░▒░▒░ ░ ▒▓ ░▒▓░░ ▒░▒░▒░  
░▒ ░       ░ ▒ ▒░ ░  ░      ░  ░ ▒ ▒░  ░ ▒  ▒   ░ ▒ ▒░   ░▒ ░ ▒░  ░ ▒ ▒░  
░░       ░ ░ ░ ▒  ░      ░   ░ ░ ░ ▒   ░ ░  ░ ░ ░ ░ ▒    ░░   ░ ░ ░ ░ ▒   
             ░ ░         ░       ░ ░     ░        ░ ░     ░         ░ ░   
                                       ░                                  
`

    fmt.Print(logo)

}

func startPomodoro() {
    fmt.Println("Pomodoro Started (25 minutes)\n")

    // Set the Pomodoro duration to 25 minutes
    pomodoroDuration := 25 * time.Minute

    // Get the current time to track the start time
    startTime := time.Now()

    for {
        elapsedTime := time.Since(startTime)
        remainingTime := pomodoroDuration - elapsedTime

        if remainingTime <= 0 {
            fmt.Println("\nPomodoro Completed! Take a break.")
            return
        }

        printProgressBar(pomodoroDuration, remainingTime)
        time.Sleep(1 * time.Second)
    }
}

func shortBreak() {
    fmt.Println("Short Break Started (5 minutes)\n")

    // Set the break duration to 5 minutes
    breakDuration := 5 * time.Minute

    // Get the current time to track the start time
    startTime := time.Now()

    for {
        elapsedTime := time.Since(startTime)
        remainingTime := breakDuration - elapsedTime

        if remainingTime <= 0 {
            fmt.Println("\nBreak Completed! Ready for the next Pomodoro.")
            return
        }

        printProgressBar(breakDuration, remainingTime)
        time.Sleep(1 * time.Second)
    }
}


func longBreak() {
    fmt.Println("Long Break Started (15 minutes)\n")

    // Set the break duration to 5 minutes
    breakDuration := 15 * time.Minute

    // Get the current time to track the start time
    startTime := time.Now()
    fmt.Print("\r")

    for {
        elapsedTime := time.Since(startTime)
        remainingTime := breakDuration - elapsedTime

        if remainingTime <= 0 {
            fmt.Println("\nBreak Completed! Ready for the next Pomodoro.")
            return
        }

        printProgressBar(breakDuration, remainingTime)
        time.Sleep(1 * time.Second)
    }
}


func printProgressBar(total time.Duration, remaining time.Duration) {
	progress := int((1 - float64(remaining)/float64(total)) * 30)
	minutes := int(remaining.Minutes())
	seconds := int(remaining.Seconds()) % 60


    color.Set(color.FgGreen)
    fmt.Print("\r")

	for i := 0; i < 30; i++ {
		if i < progress {
            fmt.Print("█")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Printf("%02d:%02d", minutes, seconds)
}

