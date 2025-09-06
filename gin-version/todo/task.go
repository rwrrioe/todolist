package todo

import "time"

type Task struct {
	Title       string
	Description string
	Completed   bool

	CreatedAt   time.Time
	CompletedAt *time.Time
}

func NewTask(title string, desciption string) Task {
	return Task{
		Title:       title,
		Description: desciption,
		Completed:   false,

		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
}

func (t *Task) Complete() {
	completeTime := time.Now()

	t.Completed = true
	t.CompletedAt = &completeTime
}

func (t *Task) Uncomplete() {
	t.Completed = false
	t.CompletedAt = nil
}
