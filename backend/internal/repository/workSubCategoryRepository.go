package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafaameen91/project-managment/backend/internal/models"
)

type WorkSubCategoryRepositoryInterface interface {
	GetAll(ctx context.Context, limit, offset int) ([]models.WorkSubCategory, int64, error)
	GetByID(ctx context.Context, id int64) (*models.WorkSubCategory, error)
	GetByCategoryID(ctx context.Context, categoryID int64) ([]models.WorkSubCategory, error)
	Create(ctx context.Context, subCategory *models.WorkSubCategory) (*models.WorkSubCategory, error)
	Update(ctx context.Context, id int64, subCategory *models.WorkSubCategory) (*models.WorkSubCategory, error)
	Delete(ctx context.Context, id int64) error
}

type WorkSubCategoryRepository struct {
	db *pgxpool.Pool
}

func NewWorkSubCategoryRepository(db *pgxpool.Pool) *WorkSubCategoryRepository {
	return &WorkSubCategoryRepository{db: db}
}

func (r *WorkSubCategoryRepository) GetAll(ctx context.Context, limit, offset int) ([]models.WorkSubCategory, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM workSubCategories`
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, categoryId, name, description, percentage, status
		FROM workSubCategories
		ORDER BY createdAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var subCategories []models.WorkSubCategory
	for rows.Next() {
		var s models.WorkSubCategory
		err := rows.Scan(
			&s.ID, &s.CategoryID, &s.Name, &s.Description,
			&s.Percentage, &s.Status,
		)
		if err != nil {
			return nil, 0, err
		}
		subCategories = append(subCategories, s)
	}

	return subCategories, total, rows.Err()
}

func (r *WorkSubCategoryRepository) GetByID(ctx context.Context, id int64) (*models.WorkSubCategory, error) {
	query := `
		SELECT id, categoryId, name, description, percentage, status, createdBy, createdAt
		FROM workSubCategories
		WHERE id = $1
	`

	var s models.WorkSubCategory
	err := r.db.QueryRow(ctx, query, id).Scan(
		&s.ID, &s.CategoryID, &s.Name, &s.Description,
		&s.Percentage, &s.Status, &s.CreatedBy, &s.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *WorkSubCategoryRepository) GetByCategoryID(ctx context.Context, categoryID int64) ([]models.WorkSubCategory, error) {
	query := `
		SELECT id, categoryId, name, description, percentage, status
		FROM workSubCategories
		WHERE categoryId = $1
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subCategories []models.WorkSubCategory
	for rows.Next() {
		var s models.WorkSubCategory
		err := rows.Scan(
			&s.ID, &s.CategoryID, &s.Name, &s.Description,
			&s.Percentage, &s.Status,
		)
		if err != nil {
			return nil, err
		}
		subCategories = append(subCategories, s)
	}

	return subCategories, rows.Err()
}

func (r *WorkSubCategoryRepository) Create(ctx context.Context, subCategory *models.WorkSubCategory) (*models.WorkSubCategory, error) {
	query := `
		INSERT INTO workSubCategories (categoryId, name, description, percentage, status, createdBy)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		subCategory.CategoryID, subCategory.Name, subCategory.Description,
		subCategory.Percentage, subCategory.Status, subCategory.CreatedBy,
	).Scan(&subCategory.ID, &subCategory.CreatedAt)

	if err != nil {
		return nil, err
	}

	return subCategory, nil
}

func (r *WorkSubCategoryRepository) Update(ctx context.Context, id int64, subCategory *models.WorkSubCategory) (*models.WorkSubCategory, error) {
	query := `
		UPDATE workSubCategories
		SET name = $1, description = $2, percentage = $3, status = $4
		WHERE id = $5
		RETURNING id, categoryId, name, description, percentage, status, createdBy, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		subCategory.Name, subCategory.Description,
		subCategory.Percentage, subCategory.Status, id,
	).Scan(
		&subCategory.ID, &subCategory.CategoryID, &subCategory.Name, &subCategory.Description,
		&subCategory.Percentage, &subCategory.Status, &subCategory.CreatedBy, &subCategory.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return subCategory, nil
}

func (r *WorkSubCategoryRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM workSubCategories WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
