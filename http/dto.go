package http

import (
	"encoding/json"
	"errors"
	"time"
)

type TaskDTO struct {
	Title       string
	Description string
}

func (t TaskDTO) ValidateForCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}
	if t.Description == "" {
		return errors.New("description is empty")
	}

	return nil
}

type CompleteDTO struct {
	Complete bool
}

type ErrorDTO struct {
	Message string
	Time    time.Time
}

func (e ErrorDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func newErrorDTO(err error, time time.Time) ErrorDTO {
	return ErrorDTO{
		Message: err.Error(),
		Time:    time,
	}
}
