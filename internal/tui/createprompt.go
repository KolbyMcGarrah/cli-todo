package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kolbymcgarrah/cli-todo/internal/db"
	"github.com/kolbymcgarrah/cli-todo/internal/task"
)

const (
	// input map keys
	taskKey    string = "TASK"
	projectKey string = "PROJECT"

	// tea model key
	AddTaskKey = "ADDTASK"
)

type CreateModel struct {
	inputs map[string]textinput.Model
	client *db.TaskDB
	models map[string]tea.Model
}

func newInput(placeholderText string) textinput.Model {
	input := textinput.New()
	input.Placeholder = placeholderText
	input.Prompt = ""
	input.Focus()
	return input
}

func NewCreateModel(client *db.TaskDB) *CreateModel {
	inputs := make(map[string]textinput.Model)
	inputs[taskKey] = newInput("take over the world...")
	inputs[projectKey] = newInput("enter project details")
	models := make(map[string]tea.Model)
	return &CreateModel{
		inputs: inputs,
		client: client,
		models: models,
	}
}

func (cm *CreateModel) ResetInputs() {
	for k, v := range cm.inputs {
		v.Reset()
		cm.inputs[k] = v
	}
}

func (cm *CreateModel) AddModel(key string, model tea.Model) {
	cm.models[key] = model
}

func (cm *CreateModel) Init() tea.Cmd {
	return textinput.Blink
}

func (cm *CreateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return cm.models[ListKey], tea.Quit
		case "enter":
			_ = cm.client.AddTask(task.Task{Name: cm.inputs[taskKey].Value()})
			return cm.models[ListKey], nil
		}
	}
	cm.updateInputs(msg)
	return cm, nil
}

func (cm *CreateModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, 0)
	for k := range cm.inputs {
		if cm.inputs[k].Focused() {
			m, cmd := cm.inputs[k].Update(msg)
			cm.inputs[k] = m
			cmds = append(cmds, cmd)
		}
	}
	return tea.Batch(cmds...)
}

func (cm *CreateModel) View() string {
	output := strings.Builder{}
	output.WriteString("\n" + inputStyle.Render("Enter task details: "))
	output.WriteString(cm.inputs[taskKey].View())
	output.WriteString("\n\n\n\n" + helperTextStyle.Render("press enter to save"))
	return output.String()
}
