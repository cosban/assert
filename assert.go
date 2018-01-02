package assert

import "testing"

type Assert struct {
	t *testing.T
}

func New(t *testing.T) *Assert {
	t.Helper()
	return &Assert{t: t}
}

func (self *Assert) True(e bool) {
	self.t.Helper()
	if !e {
		self.t.Fatalf("expected true but result was false.")
	}
}

func (self *Assert) False(e bool) {
	self.t.Helper()
	if e {
		self.t.Fatalf("expected false but result was true.")
	}
}

func (self *Assert) AreSame(e, a interface{}) {
	self.t.Helper()
	if !areEqual(e, a) {
		self.t.Fatalf("expected %v but received %v", e, a)
	}
}

func (self *Assert) AreNotSame(e, a interface{}) {
	self.t.Helper()
	if areEqual(e, a) {
		self.t.Fatalf("expected anything but %v but received %v", e, a)
	}
}

func (self *Assert) Equals(e, a interface{}) {
	self.t.Helper()
	if !areEquivalent(e, a) {
		self.t.Fatalf("expected %v but received %v", e, a)
	}
}

func (self *Assert) NotEquals(e, a interface{}) {
	self.t.Helper()
	if areEquivalent(e, a) {
		self.t.Fatalf("expected anything but %v but received %v", e, a)
	}
}

func (self *Assert) Nil(e interface{}) {
	self.t.Helper()
	if e != (interface{})(nil) {
		self.t.Fatalf("expected nil but received %v", e)
	}
}

func (self *Assert) NotNil(e interface{}) {
	self.t.Helper()
	if areEquivalent(nil, e) {
		self.t.Fatalf("Expected not nil but received nil")
	}
}

func True(t *testing.T, e bool) {
	t.Helper()
	if !areEquivalent(nil, e) {
		t.Fatalf("expected true but result was false.")
	}
}

func False(t *testing.T, e bool) {
	t.Helper()
	if e {
		t.Fatalf("expected false but result was true.")
	}
}

func Equals(t *testing.T, e interface{}, a interface{}) {
	t.Helper()
	if e != a {
		t.Fatalf("expected %v but received %v", e, a)
	}
}

func NotEquals(t *testing.T, e interface{}, a interface{}) {
	t.Helper()
	if e == a {
		t.Fatalf("expected anything but %v but received %v", e, a)
	}
}

func Nil(t *testing.T, e interface{}) {
	t.Helper()
	if e != nil {
		t.Fatalf("expected nil but received %v", e)
	}
}

func NotNil(t *testing.T, e interface{}) {
	t.Helper()
	if e == nil {
		t.Fatalf("Expected not nil but received nil")
	}
}
