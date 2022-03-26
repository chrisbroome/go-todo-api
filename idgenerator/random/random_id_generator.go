package random

import (
	"crypto/rand"
	"encoding/hex"
)

type IdGenerator struct{}

func NewIdGenerator() *IdGenerator {
	return &IdGenerator{}
}

func (r *IdGenerator) GenerateID() string {
	buf := make([]byte, 12)
	_, _ = rand.Read(buf)
	return hex.EncodeToString(buf)
}
