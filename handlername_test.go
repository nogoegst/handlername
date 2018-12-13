package handlername

import (
	"net/http"
	"testing"

	"github.com/matryer/is"
	"github.com/nogoegst/handlername/.pkg-test"
)

func TestHandlerName(t *testing.T) {
	is := is.New(t)
	is.Equal(HandlerName(http.HandlerFunc(pkgtest.HandlerFunction)), "HandlerFunction")
	is.Equal(HandlerName(http.HandlerFunc((&pkgtest.HandlerStruct{}).HandlerMethod)), "HandlerMethod")
	is.Equal(HandlerName(&pkgtest.HandlerStruct{}), "HandlerStruct")
	is.Equal(HandlerName(pkgtest.HandlerStructDirect{}), "HandlerStructDirect")
}
