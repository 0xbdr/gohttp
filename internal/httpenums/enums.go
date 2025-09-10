package httpenums

import (
	"fmt"
)

// HTTPMethod represents an HTTP method
type HTTPMethod string

// HTTP methods
const (
	GET     HTTPMethod = "GET"
	POST    HTTPMethod = "POST"
	PUT     HTTPMethod = "PUT"
	DELETE  HTTPMethod = "DELETE"
	PATCH   HTTPMethod = "PATCH"
	HEAD    HTTPMethod = "HEAD"
	OPTIONS HTTPMethod = "OPTIONS"
)

// String returns the string representation of the HTTP method
func (m HTTPMethod) String() string {
	return string(m)
}

// HTTPStatusCode represents an HTTP status code
type HTTPStatusCode int

// HTTP status codes
const (
	OK                     HTTPStatusCode = 200
	Created                HTTPStatusCode = 201
	BadRequest             HTTPStatusCode = 400
	Unauthorized           HTTPStatusCode = 401
	Forbidden              HTTPStatusCode = 403
	NotFound               HTTPStatusCode = 404
	InternalServerError    HTTPStatusCode = 500
	ServiceUnavailable     HTTPStatusCode = 503
	HTTPVersionNotSupported HTTPStatusCode = 505
)

// String returns the string representation of the HTTP status code
func (c HTTPStatusCode) String() string {
	return fmt.Sprintf("%d %s", c, httpStatusCodeMessages[c])
}

var httpStatusCodeMessages = map[HTTPStatusCode]string{
	OK:                     "OK",
	Created:                "Created",
	BadRequest:             "Bad Request",
	Unauthorized:           "Unauthorized",
	Forbidden:              "Forbidden",
	NotFound:               "Not Found",
	InternalServerError:    "Internal Server Error",
	ServiceUnavailable:     "Service Unavailable",
	HTTPVersionNotSupported: "HTTP Version Not Supported",
}

// ContentType represents an HTTP content type
type ContentType string

// Content types
const (
	ApplicationJSON ContentType = "application/json"
	ApplicationXML  ContentType = "application/xml"
	TextHTML        ContentType = "text/html"
	TextPlain       ContentType = "text/plain"
	ApplicationForm ContentType = "application/x-www-form-urlencoded"
	MultipartForm   ContentType = "multipart/form-data"
)

// String returns the string representation of the content type
func (c ContentType) String() string {
	return string(c)
}
