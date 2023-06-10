package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Customer struct {
	bun.BaseModel `bun:"customers" swaggerignore:"true"`

	ID int `bun:",pk,autoincrement" json:"id"`
	// personal information more informations can be added
	FirstName   string    `bun:",notnull" json:"firstName"`
	LastName    string    `bun:",notnull" json:"lastName"`
	DOB         time.Time `bun:"type:date,notnull" json:"DOB"` // Date Of birth |try to use time.time to use date info
	Address     string    `bun:",notnull" json:"Address"`
	City        string    `bun:",notnull" json:"City"`
	PostalCode  int       `bun:",notnull" json:"PostalCode"`
	NumberPhone int       `bun:",notnull" json:"NumberPhone"`
	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp"` // try timestamptz later
	DeletedAt   time.Time `bun:",soft_delete,nullzero"`
}
