// Package entities containing all data access objects (models that relate to
// the how the data is stored in the database)
package entities

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	// blank import to load postgresql drivers
	"github.com/BetaLixT/gosearch/pkg/infra/roachdb"
	_ "github.com/lib/pq"
)

// =============================================================================
// Common DAO models
// =============================================================================

// - JSONObj dao for objects to be stored as json
type JSONObj map[string]interface{}

var _ driver.Value = (*JSONObj)(nil)

// Value for db writes
func (a JSONObj) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan for db reads
func (a *JSONObj) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

// - JSONMapString  dao for map[string]string to be stored as json
type JSONMapString map[string]string

var _ driver.Value = (*JSONObj)(nil)

// Value fo db writes
func (a JSONMapString) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan for db reads
func (a *JSONMapString) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type IDVersion struct {
	ID      uint64
	Version uint64
}

func ToArrayContract[in any, out any](
	e []in,
	conv func(in) *out,
) []*out {
	res := make([]*out, 0, len(e))
	for idx := range e {
		res = append(res, conv(e[idx]))
	}
	return res
}

func ToArrayContractErrorable[in any, out any](
	e []in,
	conv func(*in) (*out, error),
) ([]*out, error) {
	res := make([]*out, 0, len(e))
	for idx := range e {
		c, err := conv(&e[idx])
		if err != nil {
			return nil, err
		}
		res = append(res, c)
	}
	return res, nil
}

func ToArrayContractErrorableFromPointer[in any, out any](
	e []*in,
	conv func(*in) (*out, error),
) ([]*out, error) {
	res := make([]*out, 0, len(e))
	for idx := range e {
		c, err := conv(e[idx])
		if err != nil {
			return nil, err
		}
		res = append(res, c)
	}
	return res, nil
}

// =============================================================================
// Migrations
// =============================================================================

// GetMigrationScripts provides all the migration scripts required for the
// application
func GetMigrationScripts() []roachdb.MigrationScript {
	migrationScripts := []roachdb.MigrationScript{
		{
			Key: "initial",
			Up: `
			  CREATE TABLE SearchIndex (
				  key string PRIMARY KEY,
				  documents bigint[]
			  );

			  CREATE TABLE Document (
				  id bigint PRIMARY KEY,
				  document jsonb NOT NULL,
				  keys string[]
			  );
			`,
			Down: `
			  DELETE TABLE SearchIndex;
			  DELETE TABLE Document;
			`,
		},
	}
	return migrationScripts
}
