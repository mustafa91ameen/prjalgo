package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafaameen91/project-managment/backend/internal/models"
)

type UserRoleRepositoryInterface interface {
	GetAll(ctx context.Context) ([]models.UserRole, error)
	GetByID(ctx context.Context, id int64) (*models.UserRole, error)
	GetByUserID(ctx context.Context, userID int64) ([]models.UserRole, error)
	Create(ctx context.Context, userRole *models.UserRole) (*models.UserRole, error)
	Delete(ctx context.Context, id int64) error
}

type UserRoleRepository struct {
	db *pgxpool.Pool
}

func NewUserRoleRepository(db *pgxpool.Pool) *UserRoleRepository {
	return &UserRoleRepository{db: db}
}

func (r *UserRoleRepository) GetAll(ctx context.Context) ([]models.UserRole, error) {
	query := `
		SELECT id, userId, roleId, createdAt
		FROM userRoles
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userRoles []models.UserRole
	for rows.Next() {
		var ur models.UserRole
		err := rows.Scan(&ur.ID, &ur.UserID, &ur.RoleID, &ur.CreatedAt)
		if err != nil {
			return nil, err
		}
		userRoles = append(userRoles, ur)
	}

	return userRoles, rows.Err()
}

func (r *UserRoleRepository) GetByID(ctx context.Context, id int64) (*models.UserRole, error) {
	query := `
		SELECT id, userId, roleId, createdAt
		FROM userRoles
		WHERE id = $1
	`

	var ur models.UserRole
	err := r.db.QueryRow(ctx, query, id).Scan(
		&ur.ID, &ur.UserID, &ur.RoleID, &ur.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &ur, nil
}

func (r *UserRoleRepository) GetByUserID(ctx context.Context, userID int64) ([]models.UserRole, error) {
	query := `
		SELECT id, userId, roleId, createdAt
		FROM userRoles
		WHERE userId = $1
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userRoles []models.UserRole
	for rows.Next() {
		var ur models.UserRole
		err := rows.Scan(&ur.ID, &ur.UserID, &ur.RoleID, &ur.CreatedAt)
		if err != nil {
			return nil, err
		}
		userRoles = append(userRoles, ur)
	}

	return userRoles, rows.Err()
}

func (r *UserRoleRepository) Create(ctx context.Context, userRole *models.UserRole) (*models.UserRole, error) {
	query := `
		INSERT INTO userRoles (userId, roleId)
		VALUES ($1, $2)
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		userRole.UserID, userRole.RoleID,
	).Scan(&userRole.ID, &userRole.CreatedAt)

	if err != nil {
		return nil, err
	}

	return userRole, nil
}

func (r *UserRoleRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM userRoles WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
