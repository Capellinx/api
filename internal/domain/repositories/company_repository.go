package repositories

import (
	"api/internal/domain/entities"
	"context"
)

type CompanyRepository interface {
	FindByCnpj(ctx context.Context, cnpj string) (*entities.Company, error)
	FindAll(ctx context.Context) ([]*entities.Company, error)
	FindByID(ctx context.Context, id string) (*entities.Company, error)
	Desactive(ctx context.Context, company *entities.Company) error
	Create(ctx context.Context, company *entities.Company) error
}
