package repository

import (
	"api/internal/domain/entity"
	"api/internal/mapper"
	"context"
	"errors"
	"gorm.io/gorm"
)

type CompanyRepositoryPostgres struct {
	db *gorm.DB
}

func NewCompanyRepositoryPostgres(db *gorm.DB) *CompanyRepositoryPostgres {
	return &CompanyRepositoryPostgres{db: db}
}

func (rp *CompanyRepositoryPostgres) FindByCnpj(ctx context.Context, cnpj string) (*entity.Company, error) {
	var companyDB mapper.CompanyDB
	if err := rp.db.WithContext(ctx).Where("cnpj = ?", cnpj).First(&companyDB).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return companyDB.ToEntity(), nil
}

func (rp *CompanyRepositoryPostgres) Create(ctx context.Context, company *entity.Company) error {
	dbCompany := mapper.FromEntity(company)
	if err := rp.db.WithContext(ctx).Create(dbCompany).Error; err != nil {
		return err
	}
	return nil
}
