// package assert is available to assist with the testing portion of your project. It has been written to be fairly
// simple to use and requires no configuration.
package assert

import (
	"reflect"
	"testing"
)

// Assert wraps a *testing.T so that it does not need to be passed in to every assertion call
type Assert struct {
	t *testing.T
}

// New returns a new *Assert that contains a *testing.T
func New(t *testing.T) *Assert {
	t.Helper()
	return &Assert{t: t}
}

// True checks to see if actual is true. If it is not, it will cause the test to fail and log as such why it failed.
func (assert *Assert) True(actual bool) {
	assert.t.Helper()
	True(assert.t, actual)
}

// False checks to see if actual is false. If it is not, it will cause the test to fail and log as such why it failed.
func (assert *Assert) False(actual bool) {
	assert.t.Helper()
	False(assert.t, actual)
}

// AreIdentical checks whether expected and actual are both the same type and are equivalent, failing the test if they aren't
func (assert *Assert) AreIdentical(expected, actual interface{}) {
	assert.t.Helper()
	AreIdentical(assert.t, expected, actual)
}

// AreNotIdentical checks whether expected and actual are both the same type and are equivalent, failing the test if they are
func (assert *Assert) AreNotIdentical(expected, actual interface{}) {
	assert.t.Helper()
	AreNotIdentical(assert.t, expected, actual)
}

// Equals checks whether expected and actual are equivalent, failing the test if they aren't
func (assert *Assert) Equals(expected, actual interface{}) {
	assert.t.Helper()
	Equals(assert.t, expected, actual)
}

// NotEquals checks whether expected and actual are equivalent, failing the test if they are
func (assert *Assert) NotEquals(expected, actual interface{}) {
	assert.t.Helper()
	NotEquals(assert.t, expected, actual)
}

// Nil checks whether actual is nil, failing if it is not
func (assert *Assert) Nil(actual interface{}) {
	assert.t.Helper()
	Nil(assert.t, actual)
}

// NotNil checks actual is nil, failing if it is
func (assert *Assert) NotNil(actual interface{}) {
	assert.t.Helper()
	NotNil(assert.t, actual)
}

// True checks to see if actual is true. If it is not, it will cause the test to fail and log as such why it failed.
func True(t *testing.T, actual bool) {
	t.Helper()
	if !actual {
		t.Errorf("expected false but result was true.")
	}
}

// False checks to see if actual is true. If it is, it will cause the test to fail and log as such why it failed.
func False(t *testing.T, actual bool) {
	t.Helper()
	if actual {
		t.Errorf("expected false but result was true.")
	}
}

// AreIdentical checks whether expected and actual are both the same type and are equivalent, failing the test if they aren't
func AreIdentical(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if !areEqual(expected, actual) {
		t.Errorf("expected %+v but received %+v", expected, actual)
	}
}

// AreNotIdentical checks whether expected and actual are both the same type and are equivalent, failing the test if they are
func AreNotIdentical(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if areEqual(expected, actual) {
		t.Errorf("expected anything except %+v but received %+v", expected, actual)
	}
}

// Equals checks whether expected and actual are equivalent, failing the test if they aren't
func Equals(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if !areEquivalent(expected, actual) {
		t.Errorf("expected %+v but received %+v", expected, actual)
	}
}

// NotEquals checks whether expected and actual are equivalent, failing the test if they are
func NotEquals(t *testing.T, expected, actual interface{}) {
	t.Helper()
	if areEquivalent(expected, actual) {
		t.Errorf("expected anything except %+v but received %+v", expected, actual)
	}
}

// Nil checks whether actual is nil, failing if it is not
func Nil(t *testing.T, actual interface{}) {
	t.Helper()
	value := reflect.ValueOf(actual)
	if actual != nil && !(canNil(value) && value.IsNil()) {
		t.Errorf("expected nil but received %+v", actual)
	}
}

// NotNil checks whether actual is nil, failing if it is
func NotNil(t *testing.T, actual interface{}) {
	t.Helper()
	value := reflect.ValueOf(actual)
	if actual == nil || (canNil(value) && value.IsNil()) {
		t.Errorf("Expected not nil but received nil")
	}
}

func canNil(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return true
	default:
		return false
	}
}
