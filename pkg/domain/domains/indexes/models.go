package indexes

type CreateIndex struct {
	Keys []string
	doc  uint64
}

func NewCreateIndex(
	keys []string,
	doc uint64,
) *CreateIndex {
	return &CreateIndex{
		keys,
		docs,
	}
}
