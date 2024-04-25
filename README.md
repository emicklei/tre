# tre - tracing error in the call stack

Package tre has the TracingError type to collect stack information when an error is caught.

[![GoDoc](https://godoc.org/github.com/emicklei/tre?status.svg)](https://godoc.org/github.com/emicklei/tre)
[![Go Report Card](https://goreportcard.com/badge/github.com/emicklei/tre)](https://goreportcard.com/report/github.com/emicklei/tre)

## usage

```
func foo(param string) error {
    if err := someOperation(); err != nil {
        return tre.New(err,"foo failed", "param", param)
    }
}
```

(c) 2016, http://ernestmicklei.com. MIT License