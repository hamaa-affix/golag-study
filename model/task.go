package model

import "time"

type Task struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Title     string    `json:"title" gorm:"not null"`
    Status    uint8     `json:"status" gorm:"not null"`
    CreatedAt time.Time `json:"created_at"`
    UpdateAt  time.Time `json:"updated_at"`
}
