package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafaameen91/project-managment/backend/internal/models"
)

type DebtorRepositoryInterface interface {
	GetAll(ctx context.Context) ([]models.Debtor, error)
	GetByID(ctx context.Context, id int64) (*models.Debtor, error)
	Create(ctx context.Context, debtor *models.Debtor) (*models.Debtor, error)
	Update(ctx context.Context, id int64, debtor *models.Debtor) (*models.Debtor, error)
	Delete(ctx context.Context, id int64) error
}

type DebtorRepository struct {
	db *pgxpool.Pool
}

func NewDebtorRepository(db *pgxpool.Pool) *DebtorRepository {
	return &DebtorRepository{db: db}
}

func (r *DebtorRepository) GetAll(ctx context.Context) ([]models.Debtor, error) {
	query := `
		SELECT id, name, email, phone, totalDebt, currency, dueDate, status
		FROM debtors
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var debtors []models.Debtor
	for rows.Next() {
		var d models.Debtor
		err := rows.Scan(
			&d.ID, &d.Name, &d.Email, &d.Phone, &d.TotalDebt,
			&d.Currency, &d.DueDate, &d.Status,
		)
		if err != nil {
			return nil, err
		}
		debtors = append(debtors, d)
	}

	return debtors, rows.Err()
}

func (r *DebtorRepository) GetByID(ctx context.Context, id int64) (*models.Debtor, error) {
	query := `
		SELECT id, name, email, phone, totalDebt, currency, dueDate, status, notes, createdBy, createdAt
		FROM debtors
		WHERE id = $1
	`

	var d models.Debtor
	err := r.db.QueryRow(ctx, query, id).Scan(
		&d.ID, &d.Name, &d.Email, &d.Phone, &d.TotalDebt, &d.Currency,
		&d.DueDate, &d.Status, &d.Notes, &d.CreatedBy, &d.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &d, nil
}

func (r *DebtorRepository) Create(ctx context.Context, debtor *models.Debtor) (*models.Debtor, error) {
	query := `
		INSERT INTO debtors (name, email, phone, totalDebt, currency, dueDate, status, notes, createdBy)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		debtor.Name, debtor.Email, debtor.Phone, debtor.TotalDebt, debtor.Currency,
		debtor.DueDate, debtor.Status, debtor.Notes, debtor.CreatedBy,
	).Scan(&debtor.ID, &debtor.CreatedAt)

	if err != nil {
		return nil, err
	}

	return debtor, nil
}

func (r *DebtorRepository) Update(ctx context.Context, id int64, debtor *models.Debtor) (*models.Debtor, error) {
	query := `
		UPDATE debtors
		SET name = $1, email = $2, phone = $3, totalDebt = $4, currency = $5,
		    dueDate = $6, status = $7, notes = $8
		WHERE id = $9
		RETURNING id, name, email, phone, totalDebt, currency, dueDate, status, notes, createdBy, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		debtor.Name, debtor.Email, debtor.Phone, debtor.TotalDebt, debtor.Currency,
		debtor.DueDate, debtor.Status, debtor.Notes, id,
	).Scan(
		&debtor.ID, &debtor.Name, &debtor.Email, &debtor.Phone, &debtor.TotalDebt,
		&debtor.Currency, &debtor.DueDate, &debtor.Status, &debtor.Notes,
		&debtor.CreatedBy, &debtor.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return debtor, nil
}

func (r *DebtorRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM debtors WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
