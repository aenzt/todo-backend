package repository

import (
	"todo-backend/database/sql"
)

type Repository struct {
	User UserInterface
	Role RoleInterface
}

func Init(db sql.DB) *Repository {
	return &Repository{
		User: InitUser(db),
		Role: InitRole(db),
	}
}
