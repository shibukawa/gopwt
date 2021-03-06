package translatedassert

import (
	"reflect"
	"testing"
)

func TestBcomplex(t *testing.T) {
	if complex(1, 2) != Bcomplex(1, 2) {
		t.Errorf("Bcomplex behavior should be equal to complex")
	}

	if complex(1.0, 2.2) != Bcomplex(1.0, 2.2) {
		t.Errorf("Bcomplex behavior should be equal to complex, %#v %#v", complex(1.0, 2.2), Bcomplex(1.0, 2.2))
	}

	if complex(float64(1.0), float64(2.2)) != Bcomplex(float64(1.0), float64(2.2)) {
		t.Errorf("Bcomplex behavior should be equal to complex, %#v %#v", complex(1.0, 2.2), Bcomplex(1.0, 2.2))
	}

	if complex(float32(1.0), float32(2.2)) != Bcomplex(float32(1.0), float32(2.2)).(complex64) {
		t.Errorf("Bcomplex behavior should be equal to complex, %#v %#v", complex(1.0, 2.2), Bcomplex(1.0, 2.2))
	}
}

func TestBnew(t *testing.T) {
	var expected interface{}
	var got interface{}

	// builtin types
	expected = new(string)
	got = Bnew(RVOf(new(string)).Elem().Type())
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected=%#v, got=%#v", expected, got)
	}

	expected = new(uint32)
	got = Bnew(RVOf(new(uint32)).Elem().Type())
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected=%#v, got=%#v", expected, got)
	}

	expected = new(error)
	got = Bnew(RVOf(new(error)).Elem().Type())
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected=%#v, got=%#v", expected, got)
	}

	// slice, map
	expected = new([]string)
	got = Bnew(RVOf(new([]string)).Elem().Type())
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected=%#v, got=%#v", expected, got)
	}

	expected = new(map[string]string)
	got = Bnew(RVOf(new(map[string]string)).Elem().Type())
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected=%#v, got=%#v", expected, got)
	}

	// chan
	expected = new(chan int)
	got = Bnew(RVOf(new(chan int)).Elem().Type())
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected=%#v, got=%#v", expected, got)
	}

	// other type
	type foo struct{}
	expected = new(foo)
	got = Bnew(RVOf(new(foo)).Elem().Type())
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected=%#v, got=%#v", expected, got)
	}
}

func TestBmake(t *testing.T) {
	var expected interface{}
	var got interface{}

	// slice
	expected = make([]string, 12, 14)
	got = Bmake(RTOf([]string{}), 12, 14)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected=%#v, got=%#v", expected, got)
	}

	expected = make([]string, 12)
	got = Bmake(RTOf([]string{}), 12, 12)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected=%#v, got=%#v", expected, got)
	}
	if len(expected.([]string)) != len(got.([]string)) {
		t.Errorf("expected=%#v, got=%#v", len(expected.([]string)), len(got.([]string)))
	}
	if cap(expected.([]string)) != cap(got.([]string)) {
		t.Errorf("expected=%#v, got=%#v", cap(expected.([]string)), cap(got.([]string)))
	}

	// map
	expected = make(map[string]string)
	got = Bmake(RTOf(map[string]string{}))
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected=%#v, got=%#v", expected, got)
	}
}
