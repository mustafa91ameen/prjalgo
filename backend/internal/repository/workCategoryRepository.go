package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafaameen91/project-managment/backend/internal/models"
)

type WorkCategoryRepository struct {
	db *pgxpool.Pool
}

func NewWorkCategoryRepository(db *pgxpool.Pool) *WorkCategoryRepository {
	return &WorkCategoryRepository{db: db}
}

func (r *WorkCategoryRepository) GetAll(ctx context.Context) ([]models.WorkCategory, error) {
	query := `
		SELECT id, name, description, status, createdBy, createdAt, updatedAt
		FROM workCategories
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.WorkCategory
	for rows.Next() {
		var c models.WorkCategory
		err := rows.Scan(
			&c.ID, &c.Name, &c.Description, &c.Status,
			&c.CreatedBy, &c.CreatedAt, &c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, rows.Err()
}

func (r *WorkCategoryRepository) GetByID(ctx context.Context, id int64) (*models.WorkCategory, error) {
	query := `
		SELECT id, name, description, status, createdBy, createdAt, updatedAt
		FROM workCategories
		WHERE id = $1
	`

	var c models.WorkCategory
	err := r.db.QueryRow(ctx, query, id).Scan(
		&c.ID, &c.Name, &c.Description, &c.Status,
		&c.CreatedBy, &c.CreatedAt, &c.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *WorkCategoryRepository) Create(ctx context.Context, category *models.WorkCategory) (*models.WorkCategory, error) {
	query := `
		INSERT INTO workCategories (name, description, status, createdBy)
		VALUES ($1, $2, $3, $4)
		RETURNING id, createdAt, updatedAt
	`

	err := r.db.QueryRow(ctx, query,
		category.Name, category.Description, category.Status, category.CreatedBy,
	).Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *WorkCategoryRepository) Update(ctx context.Context, id int64, category *models.WorkCategory) (*models.WorkCategory, error) {
	query := `
		UPDATE workCategories
		SET name = $1, description = $2, status = $3
		WHERE id = $4
		RETURNING id, name, description, status, createdBy, createdAt, updatedAt
	`

	err := r.db.QueryRow(ctx, query,
		category.Name, category.Description, category.Status, id,
	).Scan(
		&category.ID, &category.Name, &category.Description, &category.Status,
		&category.CreatedBy, &category.CreatedAt, &category.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (r *WorkCategoryRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM workCategories WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
