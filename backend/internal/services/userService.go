package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/dtos"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
	"github.com/mustafa91ameen/prjalgo/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExists   = errors.New("user already exists")
)

type UserService struct {
	db               *pgxpool.Pool
	userRepo         repository.UserRepositoryInterface
	userRoleRepo     repository.UserRoleRepositoryInterface
	refreshTokenRepo repository.RefreshTokenRepositoryInterface
}

func NewUserService(
	db *pgxpool.Pool,
	userRepo repository.UserRepositoryInterface,
	userRoleRepo repository.UserRoleRepositoryInterface,
	refreshTokenRepo repository.RefreshTokenRepositoryInterface,
) *UserService {
	return &UserService{
		db:               db,
		userRepo:         userRepo,
		userRoleRepo:     userRoleRepo,
		refreshTokenRepo: refreshTokenRepo,
	}
}

func (s *UserService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.UserSummary], error) {
	users, total, err := s.userRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.UserSummary]{}, err
	}

	result := make([]dtos.UserSummary, len(users))
	for i, u := range users {
		result[i] = toUserSummaryDTO(u)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *UserService) GetByID(ctx context.Context, id int64) (*dtos.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	dto := toUserDTO(user)
	return &dto, nil
}

func (s *UserService) Create(ctx context.Context, req dtos.CreateUser) (*dtos.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	defaultStatus := "active"
	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		FullName: req.FullName,
		Phone:    req.Phone,
		JobTitle: req.JobTitle,
		Status:   &defaultStatus,
	}

	created, err := s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	dto := toUserDTO(created)
	return &dto, nil
}

func (s *UserService) Update(ctx context.Context, id int64, req dtos.UpdateUser) (*dtos.User, error) {
	existing, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// Apply partial updates
	if req.Username != nil {
		existing.Username = *req.Username
	}
	if req.Email != nil {
		existing.Email = *req.Email
	}
	if req.FullName != nil {
		existing.FullName = *req.FullName
	}
	if req.Phone != nil {
		existing.Phone = *req.Phone
	}
	if req.Avatar != nil {
		existing.Avatar = req.Avatar
	}
	if req.JobTitle != nil {
		existing.JobTitle = *req.JobTitle
	}
	if req.Status != nil {
		existing.Status = req.Status
	}

	updated, err := s.userRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	dto := toUserDTO(updated)
	return &dto, nil
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	err := s.userRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrUserNotFound
		}
		return err
	}
	return nil
}

func (s *UserService) GetUserRoles(ctx context.Context, userID int64) ([]models.UserRole, error) {
	return s.userRoleRepo.GetByUserID(ctx, userID)
}

func (s *UserService) AssignRole(ctx context.Context, userID, roleID int64) (*models.UserRole, error) {
	userRole := &models.UserRole{
		UserID: userID,
		RoleID: roleID,
	}
	return s.userRoleRepo.Create(ctx, userRole)
}

func (s *UserService) RemoveRole(ctx context.Context, userRoleID int64) error {
	return s.userRoleRepo.Delete(ctx, userRoleID)
}

func (s *UserService) UpdatePassword(ctx context.Context, id int64, newPassword string) error {
	_, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrUserNotFound
		}
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.userRepo.UpdatePassword(ctx, id, string(hashedPassword))
}

func (s *UserService) UpdateStatus(ctx context.Context, id int64, status string) error {
	_, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrUserNotFound
		}
		return err
	}

	// Start transaction for atomic status update and token revocation
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Update user status within transaction
	err = s.userRepo.UpdateStatusWithTx(ctx, tx, id, status)
	if err != nil {
		return err
	}

	// Revoke all refresh tokens when user is deactivated
	if status == "inactive" {
		err = s.refreshTokenRepo.RevokeAllByUserIDWithTx(ctx, tx, id)
		if err != nil {
			return err
		}
	}

	// Commit transaction
	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

// GetAllForDropdown returns lightweight user data for dropdown menus
func (s *UserService) GetAllForDropdown(ctx context.Context) ([]dtos.UserDropdown, error) {
	users, err := s.userRepo.GetAllForDropdown(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.UserDropdown, len(users))
	for i, u := range users {
		result[i] = dtos.UserDropdown{
			ID:       u.ID,
			FullName: u.FullName,
			JobTitle: u.JobTitle,
		}
	}
	return result, nil
}

// DTO conversion helpers

func toUserSummaryDTO(u models.User) dtos.UserSummary {
	return dtos.UserSummary{
		ID:        u.ID,
		UserName:  u.Username,
		FullName:  u.FullName,
		Email:     u.Email,
		Phone:     u.Phone,
		JobTitle:  u.JobTitle,
		Status:    u.Status,
		LastLogin: u.LastLogin,
	}
}

func toUserDTO(u *models.User) dtos.User {
	return dtos.User{
		ID:        u.ID,
		Username:  u.Username,
		FullName:  u.FullName,
		Email:     u.Email,
		Phone:     u.Phone,
		Avatar:    u.Avatar,
		JobTitle:  u.JobTitle,
		Status:    u.Status,
		LastLogin: u.LastLogin,
		CreatedAt: u.CreatedAt,
	}
}
