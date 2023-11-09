package task

import (
	"encoding/json"
	"fmt"
	"time"
)

type Task struct {
	ID      uint64
	Name    string
	Project string
	Status  string
	Created time.Time
}

func NewTask(name string, project string) Task {
	return Task{
		Name:    name,
		Project: project,
		Status:  "todo",
		Created: time.Now(),
	}
}

func (t Task) UpdateTaskValues(updated Task) {
	if updated.Name != "" {
		t.Name = updated.Name
	}
	if updated.Project != "" {
		t.Project = updated.Project
	}
	if updated.Status != "" {
		t.Status = updated.Status
	}
}

func (t Task) ToBytes() ([]byte, error) {
	taskBytes, err := json.Marshal(t)
	if err != nil {
		return nil, fmt.Errorf("error converting to bytes: %w", err)
	}
	return taskBytes, nil
}
