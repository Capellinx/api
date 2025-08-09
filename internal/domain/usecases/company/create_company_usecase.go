package company

import (
	"api/internal/domain/entities"
	"api/internal/domain/repositories"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type CreateCompanyInputDTO struct {
	Name         string   `json:"name" validate:"required"`
	SocialReason string   `json:"social_reason"`
	Email        string   `json:"email"`
	Hosts        []string `json:"hosts"`
	Cnpj         string   `json:"cnpj"`
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
		ID:           uuid.New().String(),
		Name:         c.Name,
		SocialReason: c.SocialReason,
		Hosts:        c.Hosts,
		Email:        c.Email,
		Cnpj:         c.Cnpj,
	}

	if err := uc.repo.Create(ctx, company); err != nil {
		return fmt.Errorf("failed to create company: %w", err)
	}

	return nil
}
