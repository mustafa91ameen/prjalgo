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
	ErrProjectNotFound = errors.New("project not found")
)

type ProjectService struct {
	projectRepo *repository.ProjectRepository
}

func NewProjectService(projectRepo *repository.ProjectRepository) *ProjectService {
	return &ProjectService{
		projectRepo: projectRepo,
	}
}

func (s *ProjectService) GetAll(ctx context.Context) ([]dtos.ProjectSummary, error) {
	projects, err := s.projectRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.ProjectSummary, len(projects))
	for i, p := range projects {
		result[i] = toProjectSummaryDTO(p)
	}
	return result, nil
}

func (s *ProjectService) GetByID(ctx context.Context, id int64) (*dtos.Project, error) {
	project, err := s.projectRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProjectNotFound
		}
		return nil, err
	}
	dto := toProjectDTO(project)
	return &dto, nil
}

func (s *ProjectService) Create(ctx context.Context, req dtos.CreateProject) (*dtos.Project, error) {
	project := &models.Project{
		Name:               req.Name,
		Type:               req.Type,
		Description:        req.Description,
		ClientPhone:        req.ClientPhone,
		Location:           req.Location,
		StartDate:          req.StartDate,
		Duration:           req.Duration,
		WarningCost:        req.WarningCost,
		TotalCost:          req.TotalCost,
		Status:             "draft",
		ProgressPercentage: 0,
		Notes:              req.Notes,
		CreatedBy:          req.CreatedBy,
	}

	created, err := s.projectRepo.Create(ctx, project)
	if err != nil {
		return nil, err
	}

	dto := toProjectDTO(created)
	return &dto, nil
}

func (s *ProjectService) Update(ctx context.Context, id int64, req dtos.UpdateProject) (*dtos.Project, error) {
	existing, err := s.projectRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrProjectNotFound
		}
		return nil, err
	}

	// Apply partial updates
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Type != nil {
		existing.Type = req.Type
	}
	if req.Description != nil {
		existing.Description = req.Description
	}
	if req.ClientPhone != nil {
		existing.ClientPhone = req.ClientPhone
	}
	if req.Location != nil {
		existing.Location = req.Location
	}
	if req.StartDate != nil {
		existing.StartDate = req.StartDate
	}
	if req.Duration != nil {
		existing.Duration = req.Duration
	}
	if req.WarningCost != nil {
		existing.WarningCost = *req.WarningCost
	}
	if req.TotalCost != nil {
		existing.TotalCost = *req.TotalCost
	}
	if req.Status != nil {
		existing.Status = *req.Status
	}
	if req.ProgressPercentage != nil {
		existing.ProgressPercentage = *req.ProgressPercentage
	}
	if req.Notes != nil {
		existing.Notes = req.Notes
	}

	updated, err := s.projectRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	dto := toProjectDTO(updated)
	return &dto, nil
}

func (s *ProjectService) Delete(ctx context.Context, id int64) error {
	err := s.projectRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrProjectNotFound
		}
		return err
	}
	return nil
}

// DTO conversion helpers
func toProjectSummaryDTO(p models.Project) dtos.ProjectSummary {
	return dtos.ProjectSummary{
		ID:                 p.ID,
		Name:               p.Name,
		Type:               p.Type,
		ClientPhone:        p.ClientPhone,
		Location:           p.Location,
		StartDate:          p.StartDate,
		Status:             p.Status,
		ProgressPercentage: p.ProgressPercentage,
		WarningCost:        p.WarningCost,
		TotalCost:          p.TotalCost,
	}
}

func toProjectDTO(p *models.Project) dtos.Project {
	return dtos.Project{
		ID:                 p.ID,
		Name:               p.Name,
		Type:               p.Type,
		Description:        p.Description,
		ClientPhone:        p.ClientPhone,
		Location:           p.Location,
		StartDate:          p.StartDate,
		Duration:           p.Duration,
		WarningCost:        p.WarningCost,
		TotalCost:          p.TotalCost,
		Status:             p.Status,
		ProgressPercentage: p.ProgressPercentage,
		Notes:              p.Notes,
		CreatedBy:          p.CreatedBy,
		CreatedAt:          p.CreatedAt,
	}
}