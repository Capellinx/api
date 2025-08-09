package company

import (
	"api/internal/domain/entities"
	"api/internal/domain/repositories"
	"context"
	"fmt"
)

type FetchAllCompanyUseCase struct {
	r repositories.CompanyRepository
}

func NewFetchAllCompanyUseCase(r repositories.CompanyRepository) *FetchAllCompanyUseCase {
	return &FetchAllCompanyUseCase{r: r}
}

func (uc *FetchAllCompanyUseCase) Execute(ctx context.Context) ([]*entities.Company, error) {
	cp, err := uc.r.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch companies: %w", err)
	}
	return cp, nil
}
