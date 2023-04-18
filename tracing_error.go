package tre

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
)

// Root path is automatically determined from the calling function's source file location.
var rootPath string

func init() {
	// Catch the calling function's source file path.
	_, file, _, _ := runtime.Caller(1)
	// Save the directory alone.
	rootPath = filepath.Dir(file)
}

// TracingError encapsulates an error and collects tracing information
// back to the point where it is handled (logged,ignored,responded...).
type TracingError struct {
	error
	callTrace []tracePoint
}

// tracePoint is a container for individual trace entries in overall call trace
type tracePoint struct {
	line     int
	filename string
	function string
	message  string
	context  map[string]interface{}
}

func (t tracePoint) printOn(b *bytes.Buffer) {
	fmt.Fprintf(b, "%s:%d %s:%s", t.filename, t.line, t.function, t.message)
	for k, v := range t.context {
		fmt.Fprintf(b, " %s=%v ", k, v)
	}
}

// New creates a TracingError with a failure message and optional context information.
// It accepts either an error, a TracingError or nil. If the error is nil then return nil.
func New(err error, msg string, kv ...interface{}) error {
	if err == nil {
		return nil
	}
	var terror *TracingError
	tp := newTracePoint()
	tp.message = msg
	if ewc, ok := err.(*TracingError); ok {
		ewc.callTrace = append(ewc.callTrace, tp)
		terror = ewc
	} else {
		terror = &TracingError{
			error:     err,
			callTrace: []tracePoint{tp},
		}
	}
	for i := 0; i < len(kv); i += 2 {
		// handle odd key count
		if i == len(kv)-1 {
			break
		}
		// expect string keys
		k, ok := kv[i].(string)
		if !ok {
			// convert if not
			k = fmt.Sprintf("%v", kv[i])
		}
		tp.context[k] = kv[i+1]
	}
	return terror
}

func newTracePoint() tracePoint {
	pc, file, line, _ := runtime.Caller(2)
	function := runtime.FuncForPC(pc).Name()
	_, function = filepath.Split(function)
	strip := lengthOfLargestMatchingPrefix(file, rootPath+string(os.PathSeparator))
	file = file[strip:] // trims project's root path.
	return tracePoint{line: line, filename: file, function: function, context: map[string]interface{}{}}
}

// Error returns a pretty report of this error.
func (e TracingError) Error() string {
	buf := new(bytes.Buffer)
	ctx := e.LoggingContext()
	keys := []string{}
	for k := range ctx {
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys))
	for _, each := range keys {
		v := ctx[each]
		fmt.Fprintf(buf, "%s=%v\n", each, v)
	}
	return buf.String()
}

// LoggingContext collects all data for context aware logging purposes.
// Fixed keys are {err,line,func,file,stack} unless the value is empty.
func (e TracingError) LoggingContext() map[string]interface{} {
	ctx := map[string]interface{}{}
	// start with context ; could have reserved keys
	for _, each := range e.callTrace {
		for k, v := range each.context {
			ctx[k] = v
		}
	}
	ctx["err"] = e.error
	ctx["err.type"] = fmt.Sprintf("%T", e.error)
	if len(e.callTrace) > 0 {
		caught := e.callTrace[0]
		ctx["func"] = caught.function
		ctx["loc"] = fmt.Sprintf("%s:%d", caught.filename, caught.line)
		ctx["msg"] = caught.message
	}
	return ctx
}

// Cause returns the initiating error.
func (e TracingError) Cause() error {
	return e.error
}

// As assists with errors.Unwrap
func (e TracingError) As(target interface{}) bool {
	_, ok := target.(TracingError)
	return ok
}

// Unwrap to be able to easily nest/unnest
func (e TracingError) Unwrap() error {
	return e.error
}

// Cause returns the initiating error by recursively seeking it.
func Cause(err error) error {
	if te, ok := err.(*TracingError); ok {
		return Cause(te.Cause())
	}
	return err
}

func lengthOfLargestMatchingPrefix(s1, s2 string) int {
	if len(s1) == 0 {
		return 0
	}
	if len(s2) == 0 {
		return 0
	}
	r1 := []rune(s1)
	r2 := []rune(s2)
	for i, each := range r1 {
		if i == len(r2) {
			return i
		}
		if each != r2[i] {
			return i
		}
	}
	return len(r1)
}
