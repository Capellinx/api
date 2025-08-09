package external

import "golang.org/x/crypto/bcrypt"

type BcryptExternal struct {
	cost int
}

func NewBcryptExternal(cost int) *BcryptExternal {
	return &BcryptExternal{cost: cost}
}

func (b *BcryptExternal) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	return string(bytes), err
}

func (b *BcryptExternal) Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
