package repo

import (
	"context"
	"fmt"
	"rest-notes-api/internal/domain"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func NewPostgresRepo(dsn string) (*PostgresRepo, error) {
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return &PostgresRepo{}, err
	}

	return &PostgresRepo{db: db}, nil
}

func (p *PostgresRepo) Create(ctx context.Context, note domain.Note) (string, error) {
	q := `INSERT INTO notes (id, title, content, created_at)
          VALUES (gen_random_uuid(), $1, $2, $3)
          RETURNING id`

	var id string

	err := p.db.QueryRowContext(ctx, q, note.Title, note.Content, note.CreateAt).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("failed to insert note: %w", err)
	}

	return id, nil
}

func (p *PostgresRepo) GetAll(ctx context.Context) ([]domain.Note, error) {
	var notes []domain.Note
	q := `SELECT id, title, content, created_at FROM notes ORDER BY created_at DESC`

	if err := p.db.SelectContext(ctx, &notes, q); err != nil {
		return nil, fmt.Errorf("failed to get notes: %w", err)
	}

	return notes, nil
}

func (r *PostgresRepo) Update(ctx context.Context, note domain.Note) error {
	query := `
		UPDATE notes 
		SET title = $1, content = $2 
		WHERE id = $3
	`

	res, err := r.db.ExecContext(ctx, query, note.Title, note.Content, note.ID)
	if err != nil {
		return fmt.Errorf("failed to update note: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("note not found")
	}

	return nil
}

func (p *PostgresRepo) Delete(ctx context.Context, id string) error {
	q := `DELETE FROM notes WHERE id = $1`

	res, err := p.db.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("failed to delete note: %w", err)
	}

	row, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if row == 0 {
		return fmt.Errorf("note not found")
	}

	return nil
}

func (r *PostgresRepo) FindByID(ctx context.Context, id string) (domain.Note, error) {
	var note domain.Note

	query := `SELECT id, title, content, created_at FROM notes WHERE id = $1`

	if err := r.db.GetContext(ctx, &note, query, id); err != nil {
		return domain.Note{}, fmt.Errorf("failed to get note by id: %w", err)
	}

	return note, nil
}

func (p *PostgresRepo) Init(ctx context.Context) error {
	schema := `
	CREATE TABLE IF NOT EXISTS notes (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT,
    created_at TIMESTAMP NOT NULL
	);`

	_, err := p.db.ExecContext(ctx, schema)
	return err
}

func (p *PostgresRepo) Close() error {
	return p.db.Close()
}
