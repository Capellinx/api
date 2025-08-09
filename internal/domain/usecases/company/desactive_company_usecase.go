package company

import (
	"api/internal/domain/repositories"
	"context"
	"errors"
	"log/slog"
	"time"
)

type DesactiveCompanyInputDTO struct {
	ID string `json:"id"`
}
type DesactiveCompanyUseCase struct {
	CompanyRepository repositories.CompanyRepository
}

func NewDesactiveCompanyUseCase(companyRepository repositories.CompanyRepository) *DesactiveCompanyUseCase {
	return &DesactiveCompanyUseCase{CompanyRepository: companyRepository}
}

func (uc *DesactiveCompanyUseCase) Execute(ctx context.Context, input DesactiveCompanyInputDTO) error {
	co, err := uc.CompanyRepository.FindByID(ctx, input.ID)
	if err != nil {
		slog.Error("Error finding company by ID", "error", err)
		return err
	}

	if co == nil {
		slog.Error("Company not found", "id", input.ID)
		return errors.New("company not found")
	}

	now := time.Now()

	co.Active = false
	co.DeletedAt = &now

	if err := uc.CompanyRepository.Desactive(ctx, co); err != nil {
		slog.Error("Error deactivating company", "error", err)
		return err
	}

	return nil
}
