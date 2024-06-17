package models

import "time"

type Project struct {
	ID          string
	ProjectName string
	CreatedAt   time.Time
}
