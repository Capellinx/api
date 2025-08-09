package repositories

import "api/internal/domain/entities"

type CollaboratorRepository interface {
	FindByEmail(email string) (*entities.Collaborator, error)
	Create(collaborator *entities.Collaborator) error
}
