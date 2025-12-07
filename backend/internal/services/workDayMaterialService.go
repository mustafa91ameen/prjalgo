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
	ErrWorkDayMaterialNotFound = errors.New("work day material not found")
)

type WorkDayMaterialService struct {
	materialRepo repository.WorkDayMaterialRepositoryInterface
}

func NewWorkDayMaterialService(materialRepo repository.WorkDayMaterialRepositoryInterface) *WorkDayMaterialService {
	return &WorkDayMaterialService{
		materialRepo: materialRepo,
	}
}

func (s *WorkDayMaterialService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.WorkDayMaterial], error) {
	materials, total, err := s.materialRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.WorkDayMaterial]{}, err
	}

	result := make([]dtos.WorkDayMaterial, len(materials))
	for i, m := range materials {
		result[i] = toWorkDayMaterialDTO(m)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *WorkDayMaterialService) GetByWorkDayID(ctx context.Context, workDayID int64) ([]dtos.WorkDayMaterial, error) {
	materials, err := s.materialRepo.GetByWorkDayID(ctx, workDayID)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.WorkDayMaterial, len(materials))
	for i, m := range materials {
		result[i] = toWorkDayMaterialDTO(m)
	}
	return result, nil
}

func (s *WorkDayMaterialService) GetByID(ctx context.Context, id int64) (*dtos.WorkDayMaterial, error) {
	material, err := s.materialRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayMaterialNotFound
		}
		return nil, err
	}
	dto := toWorkDayMaterialDTO(*material)
	return &dto, nil
}

func (s *WorkDayMaterialService) Create(ctx context.Context, req dtos.CreateWorkDayMaterial) (*dtos.WorkDayMaterial, error) {
	material := &models.WorkDayMaterial{
		WorkDayID:    req.WorkDayID,
		MaterialName: req.MaterialName,
		Quantity:     req.Quantity,
		Cost:         req.Cost,
		Notes:        req.Notes,
	}

	created, err := s.materialRepo.Create(ctx, material)
	if err != nil {
		return nil, err
	}

	dto := toWorkDayMaterialDTO(*created)
	return &dto, nil
}

func (s *WorkDayMaterialService) Update(ctx context.Context, id int64, req dtos.UpdateWorkDayMaterial) (*dtos.WorkDayMaterial, error) {
	existing, err := s.materialRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayMaterialNotFound
		}
		return nil, err
	}

	// Apply partial updates
	if req.MaterialName != nil {
		existing.MaterialName = *req.MaterialName
	}
	if req.Quantity != nil {
		existing.Quantity = *req.Quantity
	}
	if req.Cost != nil {
		existing.Cost = *req.Cost
	}
	if req.Notes != nil {
		existing.Notes = req.Notes
	}

	updated, err := s.materialRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	dto := toWorkDayMaterialDTO(*updated)
	return &dto, nil
}

func (s *WorkDayMaterialService) Delete(ctx context.Context, id int64) error {
	err := s.materialRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrWorkDayMaterialNotFound
		}
		return err
	}
	return nil
}

// DTO conversion helper
func toWorkDayMaterialDTO(m models.WorkDayMaterial) dtos.WorkDayMaterial {
	return dtos.WorkDayMaterial{
		ID:           m.ID,
		WorkDayID:    m.WorkDayID,
		MaterialName: m.MaterialName,
		Quantity:     m.Quantity,
		Cost:         m.Cost,
		Notes:        m.Notes,
		CreatedAt:    m.CreatedAt,
	}
}
