package main

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/muesli/reflow/wordwrap"
	"golang.org/x/term"
)

type glib struct {
	messages []string
	input    string
}

func (m glib) Init() tea.Cmd {
	authenticate()
	go listen()
	return nil
}

func (m glib) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		if idx := strings.Index(msg, USERNAME); idx != -1 {
			l := msg[:idx]
			r := msg[idx+len(USERNAME):]
			msg = l + color.New(color.FgBlack, color.BgYellow).Sprint(USERNAME) + r
		}

		m.messages = append(m.messages, msg)
		return m, nil
	}

	return m, nil
}

func (m glib) View() string {
	var b strings.Builder

	terminalWidth, terminalHeight, err := term.GetSize(0)
	if err != nil {
		log.Fatal("error getting terminal size: ", err)
	}
	numMessages := terminalHeight - 2 // -2 for the input prompt and cursor

	padding := numMessages - len(m.messages)
	if padding < 0 {
		padding = 0
	}

	for i := 0; i < padding; i++ {
		b.WriteRune('\n')
	}

	for _, msg := range m.messages {
		b.WriteString(wordwrap.String(string(msg), terminalWidth))
		b.WriteRune('\n')
	}

	b.WriteString("\033[32m>\033[0m ")
	b.WriteString(wordwrap.String(m.input, terminalWidth-2))

	return b.String()
}
