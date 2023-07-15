package handlers

import (
	"context"
	"fmt"
	"github.com/BetaLixT/gex/pkg/app/server/common"
	srvcontracts "github.com/BetaLixT/gex/pkg/app/server/contracts"
	"github.com/BetaLixT/gex/pkg/domain/base/cntxt"
	"github.com/BetaLixT/gex/pkg/domain/base/logger"
	"github.com/BetaLixT/gex/pkg/domain/contracts"
	"github.com/BetaLixT/gex/pkg/domain/usecases"
	"github.com/betalixt/gorr"
	"go.uber.org/zap"
	"time"
)

type DocumentsHandler struct {
	srvcontracts.UnimplementedDocumentsServer
	lgrf logger.IFactory
	uscs *usecases.UseCases
}

var _ srvcontracts.DocumentsServer = (*DocumentsHandler)(nil)

func NewDocumentsHandler(
	lgrf logger.IFactory,
	uscs *usecases.UseCases,
) *DocumentsHandler {
	return &DocumentsHandler{
		lgrf: lgrf,
		uscs: uscs,
	}
}
func (h *DocumentsHandler) Create(
	c context.Context,
	cmd *contracts.CreateIndexedDocumentCommand,
) (res *contracts.DocumentCreatedResponse, err error) {
	ctx, ok := c.(cntxt.IContext)
	if !ok {
		return nil, common.NewInvalidContextProvidedToHandlerError()
	}
	ctx.SetTimeout(2 * time.Minute)
	lgr := h.lgrf.Create(ctx)
	lgr.Info(
		"handling",
		zap.Any("cmd", cmd),
	)
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = gorr.NewUnexpectedError(fmt.Errorf("%v", r))
				lgr.Error(
					"root panic recovered handling request",
					zap.Any("panic", r),
					zap.Stack("stack"),
				)
			} else {
				lgr.Error(
					"root panic recovered handling request",
					zap.Error(err),
					zap.Stack("stack"),
				)
			}
			ctx.Cancel()
		}
		if err != nil {
			if _, ok := err.(*gorr.Error); !ok {
				err = gorr.NewUnexpectedError(err)
			}
		}
	}()
	res, err = h.uscs.DocumentsCreate(
		ctx,
		cmd,
	)
	if err != nil {
		lgr.Error(
			"command handling failed",
			zap.Error(err),
		)
	}
	ctx.Cancel()
	return
}

func (h *DocumentsHandler) Query(
	c context.Context,
	qry *contracts.SearchQuery,
) (res *contracts.SearchResponse, err error) {
	ctx, ok := c.(cntxt.IContext)
	if !ok {
		return nil, common.NewInvalidContextProvidedToHandlerError()
	}
	ctx.SetTimeout(2 * time.Minute)
	lgr := h.lgrf.Create(ctx)
	lgr.Info(
		"handling",
		zap.Any("qry", qry),
	)
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = gorr.NewUnexpectedError(fmt.Errorf("%v", r))
				lgr.Error(
					"root panic recovered handling request",
					zap.Any("panic", r),
					zap.Stack("stack"),
				)
			} else {
				lgr.Error(
					"root panic recovered handling request",
					zap.Error(err),
					zap.Stack("stack"),
				)
			}
			ctx.Cancel()
		}
		if err != nil {
			if _, ok := err.(*gorr.Error); !ok {
				err = gorr.NewUnexpectedError(err)
			}
		}
	}()
	res, err = h.uscs.DocumentsQuery(
		ctx,
		qry,
	)
	if err != nil {
		lgr.Error(
			"command handling failed",
			zap.Error(err),
		)
	}
	ctx.Cancel()
	return
}
