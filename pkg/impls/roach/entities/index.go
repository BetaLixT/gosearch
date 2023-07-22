package entities

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type IndexEntity struct {
	Key       string      `db:"key"`
	Documents UInt64Array `db:"documents"`
}

var _ sql.Scanner = (*UInt64Array)(nil)
var _ driver.Valuer = (*UInt64Array)(nil)

type UInt64Array []uint64

func (v *UInt64Array) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &v)
}

func (v *UInt64Array) Value() (driver.Value, error) {
	return json.Marshal(v)
}
