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
	roleRepo repository.RoleRepositoryInterface
}

func NewRoleService(roleRepo repository.RoleRepositoryInterface) *RoleService {
	return &RoleService{
		roleRepo: roleRepo,
	}
}

func (s *RoleService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.Role], error) {
	roles, total, err := s.roleRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.Role]{}, err
	}

	result := make([]dtos.Role, len(roles))
	for i, r := range roles {
		result[i] = toRoleDTO(r)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
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
