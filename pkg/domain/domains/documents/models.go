package documents

type CreateDocument struct {
	ID       uint64
	Document map[string]interface{}
	Indexed  []string
}

func NewCreateDocument(
	id uint64,
	doc map[string]interface{},
	indexed []string,
) *CreateDocument {
	return &CreateDocument{
		id,
		doc,
		indexed,
	}
}
