package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/mustafaameen91/project-managment/backend/internal/dtos"
	"github.com/mustafaameen91/project-managment/backend/internal/models"
	"github.com/mustafaameen91/project-managment/backend/internal/repository"
)

var (
	ErrUserRoleNotFound = errors.New("user role not found")
)

type UserRoleService struct {
	userRoleRepo *repository.UserRoleRepository
}

func NewUserRoleService(userRoleRepo *repository.UserRoleRepository) *UserRoleService {
	return &UserRoleService{
		userRoleRepo: userRoleRepo,
	}
}

func (s *UserRoleService) GetAll(ctx context.Context) ([]dtos.UserRole, error) {
	userRoles, err := s.userRoleRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.UserRole, len(userRoles))
	for i, ur := range userRoles {
		result[i] = toUserRoleDTO(ur)
	}
	return result, nil
}

func (s *UserRoleService) GetByID(ctx context.Context, id int64) (*dtos.UserRole, error) {
	userRole, err := s.userRoleRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserRoleNotFound
		}
		return nil, err
	}
	dto := toUserRoleDTO(*userRole)
	return &dto, nil
}

func (s *UserRoleService) GetByUserID(ctx context.Context, userID int64) ([]dtos.UserRole, error) {
	userRoles, err := s.userRoleRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.UserRole, len(userRoles))
	for i, ur := range userRoles {
		result[i] = toUserRoleDTO(ur)
	}
	return result, nil
}

func (s *UserRoleService) Create(ctx context.Context, req dtos.CreateUserRole) (*dtos.UserRole, error) {
	userRole := &models.UserRole{
		UserID: req.UserID,
		RoleID: req.RoleID,
	}

	created, err := s.userRoleRepo.Create(ctx, userRole)
	if err != nil {
		return nil, err
	}

	dto := toUserRoleDTO(*created)
	return &dto, nil
}

func (s *UserRoleService) Delete(ctx context.Context, id int64) error {
	err := s.userRoleRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrUserRoleNotFound
		}
		return err
	}
	return nil
}

// DTO conversion helper
func toUserRoleDTO(ur models.UserRole) dtos.UserRole {
	return dtos.UserRole{
		ID:        ur.ID,
		UserID:    ur.UserID,
		RoleID:    ur.RoleID,
		CreatedAt: ur.CreatedAt,
	}
}
