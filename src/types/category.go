package types

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Category struct {
	ID            string
	Name          string
	Description   string
	Creation_Date timestamp.Timestamp
	Last_Update   timestamp.Timestamp
}
