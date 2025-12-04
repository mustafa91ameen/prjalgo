package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafaameen91/project-managment/backend/internal/models"
)

type WorkDayRepository struct {
	db *pgxpool.Pool
}

func NewWorkDayRepository(db *pgxpool.Pool) *WorkDayRepository {
	return &WorkDayRepository{db: db}
}

func (r *WorkDayRepository) GetAll(ctx context.Context) ([]models.WorkDay, error) {
	query := `
		SELECT id, projectId, workSubCategoryId, workDate, description,
		       status, totalCost, notes, createdBy, createdAt, updatedAt
		FROM workDays
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workDays []models.WorkDay
	for rows.Next() {
		var w models.WorkDay
		err := rows.Scan(
			&w.ID, &w.ProjectID, &w.WorkSubCategoryID, &w.WorkDate, &w.Description,
			&w.Status, &w.TotalCost, &w.Notes, &w.CreatedBy, &w.CreatedAt, &w.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		workDays = append(workDays, w)
	}

	return workDays, rows.Err()
}

func (r *WorkDayRepository) GetByID(ctx context.Context, id int64) (*models.WorkDay, error) {
	query := `
		SELECT id, projectId, workSubCategoryId, workDate, description,
		       status, totalCost, notes, createdBy, createdAt, updatedAt
		FROM workDays
		WHERE id = $1
	`

	var w models.WorkDay
	err := r.db.QueryRow(ctx, query, id).Scan(
		&w.ID, &w.ProjectID, &w.WorkSubCategoryID, &w.WorkDate, &w.Description,
		&w.Status, &w.TotalCost, &w.Notes, &w.CreatedBy, &w.CreatedAt, &w.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &w, nil
}

func (r *WorkDayRepository) GetByProjectID(ctx context.Context, projectID int64) ([]models.WorkDay, error) {
	query := `
		SELECT id, projectId, workSubCategoryId, workDate, description,
		       status, totalCost, notes, createdBy, createdAt, updatedAt
		FROM workDays
		WHERE projectId = $1
		ORDER BY workDate DESC
	`

	rows, err := r.db.Query(ctx, query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var workDays []models.WorkDay
	for rows.Next() {
		var w models.WorkDay
		err := rows.Scan(
			&w.ID, &w.ProjectID, &w.WorkSubCategoryID, &w.WorkDate, &w.Description,
			&w.Status, &w.TotalCost, &w.Notes, &w.CreatedBy, &w.CreatedAt, &w.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		workDays = append(workDays, w)
	}

	return workDays, rows.Err()
}

func (r *WorkDayRepository) Create(ctx context.Context, workDay *models.WorkDay) (*models.WorkDay, error) {
	query := `
		INSERT INTO workDays (projectId, workSubCategoryId, workDate, description,
		                      status, totalCost, notes, createdBy)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, createdAt, updatedAt
	`

	err := r.db.QueryRow(ctx, query,
		workDay.ProjectID, workDay.WorkSubCategoryID, workDay.WorkDate, workDay.Description,
		workDay.Status, workDay.TotalCost, workDay.Notes, workDay.CreatedBy,
	).Scan(&workDay.ID, &workDay.CreatedAt, &workDay.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return workDay, nil
}

func (r *WorkDayRepository) Update(ctx context.Context, id int64, workDay *models.WorkDay) (*models.WorkDay, error) {
	query := `
		UPDATE workDays
		SET projectId = $1, workSubCategoryId = $2, workDate = $3, description = $4,
		    status = $5, totalCost = $6, notes = $7
		WHERE id = $8
		RETURNING id, projectId, workSubCategoryId, workDate, description,
		          status, totalCost, notes, createdBy, createdAt, updatedAt
	`

	err := r.db.QueryRow(ctx, query,
		workDay.ProjectID, workDay.WorkSubCategoryID, workDay.WorkDate, workDay.Description,
		workDay.Status, workDay.TotalCost, workDay.Notes, id,
	).Scan(
		&workDay.ID, &workDay.ProjectID, &workDay.WorkSubCategoryID, &workDay.WorkDate,
		&workDay.Description, &workDay.Status, &workDay.TotalCost, &workDay.Notes,
		&workDay.CreatedBy, &workDay.CreatedAt, &workDay.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return workDay, nil
}

func (r *WorkDayRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM workDays WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
