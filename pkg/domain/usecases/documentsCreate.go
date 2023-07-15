package usecases

import (
	"github.com/BetaLixT/gex/pkg/domain/base/cntxt"
	"github.com/BetaLixT/gex/pkg/domain/contracts"
	"github.com/betalixt/gorr"
	"go.uber.org/zap"
)

// DocumentsCreate usecase for
func (u *UseCases) DocumentsCreate(
	ctx cntxt.IUseCaseContext,
	cmd *contracts.CreateIndexedDocumentCommand,
) (res *contracts.DocumentCreatedResponse, err error) {
	lgr := u.lgrf.Create(ctx)
	lgr.Info(
		"running usecase",
		zap.String("resource", "DocumentsHandler"),
		zap.String("usecase", "Create"),
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
