package types

import (
	"github.com/google/uuid"
)

type ProductBasic struct {
	ID     uuid.UUID
	Name   string
	Rating float32
	Price  Money
	Badge  string
}
