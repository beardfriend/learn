package models

import (
	"database/sql/driver"
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt Date
}

type Date time.Time

func (date Date) Value() (driver.Value, error) {
	h := time.Time(date).Hour()
	min := time.Time(date).Minute()
	s := time.Time(date).Second()
	y, m, d := time.Time(date).Date()
	return time.Date(y, m, d, h, min, s, 0, time.Time(date).Location()), nil
}
