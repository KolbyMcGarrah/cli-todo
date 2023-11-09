package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kolbymcgarrah/cli-todo/internal/db"
	"github.com/kolbymcgarrah/cli-todo/internal/tui"
	"github.com/spf13/cobra"
)

func Add() *cobra.Command {
	return &cobra.Command{
		Use:     "add",
		Short:   "add a command",
		Args:    cobra.NoArgs,
		Example: "todo add -p myproject -m create a new task",
		RunE:    addCommand,
	}
}

func addCommand(cmd *cobra.Command, args []string) error {
	client, err := db.OpenDB()
	if err != nil {
		return fmt.Errorf("error opening db")
	}

	// flags := cmd.Flags()
	// project, _ := flags.GetString("project")
	// message, err := flags.GetString("message")
	// if err != nil {
	// 	return fmt.Errorf("error getting message: %w", err)
	// }
	// if message == "" {
	// 	return fmt.Errorf("no message provided")
	// }
	// task := task.NewTask(message, project)
	// return client.AddTask(task)
	createModel := tui.NewCreateModel(client)
	listModel, err := tui.NewListModel(client)
	if err != nil {
		return fmt.Errorf("error creating list model: %w", err)
	}
	createModel.AddModel(tui.ListKey, listModel)
	p := tea.NewProgram(createModel)
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error starting program: %w", err)
	}
	return nil
}
