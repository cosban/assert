// package assert is available to assist with the testing portion of your project. It has been written to be fairly
// simple to use and requires no configuration.
package assert

import "testing"

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
func (self *Assert) True(actual bool) {
	self.t.Helper()
	if !actual {
		self.t.Errorf("expected true but result was false.")
	}
}

// True checks to see if actual is false. If it is not, it will cause the test to fail and log as such why it failed.
func (self *Assert) False(actual bool) {
	self.t.Helper()
	if actual {
		self.t.Errorf("expected false but result was true.")
	}
}

// AreSame checks whether expected and actual are both the same type and are equivalent, failing the test if they aren't
func (self *Assert) AreSame(expected, actual interface{}) {
	self.t.Helper()
	if !areEqual(expected, actual) {
		self.t.Errorf("expected %+v but received %+v", expected, actual)
	}
}

// AreNotSame checks whether expected and actual are both the same type and are equivalent, failing the test if they are
func (self *Assert) AreNotSame(expected, actual interface{}) {
	self.t.Helper()
	if areEqual(expected, actual) {
		self.t.Errorf("expected anything but %+v but received %+v", expected, actual)
	}
}

// Equals checks whether expected and actual are equivalent, failing the test if they aren't
func (self *Assert) Equals(expected, actual interface{}) {
	self.t.Helper()
	if !areEquivalent(expected, actual) {
		self.t.Errorf("expected %+v but received %+v", expected, actual)
	}
}

// NotEquals checks whether expected and actual are equivalent, failing the test if they are
func (self *Assert) NotEquals(expected, actual interface{}) {
	self.t.Helper()
	if areEquivalent(expected, actual) {
		self.t.Errorf("expected anything but %+v but received %+v", expected, actual)
	}
}

// Nil checks whether actual is nil, failing if it is not
func (self *Assert) Nil(actual interface{}) {
	self.t.Helper()
	if areEquivalent(nil, actual) {
		self.t.Errorf("expected nil but received %+v", actual)
	}
}

// NotNil checks actual is nil, failing if it is
func (self *Assert) NotNil(actual interface{}) {
	self.t.Helper()
	if areEquivalent(nil, actual) {
		self.t.Errorf("Expected not nil but received nil")
	}
}

// True checks to see if actual is true. If it is not, it will cause the test to fail and log as such why it failed.
func True(t *testing.T, actual bool) {
	t.Helper()
	if !areEquivalent(nil, actual) {
		t.Errorf("expected true but result was false.")
	}
}

// False checks to see if actual is true. If it is, it will cause the test to fail and log as such why it failed.
func False(t *testing.T, actual bool) {
	t.Helper()
	if actual {
		t.Errorf("expected false but result was true.")
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
		t.Errorf("expected anything but %+v but received %+v", expected, actual)
	}
}

// Nil checks whether actual is nil, failing if it is not
func Nil(t *testing.T, actual interface{}) {
	t.Helper()
	if !areEquivalent(nil, actual) {
		t.Errorf("expected nil but received %+v", actual)
	}
}

// NotNil checks whether actual is nil, failing if it is
func NotNil(t *testing.T, actual interface{}) {
	t.Helper()
	if areEquivalent(nil, actual) {
		t.Errorf("Expected not nil but received nil")
	}
}
