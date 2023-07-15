// Package repos implements the interfaces defined on the domain layer
package base

import (
	"context"

	"github.com/BetaLixT/gosearch/pkg/domain/base/logger"
	"github.com/BetaLixT/gosearch/pkg/impls/base/cntxt"
	"github.com/BetaLixT/gosearch/pkg/impls/roach/common"

	"github.com/BetaLixT/tsqlx"
)

// BaseDataRepository is the base repository containing common database
// functionality to be embeded by other repos that implement persistence to the
// database
type BaseDataRepository struct {
	dbctx *tsqlx.TracedDB
	lgrf  logger.IFactory
}

// NewBaseDataRepository Constructs a new base data repository
func New(
	dbctx *tsqlx.TracedDB,
	lgrf logger.IFactory,
) *BaseDataRepository {
	return &BaseDataRepository{
		dbctx: dbctx,
		lgrf:  lgrf,
	}
}

func (r *BaseDataRepository) Lgrf() logger.IFactory {
	return r.lgrf
}

// Gets database context
// use this for reads only (SELECT)
func (r *BaseDataRepository) DBCtx() *tsqlx.TracedDB {
	return r.dbctx
}

// Gets an existing database transaction if exists
// else creates and returns database transaction
// use this for writes (INSERT, UPDATE, DELETE)
func (r *BaseDataRepository) GetDBTx(
	ctx cntxt.IContext,
) (*tsqlx.TracedTx, error) {
	idbtx, nw, err := ctx.GetTransactionObject(
		common.SqlTransactionObjectKey,
		func() (interface{}, error) {
			return r.dbctx.Beginx()
		},
	)
	if err != nil {
		return nil, err
	}

	dbtx, ok := idbtx.(*tsqlx.TracedTx)
	if !ok {
		return nil, common.NewFailedToAssertDatabaseCtxTypeError()
	}
	if nw {
		ctx.RegisterCommitAction(func(ctx context.Context) error {
			return dbtx.Commit()
		})
		ctx.RegisterCompensatoryAction(func(ctx context.Context) error {
			return dbtx.Rollback()
		})
	}
	return dbtx, nil
}

// GetValueOrDefault either returns the value if the provided pointer is not nil
// else it provides the default value
func GetValueOrDefault[v comparable](in *v) (out v) {
	if in != nil {
		out = *in
	}
	return out
}
