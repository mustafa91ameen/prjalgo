package container

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mustafaameen91/project-managment/backend/internal/handlers"
	"github.com/mustafaameen91/project-managment/backend/internal/repository"
	"github.com/mustafaameen91/project-managment/backend/internal/services"
)

// Container holds all application dependencies
type Container struct {
	// Project
	ProjectHandler *handlers.ProjectHandler

	// WorkDay
	WorkDayHandler          *handlers.WorkDayHandler
	WorkDayMaterialHandler  *handlers.WorkDayMaterialHandler
	WorkDayLaborHandler     *handlers.WorkDayLaborHandler
	WorkDayEquipmentHandler *handlers.WorkDayEquipmentHandler

	// Work Categories
	WorkCategoryHandler    *handlers.WorkCategoryHandler
	WorkSubCategoryHandler *handlers.WorkSubCategoryHandler

	// Finance
	ExpenseHandler *handlers.ExpenseHandler
	IncomeHandler  *handlers.IncomeHandler

	// Debtor
	DebtorHandler *handlers.DebtorHandler

	// User
	UserHandler     *handlers.UserHandler
	UserRoleHandler *handlers.UserRoleHandler

	// Role & Permissions
	RoleHandler     *handlers.RoleHandler
	PageHandler     *handlers.PageHandler
	RolePageHandler *handlers.RolePageHandler
}

// New creates a new container with all dependencies wired up
func New(db *pgxpool.Pool) *Container {
	// Repositories
	projectRepo := repository.NewProjectRepository(db)
	workDayRepo := repository.NewWorkDayRepository(db)
	workDayMaterialRepo := repository.NewWorkDayMaterialRepository(db)
	workDayLaborRepo := repository.NewWorkDayLaborRepository(db)
	workDayEquipmentRepo := repository.NewWorkDayEquipmentRepository(db)
	workCategoryRepo := repository.NewWorkCategoryRepository(db)
	workSubCategoryRepo := repository.NewWorkSubCategoryRepository(db)
	expenseRepo := repository.NewExpenseRepository(db)
	incomeRepo := repository.NewIncomeRepository(db)
	debtorRepo := repository.NewDebtorRepository(db)
	userRepo := repository.NewUserRepository(db)
	userRoleRepo := repository.NewUserRoleRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	pageRepo := repository.NewPageRepository(db)
	rolePageRepo := repository.NewRolePageRepository(db)

	// Services
	projectService := services.NewProjectService(projectRepo)
	workDayService := services.NewWorkDayService(workDayRepo)
	workDayMaterialService := services.NewWorkDayMaterialService(workDayMaterialRepo)
	workDayLaborService := services.NewWorkDayLaborService(workDayLaborRepo)
	workDayEquipmentService := services.NewWorkDayEquipmentService(workDayEquipmentRepo)
	workCategoryService := services.NewWorkCategoryService(workCategoryRepo)
	workSubCategoryService := services.NewWorkSubCategoryService(workSubCategoryRepo)
	expenseService := services.NewExpenseService(expenseRepo)
	incomeService := services.NewIncomeService(incomeRepo)
	debtorService := services.NewDebtorService(debtorRepo)
	userService := services.NewUserService(userRepo, userRoleRepo)
	userRoleService := services.NewUserRoleService(userRoleRepo)
	roleService := services.NewRoleService(roleRepo)
	pageService := services.NewPageService(pageRepo)
	rolePageService := services.NewRolePageService(rolePageRepo)

	// Handlers
	return &Container{
		ProjectHandler:          handlers.NewProjectHandler(projectService),
		WorkDayHandler:          handlers.NewWorkDayHandler(workDayService),
		WorkDayMaterialHandler:  handlers.NewWorkDayMaterialHandler(workDayMaterialService),
		WorkDayLaborHandler:     handlers.NewWorkDayLaborHandler(workDayLaborService),
		WorkDayEquipmentHandler: handlers.NewWorkDayEquipmentHandler(workDayEquipmentService),
		WorkCategoryHandler:     handlers.NewWorkCategoryHandler(workCategoryService),
		WorkSubCategoryHandler:  handlers.NewWorkSubCategoryHandler(workSubCategoryService),
		ExpenseHandler:          handlers.NewExpenseHandler(expenseService),
		IncomeHandler:           handlers.NewIncomeHandler(incomeService),
		DebtorHandler:           handlers.NewDebtorHandler(debtorService),
		UserHandler:             handlers.NewUserHandler(userService),
		UserRoleHandler:         handlers.NewUserRoleHandler(userRoleService),
		RoleHandler:             handlers.NewRoleHandler(roleService),
		PageHandler:             handlers.NewPageHandler(pageService),
		RolePageHandler:         handlers.NewRolePageHandler(rolePageService),
	}
}
