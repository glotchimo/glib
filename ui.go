package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	messages []string
	input    string
}

func (m model) Init() tea.Cmd {
	authenticate()
	go listen()
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEscape:
			return m, tea.Quit

		case tea.KeyEnter:
			m.messages = append(m.messages, fmt.Sprintf("\033[34m%s:\033[0m %s", USERNAME, m.input))
			IRC.Say(CHANNEL, m.input)
			m.input = ""
			return m, nil

		case tea.KeyBackspace, tea.KeyDelete:
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
			return m, nil

		case tea.KeyCtrlW:
			s := strings.TrimRight(m.input, " ")
			i := strings.LastIndex(s, " ")
			if i == -1 {
				m.input = ""
			} else {
				m.input = s[:i]
			}
			return m, nil

		case tea.KeyCtrlU:
			m.input = ""
			return m, nil

		case tea.KeyTab:
			m.input += "\t"

		default:
			m.input += msg.String()
			return m, nil
		}

	case string:
		m.messages = append(m.messages, msg)
		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	var b strings.Builder

	for _, msg := range m.messages {
		b.WriteString(string(msg))
		b.WriteRune('\n')
	}

	b.WriteString("> ")
	b.WriteString(m.input)

	return b.String()
}
