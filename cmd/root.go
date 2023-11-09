package cmd

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kolbymcgarrah/cli-todo/internal/db"
	"github.com/kolbymcgarrah/cli-todo/internal/tui"
	"github.com/spf13/cobra"
)

func Execute() error {
	var rootCmd = &cobra.Command{
		Use:   "todo",
		Short: "create, complete and list todo items from the terminal",
		Long:  "Create, complete and List todo items from the terminal",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := db.OpenDB()
			if err != nil {
				return fmt.Errorf("error opening db: %w", err)
			}
			client.CreateTaskBucket()
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
	return rootCmd.ExecuteContext(context.Background())
}
