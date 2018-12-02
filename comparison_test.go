package assert

import (
	"testing"
)

type superItem struct {
	item
	other
}

type item struct {
	Map     map[string]interface{}
	Slice   []interface{}
	String  string
	Float64 float64
	Int64   int64
	Bool    bool
}

type other struct {
	Map     map[string]interface{}
	Slice   []interface{}
	String  string
	Float64 float64
	Int64   int64
	Bool    bool
}

func TestZeroEqualWithSelf(t *testing.T) {
	actual := areEqual(item{}, item{})

	if !actual {
		t.Error("Expected two empty structs of the same type to be equal")
		t.Fail()
	}
}

func TestSimilarObjectsNotEqual(t *testing.T) {
	item := item{}
	other := other{}

	actual := areEqual(item, other)
	if actual {
		t.Error("Expected two similar, but different items to not be equal")
		t.Fail()
	}
}

func TestInitializedAgainstUninitializedMapsNotEqual(t *testing.T) {
	given := item{
		Map: map[string]interface{}{},
	}

	actual := areEqual(given, item{})
	if actual {
		t.Error("Items with initialized maps should not be equal to items with uninitialized maps")
		t.Fail()
	}
}

func TestInitializedAgainstUninitializedSlices(t *testing.T) {
	given := item{
		Slice: []interface{}{},
	}

	actual := areEqual(given, item{})
	if actual {
		t.Error("Items with initialized maps should not be equal to items with uninitialized maps")
		t.Fail()
	}
}

func TestSimilarSlicesEqual(t *testing.T) {
	given := []string{"boop"}

	other := []string{"boop"}

	actual := areEqual(given, other)
	if !actual {
		t.Error("Slices should only be equal if they are the same slice")
		t.Fail()
	}
}
