package httpenums

import (
	"fmt"
	"strings"
)

type HTTPMethod string

const (
	GET     HTTPMethod = "GET"
	POST    HTTPMethod = "POST"
	PUT     HTTPMethod = "PUT"
	DELETE  HTTPMethod = "DELETE"
	PATCH   HTTPMethod = "PATCH"
	HEAD    HTTPMethod = "HEAD"
	OPTIONS HTTPMethod = "OPTIONS"
	CONNECT HTTPMethod = "CONNECT"
	TRACE   HTTPMethod = "TRACE"
)

func (m HTTPMethod) String() string {
	return string(m)
}

func IsSafeMethod(m HTTPMethod) bool {
	return m == GET || m == HEAD || m == OPTIONS || m == TRACE
}

func IsIdempotentMethod(m HTTPMethod) bool {
	return IsSafeMethod(m) || m == PUT || m == DELETE
}

type HTTPStatusCode int

const (
	OK                      HTTPStatusCode = 200
	Created                 HTTPStatusCode = 201
	Accepted                HTTPStatusCode = 202
	NoContent               HTTPStatusCode = 204
	BadRequest              HTTPStatusCode = 400
	Unauthorized            HTTPStatusCode = 401
	Forbidden               HTTPStatusCode = 403
	NotFound                HTTPStatusCode = 404
	MethodNotAllowed        HTTPStatusCode = 405
	Conflict                HTTPStatusCode = 409
	UnsupportedMediaType    HTTPStatusCode = 415
	TooManyRequests         HTTPStatusCode = 429
	InternalServerError     HTTPStatusCode = 500
	NotImplemented          HTTPStatusCode = 501
	BadGateway              HTTPStatusCode = 502
	ServiceUnavailable      HTTPStatusCode = 503
	GatewayTimeout          HTTPStatusCode = 504
	HTTPVersionNotSupported HTTPStatusCode = 505
)

var httpStatusText = map[HTTPStatusCode]string{
	OK:                      "OK",
	Created:                 "Created",
	Accepted:                "Accepted",
	NoContent:               "No Content",
	BadRequest:              "Bad Request",
	Unauthorized:            "Unauthorized",
	Forbidden:               "Forbidden",
	NotFound:                "Not Found",
	MethodNotAllowed:        "Method Not Allowed",
	Conflict:                "Conflict",
	UnsupportedMediaType:    "Unsupported Media Type",
	TooManyRequests:         "Too Many Requests",
	InternalServerError:     "Internal Server Error",
	NotImplemented:          "Not Implemented",
	BadGateway:              "Bad Gateway",
	ServiceUnavailable:      "Service Unavailable",
	GatewayTimeout:          "Gateway Timeout",
	HTTPVersionNotSupported: "HTTP Version Not Supported",
}

func (c HTTPStatusCode) String() string {
	msg, ok := httpStatusText[c]
	if !ok {
		msg = "Unknown Status"
	}
	return fmt.Sprintf("%d %s", c, msg)
}

func (c HTTPStatusCode) IsError() bool {
	return c == 400
}

func (c HTTPStatusCode) IsServerError() bool {
	return c == 500
}

type ContentType string

const (
	ApplicationJSON        ContentType = "application/json"
	ApplicationXML         ContentType = "application/xml"
	TextHTML               ContentType = "text/html"
	TextPlain              ContentType = "text/plain"
	ApplicationForm        ContentType = "application/x-www-form-urlencoded"
	MultipartForm          ContentType = "multipart/form-data"
	ApplicationOctetStream ContentType = "application/octet-stream"
)

func (c ContentType) String() string {
	return string(c)
}

func (c ContentType) IsText() bool {
	return strings.HasPrefix(string(c), "text/") || c == ApplicationJSON || c == ApplicationXML
}

func (c ContentType) IsBinary() bool {
	return c == ApplicationOctetStream || c == MultipartForm
}
