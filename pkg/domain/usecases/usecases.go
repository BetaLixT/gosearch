package usecases

import (
	"github.com/BetaLixT/gosearch/pkg/domain/base/logger"
	"github.com/BetaLixT/gosearch/pkg/domain/base/uids"
	"github.com/BetaLixT/gosearch/pkg/domain/domains/documents"
	"github.com/BetaLixT/gosearch/pkg/domain/domains/indexes"
)

type UseCases struct {
	lgrf    logger.IFactory
	uidrepo uids.IRepository
	docrepo documents.IRepository
	idxrepo indexes.IRepository
}

func NewUseCases(
	lgrf logger.IFactory,
	uidrepo uids.IRepository,
	docrepo documents.IRepository,
	idxrepo indexes.IRepository,
) *UseCases {
	return &UseCases{
		lgrf,
		uidrepo,
		docrepo,
		idxrepo,
	}
}
