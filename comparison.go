package assert

import (
	"bytes"
	"reflect"
)

func areEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}
	if e, ok := expected.([]byte); ok {
		if a, ok := actual.([]byte); !ok {
			return false
		} else if e == nil || a == nil {
			return e == nil && a == nil
		} else {
			return bytes.Equal(e, a)
		}
	}
	return reflect.DeepEqual(expected, actual)

}

func areEquivalent(expected, actual interface{}) bool {
	if areEqual(expected, actual) {
		return true
	}
	at := reflect.TypeOf(actual)
	if at == nil {
		return false
	}
	ev := reflect.ValueOf(expected)
	if ev.IsValid() {
		if ev.Type().ConvertibleTo(at) {
			// Attempt comparison after type conversion
			return reflect.DeepEqual(ev.Convert(at).Interface(), actual)
		}
	}
	return false
}
