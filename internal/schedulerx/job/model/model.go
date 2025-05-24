package job_model

import (
	"gorm.io/gorm"
)

const (
	STATUS_SCHEDULED = "scheduled"
	STATUS_PICKED    = "picked"
	STATUS_COMPLETED = "completed"
	STATUS_FAILED    = "failed"
)

type Job struct {
	gorm.Model

	CommandId string `json:"command_id" gorm:"command_id"`
	Status    string `json:"status" gorm:"status"`
	Schedule  string `json:"schedule" gorm:"schedule"`
	StartedAt int    `json:"started_at" gorm:"started_at"`
	EndedAt   int    `json:"ended_at" gorm:"ended_at"`
	Active    bool   `json:"active" gorm:"active"`
}
