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
	ErrWorkSubCategoryNotFound = errors.New("work sub category not found")
)

type WorkSubCategoryService struct {
	subCategoryRepo *repository.WorkSubCategoryRepository
}

func NewWorkSubCategoryService(subCategoryRepo *repository.WorkSubCategoryRepository) *WorkSubCategoryService {
	return &WorkSubCategoryService{
		subCategoryRepo: subCategoryRepo,
	}
}

func (s *WorkSubCategoryService) GetAll(ctx context.Context) ([]dtos.WorkSubCategorySummary, error) {
	subCategories, err := s.subCategoryRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.WorkSubCategorySummary, len(subCategories))
	for i, sc := range subCategories {
		result[i] = toWorkSubCategorySummaryDTO(sc)
	}
	return result, nil
}

func (s *WorkSubCategoryService) GetByID(ctx context.Context, id int64) (*dtos.WorkSubCategory, error) {
	subCategory, err := s.subCategoryRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkSubCategoryNotFound
		}
		return nil, err
	}
	dto := toWorkSubCategoryDTO(subCategory)
	return &dto, nil
}

func (s *WorkSubCategoryService) GetByCategoryID(ctx context.Context, categoryID int64) ([]dtos.WorkSubCategorySummary, error) {
	subCategories, err := s.subCategoryRepo.GetByCategoryID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.WorkSubCategorySummary, len(subCategories))
	for i, sc := range subCategories {
		result[i] = toWorkSubCategorySummaryDTO(sc)
	}
	return result, nil
}

func (s *WorkSubCategoryService) Create(ctx context.Context, req dtos.CreateWorkSubCategory) (*dtos.WorkSubCategory, error) {
	subCategory := &models.WorkSubCategory{
		CategoryID:  req.CategoryID,
		Name:        req.Name,
		Description: req.Description,
		Percentage:  req.Percentage,
		Status:      req.Status,
		CreatedBy:   req.CreatedBy,
	}

	created, err := s.subCategoryRepo.Create(ctx, subCategory)
	if err != nil {
		return nil, err
	}

	dto := toWorkSubCategoryDTO(created)
	return &dto, nil
}

func (s *WorkSubCategoryService) Update(ctx context.Context, id int64, req dtos.UpdateWorkSubCategory) (*dtos.WorkSubCategory, error) {
	existing, err := s.subCategoryRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkSubCategoryNotFound
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
	if req.Percentage != nil {
		existing.Percentage = req.Percentage
	}
	if req.Status != nil {
		existing.Status = req.Status
	}

	updated, err := s.subCategoryRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	dto := toWorkSubCategoryDTO(updated)
	return &dto, nil
}

func (s *WorkSubCategoryService) Delete(ctx context.Context, id int64) error {
	err := s.subCategoryRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrWorkSubCategoryNotFound
		}
		return err
	}
	return nil
}

// DTO conversion helpers
func toWorkSubCategorySummaryDTO(sc models.WorkSubCategory) dtos.WorkSubCategorySummary {
	return dtos.WorkSubCategorySummary{
		ID:          sc.ID,
		CategoryID:  sc.CategoryID,
		Name:        sc.Name,
		Description: sc.Description,
		Percentage:  sc.Percentage,
		Status:      sc.Status,
	}
}

func toWorkSubCategoryDTO(sc *models.WorkSubCategory) dtos.WorkSubCategory {
	return dtos.WorkSubCategory{
		ID:          sc.ID,
		CategoryID:  sc.CategoryID,
		Name:        sc.Name,
		Description: sc.Description,
		Percentage:  sc.Percentage,
		Status:      sc.Status,
		CreatedBy:   sc.CreatedBy,
		CreatedAt:   sc.CreatedAt,
	}
}
