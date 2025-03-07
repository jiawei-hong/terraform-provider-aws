package types

import (
	"reflect"
)

// IsZero returns true if `v` is `nil` or points to the zero value of `T`.
func IsZero[T any](v *T) bool {
	return v == nil || reflect.ValueOf(*v).IsZero()
}
