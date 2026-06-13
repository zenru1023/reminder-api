package model

import "time"

type Status string

const (
	StatusPending   Status = "pending"
	StatusDone      Status = "done"
	StatusCancelled Status = "cancelled"
)

type Reminder struct {
	ID          uint64    `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	RemindAt    time.Time `db:"remind_at" json:"remind_at"`
	Status      Status    `db:"status" json:"status"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

// CreateReminderRequest - POST /v1/reminders
type CreateReminderRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	RemindAt    time.Time `json:"remind_at"`
}
