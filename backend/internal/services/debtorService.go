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
	ErrDebtorNotFound      = errors.New("debtor not found")
	ErrDebtorStatsNotFound = errors.New("debtor stats not found")
)

type DebtorService struct {
	debtorRepo  repository.DebtorRepositoryInterface
	expenseRepo repository.ExpenseRepositoryInterface
}

func NewDebtorService(debtorRepo repository.DebtorRepositoryInterface, expenseRepo repository.ExpenseRepositoryInterface) *DebtorService {
	return &DebtorService{
		debtorRepo:  debtorRepo,
		expenseRepo: expenseRepo,
	}
}

func (s *DebtorService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.DebtorSummary], error) {
	debtors, total, err := s.debtorRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.DebtorSummary]{}, err
	}

	result := make([]dtos.DebtorSummary, len(debtors))
	for i, d := range debtors {
		// Get paid amount for each debtor
		paidAmount, _ := s.expenseRepo.GetTotalPaidByDebtorID(ctx, d.ID)
		result[i] = toDebtorSummaryDTOWithPaid(d, paidAmount)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *DebtorService) GetByID(ctx context.Context, id int64) (*dtos.Debtor, error) {
	debtor, err := s.debtorRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrDebtorNotFound
		}
		return nil, err
	}
	// Get paid amount for this debtor
	paidAmount, _ := s.expenseRepo.GetTotalPaidByDebtorID(ctx, id)
	dto := toDebtorDTOWithPaid(debtor, paidAmount)
	return &dto, nil
}

func (s *DebtorService) Create(ctx context.Context, req dtos.CreateDebtor) (*dtos.Debtor, error) {
	// Default status to "active" if not provided or invalid
	status := req.Status
	if status == nil || (*status != "active" && *status != "paid") {
		defaultStatus := "active"
		status = &defaultStatus
	}

	debtor := &models.Debtor{
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		TotalDebt: req.TotalDebt,
		Currency:  req.Currency,
		DueDate:   req.DueDate,
		Status:    status,
		Notes:     req.Notes,
		CreatedBy: req.CreatedBy,
	}

	created, err := s.debtorRepo.Create(ctx, debtor)
	if err != nil {
		return nil, err
	}

	// New debtor has 0 paid amount
	dto := toDebtorDTOWithPaid(created, 0)
	return &dto, nil
}

func (s *DebtorService) Update(ctx context.Context, id int64, req dtos.UpdateDebtor) (*dtos.Debtor, error) {
	existing, err := s.debtorRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrDebtorNotFound
		}
		return nil, err
	}

	// Apply partial updates
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Email != nil {
		existing.Email = req.Email
	}
	if req.Phone != nil {
		existing.Phone = req.Phone
	}
	if req.TotalDebt != nil {
		existing.TotalDebt = *req.TotalDebt
	}
	if req.Currency != nil {
		existing.Currency = *req.Currency
	}
	if req.DueDate != nil {
		existing.DueDate = req.DueDate
	}
	if req.Status != nil {
		// Validate status is a valid debtor_status enum value
		if *req.Status == "active" || *req.Status == "paid" {
			existing.Status = req.Status
		}
		// Invalid status values are ignored, keeping the existing status
	}
	if req.Notes != nil {
		existing.Notes = req.Notes
	}

	updated, err := s.debtorRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	// Get paid amount for this debtor
	paidAmount, _ := s.expenseRepo.GetTotalPaidByDebtorID(ctx, id)
	dto := toDebtorDTOWithPaid(updated, paidAmount)
	return &dto, nil
}

func (s *DebtorService) Delete(ctx context.Context, id int64) error {
	err := s.debtorRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrDebtorNotFound
		}
		return err
	}
	return nil
}

func (s *DebtorService) GetStats(ctx context.Context, period string) (*dtos.DebtorStatsResponse, error) {
	stats, err := s.debtorRepo.GetStats(ctx, period)
	if err != nil {
		return nil, err
	}

	return &dtos.DebtorStatsResponse{
		Total:       stats.Total,
		Active:      stats.Active,
		Paid:        stats.Paid,
		TotalDebt:   stats.TotalDebt,
		ActiveDebt:  stats.ActiveDebt,
		AverageDebt: stats.AverageDebt,
	}, nil
}

// GetActiveDebtorsWithRemaining returns active debtors with remaining debt (for expense form dropdown)
func (s *DebtorService) GetActiveDebtorsWithRemaining(ctx context.Context) ([]dtos.DebtorWithRemaining, error) {
	debtors, err := s.debtorRepo.GetActiveDebtorsWithRemaining(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.DebtorWithRemaining, len(debtors))
	for i, d := range debtors {
		remainingDebt := d.TotalDebt - d.PaidAmount
		if remainingDebt < 0 {
			remainingDebt = 0
		}
		result[i] = dtos.DebtorWithRemaining{
			ID:            d.ID,
			Name:          d.Name,
			TotalDebt:     d.TotalDebt,
			PaidAmount:    d.PaidAmount,
			RemainingDebt: remainingDebt,
			Currency:      d.Currency,
		}
	}

	return result, nil
}

// DTO conversion helpers
func toDebtorSummaryDTOWithPaid(d models.Debtor, paidAmount float64) dtos.DebtorSummary {
	remainingDebt := d.TotalDebt - paidAmount
	if remainingDebt < 0 {
		remainingDebt = 0
	}
	return dtos.DebtorSummary{
		ID:            d.ID,
		Name:          d.Name,
		Email:         d.Email,
		Phone:         d.Phone,
		TotalDebt:     d.TotalDebt,
		PaidAmount:    paidAmount,
		RemainingDebt: remainingDebt,
		Currency:      d.Currency,
		DueDate:       d.DueDate,
		Status:        d.Status,
		Notes:         d.Notes,
	}
}

func toDebtorDTOWithPaid(d *models.Debtor, paidAmount float64) dtos.Debtor {
	remainingDebt := d.TotalDebt - paidAmount
	if remainingDebt < 0 {
		remainingDebt = 0
	}
	return dtos.Debtor{
		ID:            d.ID,
		Name:          d.Name,
		Email:         d.Email,
		Phone:         d.Phone,
		TotalDebt:     d.TotalDebt,
		PaidAmount:    paidAmount,
		RemainingDebt: remainingDebt,
		Currency:      d.Currency,
		DueDate:       d.DueDate,
		Status:        d.Status,
		Notes:         d.Notes,
		CreatedBy:     d.CreatedBy,
		CreatedAt:     d.CreatedAt,
	}
}
