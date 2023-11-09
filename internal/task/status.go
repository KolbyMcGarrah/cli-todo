package task

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

func (s Status) String() string {
	return [...]string{"todo", "in progress", "done"}[s]
}

// Move the status to the next status or restart if at done
func (s Status) Next() int {
	if s == Done {
		return Todo.Int()
	}
	return int(s + 1)
}

// move the status back a status or the end if at the start
func (s Status) Prev() int {
	if s == Todo {
		return Done.Int()
	}
	return int(s - 1)
}

func (s Status) Int() int {
	return int(s)
}
