package domain

import "time"

type Role struct {
	ID        int       `db:"id" json:"id"`
	Slug      string    `db:"slug" json:"slug"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

const (
	RoleCustomer = 1
	RoleSeller   = 2
	RoleAdmin    = 3
)
