package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kolbymcgarrah/cli-todo/internal/db"
	"github.com/kolbymcgarrah/cli-todo/internal/task"
)

const (
	// tea model key
	ListKey = "LIST"
)

type ListModel struct {
	tasks    []task.Task      // the list of tasks
	cursor   int              // which task our cursor is pointing to
	selected map[int]struct{} // which tasks are selected
	client   *db.TaskDB
	models   map[string]tea.Model
}

func NewListModel(client *db.TaskDB) (*ListModel, error) {

	tasks, err := client.GetTasks()
	if err != nil {
		return nil, fmt.Errorf("error getting tasks: %w", err)
	}
	models := make(map[string]tea.Model)
	return &ListModel{
		tasks:    tasks,
		selected: make(map[int]struct{}),
		client:   client,
		models:   models,
	}, nil
}

func (lm *ListModel) AddModel(key string, model tea.Model) {
	lm.models[key] = model
}

func (cm *ListModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (cm *ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return cm, tea.Quit
		case "up", "k":
			if cm.cursor > 0 {
				cm.cursor--
			} else {
				cm.cursor = len(cm.tasks) - 1
			}
		case "down", "j":
			if cm.cursor < len(cm.tasks)-1 {
				cm.cursor++
			} else {
				cm.cursor = 0
			}
		case "enter", " ":
			{
				if _, found := cm.selected[cm.cursor]; found {
					delete(cm.selected, cm.cursor)
				} else {
					cm.selected[cm.cursor] = struct{}{}
				}
			}
		case "d":
			{
				for k := range cm.selected {
					tsk := cm.tasks[k]
					err := cm.client.DeleteTask(tsk.ID)
					if err != nil {
						fmt.Printf("error deleting task: %s\n", err.Error())
					}
					cm.cursor = 0
					delete(cm.selected, k)
				}
			}
		case "a":
			createModel := cm.models[AddTaskKey].(*CreateModel)
			createModel.ResetInputs()
			return createModel, nil
		}
	}
	return cm, nil
}

func (cm *ListModel) View() string {
	var err error
	cm.tasks, err = cm.client.GetTasks()
	if err != nil {
		fmt.Printf("error getting updated tasks: %s", err)
	}
	s := "Tasks\n\n"
	for i, task := range cm.tasks {
		cursor := "  " // no cursor
		if cm.cursor == i {
			cursor = cursorStyle.Render(" >")
		}
		// Is this choice selected?
		checked := " " // not selected
		if _, found := cm.selected[i]; found {
			checked = xStyle.Render("x")
		}

		// render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursorStyle.Render(cursor), checked, taskText.Render(task.Name))
	}

	ht := helperTextStyle.Render("\nPress 'q' to quit\nPress 'd' to delete selected tasks\nPress 'a' to add new task\n")
	return fmt.Sprintf("%s%s", s, ht)
}
