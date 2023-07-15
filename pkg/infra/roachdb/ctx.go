package roachdb

import (
	"github.com/BetaLixT/tsqlx"
	"github.com/jmoiron/sqlx"
)

// NewDatabaseContext creates a new database context
func NewDatabaseContext(
	tracer tsqlx.ITracer,
	optn *DatabaseOptions,
) (*tsqlx.TracedDB, error) {
	db, err := sqlx.Open("postgres", optn.ConnectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return tsqlx.NewTracedDB(
		db,
		tracer,
		optn.DatabaseServiceName,
	), nil
}
