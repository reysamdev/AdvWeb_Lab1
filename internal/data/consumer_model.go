package data

import (
	"context"
	"database/sql"
	"time"
)

type Consumer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	Version   int       `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ConsumerModel struct {
	DB *sql.DB
}

func (m ConsumerModel) Insert(c *Consumer) error {
	query := `
		INSERT INTO consumers (name, email, status)
		VALUES ($1, $2, 'active')
		RETURNING id, version, created_at, updated_at`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, c.Name, c.Email).Scan(
		&c.ID, &c.Version, &c.CreatedAt, &c.UpdatedAt,
	)
}