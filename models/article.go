package models

import "time"

type Article struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	TITLE     string    `gorm:"type:varchar(200)" json:"title"`
	CONTENT   string    `gorm:"type:text" json:"content"`
	CATEGORY  string    `gorm:"type:varchar(100)" json:"category"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	STATUS    string    `gorm:"type:varchar(200)" json:"status"`
}
