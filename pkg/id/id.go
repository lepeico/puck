package id

import (
	"crypto/rand"
	"math/big"
)

var (
	validCharacters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
	validRegexp     = "^[A-Za-z0-9_-]+$"
)

// ID is a byte form of id
type ID []byte

// New returns ID of lenght
func New(lenght int) ID {
	id := make(ID, lenght)
	for i := range id {
		if n, err := rand.Int(rand.Reader, big.NewInt(64)); err == nil {
			id[i] = validCharacters[n.Int64()]
		}
	}

	return id
}

func FromString(str string) ID {
	return []byte(str)
}

func (id ID) String() string {
	return string(id[:])
}
