package model

import (
	"github.com/uptrace/bun"
)

type Revenue struct {
	bun.BaseModel `bun:"revenue,alias:r"`

	Month   string `json:"month" bun:",type:varchar(4),notnull,unique"`
	Revenue int    `json:"revenue" bun:",notnull"`
}
