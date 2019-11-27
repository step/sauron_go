package flowidgenerator

import (
	"github.com/google/uuid"
)

// UUIDGenerator is type of flowidgenerator that
// generates a new uuid
type UUIDGenerator struct{}

// New returns a string representation of a new UUID
func (u UUIDGenerator) New() string {
	return uuid.New().String()
}

// NewUUIDGenerator should be called when a new
// UUIDGenerator needs to be created
func NewUUIDGenerator() UUIDGenerator {
	return UUIDGenerator{}
}