package dto

import (
	"fmt"
	"time"
)

type Identity struct {
	ID     int
	Mail   string
	Date   time.Time
	Active bool
}

func (i *Identity) String() string {
	return fmt.Sprintf("Identity [ID: %d, Mail: %s, Date: %s, Active: %t]",
		i.ID, i.Mail, i.Date, i.Active)
}

type CreateIdentity struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

func (c *CreateIdentity) String() string {
	return fmt.Sprintf("CreateIdentity [Mail: %s, Password: %s]", c.Mail, c.Password)
}
