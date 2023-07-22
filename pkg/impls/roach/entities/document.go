package entities

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type DocumentEntity struct {
	ID       uint64      `db:"id"`
	Document JSONObj     `db:"document"`
	Keys     StringArray `db:"keys"`
}

var _ sql.Scanner = (*StringArray)(nil)
var _ driver.Valuer = (*StringArray)(nil)

type StringArray []uint64

func (v *StringArray) Scan(value any) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &v)
}

func (v *StringArray) Value() (driver.Value, error) {
	return json.Marshal(v)
}
