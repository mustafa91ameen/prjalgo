package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafaameen91/project-managment/backend/internal/models"
)

type PageRepository struct {
	db *pgxpool.Pool
}

func NewPageRepository(db *pgxpool.Pool) *PageRepository {
	return &PageRepository{db: db}
}

func (r *PageRepository) GetAll(ctx context.Context) ([]models.Page, error) {
	query := `
		SELECT id, name, icon, route, status, createdAt
		FROM pages
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pages []models.Page
	for rows.Next() {
		var p models.Page
		err := rows.Scan(&p.ID, &p.Name, &p.Icon, &p.Route, &p.Status, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		pages = append(pages, p)
	}

	return pages, rows.Err()
}

func (r *PageRepository) GetByID(ctx context.Context, id int64) (*models.Page, error) {
	query := `
		SELECT id, name, icon, route, status, createdAt
		FROM pages
		WHERE id = $1
	`

	var p models.Page
	err := r.db.QueryRow(ctx, query, id).Scan(
		&p.ID, &p.Name, &p.Icon, &p.Route, &p.Status, &p.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PageRepository) GetActivePages(ctx context.Context) ([]models.Page, error) {
	query := `
		SELECT id, name, icon, route, status, createdAt
		FROM pages
		WHERE status = 'active'
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pages []models.Page
	for rows.Next() {
		var p models.Page
		err := rows.Scan(&p.ID, &p.Name, &p.Icon, &p.Route, &p.Status, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		pages = append(pages, p)
	}

	return pages, rows.Err()
}

func (r *PageRepository) Create(ctx context.Context, page *models.Page) (*models.Page, error) {
	query := `
		INSERT INTO pages (name, icon, route, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		page.Name, page.Icon, page.Route, page.Status,
	).Scan(&page.ID, &page.CreatedAt)

	if err != nil {
		return nil, err
	}

	return page, nil
}

func (r *PageRepository) Update(ctx context.Context, id int64, page *models.Page) (*models.Page, error) {
	query := `
		UPDATE pages
		SET name = $1, icon = $2, route = $3, status = $4
		WHERE id = $5
		RETURNING id, name, icon, route, status, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		page.Name, page.Icon, page.Route, page.Status, id,
	).Scan(&page.ID, &page.Name, &page.Icon, &page.Route, &page.Status, &page.CreatedAt)

	if err != nil {
		return nil, err
	}

	return page, nil
}

func (r *PageRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM pages WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
