package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafaameen91/project-managment/backend/internal/models"
)

type RoleRepository struct {
	db *pgxpool.Pool
}

func NewRoleRepository(db *pgxpool.Pool) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) GetAll(ctx context.Context) ([]models.Role, error) {
	query := `
		SELECT id, name, description, createdAt, updatedAt
		FROM roles
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		err := rows.Scan(
			&role.ID, &role.Name, &role.Description,
			&role.CreatedAt, &role.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, rows.Err()
}

func (r *RoleRepository) GetByID(ctx context.Context, id int64) (*models.Role, error) {
	query := `
		SELECT id, name, description, createdAt, updatedAt
		FROM roles
		WHERE id = $1
	`

	var role models.Role
	err := r.db.QueryRow(ctx, query, id).Scan(
		&role.ID, &role.Name, &role.Description,
		&role.CreatedAt, &role.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *RoleRepository) Create(ctx context.Context, role *models.Role) (*models.Role, error) {
	query := `
		INSERT INTO roles (name, description)
		VALUES ($1, $2)
		RETURNING id, createdAt, updatedAt
	`

	err := r.db.QueryRow(ctx, query,
		role.Name, role.Description,
	).Scan(&role.ID, &role.CreatedAt, &role.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *RoleRepository) Update(ctx context.Context, id int64, role *models.Role) (*models.Role, error) {
	query := `
		UPDATE roles
		SET name = $1, description = $2
		WHERE id = $3
		RETURNING id, name, description, createdAt, updatedAt
	`

	err := r.db.QueryRow(ctx, query,
		role.Name, role.Description, id,
	).Scan(
		&role.ID, &role.Name, &role.Description,
		&role.CreatedAt, &role.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *RoleRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM roles WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
