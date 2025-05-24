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

	ID        string `json:"id" gorm:"primaryKey;unique;not null"`
	Command   string `json:"command_id" gorm:"command;index:cmd_schedule_idx;uniqueIndex:cmd_schedule_unique"`
	Status    string `json:"status" gorm:"status;default:'scheduled'"`
	Schedule  int64  `json:"schedule" gorm:"schedule;index:cmd_schedule_idx;uniqueIndex:cmd_schedule_unique"`
	StartedAt int    `json:"started_at" gorm:"started_at;default:0"`
	EndedAt   int    `json:"ended_at" gorm:"ended_at;default:0"`
	Params    string `json:"params" gorm:"params;not null"`
}
