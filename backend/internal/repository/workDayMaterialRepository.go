package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafaameen91/project-managment/backend/internal/models"
)

type WorkDayMaterialRepository struct {
	db *pgxpool.Pool
}

func NewWorkDayMaterialRepository(db *pgxpool.Pool) *WorkDayMaterialRepository {
	return &WorkDayMaterialRepository{db: db}
}

func (r *WorkDayMaterialRepository) GetAll(ctx context.Context) ([]models.WorkDayMaterial, error) {
	query := `
		SELECT id, workDayId, materialName, quantity, cost, notes, createdAt, updatedAt
		FROM workDayMaterials
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var materials []models.WorkDayMaterial
	for rows.Next() {
		var m models.WorkDayMaterial
		err := rows.Scan(
			&m.ID, &m.WorkDayID, &m.MaterialName, &m.Quantity,
			&m.Cost, &m.Notes, &m.CreatedAt, &m.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		materials = append(materials, m)
	}

	return materials, rows.Err()
}

func (r *WorkDayMaterialRepository) GetByID(ctx context.Context, id int64) (*models.WorkDayMaterial, error) {
	query := `
		SELECT id, workDayId, materialName, quantity, cost, notes, createdAt, updatedAt
		FROM workDayMaterials
		WHERE id = $1
	`

	var m models.WorkDayMaterial
	err := r.db.QueryRow(ctx, query, id).Scan(
		&m.ID, &m.WorkDayID, &m.MaterialName, &m.Quantity,
		&m.Cost, &m.Notes, &m.CreatedAt, &m.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *WorkDayMaterialRepository) GetByWorkDayID(ctx context.Context, workDayID int64) ([]models.WorkDayMaterial, error) {
	query := `
		SELECT id, workDayId, materialName, quantity, cost, notes, createdAt, updatedAt
		FROM workDayMaterials
		WHERE workDayId = $1
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query, workDayID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var materials []models.WorkDayMaterial
	for rows.Next() {
		var m models.WorkDayMaterial
		err := rows.Scan(
			&m.ID, &m.WorkDayID, &m.MaterialName, &m.Quantity,
			&m.Cost, &m.Notes, &m.CreatedAt, &m.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		materials = append(materials, m)
	}

	return materials, rows.Err()
}

func (r *WorkDayMaterialRepository) Create(ctx context.Context, material *models.WorkDayMaterial) (*models.WorkDayMaterial, error) {
	query := `
		INSERT INTO workDayMaterials (workDayId, materialName, quantity, cost, notes)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, createdAt, updatedAt
	`

	err := r.db.QueryRow(ctx, query,
		material.WorkDayID, material.MaterialName, material.Quantity,
		material.Cost, material.Notes,
	).Scan(&material.ID, &material.CreatedAt, &material.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return material, nil
}

func (r *WorkDayMaterialRepository) Update(ctx context.Context, id int64, material *models.WorkDayMaterial) (*models.WorkDayMaterial, error) {
	query := `
		UPDATE workDayMaterials
		SET workDayId = $1, materialName = $2, quantity = $3, cost = $4, notes = $5
		WHERE id = $6
		RETURNING id, workDayId, materialName, quantity, cost, notes, createdAt, updatedAt
	`

	err := r.db.QueryRow(ctx, query,
		material.WorkDayID, material.MaterialName, material.Quantity,
		material.Cost, material.Notes, id,
	).Scan(
		&material.ID, &material.WorkDayID, &material.MaterialName, &material.Quantity,
		&material.Cost, &material.Notes, &material.CreatedAt, &material.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return material, nil
}

func (r *WorkDayMaterialRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM workDayMaterials WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
