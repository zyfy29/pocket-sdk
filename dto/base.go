package dto

import "fmt"

type Resp[T any] struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Content T      `json:"content"`
}

func (r Resp[T]) ErrorFailed() error {
	return fmt.Errorf("pocket returns %d, due to: %s", r.Status, r.Message)
}
