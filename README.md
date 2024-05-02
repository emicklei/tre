# tre - tracing error in the call stack

Package tre has the TracingError type to collect stack information when an error is caught.

[![GoDoc](https://godoc.org/github.com/emicklei/tre?status.svg)](https://godoc.org/github.com/emicklei/tre)
[![Go Report Card](https://goreportcard.com/badge/github.com/emicklei/tre)](https://goreportcard.com/report/github.com/emicklei/tre)
[![codecov](https://codecov.io/gh/emicklei/tre/branch/master/graph/badge.svg)](https://codecov.io/gh/emicklei/tre)

## usage

```
package main

import (
	"errors"
	"fmt"

	"github.com/emicklei/tre"
)

func main() {
	err := doThis("sing")
	fmt.Println(err.Error())
}

func doThis(task string) error {
	err := doMore("prepare")
	return tre.New(err, "failed to do this", "task", task)
}

func doMore(task string) error {
	ctx := map[string]any{"task": task, "guality": 42}
	return tre.New(doThat(task), "cannot do more", ctx)
}

func doThat(task string) error {
	return tre.New(errors.New("bummer"), "doing that failed", "task", task)
}
```
see `examples/main.go`


(c) 2016, http://ernestmicklei.com. MIT License