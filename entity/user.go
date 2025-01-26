package entity

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"users,alias:u"`

	ID       uuid.UUID `json:"id" bun:"type:char(36),default:uuid(),pk"`
	Name     string    `json:"name" bun:",notnull,type:varchar(45)"`
	Email    string    `json:"email" bun:",notnull,type:varchar(255)"`
	Password string    `json:"password" bun:",notnull,type:varchar(255)"`
}

type UserResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type LoginResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Token string    `json:"token"`
}
