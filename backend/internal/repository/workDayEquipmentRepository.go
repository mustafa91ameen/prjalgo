package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafaameen91/project-managment/backend/internal/models"
)

type WorkDayEquipmentRepositoryInterface interface {
	GetAll(ctx context.Context) ([]models.WorkDayEquipment, error)
	GetByID(ctx context.Context, id int64) (*models.WorkDayEquipment, error)
	GetByWorkDayID(ctx context.Context, workDayID int64) ([]models.WorkDayEquipment, error)
	Create(ctx context.Context, equipment *models.WorkDayEquipment) (*models.WorkDayEquipment, error)
	Update(ctx context.Context, id int64, equipment *models.WorkDayEquipment) (*models.WorkDayEquipment, error)
	Delete(ctx context.Context, id int64) error
}

type WorkDayEquipmentRepository struct {
	db *pgxpool.Pool
}

func NewWorkDayEquipmentRepository(db *pgxpool.Pool) *WorkDayEquipmentRepository {
	return &WorkDayEquipmentRepository{db: db}
}

func (r *WorkDayEquipmentRepository) GetAll(ctx context.Context) ([]models.WorkDayEquipment, error) {
	query := `
		SELECT id, workDayId, equipmentName, quantity, cost, notes, createdAt
		FROM workDayEquipment
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipments []models.WorkDayEquipment
	for rows.Next() {
		var e models.WorkDayEquipment
		err := rows.Scan(
			&e.ID, &e.WorkDayID, &e.EquipmentName, &e.Quantity,
			&e.Cost, &e.Notes, &e.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		equipments = append(equipments, e)
	}

	return equipments, rows.Err()
}

func (r *WorkDayEquipmentRepository) GetByID(ctx context.Context, id int64) (*models.WorkDayEquipment, error) {
	query := `
		SELECT id, workDayId, equipmentName, quantity, cost, notes, createdAt
		FROM workDayEquipment
		WHERE id = $1
	`

	var e models.WorkDayEquipment
	err := r.db.QueryRow(ctx, query, id).Scan(
		&e.ID, &e.WorkDayID, &e.EquipmentName, &e.Quantity,
		&e.Cost, &e.Notes, &e.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (r *WorkDayEquipmentRepository) GetByWorkDayID(ctx context.Context, workDayID int64) ([]models.WorkDayEquipment, error) {
	query := `
		SELECT id, workDayId, equipmentName, quantity, cost, notes, createdAt
		FROM workDayEquipment
		WHERE workDayId = $1
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query, workDayID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipments []models.WorkDayEquipment
	for rows.Next() {
		var e models.WorkDayEquipment
		err := rows.Scan(
			&e.ID, &e.WorkDayID, &e.EquipmentName, &e.Quantity,
			&e.Cost, &e.Notes, &e.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		equipments = append(equipments, e)
	}

	return equipments, rows.Err()
}

func (r *WorkDayEquipmentRepository) Create(ctx context.Context, equipment *models.WorkDayEquipment) (*models.WorkDayEquipment, error) {
	query := `
		INSERT INTO workDayEquipment (workDayId, equipmentName, quantity, cost, notes)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		equipment.WorkDayID, equipment.EquipmentName, equipment.Quantity,
		equipment.Cost, equipment.Notes,
	).Scan(&equipment.ID, &equipment.CreatedAt)

	if err != nil {
		return nil, err
	}

	return equipment, nil
}

func (r *WorkDayEquipmentRepository) Update(ctx context.Context, id int64, equipment *models.WorkDayEquipment) (*models.WorkDayEquipment, error) {
	query := `
		UPDATE workDayEquipment
		SET equipmentName = $1, quantity = $2, cost = $3, notes = $4
		WHERE id = $5
		RETURNING id, workDayId, equipmentName, quantity, cost, notes, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		equipment.EquipmentName, equipment.Quantity,
		equipment.Cost, equipment.Notes, id,
	).Scan(
		&equipment.ID, &equipment.WorkDayID, &equipment.EquipmentName, &equipment.Quantity,
		&equipment.Cost, &equipment.Notes, &equipment.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return equipment, nil
}

func (r *WorkDayEquipmentRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM workDayEquipment WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
