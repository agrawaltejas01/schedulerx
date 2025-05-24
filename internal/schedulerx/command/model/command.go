package command_model

import (
	"gorm.io/gorm"
)

type Command struct {
	gorm.Model

	ID      string   `json:"id" gorm:"primaryKey;unique;not null"`
	Command string   `json:"command" gorm:"command;unique;not null"`
	Cron    string   `json:"cron" gorm:"cron;not null"`
	Active  bool     `json:"active" gorm:"active;not null;default:true"`
	Params  []string `json:"params" gorm:"-"`
}
