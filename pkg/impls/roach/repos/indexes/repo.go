package indexes

import (
	"context"
	"fmt"
	"strings"

	"github.com/BetaLixT/gosearch/pkg/domain/domains/indexes"
	"github.com/BetaLixT/gosearch/pkg/impls/roach/entities"
	"github.com/BetaLixT/gosearch/pkg/impls/roach/repos/base"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

var _ (indexes.IRepository) = (*Repository)(nil)

func New(b *base.BaseDataRepository) *Repository {
	return &Repository{
		b,
	}
}

type Repository struct {
	*base.BaseDataRepository
}

func (r *Repository) Create(
	ctx context.Context,
	idxs []indexes.CreateIndex,
) error {

	lgr := r.Lgrf().Create(ctx)
	lgr.Info("creating index", zap.Any("indexes", idxs))

	pbeg := 1
	values := make([]interface{}, 0, len(idxs)*2)
	valsqry := make([]string, 0, len(idxs))

	// TODO: can be optimized
	for idx := range idxs {
		for _, key := range idxs[idx].Keys {
			valsqry = append(
				valsqry,
				fmt.Sprintf("($%d, $%d)", pbeg, pbeg+1),
			)
			values = append(
				values,
				key,
				[]uint64{idxs[idx].Document},
			)
			pbeg += 2
		}
	}

	_, err := r.DBCtx().Exec(
		ctx,
		fmt.Sprintf(upsert, strings.Join(valsqry, ",")),
		values...,
	)
	if err != nil {
		lgr.Error("failed to insert", zap.Error(err))
		return err
	}
	return nil
}

func (r *Repository) GetDocs(
	ctx context.Context,
	keys []string,
) ([]uint64, error) {
	lgr := r.Lgrf().Create(ctx)
	lgr.Info("fetching indexed docs", zap.Strings("keys", keys))
	res := []entities.IndexEntity{}
	err := r.DBCtx().Select(
		ctx,
		&res,
		get,
		pq.Array(keys),
	)

	if err != nil {
		lgr.Error("failed to select", zap.Error(err))
		return nil, err
	}

	unqs := map[uint64]struct{}{}
	docs := []uint64{} // TODO optimize capacity
	for idx := range res {
		for jidx := range res[idx].Documents {
			if _, ok := unqs[res[idx].Documents[jidx]]; !ok {
				unqs[res[idx].Documents[jidx]] = struct{}{}
				docs = append(docs, res[idx].Documents[jidx])
			}
		}
	}

	return docs, nil
}

const (
	upsert = `
	  INSERT INTO SearchIndex (key, documents)
	  VALUES %s
	  ON CONFLICT (key)
	  DO UPDATE SET documents = array_cat(documents, excluded.documents)
	  RETURNING *
	`

	get = `
	  SELECT * FROM SearchIndex WHERE key = ANY($1) 
	`
)
