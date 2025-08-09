package mapper

import (
	"api/internal/domain/entity"
	"github.com/lib/pq"
	"time"
)

type CompanyDB struct {
	ID           string         `gorm:"type:uuid;primaryKey"`
	Name         string         `gorm:"column:name"`
	SocialReason string         `gorm:"column:social_reason"`
	Email        string         `gorm:"column:email"`
	Cnpj         string         `gorm:"column:cnpj"`
	Hosts        pq.StringArray `gorm:"type:text[]"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    *time.Time     `gorm:"column:deleted_at"`
}

func (CompanyDB) TableName() string {
	return "company"
}

func (db *CompanyDB) ToEntity() *entity.Company {
	return &entity.Company{
		ID:           db.ID,
		Name:         db.Name,
		SocialReason: db.SocialReason,
		Email:        db.Email,
		Cnpj:         db.Cnpj,
		Hosts:        []string(db.Hosts),
		CreatedAt:    db.CreatedAt,
		UpdatedAt:    db.UpdatedAt,
		DeletedAt:    db.DeletedAt,
	}
}

func FromEntity(company *entity.Company) *CompanyDB {
	return &CompanyDB{
		ID:           company.ID,
		Name:         company.Name,
		SocialReason: company.SocialReason,
		Email:        company.Email,
		Cnpj:         company.Cnpj,
		Hosts:        pq.StringArray(company.Hosts),
		CreatedAt:    company.CreatedAt,
		UpdatedAt:    company.UpdatedAt,
		DeletedAt:    company.DeletedAt,
	}
}
