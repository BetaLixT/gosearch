package documents

import (
	"context"
	"fmt"
	"strings"

	"github.com/BetaLixT/gosearch/pkg/domain/domains/documents"
	"github.com/BetaLixT/gosearch/pkg/impls/roach/entities"
	"github.com/BetaLixT/gosearch/pkg/impls/roach/repos/base"
	"github.com/lib/pq"
	"go.uber.org/zap"
)

var _ (documents.IRepository) = (*Repository)(nil)

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
	docs []documents.CreateDocument,
) error {
	lgr := r.Lgrf().Create(ctx)
	lgr.Info("creating documents")

	pbeg := 1
	values := make([]interface{}, 0, len(docs)*3)
	valsqry := make([]string, 0, len(docs))

	// TODO: can be optimized
	for idx := range docs {
		valsqry = append(
			valsqry,
			fmt.Sprintf("($%d, $%d, $%d)", pbeg, pbeg+1, pbeg+2),
		)
		values = append(
			values,
			docs[idx].ID,
			entities.JSONObj(docs[idx].Document),
			pq.Array(docs[idx].Indexed),
		)
		pbeg += 3
	}

	_, err := r.DBCtx().Exec(
		ctx,
		fmt.Sprintf(insert, strings.Join(valsqry, ",")),
		values...,
	)
	if err != nil {
		lgr.Error("failed to insert", zap.Error(err))
		return err
	}
	return nil
}

func (r *Repository) Get(
	ctx context.Context,
	ids []uint64,
) ([]map[string]interface{}, error) {
	lgr := r.Lgrf().Create(ctx)
	lgr.Info("fetching documents", zap.Uint64s("ids", ids))
	res := []entities.DocumentEntity{}
	err := r.DBCtx().Select(
		ctx,
		&res,
		get,
		pq.Array(ids),
	)

	if err != nil {
		lgr.Error("failed to select", zap.Error(err))
		return nil, err
	}

	docs := make([]map[string]interface{}, 0, len(res))
	for idx := range res {
		docs = append(docs, res[idx].Document)
	}

	return docs, nil
}

const (
	insert = `
	  INSERT INTO Document (id, document, keys)
	  VALUES %s
	  RETURNING *
	`

	get = `
	  SELECT * FROM Document WHERE id = ANY($1) 
	`
)
