package company

import (
	"api/internal/domain/entities"
	"api/internal/domain/repositories"
	"context"
	"log/slog"
)

type FindOneCompanyUseCase struct {
	r repositories.CompanyRepository
}

func NewFindOneCompanyUseCase(r repositories.CompanyRepository) *FindOneCompanyUseCase {
	return &FindOneCompanyUseCase{r: r}
}

func (uc *FindOneCompanyUseCase) Execute(ctx context.Context, id string) (*entities.Company, error) {
	cp, err := uc.r.FindByID(ctx, id)
	if err != nil {
		slog.Error("Company repository FindByID", "error", err)
		return nil, err
	}

	if cp == nil {
		slog.Error("Company not found", "id", id)
		return nil, err
	}

	return cp, nil
}
