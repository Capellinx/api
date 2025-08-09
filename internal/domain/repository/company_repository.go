package repository

import (
	"api/internal/domain/entity"
	"context"
)

type CompanyRepository interface {
	FindByCnpj(ctx context.Context, cnpj string) (*entity.Company, error)
	Create(ctx context.Context, company *entity.Company) error
}
