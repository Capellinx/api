package entities

import (
	"github.com/google/uuid"
	"time"
)

type Collaborator struct {
	ID        string     `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(100);not null" json:"name"`
	Email     string     `gorm:"type:varchar(100);not null;unique" json:"email"`
	Position  string     `gorm:"type:varchar(100);not null" json:"position"`
	Phone     string     `gorm:"type:varchar(15);not null" json:"phone"`
	Active    bool       `gorm:"default:true" json:"active"`
	CompanyID string     `gorm:"type:uuid;not null" json:"company_id"`
	Role      string     `gorm:"type:varchar(50);null" json:"role"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

func NewCollaborator(name, email, position, phone, companyID, role string, active bool) *Collaborator {
	return &Collaborator{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		Position:  position,
		Phone:     phone,
		Active:    active,
		CompanyID: companyID,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
