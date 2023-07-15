package cntxt

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/BetaLixT/gex/pkg/domain/base/cntxt"
	domcntxt "github.com/BetaLixT/gex/pkg/domain/base/cntxt"
	"github.com/BetaLixT/gex/pkg/domain/base/logger"
	implcntxt "github.com/BetaLixT/gex/pkg/impls/base/cntxt"
	"github.com/BetaLixT/gex/pkg/impls/base/common"

	"github.com/BetaLixT/go-resiliency/retrier"
	"go.uber.org/zap"
)

// =============================================================================
// Context factory and trace parsing logic
// =============================================================================

var _ cntxt.IFactory = (*ContextFactory)(nil)

// ContextFactory to create new contexts
type ContextFactory struct {
	lgrf logger.IFactory
}

// NewContextFactory constructor for context factory
func NewContextFactory(
	lgrf logger.IFactory,
) *ContextFactory {
	return &ContextFactory{
		lgrf,
	}
}

// Create creates a new context with timeout, transactions and trace info
func (f *ContextFactory) Create(
	traceparent string,
) domcntxt.IContext {
	ver, tid, pid, rid, flg, err := parseTraceParent(traceparent)
	if err != nil {
		lgr := f.lgrf.Create(context.Background())
		lgr.Error("failed to generate trace info", zap.Error(err))
	}
	c := &internalContext{
		f.lgrf,

		&sync.Mutex{},
		nil,
		make(chan struct{}, 1),
		time.Time{},

		*retrier.New(retrier.ExponentialBackoff(
			5,
			500*time.Millisecond,
		),
			retrier.DefaultClassifier{},
		),
		[]implcntxt.Action{},
		[]implcntxt.Action{},
		map[string]interface{}{},
		false,
		false,
		&sync.Mutex{},
		ver,
		tid,
		pid,
		rid,
		flg,

		false,
		0,

		sync.RWMutex{},
		map[any]any{},
	}

	return c
}

// parseTraceParent parses and or generates trace information
func parseTraceParent(
	traceprnt string,
) (ver, tid, pid, rid, flg string, err error) {
	ver, tid, pid, flg, err = decodeTraceparent(traceprnt)
	// If the header could not be decoded, generate a new header
	if err != nil {
		ver, flg = "00", "01"
		if tid, err = generateRadomHexString(16); err != nil {
			return "", "", "", "", "", common.NewHexStringGenerationFailedError(err)
		}
	}

	// Generate a new resource id
	rid, err = generateRadomHexString(8)
	if err != nil {
		return "", "", "", "", "", common.NewHexStringGenerationFailedError(err)
	}
	return
}

func generateRadomHexString(n int) (string, error) {
	buff := make([]byte, n)
	if _, err := rand.Read(buff); err != nil {
		return "", err
	}
	return hex.EncodeToString(buff), nil
}

func decodeTraceparent(traceparent string) (string, string, string, string, error) {
	// Fast fail for common case of empty string
	if traceparent == "" {
		return "", "", "", "", fmt.Errorf("traceparent is empty string")
	}

	hexfmt, err := regexp.Compile("^[0-9A-Fa-f]*$")
	vals := strings.Split(traceparent, "-")

	if len(vals) == 4 {
		ver, tid, pid, flg := vals[0], vals[1], vals[2], vals[3]
		if !hexfmt.MatchString(ver) || len(ver) != 2 {
			err = fmt.Errorf("invalid traceparent version")
		} else if !hexfmt.MatchString(pid) || len(pid) != 16 {
			err = fmt.Errorf("invalid traceparent parent id")
		} else if !hexfmt.MatchString(flg) || len(flg) != 2 {
			err = fmt.Errorf("invalid traceparent flag")
		} else if !hexfmt.MatchString(tid) || len(tid) != 32 {
			err = fmt.Errorf("invalid traceparent trace id")
		} else if tid == "00000000000000000000000000000000" {
			err = fmt.Errorf("traceparent trace id value is zero")
		} else {
			return ver, tid, pid, flg, nil
		}
	} else {
		err = fmt.Errorf("invalid traceparent trace id")
	}

	return "", "", "", "", err
}
