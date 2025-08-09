package entities

import (
	"github.com/google/uuid"
	"time"
)

type Company struct {
	ID           string     `gorm:"type:uuid;primaryKey" json:"id"`
	Name         string     `json:"name"`
	SocialReason string     `json:"social_reason"`
	Email        string     `json:"email"`
	Cnpj         string     `json:"cnpj"`
	Hosts        []string   `gorm:"type:text[]" json:"hosts"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

func NewCompany(name, socialReason, email, cnpj string, hosts []string) *Company {
	now := time.Now()
	return &Company{
		ID:           uuid.New().String(),
		Name:         name,
		SocialReason: socialReason,
		Email:        email,
		Hosts:        hosts,
		Cnpj:         cnpj,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func (Company) TableName() string {
	return "company"
}
