package types

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/shopspring/decimal"
)

type ProductBasic struct {
	ID     string
	Name   string
	Rating float32
	Price  decimal.Decimal
}

type Product struct {
	BasicInfo     ProductBasic
	Badge         string
	Description   string
	InStock       int
	Category      Category
	Creation_Date timestamp.Timestamp
	Last_Update   timestamp.Timestamp
}
