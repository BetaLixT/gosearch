package indexes

import (
	"context"

	"github.com/BetaLixT/gosearch/pkg/domain/domains/indexes"
	"github.com/BetaLixT/gosearch/pkg/impls/roach/repos/base"
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

}

func (r *Repository) GetDocs(
	ctx context.Context,
	keys []string,
) ([]uint64, error) {

}
