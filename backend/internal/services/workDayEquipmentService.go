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
	ErrWorkDayEquipmentNotFound = errors.New("work day equipment not found")
)

type WorkDayEquipmentService struct {
	equipmentRepo repository.WorkDayEquipmentRepositoryInterface
}

func NewWorkDayEquipmentService(equipmentRepo repository.WorkDayEquipmentRepositoryInterface) *WorkDayEquipmentService {
	return &WorkDayEquipmentService{
		equipmentRepo: equipmentRepo,
	}
}

func (s *WorkDayEquipmentService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.WorkDayEquipment], error) {
	equipments, total, err := s.equipmentRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.WorkDayEquipment]{}, err
	}

	result := make([]dtos.WorkDayEquipment, len(equipments))
	for i, e := range equipments {
		result[i] = toWorkDayEquipmentDTO(e)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *WorkDayEquipmentService) GetByWorkDayID(ctx context.Context, workDayID int64) ([]dtos.WorkDayEquipment, error) {
	equipments, err := s.equipmentRepo.GetByWorkDayID(ctx, workDayID)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.WorkDayEquipment, len(equipments))
	for i, e := range equipments {
		result[i] = toWorkDayEquipmentDTO(e)
	}
	return result, nil
}

func (s *WorkDayEquipmentService) GetByID(ctx context.Context, id int64) (*dtos.WorkDayEquipment, error) {
	equipment, err := s.equipmentRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayEquipmentNotFound
		}
		return nil, err
	}
	dto := toWorkDayEquipmentDTO(*equipment)
	return &dto, nil
}

func (s *WorkDayEquipmentService) Create(ctx context.Context, req dtos.CreateWorkDayEquipment) (*dtos.WorkDayEquipment, error) {
	equipment := &models.WorkDayEquipment{
		WorkDayID:     req.WorkDayID,
		EquipmentName: req.EquipmentName,
		Quantity:      req.Quantity,
		Cost:          req.Cost,
		Notes:         req.Notes,
	}

	created, err := s.equipmentRepo.Create(ctx, equipment)
	if err != nil {
		return nil, err
	}

	dto := toWorkDayEquipmentDTO(*created)
	return &dto, nil
}

func (s *WorkDayEquipmentService) Update(ctx context.Context, id int64, req dtos.UpdateWorkDayEquipment) (*dtos.WorkDayEquipment, error) {
	existing, err := s.equipmentRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrWorkDayEquipmentNotFound
		}
		return nil, err
	}

	// Apply partial updates
	if req.EquipmentName != nil {
		existing.EquipmentName = *req.EquipmentName
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

	updated, err := s.equipmentRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	dto := toWorkDayEquipmentDTO(*updated)
	return &dto, nil
}

func (s *WorkDayEquipmentService) Delete(ctx context.Context, id int64) error {
	err := s.equipmentRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrWorkDayEquipmentNotFound
		}
		return err
	}
	return nil
}

// DTO conversion helper
func toWorkDayEquipmentDTO(e models.WorkDayEquipment) dtos.WorkDayEquipment {
	return dtos.WorkDayEquipment{
		ID:            e.ID,
		WorkDayID:     e.WorkDayID,
		EquipmentName: e.EquipmentName,
		Quantity:      e.Quantity,
		Cost:          e.Cost,
		Notes:         e.Notes,
		CreatedAt:     e.CreatedAt,
	}
}
