package errors

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/valyala/fasthttp"
)

const publicTextPrefix = "arm.errors"

var (
	AccessDeniedError   = func() *Error { return New("access denied", fasthttp.StatusForbidden) }
	NotFound            = func() *Error { return New("not found", fasthttp.StatusNotFound) }
	BadMetaValue        = func() *Error { return New("bad meta", fasthttp.StatusBadRequest) }
	ConflictError       = func() *Error { return New("conflict creating operation. id already exsist", fasthttp.StatusConflict) }
	BadRequestError     = func() *Error { return New("bad request", fasthttp.StatusBadRequest) }
	ErrCarAlreadyExists = func() *Error { return New("car already exists", fasthttp.StatusConflict) }
)

func GetWrappedErrorForConfig(err error, paramName, via string) error {
	return Wrap(err, fmt.Sprintf("BrokenConfig - %s has wrong format, must be parsed via `%s`", paramName, via))
}

func buildPublicText(serviceName, subsystem, msg string) string {
	parts := []string{publicTextPrefix}
	if serviceName != "" {
		parts = append(parts, serviceName)
	}
	if subsystem != "" {
		parts = append(parts, subsystem)
	}
	parts = append(parts, toCamelCase(msg))
	return strings.Join(parts, ".")
}

func toCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	if len(words) == 0 {
		return s
	}

	result := strings.ToLower(words[0])
	for _, w := range words[1:] {
		result += capitalize(strings.ToLower(w))
	}

	return result
}

func capitalize(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func getCustomError(err1 error, err2 error) (*Error, *Error, bool) {
	var e1, e2 *Error

	var jsonErr1 *JsonRPCError
	ok := errors.As(err1, &jsonErr1)
	if ok {
		e1 = jsonErr1.Data
	} else {
		ok = errors.As(err1, &e1)
		if !ok {
			return e1, e2, false
		}
	}

	var jsonErr2 *JsonRPCError
	ok = errors.As(err2, &jsonErr2)
	if ok {
		e2 = jsonErr2.Data
	} else {
		ok = errors.As(err2, &e2)
		if !ok {
			return e1, e2, false
		}
	}

	return e1, e2, true
}

func (e *Error) addCause(publicReason, internalReason string) *Error {
	if e.Cause == nil {
		e.Cause = make(map[string]any)
	}
	e.Cause[publicReason] = internalReason
	return e
}
