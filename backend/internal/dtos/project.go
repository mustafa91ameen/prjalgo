package dtos

import "time"

// Response DTOs
type ProjectSummary struct {
	ID                 int64      `json:"id"`
	Name               string     `json:"name"`
	Type               *string    `json:"type"`
	ClientPhone        *string    `json:"clientPhone"`
	Location           *string    `json:"location"`
	StartDate          *time.Time `json:"startDate"`
	Status             string     `json:"status"`
	ProgressPercentage float64    `json:"progressPercentage"`
	WarningCost        float64    `json:"warningCost"`
	TotalCost          float64    `json:"totalCost"`
}

type Project struct {
	ID                 int64      `json:"id"`
	Name               string     `json:"name"`
	Type               *string    `json:"type"`
	Description        *string    `json:"description"`
	ClientPhone        *string    `json:"clientPhone"`
	Location           *string    `json:"location"`
	StartDate          *time.Time `json:"startDate"`
	Duration           *int       `json:"duration"`
	WarningCost        float64    `json:"warningCost"`
	TotalCost          float64    `json:"totalCost"`
	Status             string     `json:"status"`
	ProgressPercentage float64    `json:"progressPercentage"`
	Notes              *string    `json:"notes"`
	CreatedBy          *int64     `json:"createdBy"`
	CreatedAt          time.Time  `json:"createdAt"`
}

// Request DTOs
type CreateProject struct {
	Name        string     `json:"name" validate:"required"`
	Type        *string    `json:"type"`
	Description *string    `json:"description"`
	ClientPhone *string    `json:"clientPhone"`
	Location    *string    `json:"location"`
	StartDate   *time.Time `json:"startDate"`
	Duration    *int       `json:"duration"`
	WarningCost float64    `json:"warningCost" validate:"required"`
	TotalCost   float64    `json:"totalCost" validate:"required"`
	Notes       *string    `json:"notes"`
	CreatedBy   *int64     `json:"createdBy"`
}

type UpdateProject struct {
	Name               *string    `json:"name"`
	Type               *string    `json:"type"`
	Description        *string    `json:"description"`
	ClientPhone        *string    `json:"clientPhone"`
	Location           *string    `json:"location"`
	StartDate          *time.Time `json:"startDate"`
	Duration           *int       `json:"duration"`
	WarningCost        *float64   `json:"warningCost"`
	TotalCost          *float64   `json:"totalCost"`
	Status             *string    `json:"status"`
	ProgressPercentage *float64   `json:"progressPercentage"`
	Notes              *string    `json:"notes"`
}
