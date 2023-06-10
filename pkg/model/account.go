package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Account struct {
	bun.BaseModel `bun:"accounts" swaggerignore:"true"`

	ID            int    `bun:",pk,autoincrement" json:"id"`
	Type          int    `bun:",pk,notnull" json:"accountType"` // 1=check 2=saving 3=etf...
	BankAccountID string `bun:",notnull" json:"bankAccountID"`
	UserId        string `bun:",notnull" json:"userId"`
	Country       string `bun:",notnull" json:"country"`
	// paramater credit, max spending ...
	Balance   float32   `bun:",notnull,default:0" json:"balance"`
	CreatedAt time.Time `validate:"-" bun:"type:timestamptz,nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `validate:"-" bun:"type:timestamptz,nullzero,soft_delete,"`
	Disabled  bool      `validate:"-" bun:",notnull,default:false"` // fraud activity
}
