// All of the implementation level errors
// first digit identifies the layer (3 = infra)
// the first two digit identify the domain the error
// was created for 99 refers to a non domain specific error

package common

import "github.com/betalixt/gorr"

const (
	HexStringGenerationFailedErrorCode    = 3_00_000
	HexStringGenerationFailedErrorMessage = "HexStringGenerationFailedError"
)

func NewHexStringGenerationFailedError(err error) *gorr.Error {
	return gorr.NewError(
		gorr.ErrorCode{
			Code:    HexStringGenerationFailedErrorCode,
			Message: HexStringGenerationFailedErrorMessage,
		},
		500,
		err.Error(),
	)
}
