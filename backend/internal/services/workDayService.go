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
	ErrWorkDayNotFound = errors.New("work day not found")
)

type WorkDayService struct {
	workDayRepo *repository.WorkDayRepository
}

func NewWorkDayService(workDayRepo *repository.WorkDayRepository) *WorkDayService {
	return &WorkDayService{
		workDayRepo: workDayRepo,
	}
}

func (s *WorkDayService) GetAll(ctx context.Context) ([]dtos.WorkDaySummary, error) {
	workDays, err := s.workDayRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.WorkDaySummary, len(workDays))
	for i, w := range workDays {
		result[i] = toWorkDaySummaryDTO(w)
	}
	return result, nil
}

func (s *WorkDayService) GetByID(ctx context.Context, id int64) (*dtos.WorkDay, error) {
	workDay, err := s.workDayRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayNotFound
		}
		return nil, err
	}
	dto := toWorkDayDTO(workDay)
	return &dto, nil
}

func (s *WorkDayService) GetByProjectID(ctx context.Context, projectID int64) ([]dtos.WorkDaySummary, error) {
	workDays, err := s.workDayRepo.GetByProjectID(ctx, projectID)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.WorkDaySummary, len(workDays))
	for i, w := range workDays {
		result[i] = toWorkDaySummaryDTO(w)
	}
	return result, nil
}

func (s *WorkDayService) Create(ctx context.Context, req dtos.CreateWorkDay) (*dtos.WorkDay, error) {
	workDay := &models.WorkDay{
		ProjectID:         req.ProjectID,
		WorkSubCategoryID: req.WorkSubCategoryID,
		WorkDate:          req.WorkDate,
		Description:       req.Description,
		Status:            "draft",
		TotalCost:         0,
		Notes:             req.Notes,
		CreatedBy:         req.CreatedBy,
	}

	created, err := s.workDayRepo.Create(ctx, workDay)
	if err != nil {
		return nil, err
	}

	dto := toWorkDayDTO(created)
	return &dto, nil
}

func (s *WorkDayService) Update(ctx context.Context, id int64, req dtos.UpdateWorkDay) (*dtos.WorkDay, error) {
	existing, err := s.workDayRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayNotFound
		}
		return nil, err
	}

	// Apply partial updates
	if req.WorkSubCategoryID != nil {
		existing.WorkSubCategoryID = req.WorkSubCategoryID
	}
	if req.WorkDate != nil {
		existing.WorkDate = *req.WorkDate
	}
	if req.Description != nil {
		existing.Description = req.Description
	}
	if req.Status != nil {
		existing.Status = *req.Status
	}
	if req.TotalCost != nil {
		existing.TotalCost = *req.TotalCost
	}
	if req.Notes != nil {
		existing.Notes = req.Notes
	}

	updated, err := s.workDayRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	dto := toWorkDayDTO(updated)
	return &dto, nil
}

func (s *WorkDayService) Delete(ctx context.Context, id int64) error {
	err := s.workDayRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrWorkDayNotFound
		}
		return err
	}
	return nil
}

// DTO conversion helpers
func toWorkDaySummaryDTO(w models.WorkDay) dtos.WorkDaySummary {
	return dtos.WorkDaySummary{
		ID:                w.ID,
		ProjectID:         w.ProjectID,
		WorkSubCategoryID: w.WorkSubCategoryID,
		WorkDate:          w.WorkDate,
		Status:            w.Status,
		TotalCost:         w.TotalCost,
	}
}

func toWorkDayDTO(w *models.WorkDay) dtos.WorkDay {
	return dtos.WorkDay{
		ID:                w.ID,
		ProjectID:         w.ProjectID,
		WorkSubCategoryID: w.WorkSubCategoryID,
		WorkDate:          w.WorkDate,
		Description:       w.Description,
		Status:            w.Status,
		TotalCost:         w.TotalCost,
		Notes:             w.Notes,
		CreatedBy:         w.CreatedBy,
		CreatedAt:         w.CreatedAt,
	}
}
