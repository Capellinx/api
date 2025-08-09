package entities

import (
	"github.com/google/uuid"
	"time"
)

type Company struct {
	ID          string     `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Cnpj        string     `json:"cnpj,omitempty"`
	Industry    string     `json:"industry"`
	CompanySize string     `json:"company_size"`
	Website     string     `json:"website"`
	Address     string     `json:"address,omitempty"`
	Country     string     `json:"country,omitempty"`
	Currency    string     `json:"currency,omitempty"`
	PostalCode  string     `json:"postal_code,omitempty"`
	City        string     `json:"city,omitempty"`
	State       string     `json:"state,omitempty"`
	Active      bool       `json:"active,omitempty" gorm:"default:true"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

func NewCompany(name, email, cnpj, industry, company_zise string) *Company {
	now := time.Now()
	return &Company{
		ID:          uuid.New().String(),
		Name:        name,
		Email:       email,
		Industry:    industry,
		Address:     "",
		Country:     "",
		Currency:    "",
		PostalCode:  "",
		City:        "",
		State:       "",
		Active:      true,
		CompanySize: company_zise,
		Cnpj:        cnpj,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (Company) TableName() string {
	return "company"
}
