package assert

import "testing"

type Assert struct {
	t *testing.T
}

func New(t *testing.T) *Assert {
	return &Assert{t: t}
}

func (self *Assert) True(e bool) {
	if !e {
		self.t.Fatalf("expected true but result was false.")
	}
}

func (self *Assert) False(e bool) {
	if e {
		self.t.Fatalf("expected false but result was true.")
	}
}

func (self *Assert) Equals(e interface{}, a interface{}) {
	if e != a {
		self.t.Fatalf("expected %v but received %v", e, a)
	}
}

func (self *Assert) NotEquals(e interface{}, a interface{}) {
	if e == a {
		self.t.Fatalf("expected anything but %v but received %v", e, a)
	}
}

func (self *Assert) Nil(e interface{}) {
	if e != (interface{})(nil) {
		self.t.Fatalf("expected nil but received %v", e)
	}
}

func (self *Assert) NotNil(e interface{}) {
	if e == (interface{})(nil) {
		self.t.Fatalf("Expected not nil but received nil")
	}
}

func True(t *testing.T, e bool) {
	if !e {
		t.Fatalf("expected true but result was false.")
	}
}

func False(t *testing.T, e bool) {
	if e {
		t.Fatalf("expected false but result was true.")
	}
}

func Equals(t *testing.T, e interface{}, a interface{}) {
	if e != a {
		t.Fatalf("expected %v but received %v", e, a)
	}
}

func NotEquals(t *testing.T, e interface{}, a interface{}) {
	if e == a {
		t.Fatalf("expected anything but %v but received %v", e, a)
	}
}

func Nil(t *testing.T, e interface{}) {
	if e != nil {
		t.Fatalf("expected nil but received %v", e)
	}
}

func NotNil(t *testing.T, e interface{}) {
	if e == nil {
		t.Fatalf("Expected not nil but received nil")
	}
}
