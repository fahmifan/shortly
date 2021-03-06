package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/fahmifan/ulids"
	"gopkg.in/guregu/null.v4"
)

type actor struct {
	CreatedBy ulids.ULID
	UpdatedBy ulids.ULID
	DeletedBy ulids.Null
}

type timestamp struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt null.Time
}

type RowScanner interface {
	Scan(dest ...interface{}) error
}

type URL struct {
	ID       ulids.ULID `json:"id,omitempty"`
	IsPublic bool       `json:"isPublic"`
	Original string     `json:"original"`
	Shorten  string     `json:"shorten,omitempty"`
	actor
	timestamp
}
type URLRepository struct {
	DB *sql.DB
}

func NewURLRepository(u *URLRepository) *URLRepository {
	return u
}

func (u *URLRepository) columns() []string {
	return []string{
		"id",
		"is_public",
		"original",
		"shorten",
		"created_by",
		"updated_by",
		"deleted_by",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

func (u *URLRepository) scanRow(url *URL, scanner RowScanner) error {
	if url == nil {
		return nil
	}

	return scanner.Scan(
		&url.ID,
		&url.IsPublic,
		&url.Original,
		&url.Shorten,
		&url.CreatedBy,
		&url.UpdatedBy,
		&url.DeletedBy,
		&url.CreatedAt,
		&url.UpdatedAt,
		&url.DeletedAt,
	)
}

func (u *URLRepository) scanRows(rows *sql.Rows) (urls []URL, err error) {
	for rows.Next() {
		url := URL{}
		err := u.scanRow(&url, rows)
		if err != nil {
			return nil, fmt.Errorf("scan url rows: %w", err)
		}
		urls = append(urls, url)
	}
	return
}

func (u *URLRepository) Create(ctx context.Context, url *URL) error {
	_, err := sq.Insert("urls").Columns(u.columns()...).
		Values(
			url.ID,
			url.IsPublic,
			url.Original,
			url.Shorten,
			url.CreatedBy,
			url.UpdatedBy,
			url.DeletedBy,
			url.CreatedAt,
			url.UpdatedAt,
			url.DeletedAt,
		).
		RunWith(u.DB).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("create urls: %w", err)
	}

	return nil
}

type ListFilter struct {
	Cursor ulids.ULID
}

func (u *URLRepository) ListByUserID(ctx context.Context, userID ulids.ULID, filter ListFilter) ([]URL, error) {
	rows, err := sq.Select(u.columns()...).From("urls").Where(sq.Eq{
		"created_by": userID,
		"deleted_at": nil,
	}).RunWith(u.DB).Query()
	if err != nil {
		return nil, fmt.Errorf("urlRepo ListByUserID query: %w", err)
	}
	defer rows.Close()

	urls, err := u.scanRows(rows)
	if err != nil {
		return nil, fmt.Errorf("urlRepo ListByUserID scanRows: %w", err)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("urlRepo ListByUserID rowsErr: %w", err)
	}

	return urls, nil
}
