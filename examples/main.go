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
