package models

import (
	"time"
)

type Schedule struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	NodeID    uint      `gorm:"index;not null" json:"node_id"`
	Node      Node      `gorm:"foreignKey:NodeID" json:"node,omitempty"`
	Task      string    `gorm:"not null" json:"task"`
	CronExpr  string    `json:"cron"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
