// All of the domain level errors (errors used by and known to the domain)
// first digit identifies the layer (2 = domain)
// the first two digit identify the domain the error was created for 00 refers
// to the acl domain, 01 to the foreigns domain, 02 to the uniques domain and
// 99 refers to a non domain specific error,

package common

import "github.com/betalixt/gorr"

const (
	InvalidContextErrorCode    = 2_99_000
	InvalidContextErrorMessage = "InvalidContextError"

	UserAccessForbiddenErrorCode    = 2_99_001
	UserAccessForbiddenErrorMessage = "UserAccessForbiddenError"

	NoUpdatesErrorCode    = 2_99_002
	NoUpdatesErrorMessage = "NoUpdatesErrorCode"

	ResourceNotFoundCode    = 2_99_003
	ResourceNotFoundMessage = "ResourceNotFound"

	NoItemsProvidedErrorCode    = 2_99_004
	NoItemsProvidedErrorMessage = "NoItemsProvidedError"

	DataEmptyErrorCode    = 2_99_005
	DataEmptyErrorMessage = "DataEmptyError"

	UnexpectedChangedCountErrorCode    = 2_99_006
	UnexpectedChangedCountErrorMessage = "UnexpectedChangedCountError"

	OneOrMoreMisingIDsErrorCode    = 2_99_007
	OneOrMoreMisingIDsErrorMessage = "OneOrMoreMisingIDsError"
)

func NewInvalidContextError() *gorr.Error {
	return gorr.NewError(
		gorr.ErrorCode{
			Code:    InvalidContextErrorCode,
			Message: InvalidContextErrorMessage,
		},
		500,
		"",
	)
}

func NewUserAccessForbiddenError() *gorr.Error {
	return gorr.NewError(
		gorr.ErrorCode{
			Code:    UserAccessForbiddenErrorCode,
			Message: UserAccessForbiddenErrorMessage,
		},
		403,
		"",
	)
}

func NewNoUpdatesError() *gorr.Error {
	return gorr.NewError(
		gorr.ErrorCode{
			Code:    NoUpdatesErrorCode,
			Message: NoUpdatesErrorMessage,
		},
		400,
		"",
	)
}

func NewResourceNotFoundError() *gorr.Error {
	return gorr.NewError(
		gorr.ErrorCode{
			Code:    ResourceNotFoundCode,
			Message: ResourceNotFoundMessage,
		},
		404,
		"",
	)
}

func NewNoItemsProvidedError() *gorr.Error {
	return gorr.NewError(
		gorr.ErrorCode{
			Code:    NoItemsProvidedErrorCode,
			Message: NoItemsProvidedErrorMessage,
		},
		400,
		"",
	)
}

func NewDataEmptyError() *gorr.Error {
	return gorr.NewError(
		gorr.ErrorCode{
			Code:    DataEmptyErrorCode,
			Message: DataEmptyErrorMessage,
		},
		400,
		"",
	)
}

func NewUnexpectedChangedCountError() *gorr.Error {
	return gorr.NewError(
		gorr.ErrorCode{
			Code:    UnexpectedChangedCountErrorCode,
			Message: UnexpectedChangedCountErrorMessage,
		},
		500,
		"",
	)
}

func NewOneOrMoreMisingIDsError() *gorr.Error {
	return gorr.NewError(
		gorr.ErrorCode{
			Code:    OneOrMoreMisingIDsErrorCode,
			Message: OneOrMoreMisingIDsErrorMessage,
		},
		404,
		"",
	)
}
