package sqlite

import "database/sql"

type URLRepository struct {
	DB *sql.DB
}

func NewURLRepository(u *URLRepository) *URLRepository {
	return u
}
