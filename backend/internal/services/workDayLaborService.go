package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/mustafa91ameen/prjalgo/backend/internal/dtos"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
	"github.com/mustafa91ameen/prjalgo/backend/internal/repository"
)

var (
	ErrWorkDayLaborNotFound = errors.New("work day labor not found")
)

type WorkDayLaborService struct {
	laborRepo repository.WorkDayLaborRepositoryInterface
}

func NewWorkDayLaborService(laborRepo repository.WorkDayLaborRepositoryInterface) *WorkDayLaborService {
	return &WorkDayLaborService{
		laborRepo: laborRepo,
	}
}

func (s *WorkDayLaborService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.WorkDayLabor], error) {
	labors, total, err := s.laborRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.WorkDayLabor]{}, err
	}

	result := make([]dtos.WorkDayLabor, len(labors))
	for i, l := range labors {
		result[i] = toWorkDayLaborDTO(l)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *WorkDayLaborService) GetByWorkDayID(ctx context.Context, workDayID int64) ([]dtos.WorkDayLabor, error) {
	labors, err := s.laborRepo.GetByWorkDayID(ctx, workDayID)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.WorkDayLabor, len(labors))
	for i, l := range labors {
		result[i] = toWorkDayLaborDTO(l)
	}
	return result, nil
}

func (s *WorkDayLaborService) GetByID(ctx context.Context, id int64) (*dtos.WorkDayLabor, error) {
	labor, err := s.laborRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayLaborNotFound
		}
		return nil, err
	}
	dto := toWorkDayLaborDTO(*labor)
	return &dto, nil
}

func (s *WorkDayLaborService) Create(ctx context.Context, req dtos.CreateWorkDayLabor) (*dtos.WorkDayLabor, error) {
	labor := &models.WorkDayLabor{
		WorkDayID:  req.WorkDayID,
		WorkerName: req.WorkerName,
		JobTitle:   req.JobTitle,
		Phone:      req.Phone,
		Address:    req.Address,
		Quantity:   req.Quantity,
		Cost:       req.Cost,
		Notes:      req.Notes,
	}

	created, err := s.laborRepo.Create(ctx, labor)
	if err != nil {
		return nil, err
	}

	dto := toWorkDayLaborDTO(*created)
	return &dto, nil
}

func (s *WorkDayLaborService) Update(ctx context.Context, id int64, req dtos.UpdateWorkDayLabor) (*dtos.WorkDayLabor, error) {
	existing, err := s.laborRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayLaborNotFound
		}
		return nil, err
	}

	// Apply partial updates
	if req.WorkerName != nil {
		existing.WorkerName = *req.WorkerName
	}
	if req.JobTitle != nil {
		existing.JobTitle = req.JobTitle
	}
	if req.Phone != nil {
		existing.Phone = req.Phone
	}
	if req.Address != nil {
		existing.Address = req.Address
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

	updated, err := s.laborRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	dto := toWorkDayLaborDTO(*updated)
	return &dto, nil
}

func (s *WorkDayLaborService) Delete(ctx context.Context, id int64) error {
	err := s.laborRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrWorkDayLaborNotFound
		}
		return err
	}
	return nil
}

// DTO conversion helper
func toWorkDayLaborDTO(l models.WorkDayLabor) dtos.WorkDayLabor {
	return dtos.WorkDayLabor{
		ID:         l.ID,
		WorkDayID:  l.WorkDayID,
		WorkerName: l.WorkerName,
		JobTitle:   l.JobTitle,
		Phone:      l.Phone,
		Address:    l.Address,
		Quantity:   l.Quantity,
		Cost:       l.Cost,
		Notes:      l.Notes,
		CreatedAt:  l.CreatedAt,
	}
}