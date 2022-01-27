package model

import (
	"gorm.io/gorm"
)

type ActiveLog struct {
	gorm.Model
	ReceiveMessage string `gorm:"not null" json:"receivemessage"`
}
