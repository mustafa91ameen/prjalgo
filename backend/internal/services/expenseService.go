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
	ErrExpenseNotFound      = errors.New("expense not found")
	ErrExpenseStatsNotFound = errors.New("expense stats not found")
)

type ExpenseService struct {
	expenseRepo repository.ExpenseRepositoryInterface
	debtorRepo  repository.DebtorRepositoryInterface
}

func NewExpenseService(expenseRepo repository.ExpenseRepositoryInterface, debtorRepo repository.DebtorRepositoryInterface) *ExpenseService {
	return &ExpenseService{
		expenseRepo: expenseRepo,
		debtorRepo:  debtorRepo,
	}
}

func (s *ExpenseService) GetAll(ctx context.Context, limit, offset int) (dtos.PaginatedResponse[dtos.ExpenseSummary], error) {
	expenses, total, err := s.expenseRepo.GetAll(ctx, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.ExpenseSummary]{}, err
	}

	result := make([]dtos.ExpenseSummary, len(expenses))
	for i, e := range expenses {
		result[i] = toExpenseSummaryDTO(e)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *ExpenseService) GetByID(ctx context.Context, id int64) (*dtos.Expense, error) {
	expense, err := s.expenseRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrExpenseNotFound
		}
		return nil, err
	}
	dto := toExpenseDTO(expense)
	return &dto, nil
}

func (s *ExpenseService) GetByProjectID(ctx context.Context, projectID int64) ([]dtos.ExpenseSummary, error) {
	expenses, err := s.expenseRepo.GetByProjectID(ctx, projectID)
	if err != nil {
		return nil, err
	}

	result := make([]dtos.ExpenseSummary, len(expenses))
	for i, e := range expenses {
		result[i] = toExpenseSummaryDTO(e)
	}
	return result, nil
}

func (s *ExpenseService) GetByProjectIDPaginated(ctx context.Context, projectID int64, limit, offset int) (dtos.PaginatedResponse[dtos.ExpenseSummary], error) {
	expenses, total, err := s.expenseRepo.GetByProjectIDPaginated(ctx, projectID, limit, offset)
	if err != nil {
		return dtos.PaginatedResponse[dtos.ExpenseSummary]{}, err
	}

	result := make([]dtos.ExpenseSummary, len(expenses))
	for i, e := range expenses {
		result[i] = toExpenseSummaryDTO(e)
	}

	page := (offset / limit) + 1
	return dtos.NewPaginatedResponse(result, total, page, limit), nil
}

func (s *ExpenseService) Create(ctx context.Context, req dtos.CreateExpense) (*dtos.Expense, error) {
	// Default status to "pending" if not provided
	status := req.Status
	if status == nil {
		defaultStatus := "pending"
		status = &defaultStatus
	}

	expense := &models.Expense{
		Name:        req.Name,
		Amount:      req.Amount,
		Type:        req.Type,
		ExpenseDate: req.ExpenseDate,
		ProjectID:   req.ProjectID,
		DebtorID:    req.DebtorID,
		Status:      status,
		Notes:       req.Notes,
		CreatedBy:   req.CreatedBy,
	}

	created, err := s.expenseRepo.Create(ctx, expense)
	if err != nil {
		return nil, err
	}

	// Check if this is a debt payment and auto-update debtor status if fully paid
	if created.DebtorID != nil && created.Status != nil && *created.Status == "approved" {
		s.checkAndUpdateDebtorStatus(ctx, *created.DebtorID)
	}

	dto := toExpenseDTO(created)
	return &dto, nil
}

func (s *ExpenseService) Update(ctx context.Context, id int64, req dtos.UpdateExpense) (*dtos.Expense, error) {
	existing, err := s.expenseRepo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrExpenseNotFound
		}
		return nil, err
	}

	// Apply partial updates
	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Amount != nil {
		existing.Amount = *req.Amount
	}
	if req.Type != nil {
		existing.Type = req.Type
	}
	if req.ExpenseDate != nil {
		existing.ExpenseDate = *req.ExpenseDate
	}
	if req.ProjectID != nil {
		existing.ProjectID = req.ProjectID
	}
	if req.DebtorID != nil {
		existing.DebtorID = req.DebtorID
	}
	if req.Status != nil {
		existing.Status = req.Status
	}
	if req.Notes != nil {
		existing.Notes = req.Notes
	}

	updated, err := s.expenseRepo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	// Check if this is a debt payment and auto-update debtor status if fully paid
	if updated.DebtorID != nil && updated.Status != nil && *updated.Status == "approved" {
		s.checkAndUpdateDebtorStatus(ctx, *updated.DebtorID)
	}

	dto := toExpenseDTO(updated)
	return &dto, nil
}

func (s *ExpenseService) Delete(ctx context.Context, id int64) error {
	err := s.expenseRepo.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrExpenseNotFound
		}
		return err
	}
	return nil
}

func (s *ExpenseService) GetStats(ctx context.Context, period string) (*dtos.ExpenseStatsResponse, error) {
	stats, err := s.expenseRepo.GetStats(ctx, period)
	if err != nil {
		return nil, err
	}

	return &dtos.ExpenseStatsResponse{
		Total:         stats.Total,
		TotalAmount:   stats.TotalAmount,
		Pending:       stats.Pending,
		Approved:      stats.Approved,
		AverageAmount: stats.AverageAmount,
	}, nil
}

// DTO conversion helpers
func toExpenseSummaryDTO(e models.Expense) dtos.ExpenseSummary {
	return dtos.ExpenseSummary{
		ID:          e.ID,
		Name:        e.Name,
		Amount:      e.Amount,
		Type:        e.Type,
		ExpenseDate: e.ExpenseDate,
		ProjectID:   e.ProjectID,
		DebtorID:    e.DebtorID,
		Status:      e.Status,
		Notes:       e.Notes,
	}
}

func toExpenseDTO(e *models.Expense) dtos.Expense {
	return dtos.Expense{
		ID:          e.ID,
		Name:        e.Name,
		Amount:      e.Amount,
		Type:        e.Type,
		ExpenseDate: e.ExpenseDate,
		ProjectID:   e.ProjectID,
		DebtorID:    e.DebtorID,
		Status:      e.Status,
		Notes:       e.Notes,
		CreatedBy:   e.CreatedBy,
		CreatedAt:   e.CreatedAt,
	}
}

// checkAndUpdateDebtorStatus checks if debtor is fully paid and updates status to 'paid'
func (s *ExpenseService) checkAndUpdateDebtorStatus(ctx context.Context, debtorID int64) {
	// Get debtor details
	debtor, err := s.debtorRepo.GetByID(ctx, debtorID)
	if err != nil {
		return
	}

	// Get total paid amount
	paidAmount, err := s.expenseRepo.GetTotalPaidByDebtorID(ctx, debtorID)
	if err != nil {
		return
	}

	// If fully paid, update debtor status to 'paid'
	if paidAmount >= debtor.TotalDebt {
		s.debtorRepo.UpdateStatusToPaid(ctx, debtorID)
	}
}
