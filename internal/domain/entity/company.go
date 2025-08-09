package entity

import (
	"time"
)

type Company struct {
	ID           string     `gorm:"type:uuid;primaryKey" json:"id"`
	Name         string     `json:"name"`
	SocialReason string     `json:"social_reason"`
	Email        string     `json:"email"`
	Cnpj         string     `json:"cnpj"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

func NewCompany(name, socialReason, email, cnpj string) *Company {
	now := time.Now()
	return &Company{
		ID:           "3df3f2ff-6b0e-4604-857d-e850742ea21f",
		Name:         name,
		SocialReason: socialReason,
		Email:        email,
		Cnpj:         cnpj,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func (Company) TableName() string {
	return "company"
}
