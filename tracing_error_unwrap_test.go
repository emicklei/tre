package tre

import (
	"errors"
	"testing"
)

type TestStruct struct {
	Inside
}

func (t TestStruct) FindMe() string {
	return "Found"
}

type Inside interface {
	error
}

type Findme interface {
	FindMe() string
}

func TestUnWrap(t *testing.T) {

	inner := TestStruct{
		errors.New("inner"),
	}

	wrapped := New(New(inner, "wrap"), "outer wrap")
	var totrack Findme
	if !errors.As(wrapped, &totrack) {
		t.Errorf("Was unable to unwrap and detect inner error")
	}
	if totrack.FindMe() != "Found" {
		t.Errorf("Unwrap matched wrong type/item")
	}
}
