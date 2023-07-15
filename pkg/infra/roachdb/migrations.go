package roachdb

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/BetaLixT/tsqlx"
	"go.uber.org/zap"
)

func RunMigrations(
	ctx context.Context,
	lgr *zap.Logger,
	db *tsqlx.TracedDB,
	migrations []MigrationScript,
) error {

	tx := db.MustBegin()
	tables := []Table{}
	exist := false

	// - creating migration table if required
	err := tx.Select(ctx, &tables, ShowTablesQuery)
	if err != nil {
		lgr.Error(
			"Failed fetching tables",
			zap.Error(err),
		)
		panic(fmt.Errorf("Failed fetching tables"))
	}
	for _, table := range tables {
		if table.Type == "table" && table.TableName == "migrations" {
			exist = true
		}
	}
	var exMigrs []migrationEntity

	if !exist {
		lgr.Info("Creating migration table")
		tx.MustExec(migrationTable.Up)
		exMigrs = []migrationEntity{}
	} else {
		lgr.Info("Fetching migration history")
		err = tx.Select(ctx, &exMigrs, GetAllMigrations)
		if err != nil {
			lgr.Error(
				"failed to fetch migrations",
				zap.Error(err),
			)
			panic(fmt.Errorf("Failed fetching migration"))
		}
	}
	sort.Slice(exMigrs, func(i, j int) bool {
		return exMigrs[i].Index < exMigrs[j].Index
	})

	exMigrsLen := len(exMigrs)

	for idx, migr := range migrations {
		if idx < exMigrsLen {
			if migr.Key != exMigrs[idx].Key {
				panic(fmt.Errorf("migration key missmatch"))
			}
		} else {
			lgr.Info("Running migration", zap.String("migration", migr.Key))
			tx.MustExec(migr.Up)
			tx.MustExec(AddMigration, migr.Key)
		}
	}
	return tx.Commit()
}

type MigrationScript struct {
	Key  string
	Up   string
	Down string
}

type migrationEntity struct {
	Index           int        `db:"idx"`
	Key             string     `db:"key"`
	DateTimeCreated *time.Time `db:"datetimecreated"`
}

var timestampProcedures = MigrationScript{
	Up: `
		`,
	Down: `
		`,
}

var migrationTable = MigrationScript{
	Up: `
		CREATE TABLE migrations (
			idx SERIAL,
			key text PRIMARY KEY,
			datetimecreated timestamp with time zone NOT NULL DEFAULT now()
		);
		`,
	Down: `
		DROP TABLE Migrations;`,
}

const (
	ShowTablesQuery      = `SHOW TABLES`
	CheckMigrationExists = `
		SELECT EXISTS(
			SELECT * FROM pg_tables
			WHERE schemaname = 'public' AND tablename = 'migrations'
		) as exists`
	GetAllMigrations = `
		SELECT * FROM migrations`
	AddMigration = `
		INSERT INTO migrations (key) VALUES ($1)`
)

// Generic stuff
type ExistsEntity struct {
	Exists bool `db:"exists"`
}

type Table struct {
	SchemaName        string  `db:"schema_name"`
	TableName         string  `db:"table_name"`
	Type              string  `db:"type"`
	Owner             string  `db:"owner"`
	EstimatedRowCount int     `db:"estimated_row_count"`
	Locality          *string `db:"locality"`
}
