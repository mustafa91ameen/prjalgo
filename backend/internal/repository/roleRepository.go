package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
)

type RoleRepositoryInterface interface {
	GetAll(ctx context.Context, limit, offset int) ([]models.Role, int64, error)
	GetByID(ctx context.Context, id int64) (*models.Role, error)
	Create(ctx context.Context, role *models.Role) (*models.Role, error)
	Update(ctx context.Context, id int64, role *models.Role) (*models.Role, error)
	Delete(ctx context.Context, id int64) error
}

type RoleRepository struct {
	db *pgxpool.Pool
}

func NewRoleRepository(db *pgxpool.Pool) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) GetAll(ctx context.Context, limit, offset int) ([]models.Role, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM roles`
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, name, description, createdAt
		FROM roles
		ORDER BY createdAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		err := rows.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		roles = append(roles, role)
	}

	return roles, total, rows.Err()
}

func (r *RoleRepository) GetByID(ctx context.Context, id int64) (*models.Role, error) {
	query := `
		SELECT id, name, description, createdAt
		FROM roles
		WHERE id = $1
	`

	var role models.Role
	err := r.db.QueryRow(ctx, query, id).Scan(
		&role.ID, &role.Name, &role.Description, &role.CreatedAt,
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
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		role.Name, role.Description,
	).Scan(&role.ID, &role.CreatedAt)

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
		RETURNING id, name, description, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		role.Name, role.Description, id,
	).Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt)

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
