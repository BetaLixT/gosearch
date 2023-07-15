package documents

import (
	"github.com/BetaLixT/gex/pkg/domain/domains/documents"
	"github.com/BetaLixT/gex/pkg/impls/roach/repos/base"
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
