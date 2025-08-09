package company

import (
	"api/internal/domain/entities"
	"time"
)

type CompanyDB struct {
	ID          string     `gorm:"type:uuid;primaryKey"`
	Name        string     `gorm:"column:name"`
	Email       string     `gorm:"column:email"`
	Cnpj        string     `gorm:"column:cnpj"`
	Industry    string     `gorm:"column:industry"`
	CompanySize string     `gorm:"column:company_size"`
	Active      bool       `gorm:"column:active"`
	Website     string     `gorm:"column:website"`
	CreatedAt   time.Time  `gorm:"column:created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at"`
}

func (CompanyDB) TableName() string {
	return "company"
}

func (db *CompanyDB) ToEntity() *entities.Company {
	return &entities.Company{
		ID:          db.ID,
		Name:        db.Name,
		Email:       db.Email,
		Cnpj:        db.Cnpj,
		Industry:    db.Industry,
		CompanySize: db.CompanySize,
		Active:      db.Active,
		Website:     db.Website,
		CreatedAt:   db.CreatedAt,
		UpdatedAt:   db.UpdatedAt,
		DeletedAt:   db.DeletedAt,
	}
}

func FromEntity(company *entities.Company) *CompanyDB {
	return &CompanyDB{
		ID:          company.ID,
		Name:        company.Name,
		Email:       company.Email,
		Cnpj:        company.Cnpj,
		Active:      company.Active,
		Industry:    company.Industry,
		CompanySize: company.CompanySize,
		Website:     company.Website,
		CreatedAt:   company.CreatedAt,
		UpdatedAt:   company.UpdatedAt,
		DeletedAt:   company.DeletedAt,
	}
}
