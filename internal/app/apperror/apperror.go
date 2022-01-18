package apperror

import "github.com/harunnryd/btrade/internal/app/apperror/core"

var AppErrors = []error{
	core.ErrDBConn,
	core.ErrInvalidPayload,
	core.ErrInvalidHeader,
	core.ErrUnauthorized,
}
