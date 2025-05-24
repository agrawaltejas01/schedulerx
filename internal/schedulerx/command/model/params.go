package command_model

import (
	"gorm.io/gorm"
)

type Params struct {
	gorm.Model

	ID      string `json:"id" gorm:"primaryKey;unique;not null"`
	Command string `json:"command" gorm:"command;not null;index"`
	Param   string `json:"param" gorm:"param;not null"`
}
