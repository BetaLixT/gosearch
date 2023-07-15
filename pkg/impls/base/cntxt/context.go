package cntxt

import (
	"context"
)

type (
	Action      func(context.Context) error
	Constructor func() (interface{}, error)
)

type IContext interface {
	context.Context
	RegisterCompensatoryAction(...Action)
	RegisterCommitAction(...Action)
	GetTransactionObject(
		key string,
		constr Constructor,
	) (obj interface{}, isnew bool, err error)
	GetTraceInfo() (ver, tid, pid, rid, flg string)
	GetInvoker() uint64
}
