package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafaameen91/project-managment/backend/internal/models"
)

type WorkDayLaborRepository struct {
	db *pgxpool.Pool
}

func NewWorkDayLaborRepository(db *pgxpool.Pool) *WorkDayLaborRepository {
	return &WorkDayLaborRepository{db: db}
}

func (r *WorkDayLaborRepository) GetAll(ctx context.Context) ([]models.WorkDayLabor, error) {
	query := `
		SELECT id, workDayId, workerName, jobTitle, phone, address, quantity, cost, notes, createdAt, updatedAt
		FROM workDayLabor
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var labors []models.WorkDayLabor
	for rows.Next() {
		var l models.WorkDayLabor
		err := rows.Scan(
			&l.ID, &l.WorkDayID, &l.WorkerName, &l.JobTitle, &l.Phone,
			&l.Address, &l.Quantity, &l.Cost, &l.Notes, &l.CreatedAt, &l.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		labors = append(labors, l)
	}

	return labors, rows.Err()
}

func (r *WorkDayLaborRepository) GetByID(ctx context.Context, id int64) (*models.WorkDayLabor, error) {
	query := `
		SELECT id, workDayId, workerName, jobTitle, phone, address, quantity, cost, notes, createdAt, updatedAt
		FROM workDayLabor
		WHERE id = $1
	`

	var l models.WorkDayLabor
	err := r.db.QueryRow(ctx, query, id).Scan(
		&l.ID, &l.WorkDayID, &l.WorkerName, &l.JobTitle, &l.Phone,
		&l.Address, &l.Quantity, &l.Cost, &l.Notes, &l.CreatedAt, &l.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &l, nil
}

func (r *WorkDayLaborRepository) GetByWorkDayID(ctx context.Context, workDayID int64) ([]models.WorkDayLabor, error) {
	query := `
		SELECT id, workDayId, workerName, jobTitle, phone, address, quantity, cost, notes, createdAt, updatedAt
		FROM workDayLabor
		WHERE workDayId = $1
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query, workDayID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var labors []models.WorkDayLabor
	for rows.Next() {
		var l models.WorkDayLabor
		err := rows.Scan(
			&l.ID, &l.WorkDayID, &l.WorkerName, &l.JobTitle, &l.Phone,
			&l.Address, &l.Quantity, &l.Cost, &l.Notes, &l.CreatedAt, &l.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		labors = append(labors, l)
	}

	return labors, rows.Err()
}

func (r *WorkDayLaborRepository) Create(ctx context.Context, labor *models.WorkDayLabor) (*models.WorkDayLabor, error) {
	query := `
		INSERT INTO workDayLabor (workDayId, workerName, jobTitle, phone, address, quantity, cost, notes)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, createdAt, updatedAt
	`

	err := r.db.QueryRow(ctx, query,
		labor.WorkDayID, labor.WorkerName, labor.JobTitle, labor.Phone,
		labor.Address, labor.Quantity, labor.Cost, labor.Notes,
	).Scan(&labor.ID, &labor.CreatedAt, &labor.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return labor, nil
}

func (r *WorkDayLaborRepository) Update(ctx context.Context, id int64, labor *models.WorkDayLabor) (*models.WorkDayLabor, error) {
	query := `
		UPDATE workDayLabor
		SET workDayId = $1, workerName = $2, jobTitle = $3, phone = $4, address = $5, quantity = $6, cost = $7, notes = $8
		WHERE id = $9
		RETURNING id, workDayId, workerName, jobTitle, phone, address, quantity, cost, notes, createdAt, updatedAt
	`

	err := r.db.QueryRow(ctx, query,
		labor.WorkDayID, labor.WorkerName, labor.JobTitle, labor.Phone,
		labor.Address, labor.Quantity, labor.Cost, labor.Notes, id,
	).Scan(
		&labor.ID, &labor.WorkDayID, &labor.WorkerName, &labor.JobTitle, &labor.Phone,
		&labor.Address, &labor.Quantity, &labor.Cost, &labor.Notes, &labor.CreatedAt, &labor.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return labor, nil
}

func (r *WorkDayLaborRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM workDayLabor WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
