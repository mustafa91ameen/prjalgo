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
	ErrPageNotFound = errors.New("page not found")
)

type PageService struct {
	pageRepo *repository.PageRepository
}

func NewPageService(pageRepo *repository.PageRepository) *PageService {
	return &PageService{
		pageRepo: pageRepo,
	}
}

func (s *PageService) GetAll(ctx context.Context) ([]dtos.Page, error) {
	pages, err := s.pageRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.Page, len(pages))
	for i, p := range pages {
		result[i] = toPageDTO(p)
	}
	return result, nil
}

func (s *PageService) GetByID(ctx context.Context, id int64) (*dtos.Page, error) {
	page, err := s.pageRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrPageNotFound
		}
		return nil, err
	}
	dto := toPageDTO(*page)
	return &dto, nil
}

func (s *PageService) GetActivePages(ctx context.Context) ([]dtos.Page, error) {
	pages, err := s.pageRepo.GetActivePages(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.Page, len(pages))
	for i, p := range pages {
		result[i] = toPageDTO(p)
	}
	return result, nil
}

func (s *PageService) Create(ctx context.Context, req dtos.CreatePage) (*dtos.Page, error) {
	page := &models.Page{
		Name:   req.Name,
		Icon:   req.Icon,
		Route:  req.Route,
		Status: req.Status,
	}

	created, err := s.pageRepo.Create(ctx, page)
	if err != nil {
		return nil, err
	}

	dto := toPageDTO(*created)
	return &dto, nil
}

func (s *PageService) Update(ctx context.Context, id int64, req dtos.UpdatePage) (*dtos.Page, error) {
	existing, err := s.pageRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrPageNotFound
		}
		return nil, err
	}

	// Apply partial updates
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Icon != nil {
		existing.Icon = req.Icon
	}
	if req.Route != nil {
		existing.Route = *req.Route
	}
	if req.Status != nil {
		existing.Status = req.Status
	}

	updated, err := s.pageRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	dto := toPageDTO(*updated)
	return &dto, nil
}

func (s *PageService) Delete(ctx context.Context, id int64) error {
	err := s.pageRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrPageNotFound
		}
		return err
	}
	return nil
}

// DTO conversion helper
func toPageDTO(p models.Page) dtos.Page {
	return dtos.Page{
		ID:        p.ID,
		Name:      p.Name,
		Icon:      p.Icon,
		Route:     p.Route,
		Status:    p.Status,
		CreatedAt: p.CreatedAt,
	}
}
