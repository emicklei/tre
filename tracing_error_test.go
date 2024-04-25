package tre

import (
	"errors"
	"strings"
	"testing"
)

func TestTracingErrorString(t *testing.T) {
	err := errors.New("test")
	terr := New(err, "msg", "key", "value")
	suffix := `err=test
err.type=*errors.errorString
func=tre.TestTracingErrorString
key=value
loc=Users/emicklei/xProjects/tre/tracing_error_test.go:11
msg=msg
`
	if got, want := strings.Contains(flatten(terr.Error()), flatten(suffix)), true; got != want {
		t.Log(flatten(terr.Error()))
		t.Log(flatten(suffix))
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

// remove tabs and newlines and spaces
func flatten(s string) string {
	return strings.Replace((strings.Replace(s, "\n", "[n]", -1)), "\t", "[t]", -1)
}

func TestTracingError(t *testing.T) {
	if len(rootPath) == 0 {
		t.Fail()
	}
	e := New(propError1(), "prop failed", "ik", "Koen").(*TracingError)
	if got, want := len(e.callTrace), 3; got != want {
		t.Errorf("got %v want %v", got, want)
	}
	if got, want := Cause(e).Error(), "fail 1"; got != want {
		t.Errorf("got %v want %v", got, want)
	}
	t.Log(e)
}

func propError1() error {
	return New(propError2(), "call propError2()")
}

func propError2() error {
	return New(giveError(), "give failed", "a", 42)
}

func giveError() error {
	return errors.New("fail 1")
}

func TestEmptyTracingError(t *testing.T) {
	e := New(errors.New("empty"), "empty").(*TracingError)
	ctx := e.LoggingContext()
	if ctx["err"] != e.error {
		t.Error("err expected")
	}
	if ctx["err"] != e.error {
		t.Error("err expected")
	}
	if ctx["msg"] != "empty" {
		t.Error("empty expected")
	}
}

func TestLengthOfLargestMatchingPrefix(t *testing.T) {
	for _, each := range []struct {
		s1 string
		s2 string
		i  int
	}{
		{"a", "a", 1},
		{"", "a", 0},
		{"", "", 0},
		{"a", "", 0},
		{"abc", "abc", 3},
		{"abc", "ab c", 2},
	} {
		if got, want := lengthOfLargestMatchingPrefix(each.s1, each.s2), each.i; got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
