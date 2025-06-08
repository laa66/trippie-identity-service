package entity

import (
	"fmt"
	"time"
)

type Identity struct {
	ID       int `gorm:"primaryKey"`
	Mail     string
	Password string
	Date     time.Time
	Active   bool
}

func (i *Identity) TableName() string {
	return "account"
}

func (i *Identity) String() string {
	return fmt.Sprintf("Identity [ID: %d, Mail: %s, Password: %s, Date: %s, Active: %t]",
		i.ID, i.Mail, i.Password, i.Date, i.Active)
}
