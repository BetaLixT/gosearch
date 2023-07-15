package indexes

import (
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
