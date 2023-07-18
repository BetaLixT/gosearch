// Code generated by protoc-gen-goconsgen. DO NOT EDIT.
// source: proto/gosearch/contracts/models.proto

package contracts

import (
	structpb "google.golang.org/protobuf/types/known/structpb"
)

func NewStringArray(
	values []string,
) *StringArray {
	return &StringArray{
		Values: values,
	}
}

func NewEmptyResponse() *EmptyResponse {
	return &EmptyResponse{}
}

func NewCreateIndexedDocumentCommand(
	document map[string]interface{},
) (*CreateIndexedDocumentCommand, error) {
	resDocument, err := structpb.NewStruct(document)
	if err != nil {
		return nil, err
	}
	return &CreateIndexedDocumentCommand{
		Document: resDocument,
	}, nil
}

func NewSearchQuery(
	query string,
) *SearchQuery {
	return &SearchQuery{
		Query: query,
	}
}

func NewDocumentCreatedResponse(
	indexedTerms []string,
	id uint64,
) *DocumentCreatedResponse {
	return &DocumentCreatedResponse{
		IndexedTerms: indexedTerms,
		Id:           id,
	}
}

func NewSearchResponse() *SearchResponse {
	return &SearchResponse{}
}
