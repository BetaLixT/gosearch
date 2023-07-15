package uids

import (
	"context"

	"github.com/BetaLixT/gosearch/pkg/domain/base/uids"
	"github.com/bwmarrin/snowflake"
)

// UIDRepository for generating unique ids
type UIDRepository struct {
	sf *snowflake.Node
}

var _ uids.IRepository = (*UIDRepository)(nil)

// NewUIDRepository Constructs new UUIDRepository
func NewUIDRepository(
	sf *snowflake.Node,
) *UIDRepository {
	return &UIDRepository{
		sf: sf,
	}
}

// GetID generates and returns a unique id
func (r *UIDRepository) GetID(
	ctx context.Context,
) (uint64, error) {
	if ctx.Err() != nil {
		return 0, ctx.Err()
	}
	return uint64(r.sf.Generate().Int64()), nil
}
