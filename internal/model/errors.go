package model

import "errors"

var (
	ErrNotFound         = errors.New("reminder not found")
	ErrTitleRequired    = errors.New("title is required")
	ErrRemindAtRequired = errors.New("remind_at is required")
	ErrRemindAtPast     = errors.New("remind_at is must be in the future")
	ErrInvalidStatus    = errors.New("invalid status value")
)
