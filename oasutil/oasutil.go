package oasutil

import (
	"context"
	"net/http"
)

type contextKey string

const (
	varsKey contextKey = "__oasVars"
)

// Vars returns the route variables for the current request, if any.
func Body[T any](r *http.Request) *T {
	if rv := r.Context().Value(varsKey); rv != nil {
		return rv.(*T)
	}

	return nil
}

func RequestWithVars(r *http.Request, vars interface{}) *http.Request {
	ctx := context.WithValue(r.Context(), varsKey, vars)
	return r.WithContext(ctx)
}
