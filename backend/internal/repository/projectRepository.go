package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafa91ameen/prjalgo/backend/internal/models"
)

type ProjectRepositoryInterface interface {
	GetAll(ctx context.Context, limit, offset int) ([]models.Project, int64, error)
	GetByID(ctx context.Context, id int64) (*models.Project, error)
	Create(ctx context.Context, project *models.Project) (*models.Project, error)
	Update(ctx context.Context, id int64, project *models.Project) (*models.Project, error)
	Delete(ctx context.Context, id int64) error
	UpdateProgressPercentage(ctx context.Context, id int64, percentageToAdd float64) error
	UpdateProgressPercentageWithTx(ctx context.Context, tx pgx.Tx, id int64, percentageToAdd float64) error
	GetStats(ctx context.Context) (*ProjectStatsResult, error)
}

// ProjectStatsResult holds aggregated project statistics
type ProjectStatsResult struct {
	Total           int64   `json:"total"`
	Pending         int64   `json:"pending"`
	InProgress      int64   `json:"inProgress"`
	Completed       int64   `json:"completed"`
	TotalBudget     float64 `json:"totalBudget"`
	AverageProgress float64 `json:"averageProgress"`
}

type ProjectRepository struct {
	db *pgxpool.Pool
}

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) GetAll(ctx context.Context, limit, offset int) ([]models.Project, int64, error) {
	// Get total count
	var total int64
	countQuery := `SELECT COUNT(*) FROM projects`
	if err := r.db.QueryRow(ctx, countQuery).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `
		SELECT id, name, type, clientPhone, location, startDate,
		       status, progressPercentage, warningCost, totalCost
		FROM projects
		ORDER BY createdAt DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var p models.Project
		err := rows.Scan(
			&p.ID, &p.Name, &p.Type, &p.ClientPhone, &p.Location,
			&p.StartDate, &p.Status, &p.ProgressPercentage,
			&p.WarningCost, &p.TotalCost,
		)
		if err != nil {
			return nil, 0, err
		}
		projects = append(projects, p)
	}

	return projects, total, rows.Err()
}

func (r *ProjectRepository) GetByID(ctx context.Context, id int64) (*models.Project, error) {
	query := `
		SELECT id, name, type, description, clientPhone, location, startDate,
		       duration, warningCost, totalCost, status, progressPercentage,
		       notes, createdBy, createdAt
		FROM projects
		WHERE id = $1
	`

	var p models.Project
	err := r.db.QueryRow(ctx, query, id).Scan(
		&p.ID, &p.Name, &p.Type, &p.Description, &p.ClientPhone, &p.Location,
		&p.StartDate, &p.Duration, &p.WarningCost, &p.TotalCost, &p.Status,
		&p.ProgressPercentage, &p.Notes, &p.CreatedBy, &p.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *ProjectRepository) Create(ctx context.Context, project *models.Project) (*models.Project, error) {
	query := `
		INSERT INTO projects (name, type, description, clientPhone, location,
		                      startDate, duration, warningCost, totalCost, status,
		                      progressPercentage, notes, createdBy)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		project.Name, project.Type, project.Description, project.ClientPhone,
		project.Location, project.StartDate, project.Duration, project.WarningCost,
		project.TotalCost, project.Status, project.ProgressPercentage,
		project.Notes, project.CreatedBy,
	).Scan(&project.ID, &project.CreatedAt)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (r *ProjectRepository) Update(ctx context.Context, id int64, project *models.Project) (*models.Project, error) {
	query := `
		UPDATE projects
		SET name = $1, type = $2, description = $3, clientPhone = $4, location = $5,
		    startDate = $6, duration = $7, warningCost = $8, totalCost = $9, status = $10,
		    progressPercentage = $11, notes = $12
		WHERE id = $13
		RETURNING id, name, type, description, clientPhone, location, startDate,
		          duration, warningCost, totalCost, status, progressPercentage,
		          notes, createdBy, createdAt
	`

	err := r.db.QueryRow(ctx, query,
		project.Name, project.Type, project.Description, project.ClientPhone,
		project.Location, project.StartDate, project.Duration, project.WarningCost,
		project.TotalCost, project.Status, project.ProgressPercentage,
		project.Notes, id,
	).Scan(
		&project.ID, &project.Name, &project.Type, &project.Description,
		&project.ClientPhone, &project.Location, &project.StartDate, &project.Duration,
		&project.WarningCost, &project.TotalCost, &project.Status, &project.ProgressPercentage,
		&project.Notes, &project.CreatedBy, &project.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (r *ProjectRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM projects WHERE id = $1`
	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *ProjectRepository) UpdateProgressPercentage(ctx context.Context, id int64, percentageToAdd float64) error {
	query := `UPDATE projects SET progressPercentage = progressPercentage + $1 WHERE id = $2`
	result, err := r.db.Exec(ctx, query, percentageToAdd, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *ProjectRepository) UpdateProgressPercentageWithTx(ctx context.Context, tx pgx.Tx, id int64, percentageToAdd float64) error {
	query := `UPDATE projects SET progressPercentage = progressPercentage + $1 WHERE id = $2`
	result, err := tx.Exec(ctx, query, percentageToAdd, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *ProjectRepository) GetStats(ctx context.Context) (*ProjectStatsResult, error) {
	query := `
		SELECT
			COUNT(*) as total,
			COUNT(*) FILTER (WHERE status = 'pending') as pending,
			COUNT(*) FILTER (WHERE status = 'in_progress') as in_progress,
			COUNT(*) FILTER (WHERE status = 'completed') as completed,
			COALESCE(SUM(totalCost), 0) as total_budget,
			COALESCE(AVG(progressPercentage), 0) as average_progress
		FROM projects
	`

	var stats ProjectStatsResult
	err := r.db.QueryRow(ctx, query).Scan(
		&stats.Total,
		&stats.Pending,
		&stats.InProgress,
		&stats.Completed,
		&stats.TotalBudget,
		&stats.AverageProgress,
	)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}
