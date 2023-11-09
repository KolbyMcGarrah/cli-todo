package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kolbymcgarrah/cli-todo/internal/db"
	"github.com/kolbymcgarrah/cli-todo/internal/tui"
	"github.com/spf13/cobra"
)

func List() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "list all todo items",
		Example: "todo list",
		Aliases: []string{"l", "list"},
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := db.OpenDB()
			_ = client.CreateTaskBucket()
			if err != nil {
				return fmt.Errorf("error opening db: %w", err)
			}
			listModel, err := tui.NewListModel(client)
			if err != nil {
				return fmt.Errorf("error creating model: %w", err)
			}
			createModel := tui.NewCreateModel(client)
			listModel.AddModel(tui.AddTaskKey, createModel)
			createModel.AddModel(tui.ListKey, listModel)
			p := tea.NewProgram(listModel)
			if _, err := p.Run(); err != nil {
				return fmt.Errorf("error starting program: %w", err)
			}
			return nil
		},
	}
}
