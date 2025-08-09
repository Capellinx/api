package company

import (
	"api/internal/domain/entity"
	"api/internal/domain/repository"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type CreateCompanyInputDTO struct {
	Name         string `json:"name"`
	SocialReason string `json:"social_reason"`
	Email        string `json:"email"`
	Cnpj         string `json:"cnpj"`
}

type CreateCompanyUseCase struct {
	repo repository.CompanyRepository
}

func NewCreateCompanyUseCase(r repository.CompanyRepository) *CreateCompanyUseCase {
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

	company := &entity.Company{
		ID:           uuid.New().String(),
		Name:         c.Name,
		SocialReason: c.SocialReason,
		Email:        c.Email,
		Cnpj:         c.Cnpj,
	}

	if err := uc.repo.Create(ctx, company); err != nil {
		return fmt.Errorf("failed to create company: %w", err)
	}

	return nil
}
