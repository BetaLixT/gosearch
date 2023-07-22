package indexes

import "golang.org/x/net/context"

type IRepository interface {
	Create(
		ctx context.Context,
		idxs []CreateIndex,
	) error
	GetDocs(
		ctx context.Context,
		keys []string,
	) ([]uint64, error)
}
