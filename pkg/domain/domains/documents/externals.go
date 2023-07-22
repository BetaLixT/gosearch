package documents

import "context"

type IRepository interface {
	Create(
		ctx context.Context,
		docs []CreateDocument,
	) error
	Get(
		ctx context.Context,
		ids []uint64,
	) ([]map[string]interface{}, error)
}
