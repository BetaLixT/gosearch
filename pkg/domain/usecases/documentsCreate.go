package usecases

import (
	"unicode"

	"github.com/BetaLixT/gosearch/pkg/domain/base/cntxt"
	"github.com/BetaLixT/gosearch/pkg/domain/contracts"
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

	// TODO: possible optimization with alloc
	indexed := []string{}

	err = gorr.NewNotImplemented()
	if err != nil {
		lgr.Error(
			"error while running usecase",
			zap.Error(err),
		)
	}
	return
}

func findKeys(doc map[string]interface{}) (res []string) {
	for k := range doc {
		if s, ok := doc[k].(string); ok {
			if len(s) < EntireStringCacheMaxLength && !hasWhitespace(s) {
				res = append(res, s)
			}
			res = append(res, tokenize(s, SpecialCharacterBreakCheck)...)
			continue
		}
	}
}

func hasWhitespace(in string) bool {
	for _, b := range in {
		if unicode.IsSpace(rune(b)) {
			return true
		}
	}
	return false
}

const (
	EntireStringCacheMaxLength = 500
)
