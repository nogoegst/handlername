package handlername

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

func functionName(i interface{}) string {
	rawName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	l := strings.LastIndex(rawName, ".")
	r := strings.LastIndex(rawName, "-")
	if r < 0 || r < l {
		r = len(rawName)
	}
	return rawName[l+1 : r]
}

func structName(i interface{}) string {
	return reflect.Indirect(reflect.ValueOf(i)).Type().Name()
}

// HandlerName returns real name of http.Handler as it is
// present in source code.
// All types of http.Handler are supported: direct struct,
// referenced struct, http.HandlerFunc as a separate function
// and as function with left operand.
func HandlerName(h http.Handler) string {
	if _, ok := h.(http.HandlerFunc); ok {
		return functionName(h)
	}
	return structName(h)
}
