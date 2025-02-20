package util

import (
	"reflect"
)

// IsEmpty checks if the given struct is empty using reflection
func IsEmpty(v interface{}) bool {
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}
