package services

type HasherService interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}
