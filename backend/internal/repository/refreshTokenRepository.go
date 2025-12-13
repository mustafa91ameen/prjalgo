package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RefreshToken struct {
	ID        int64
	UserID    int64
	TokenHash string
	ExpiresAt time.Time
	Revoked   bool
	CreatedAt time.Time
}

type RefreshTokenRepositoryInterface interface {
	Create(ctx context.Context, userID int64, tokenHash string, expiresAt time.Time) (*RefreshToken, error)
	GetByTokenHash(ctx context.Context, tokenHash string) (*RefreshToken, error)
	UpdateTokenHash(ctx context.Context, id int64, newTokenHash string, expiresAt time.Time) error
	UpdateTokenHashWithTx(ctx context.Context, tx pgx.Tx, id int64, newTokenHash string, expiresAt time.Time) error
	Revoke(ctx context.Context, id int64) error
	RevokeAllByUserID(ctx context.Context, userID int64) error
	RevokeAllByUserIDWithTx(ctx context.Context, tx pgx.Tx, userID int64) error
	DeleteExpired(ctx context.Context) error
}

type RefreshTokenRepository struct {
	db *pgxpool.Pool
}

func NewRefreshTokenRepository(db *pgxpool.Pool) *RefreshTokenRepository {
	return &RefreshTokenRepository{db: db}
}

func (r *RefreshTokenRepository) Create(ctx context.Context, userID int64, tokenHash string, expiresAt time.Time) (*RefreshToken, error) {
	query := `
		INSERT INTO refreshTokens (userId, tokenHash, expiresAt)
		VALUES ($1, $2, $3)
		RETURNING id, userId, tokenHash, expiresAt, revoked, createdAt
	`

	var token RefreshToken
	err := r.db.QueryRow(ctx, query, userID, tokenHash, expiresAt).Scan(
		&token.ID, &token.UserID, &token.TokenHash,
		&token.ExpiresAt, &token.Revoked, &token.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *RefreshTokenRepository) GetByTokenHash(ctx context.Context, tokenHash string) (*RefreshToken, error) {
	query := `
		SELECT id, userId, tokenHash, expiresAt, revoked, createdAt
		FROM refreshTokens
		WHERE tokenHash = $1
	`

	var token RefreshToken
	err := r.db.QueryRow(ctx, query, tokenHash).Scan(
		&token.ID, &token.UserID, &token.TokenHash,
		&token.ExpiresAt, &token.Revoked, &token.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *RefreshTokenRepository) UpdateTokenHash(ctx context.Context, id int64, newTokenHash string, expiresAt time.Time) error {
	query := `
		UPDATE refreshTokens
		SET tokenHash = $1, expiresAt = $2
		WHERE id = $3
	`

	_, err := r.db.Exec(ctx, query, newTokenHash, expiresAt, id)
	return err
}

func (r *RefreshTokenRepository) UpdateTokenHashWithTx(ctx context.Context, tx pgx.Tx, id int64, newTokenHash string, expiresAt time.Time) error {
	query := `
		UPDATE refreshTokens
		SET tokenHash = $1, expiresAt = $2
		WHERE id = $3
	`

	_, err := tx.Exec(ctx, query, newTokenHash, expiresAt, id)
	return err
}

func (r *RefreshTokenRepository) Revoke(ctx context.Context, id int64) error {
	query := `UPDATE refreshTokens SET revoked = true WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *RefreshTokenRepository) RevokeAllByUserID(ctx context.Context, userID int64) error {
	query := `UPDATE refreshTokens SET revoked = true WHERE userId = $1`
	_, err := r.db.Exec(ctx, query, userID)
	return err
}

func (r *RefreshTokenRepository) RevokeAllByUserIDWithTx(ctx context.Context, tx pgx.Tx, userID int64) error {
	query := `UPDATE refreshTokens SET revoked = true WHERE userId = $1`
	_, err := tx.Exec(ctx, query, userID)
	return err
}

func (r *RefreshTokenRepository) DeleteExpired(ctx context.Context) error {
	query := `DELETE FROM refreshTokens WHERE expiresAt < NOW() OR revoked = true`
	_, err := r.db.Exec(ctx, query)
	return err
}
