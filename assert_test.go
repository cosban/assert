package assert

import "testing"

func TestTrue(t *testing.T) {
	if success := t.Run("TestTrue", func(t *testing.T) {
		t.Parallel()
		substitute := new(testing.T)
		assert := New(substitute)

		given := true
		assert.True(given)

		if substitute.Failed() {
			t.Fail()
		}
	}); !success {
		t.Error("True should not fail if true is passed in")
	}

	if success := t.Run("TestFalse", func(t *testing.T) {
		t.Parallel()
		substitute := new(testing.T)
		assert := New(substitute)

		given := false
		assert.True(given)

		if !substitute.Failed() {
			t.Fail()
		}
	}); !success {
		t.Error("True should fail if false is passed in")
	}
}

func TestFalse(t *testing.T) {
	if success := t.Run("TestTrue", func(t *testing.T) {
		t.Parallel()
		substitute := new(testing.T)
		assert := New(substitute)

		given := true
		assert.False(given)

		if !substitute.Failed() {
			t.Fail()
		}
	}); !success {
		t.Error("True should not fail if true is passed in")
	}

	if success := t.Run("TestFalse", func(t *testing.T) {
		t.Parallel()
		substitute := new(testing.T)
		assert := New(substitute)

		given := false
		assert.False(given)

		if substitute.Failed() {
			t.Fail()
		}
	}); !success {
		t.Error("True should fail if false is passed in")
	}
}

func TestEquals(t *testing.T) {
	if success := t.Run("TestNils", func(t *testing.T) {
		t.Parallel()
		substitute := new(testing.T)
		assert := New(substitute)

		assert.Equals(nil, nil)

		if substitute.Failed() {
			t.Fail()
		}
	}); !success {
		t.Error("Equals should not fail if nils are passed in")
	}

	if success := t.Run("TestSameVar", func(t *testing.T) {
		t.Parallel()
		substitute := new(testing.T)
		assert := New(substitute)

		given := true
		assert.Equals(given, given)

		if substitute.Failed() {
			t.Fail()
		}
	}); !success {
		t.Error("Equals should not fail if the same variable is passed in to both values")
	}

	if success := t.Run("TestZerosOfSameType", func(t *testing.T) {
		t.Parallel()
		substitute := new(testing.T)
		assert := New(substitute)

		type given struct{}

		assert.Equals(given{}, given{})

		if substitute.Failed() {
			t.Fail()
		}
	}); !success {
		t.Error("Equals should not fail if Zero values of the same type are passed in")
	}

	if success := t.Run("TestDifferentValue", func(t *testing.T) {
		t.Parallel()
		substitute := new(testing.T)
		assert := New(substitute)

		assert.Equals(true, false)

		if !substitute.Failed() {
			t.Fail()
		}
	}); !success {
		t.Error("Equals should fail if different values are passed in")
	}

	if success := t.Run("TestEquivalentValues", func(t *testing.T) {
		t.Parallel()
		substitute := new(testing.T)
		assert := New(substitute)

		type happy struct {
			String     string
			Slice      []string
			Boolean    bool
			unexported uint64
			Float64    float64
			Pointer    *int
		}

		givenInt := 5
		otherInt := 5
		given := happy{
			String: "Because I'm happy",
			Slice: []string{
				"Clap along if you feel",
				"Like a room without a roof",
				"Clap along if you feel",
				"Like happiness is the truth",
			},
			Boolean:    true,
			unexported: uint64(1),
			Float64:    float64(3.14),
			Pointer:    &givenInt,
		}

		other := happy{
			String: "Because I'm happy",
			Slice: []string{
				"Clap along if you feel",
				"Like a room without a roof",
				"Clap along if you feel",
				"Like happiness is the truth",
			},
			Boolean:    true,
			unexported: uint64(1),
			Float64:    float64(3.14),
			Pointer:    &otherInt,
		}

		assert.Equals(given, other)

		if substitute.Failed() {
			t.Fail()
		}
	}); !success {
		t.Error("Equals should not fail if equivalent values are passed in")
	}
}

func TestNil(t *testing.T) {
	if success := t.Run("TestUntypedNil", func(t *testing.T) {
		t.Parallel()
		substitute := new(testing.T)
		assert := New(substitute)

		assert.Nil(nil)

		if substitute.Failed() {
			t.Fail()
		}
	}); !success {
		t.Error("Nil should not fail if untyped nil is passed in")
	}

	if success := t.Run("TestNilPointer", func(t *testing.T) {
		t.Parallel()
		substitute := new(testing.T)
		assert := New(substitute)

		type happy struct{}
		var given *happy = nil

		assert.Nil(given)

		if substitute.Failed() {
			t.Fail()
		}
	}); !success {
		t.Error("Nil should not fail if nil pointer is passed in")
	}

	if success := t.Run("TestNonNilPointer", func(t *testing.T) {
		t.Parallel()
		substitute := new(testing.T)
		assert := New(substitute)

		type happy struct{}
		var given *happy = &happy{}

		assert.Nil(given)

		if !substitute.Failed() {
			t.Fail()
		}
	}); !success {
		t.Error("Nil should fail if non-nil pointer is passed in")
	}

	if success := t.Run("TestUninitializedStruct", func(t *testing.T) {
		t.Parallel()
		substitute := new(testing.T)
		assert := New(substitute)

		type happy struct{}
		var given happy
		assert.Nil(given)

		if !substitute.Failed() {
			t.Fail()
		}
	}); !success {
		t.Error("Nil should fail if uninitialized struct is passed in")
	}

	if success := t.Run("TestZeroStruct", func(t *testing.T) {
		t.Parallel()
		substitute := new(testing.T)
		assert := New(substitute)

		type given struct{}
		assert.Nil(given{})

		if !substitute.Failed() {
			t.Fail()
		}
	}); !success {
		t.Error("Nil should fail if Zero Value Struct is passed in")
	}
}
