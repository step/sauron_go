package flowIDGenerator

import (
	"github.com/google/uuid"
)

type UUIDGenerator struct{}

func (u UUIDGenerator) New() string {
	return uuid.New().String()
}

func NewUUIDGenerator() UUIDGenerator {
	return UUIDGenerator{}
}