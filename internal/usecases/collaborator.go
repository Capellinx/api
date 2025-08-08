package usecases

import "api/internal/domain"

type collaboratorUseCase struct {
	repo domain.ClientRepository
}

func CollaboratorUseCase(r domain.ClientRepository) domain.ClientUseCase {
	return &collaboratorUseCase{repo: r}
}

func (c collaboratorUseCase) GetClient(id uint64) (*domain.Client, error) {
	panic("implement me")
}

func (c collaboratorUseCase) CreateClient(client *domain.Client) error {
	panic("implement me")
}
