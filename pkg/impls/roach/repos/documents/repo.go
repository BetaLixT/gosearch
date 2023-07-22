package documents

import (
	"context"

	"github.com/BetaLixT/gosearch/pkg/domain/domains/documents"
	"github.com/BetaLixT/gosearch/pkg/impls/roach/repos/base"
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

}

func (r *Repository) Get(
	ctx context.Context,
	ids []uint64,
) ([]map[string]interface{}, error) {

}

const ()
