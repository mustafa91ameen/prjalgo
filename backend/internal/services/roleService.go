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
	ErrRoleNotFound = errors.New("role not found")
)

type RoleService struct {
	roleRepo *repository.RoleRepository
}

func NewRoleService(roleRepo *repository.RoleRepository) *RoleService {
	return &RoleService{
		roleRepo: roleRepo,
	}
}

func (s *RoleService) GetAll(ctx context.Context) ([]dtos.Role, error) {
	roles, err := s.roleRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.Role, len(roles))
	for i, r := range roles {
		result[i] = toRoleDTO(r)
	}
	return result, nil
}

func (s *RoleService) GetByID(ctx context.Context, id int64) (*dtos.Role, error) {
	role, err := s.roleRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrRoleNotFound
		}
		return nil, err
	}
	dto := toRoleDTO(*role)
	return &dto, nil
}

func (s *RoleService) Create(ctx context.Context, req dtos.CreateRole) (*dtos.Role, error) {
	role := &models.Role{
		Name:        req.Name,
		Description: req.Description,
	}

	created, err := s.roleRepo.Create(ctx, role)
	if err != nil {
		return nil, err
	}

	dto := toRoleDTO(*created)
	return &dto, nil
}

func (s *RoleService) Update(ctx context.Context, id int64, req dtos.UpdateRole) (*dtos.Role, error) {
	existing, err := s.roleRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrRoleNotFound
		}
		return nil, err
	}

	// Apply partial updates
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Description != nil {
		existing.Description = req.Description
	}

	updated, err := s.roleRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	dto := toRoleDTO(*updated)
	return &dto, nil
}

func (s *RoleService) Delete(ctx context.Context, id int64) error {
	err := s.roleRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrRoleNotFound
		}
		return err
	}
	return nil
}

// DTO conversion helper
func toRoleDTO(r models.Role) dtos.Role {
	return dtos.Role{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		CreatedAt:   r.CreatedAt,
	}
}
