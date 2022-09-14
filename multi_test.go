package tre

import (
	"errors"
	"testing"
)

func TestMulti(t *testing.T) {
	m := NewErrors()
	m.Add(errors.New("test"))
	m.Errorf("v:%v", 1)
	err := m.Err()
	t.Log(err)
}
