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

func bar() error {
    info := map[string]interface{}
    // ...doing stuff
    info["some"] = "thing"
    // ...doing more stuff
    info["more"] = 2
    // ...almost done
    info["done"] = true

    // ...then an issue arises
    if err != nil {
        // Answer error containing the full context object as key value pairs
        return tre.New(err, "bar failed", info)
    }

    return nil
}
```

(c) 2016, http://ernestmicklei.com. MIT License