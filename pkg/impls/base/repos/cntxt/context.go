package cntxt

import (
	"context"
	"fmt"
	"sync"
	"time"

	domcntxt "github.com/BetaLixT/gosearch/pkg/domain/base/cntxt"
	"github.com/BetaLixT/gosearch/pkg/domain/base/logger"
	implcntxt "github.com/BetaLixT/gosearch/pkg/impls/base/cntxt"
	infrcntxt "github.com/BetaLixT/gosearch/pkg/infra/cntxt"

	"github.com/BetaLixT/go-resiliency/retrier"
	"go.uber.org/zap"
)

// =============================================================================
// Internal context implementation
// =============================================================================

var (
	_ context.Context    = (*internalContext)(nil)
	_ domcntxt.IContext  = (*internalContext)(nil)
	_ infrcntxt.IContext = (*internalContext)(nil)
	_ implcntxt.IContext = (*internalContext)(nil)
)

type internalContext struct {
	lgrf logger.IFactory
	// deadline time.Time
	cancelmtx *sync.Mutex
	err       error
	done      chan struct{}
	dur       time.Time

	// - transaction
	rtr                 retrier.Retrier
	compensatoryActions []implcntxt.Action
	commitActions       []implcntxt.Action
	txObjs              map[string]interface{}
	isCommited          bool
	isRolledback        bool
	txmtx               *sync.Mutex

	// trace
	ver string
	tid string
	pid string
	rid string
	flg string

	// user contxt
	userValid bool
	invokedBy uint64

	// data
	mu     sync.RWMutex
	values map[any]any
}

// - Base context functions
func (c *internalContext) cancel(err error) {
	c.RollbackTransaction()
	c.cancelmtx.Lock()
	defer c.cancelmtx.Unlock()
	if c.err != nil {
		return
	}
	c.err = err
	close(c.done)
}

func (c *internalContext) Cancel() {
	c.cancel(fmt.Errorf("context manually canceled"))
}

func (c *internalContext) Deadline() (time.Time, bool) {
	return time.Now(), false
}

func (c *internalContext) Done() <-chan struct{} {
	return c.done
}

func (c *internalContext) Err() error {
	return c.err
}

func (c *internalContext) Value(key any) (v any) {
	c.mu.RLock()
	v = c.values[key]
	c.mu.RUnlock()
	return
}

func (c *internalContext) WithValue(key any, val any) {
	c.mu.Lock()
	c.values[key] = val
	c.mu.Unlock()
}

// - Transaction functions
func (c *internalContext) SetTimeout(timeout time.Duration) {
	c.dur = time.Now().Add(timeout)
	time.AfterFunc(
		time.Until(c.dur),
		func() {
			c.cancel(context.DeadlineExceeded)
		},
	)
}

func (c *internalContext) CommitTransaction() error {
	c.txmtx.Lock()
	defer c.txmtx.Unlock()
	if c.isCommited || c.isRolledback {
		return fmt.Errorf(
			"tried to commit transaction that has already been commited/rolled back",
		)
	}
	ctx := newMinimalContext(c)

	// TODO: some commits failing in between might be an issue here...
	// Running all commits
	for _, commit := range c.commitActions {
		err := commit(ctx)
		if err != nil {
			return err
		}
	}
	c.isCommited = true
	return nil
}

// TODO: better handling failed rollback transaction
func (c *internalContext) RollbackTransaction() {
	c.txmtx.Lock()
	defer c.txmtx.Unlock()
	if c.isCommited || c.isRolledback {
		return
	}

	c.isRolledback = true
	ctx := newMinimalContext(c)
	lgr := c.lgrf.Create(ctx)
	for _, cmp := range c.compensatoryActions {
		err := c.rtr.Run(func() error {
			err := cmp(ctx)
			if err != nil {
				lgr.Warn("failed to run compensatory action", zap.Error(err))
			}
			return err
		})
		if err != nil {
			lgr.Error(
				"failed to run compensatory action, max retries exceeded",
				zap.Error(err),
			)
		}
	}
}

func (c *internalContext) RegisterCompensatoryAction(
	cmp ...implcntxt.Action,
) {
	c.txmtx.Lock()
	defer c.txmtx.Unlock()
	c.compensatoryActions = append(c.compensatoryActions, cmp...)
}

func (c *internalContext) RegisterCommitAction(
	cmp ...implcntxt.Action,
) {
	c.txmtx.Lock()
	defer c.txmtx.Unlock()
	c.commitActions = append(c.commitActions, cmp...)
}

func (c *internalContext) GetTransactionObject(
	key string,
	constr implcntxt.Constructor,
) (interface{}, bool, error) {
	c.txmtx.Lock()
	defer c.txmtx.Unlock()
	intr, ok := c.txObjs[key]
	if ok {
		return intr, false, nil
	}
	intr, err := constr()
	if err != nil {
		return nil, false, err
	}
	c.txObjs[key] = intr
	return intr, true, nil
}

func (c *internalContext) GetTraceInfo() (ver, tid, pid, rid, flg string) {
	return c.ver, c.tid, c.pid, c.rid, c.flg
}

func (c *internalContext) GenerateSpanID() (string, error) {
	return generateRadomHexString(8)
}

func (c *internalContext) ValidateUserContext() {
	c.userValid = true
}

func (c *internalContext) IsUserValid() bool {
	return c.userValid
}

func (c *internalContext) SetInvoker(invoker uint64) {
	c.invokedBy = invoker
}

func (c *internalContext) GetInvoker() uint64 {
	return c.invokedBy
}

//easyjson:json
type mapStringInterface map[string]interface{}
