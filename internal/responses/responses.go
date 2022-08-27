package responses

import "github.com/jafarsirojov/mongoDB/internal/structs"

func newResponse(code int, message string) structs.Response {
	return structs.Response{Code: code, Message: message}
}

const (
	OkCode           = 200
	BadRequestCode   = 400
	UnauthorizedCode = 401
	NotFoundCode     = 404
	InternalErrCode  = 500
)

var (
	Success      = newResponse(OkCode, "Success")
	BadRequest   = newResponse(BadRequestCode, "BadRequest")
	InternalErr  = newResponse(InternalErrCode, "InternalErr")
	Unauthorized = newResponse(UnauthorizedCode, "Unauthorized")
	NotFound     = newResponse(NotFoundCode, "NotFound")
)
