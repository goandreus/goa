// Code generated by goa v3.0.10, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen calcsvc/design

package server

import (
	calc "calcsvc/gen/calc"
)

// NewAddPayload builds a calc service add endpoint payload.
func NewAddPayload(a int, b int) *calc.AddPayload {
	return &calc.AddPayload{
		A: a,
		B: b,
	}
}
