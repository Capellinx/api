package domain

type Client struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	SocialReason string  `json:"social_reason"`
	Email        string  `json:"email"`
	Cnpj         string  `json:"cnpj"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	DeletedAt    *string `json:"deleted_at,omitempty"`
}

type ClientRepository interface {
	GetByID(id uint64) (*Client, error)
	Create(user *Client) error
}

type ClientUseCase interface {
	GetClient(id uint64) (*Client, error)
	CreateClient(user *Client) error
}
