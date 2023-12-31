// Package roach CockroachDB implementation of the domain layer
package roach

import (
	"context"

	cntxtd "github.com/BetaLixT/gosearch/pkg/domain/base/cntxt"
	impld "github.com/BetaLixT/gosearch/pkg/domain/base/impl"
	uidsd "github.com/BetaLixT/gosearch/pkg/domain/base/uids"
	documentsd "github.com/BetaLixT/gosearch/pkg/domain/domains/documents"
	indexesd "github.com/BetaLixT/gosearch/pkg/domain/domains/indexes"
	"github.com/BetaLixT/gosearch/pkg/impls/base/repos/cntxt"
	"github.com/BetaLixT/gosearch/pkg/impls/base/repos/uids"
	"github.com/BetaLixT/gosearch/pkg/impls/roach/entities"
	"github.com/BetaLixT/gosearch/pkg/impls/roach/repos/base"
	"github.com/BetaLixT/gosearch/pkg/impls/roach/repos/documents"
	"github.com/BetaLixT/gosearch/pkg/impls/roach/repos/indexes"
	"github.com/BetaLixT/gosearch/pkg/infra/lgr"
	"github.com/BetaLixT/gosearch/pkg/infra/roachdb"

	"github.com/BetaLixT/tsqlx"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// DependencySet dependencies provided by the implementation
var DependencySet = wire.NewSet(
	NewImplementation,
	wire.Bind(
		new(impld.IImplementation),
		new(*Implementation),
	),

	// Repos
	base.New,
	cntxt.NewContextFactory,
	wire.Bind(
		new(cntxtd.IFactory),
		new(*cntxt.ContextFactory),
	),
	uids.NewUIDRepository,
	wire.Bind(
		new(uidsd.IRepository),
		new(*uids.UIDRepository),
	),
	documents.New,
	wire.Bind(
		new(documentsd.IRepository),
		new(*documents.Repository),
	),
	indexes.New,
	wire.Bind(
		new(indexesd.IRepository),
		new(*indexes.Repository),
	),
)

// Implementation used for graceful starting and stopping of the implementation
// layer
type Implementation struct {
	dbctx *tsqlx.TracedDB
	lgrf  *lgr.LoggerFactory
}

// NewImplementation constructor for the roach implementation
func NewImplementation(
	dbctx *tsqlx.TracedDB,
	lgrf *lgr.LoggerFactory,
) *Implementation {
	return &Implementation{
		dbctx,
		lgrf,
	}
}

// Start runs any routines that are required before the implemtation layer can
// be utilized
func (i *Implementation) Start(ctx context.Context) error {
	lgri := i.lgrf.Create(ctx)
	err := roachdb.RunMigrations(
		ctx,
		lgri,
		i.dbctx,
		entities.GetMigrationScripts(),
	)
	if err != nil {
		lgri.Error("failed to run migration", zap.Error(err))
		return err
	}
	return nil
}

// Stop runs any routines that are required for the implementation layer to
// gracefully shutdown
func (i *Implementation) Stop(ctx context.Context) error {
	i.lgrf.Close()
	return nil
}

// StatusCheck checks connections with dependencies and returns error if any
// fail
func (i *Implementation) StatusCheck(ctx context.Context) error {
	lgr := i.lgrf.Create(ctx)

	lgr.Info("pinging psql database...")
	err := i.dbctx.Ping()
	if err != nil {
		lgr.Error("failed pinging database", zap.Error(err))
		return err
	}
	lgr.Info("psql ok")
	return nil
}
