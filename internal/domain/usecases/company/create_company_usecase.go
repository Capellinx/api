package company

import (
	"api/internal/domain/entities"
	"api/internal/domain/repositories"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type CreateCompanyInputDTO struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email"`
	Industry    string `json:"industry"`
	CompanySize string `json:"company_size"`
	Cnpj        string `json:"cnpj"`
	Website     string `json:"website"`
}

type CreateCompanyUseCase struct {
	repo repositories.CompanyRepository
}

func NewCreateCompanyUseCase(r repositories.CompanyRepository) *CreateCompanyUseCase {
	return &CreateCompanyUseCase{repo: r}
}

func (uc *CreateCompanyUseCase) Execute(ctx context.Context, c CreateCompanyInputDTO) error {
	ec, err := uc.repo.FindByCnpj(ctx, c.Cnpj)
	if err != nil {
		return fmt.Errorf("failed to check company existence: %w", err)
	}

	if ec != nil {
		return fmt.Errorf("company already exists")
	}

	company := &entities.Company{
		ID:          uuid.New().String(),
		Name:        c.Name,
		Industry:    c.Industry,
		CompanySize: c.CompanySize,
		Email:       c.Email,
		Cnpj:        c.Cnpj,
		Website:     c.Website,
	}

	if err := uc.repo.Create(ctx, company); err != nil {
		return fmt.Errorf("failed to create company: %w", err)
	}

	return nil
}
