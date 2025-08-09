package company

import (
	"api/internal/domain/entities"
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

func (rp *CompanyRepositoryPostgres) FindByCnpj(ctx context.Context, cnpj string) (*entities.Company, error) {
	var companyDB CompanyDB
	if err := rp.db.WithContext(ctx).Where("cnpj = ?", cnpj).First(&companyDB).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return companyDB.ToEntity(), nil
}

func (rp *CompanyRepositoryPostgres) Create(ctx context.Context, company *entities.Company) error {
	dbCompany := FromEntity(company)
	if err := rp.db.WithContext(ctx).Create(dbCompany).Error; err != nil {
		return err
	}
	return nil
}

func (rp *CompanyRepositoryPostgres) FindAll(ctx context.Context) ([]*entities.Company, error) {
	dbCompany := make([]CompanyDB, 0)
	if err := rp.db.WithContext(ctx).Find(&dbCompany).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	companies := make([]*entities.Company, 0, len(dbCompany))
	for _, c := range dbCompany {
		companies = append(companies, c.ToEntity())
	}
	return companies, nil
}
