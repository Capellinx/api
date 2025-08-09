package repository

import (
	"api/internal/domain/entity"
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
	var company entity.Company
	if err := rp.db.WithContext(ctx).Where("cnpj = ?", cnpj).First(&company).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &company, nil
}

func (rp *CompanyRepositoryPostgres) Create(ctx context.Context, company *entity.Company) error {
	if err := rp.db.WithContext(ctx).Create(company).Error; err != nil {
		return err
	}
	return nil
}
