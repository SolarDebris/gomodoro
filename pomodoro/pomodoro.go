package pomodoro


import (
    "time"
    "github.com/charmbracelet/bubbles/help"
    "github.com/charmbracelet/bubbles/key"
    "github.com/charmbracelet/bubbles/stopwatch"
    "github.com/charmbracelet/bubbles/timer"
    tea "github.com/charmbracelet/bubbletea"
    
)

const timeout = time.Second * 5

type StopWatchModel struct {
	StopWatch stopwatch.Model
	Keymap    StopWatchKeymap
	Help      help.Model
	Quitting  bool
}

type StopWatchKeymap struct {
	Start key.Binding
	Stop  key.Binding
	Reset key.Binding
	Quit  key.Binding
}

func (m StopWatchModel) Init() tea.Cmd {
	return m.StopWatch.Init()
}

func (m StopWatchModel) View() string {
	// Note: you could further customize the time output by getting the
	// duration from m.stopwatch.Elapsed(), which returns a time.Duration, and
	// skip m.stopwatch.View() altogether.
	s := m.StopWatch.View() + "\n"
	if !m.Quitting {
		s = "Elapsed: " + s
		s += m.helpView()
	}
	return s
}

func (m StopWatchModel) helpView() string {
	return "\n" + m.Help.ShortHelpView([]key.Binding{
		m.Keymap.Start,
		m.Keymap.Stop,
		m.Keymap.Reset,
		m.Keymap.Quit,
	})
}

func (m StopWatchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keymap.Quit):
			m.Quitting = true
			return m, tea.Quit
		case key.Matches(msg, m.Keymap.Reset):
			return m, m.StopWatch.Reset()
		case key.Matches(msg, m.Keymap.Start, m.Keymap.Stop):
			m.Keymap.Stop.SetEnabled(!m.StopWatch.Running())
			m.Keymap.Start.SetEnabled(m.StopWatch.Running())
			return m, m.StopWatch.Toggle()
		}
	}
	var cmd tea.Cmd
	m.StopWatch, cmd = m.StopWatch.Update(msg)
	return m, cmd
}


type TimerModel struct {
	Timer    timer.Model
	Keymap   TimerKeymap
	Help     help.Model
	Quitting bool
}

type TimerKeymap struct {
	Start key.Binding
	Stop  key.Binding
	Reset key.Binding
	Quit  key.Binding
}

func (m TimerModel) Init() tea.Cmd {
	return m.Timer.Init()
}

func (m TimerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case timer.TickMsg:
		var cmd tea.Cmd
		m.Timer, cmd = m.Timer.Update(msg)
		return m, cmd

	case timer.StartStopMsg:
		var cmd tea.Cmd
		m.Timer, cmd = m.Timer.Update(msg)
		m.Keymap.Stop.SetEnabled(m.Timer.Running())
		m.Keymap.Start.SetEnabled(!m.Timer.Running())
		return m, cmd

	case timer.TimeoutMsg:
		m.Quitting = true
		return m, tea.Quit

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keymap.Quit):
			m.Quitting = true
			return m, tea.Quit
		case key.Matches(msg, m.Keymap.Reset):
			m.Timer.Timeout = timeout
		case key.Matches(msg, m.Keymap.Start, m.Keymap.Stop):
			return m, m.Timer.Toggle()
		}
	}

	return m, nil
}

func (m TimerModel) helpView() string {
	return "\n" + m.Help.ShortHelpView([]key.Binding{
		m.Keymap.Start,
		m.Keymap.Stop,
		m.Keymap.Reset,
		m.Keymap.Quit,
	})
}

func (m TimerModel) View() string {
	// For a more detailed timer view you could read m.timer.Timeout to get
	// the remaining time as a time.Duration and skip calling m.timer.View()
	// entirely.
	s := m.Timer.View()

	if m.Timer.Timedout() {
		s = "All done!"
	}
	s += "\n"
	if !m.Quitting {
		s = "Exiting in " + s
		s += m.helpView()
	}
	return s
}

