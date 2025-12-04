package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafaameen91/project-managment/backend/internal/models"
)

type IncomeRepository struct {
	db *pgxpool.Pool
}

func NewIncomeRepository(db *pgxpool.Pool) *IncomeRepository {
	return &IncomeRepository{db: db}
}

func (r *IncomeRepository) GetAll(ctx context.Context) ([]models.Income, error) {
	query := `
		SELECT id, name, amount, type, incomeDate, status
		FROM income
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var incomes []models.Income
	for rows.Next() {
		var i models.Income
		err := rows.Scan(
			&i.ID, &i.Name, &i.Amount, &i.Type, &i.IncomeDate, &i.Status,
		)
		if err != nil {
			return nil, err
		}
		incomes = append(incomes, i)
	}

	return incomes, rows.Err()
}

func (r *IncomeRepository) GetByID(ctx context.Context, id int64) (*models.Income, error) {
	query := `
		SELECT id, name, amount, type, incomeDate, status, notes, createdBy, createdAt
		FROM income
		WHERE id = $1
	`

	var i models.Income
	err := r.db.QueryRow(ctx, query, id).Scan(
		&i.ID, &i.Name, &i.Amount, &i.Type, &i.IncomeDate,
		&i.Status, &i.Notes, &i.CreatedBy, &i.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &i, nil
}

func (r *IncomeRepository) Create(ctx context.Context, income *models.Income) (*models.Income, error) {
	query := `
		INSERT INTO income (name, amount, type, incomeDate, status, notes, createdBy)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		income.Name, income.Amount, income.Type, income.IncomeDate,
		income.Status, income.Notes, income.CreatedBy,
	).Scan(&income.ID, &income.CreatedAt)

	if err != nil {
		return nil, err
	}

	return income, nil
}

func (r *IncomeRepository) Update(ctx context.Context, id int64, income *models.Income) (*models.Income, error) {
	query := `
		UPDATE income
		SET name = $1, amount = $2, type = $3, incomeDate = $4, status = $5, notes = $6
		WHERE id = $7
		RETURNING id, name, amount, type, incomeDate, status, notes, createdBy, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		income.Name, income.Amount, income.Type, income.IncomeDate,
		income.Status, income.Notes, id,
	).Scan(
		&income.ID, &income.Name, &income.Amount, &income.Type, &income.IncomeDate,
		&income.Status, &income.Notes, &income.CreatedBy, &income.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return income, nil
}

func (r *IncomeRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM income WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
