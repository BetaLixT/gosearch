package usecases

import (
	"github.com/BetaLixT/gosearch/pkg/domain/base/cntxt"
	"github.com/BetaLixT/gosearch/pkg/domain/contracts"
	"go.uber.org/zap"
)

// DocumentsQuery usecase for
func (u *UseCases) DocumentsQuery(
	ctx cntxt.IUseCaseContext,
	qry *contracts.SearchQuery,
) (res *contracts.SearchResponse, err error) {
	lgr := u.lgrf.Create(ctx)
	lgr.Info(
		"running usecase",
		zap.String("resource", "DocumentsHandler"),
		zap.String("usecase", "Query"),
	)

	keys := removeDuplicates(tokenize(qry.Query, SpecialCharacterBreakCheck, false))
	docids, err := u.idxrepo.GetDocs(ctx, keys)
	if err != nil {
		lgr.Error("failed to create indexed docs", zap.Error(err))
		return nil, err
	}

	docs, err := u.docrepo.Get(ctx, docids)
	if err != nil {
		lgr.Error("failed to fetch documents", zap.Error(err))
		return nil, err
	}

	res, err = contracts.NewSearchResponse(docs)
	if err != nil {
		lgr.Error(
			"error while running usecase",
			zap.Error(err),
		)
	}
	return
}
