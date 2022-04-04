package repository

import "database/sql"

// Repositories stores all repositories
type Repositories struct {
}

// NewRepositories returns a pointer to new Repositories
func NewRepositories(postgres *sql.DB) *Repositories {
	return &Repositories{}
}
