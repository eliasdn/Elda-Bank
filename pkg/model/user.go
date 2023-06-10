package model

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"users" swaggerignore:"true"`

	ID                int    `bun:",pk,autoincrement" json:"id"`
	Email             string `validate:"-" bun:",nullzero,notnull,unique" json:"email"`
	EncryptedPassword string `gorm:"not null" json:"encryptedPassword"` // need to use hash password and look convention
	//SignInCount       int                `validate:"min=0" bun:",notnull,default:0"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:",soft_delete,nullzero"`
	Disabled  bool      `validate:"-" bun:",notnull,default:false"`
}
