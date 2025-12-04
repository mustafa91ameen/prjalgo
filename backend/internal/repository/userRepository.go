package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafaameen91/project-managment/backend/internal/models"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	query := `
		SELECT id, username, email, password, fullName, phone, avatar,
		       jobTitle, status, lastLogin, createdAt, updatedAt
		FROM users
		ORDER BY createdAt DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(
			&u.ID, &u.Username, &u.Email, &u.Password, &u.FullName,
			&u.Phone, &u.Avatar, &u.JobTitle, &u.Status, &u.LastLogin,
			&u.CreatedAt, &u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, rows.Err()
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	query := `
		INSERT INTO users (username, email, password, fullName, phone, avatar, jobTitle, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, createdAt, updatedAt
	`

	err := r.db.QueryRow(ctx, query,
		user.Username, user.Email, user.Password, user.FullName,
		user.Phone, user.Avatar, user.JobTitle, user.Status,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, id int64, user *models.User) (*models.User, error) {
	query := `
		UPDATE users
		SET username = $1, email = $2, fullName = $3, phone = $4,
		    avatar = $5, jobTitle = $6, status = $7
		WHERE id = $8
		RETURNING id, username, email, password, fullName, phone, avatar,
		          jobTitle, status, lastLogin, createdAt, updatedAt
	`

	err := r.db.QueryRow(ctx, query,
		user.Username, user.Email, user.FullName, user.Phone,
		user.Avatar, user.JobTitle, user.Status, id,
	).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password, &user.FullName,
		&user.Phone, &user.Avatar, &user.JobTitle, &user.Status, &user.LastLogin,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	query := `
		SELECT id, username, email, password, fullName, phone, avatar,
		       jobTitle, status, lastLogin, createdAt, updatedAt
		FROM users
		WHERE id = $1
	`

	var user models.User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Username, &user.Email, &user.Password, &user.FullName,
		&user.Phone, &user.Avatar, &user.JobTitle, &user.Status, &user.LastLogin,
		&user.CreatedAt, &user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows  // or a custom ErrNotFound
	}
	return nil
}
