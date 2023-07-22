package usecases

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/BetaLixT/gosearch/pkg/domain/base/cntxt"
	"github.com/BetaLixT/gosearch/pkg/domain/contracts"
	"github.com/BetaLixT/gosearch/pkg/domain/domains/documents"
	"github.com/BetaLixT/gosearch/pkg/domain/domains/indexes"
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

	res = contracts.NewDocumentCreatedResponse(
		make([]*contracts.IndexedDocument, 0, len(cmd.Documents)),
	)
	docsToCreate := make([]documents.CreateDocument, 0, len(cmd.Documents))
	idxToCreate := make([]indexes.CreateIndex, 0, len(cmd.Documents))
	// TODO: possible optimization with alloc
	for idx := range cmd.Documents {
		did, err := u.uidrepo.GetID(ctx)
		if err != nil {
			lgr.Error("failed to generate uid", zap.Error(err))
			return nil, err
		}

		doc := cmd.Documents[idx].AsMap()
		indexed := removeDuplicates(findKeys(doc))

		docsToCreate = append(
			docsToCreate,
			*documents.NewCreateDocument(did, doc, indexed),
		)
		idxToCreate = append(
			idxToCreate,
			*indexes.NewCreateIndex(indexed, did),
		)
		res.Documents = append(
			res.Documents,
			contracts.NewIndexedDocument(did, indexed),
		)
	}

	err = u.docrepo.Create(
		ctx,
		docsToCreate,
	)
	if err != nil {
		lgr.Error("failed to create documents", zap.Error(err))
		return nil, err
	}

	err = u.idxrepo.Create(
		ctx,
		idxToCreate,
	)
	if err != nil {
		lgr.Error("failed to create index", zap.Error(err))
		return nil, err
	}

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
				res = append(res, strings.ToLower(s))
			}
			res = append(res, tokenize(s, SpecialCharacterBreakCheck, true)...)
		} else if n, ok := doc[k].(float64); ok {
			res = append(res, strconv.FormatInt(int64(n), 10))
		} else if m, ok := doc[k].(map[string]interface{}); ok {
			res = append(res, findKeys(m)...)
		}
	}
	return res
}

func removeDuplicates(in []string) (out []string) {
	uniques := map[string]struct{}{}
	out = make([]string, 0, len(in))
	for idx := range in {
		if _, ok := uniques[in[idx]]; !ok {
			uniques[in[idx]] = struct{}{}
			out = append(out, in[idx])
		}
	}
	return
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
