package usecases

import (
	"github.com/BetaLixT/gosearch/pkg/domain/base/cntxt"
	"github.com/BetaLixT/gosearch/pkg/domain/contracts"
	"github.com/betalixt/gorr"
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

	err = gorr.NewNotImplemented()
	if err != nil {
		lgr.Error(
			"error while running usecase",
			zap.Error(err),
		)
	}
	return
}
