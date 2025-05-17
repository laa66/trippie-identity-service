package entity

import (
	"fmt"
	"time"
)

type Identity struct {
	ID       int
	Mail     string
	Login    string
	Password string
	Date     time.Time
	Active   bool
}

func (i *Identity) TableName() string {
	return "IDENTITY"
}

func (i *Identity) String() string {
	return fmt.Sprintf("Identity [ID: %d, Mail: %s, Login: %s, Password: %s, Date: %s, Active: %t]", 
	i.ID, i.Mail, i.Login, i.Password, i.Date, i.Active)
}