package tre

import (
	"errors"
	"fmt"
	"strings"
)

// NewErrors returns a new CompositeError to collect errors and build a single error.
func NewErrors() *CompositeError {
	return new(CompositeError)
}

// CompositeError holds errors
type CompositeError struct {
	list []error
}

// Adds a an error to the list unless it is nil.
func (c *CompositeError) Add(err error) *CompositeError {
	if err == nil {
		return c
	}
	c.list = append(c.list, err)
	return c
}

// Err creates a new error with a message composed of each error
func (c *CompositeError) Err() error {
	if len(c.list) == 0 {
		return nil
	}
	return errors.New(c.message())
}

// Errorf add a new error.
func (c *CompositeError) Errorf(msg string, args ...any) *CompositeError {
	c.list = append(c.list, fmt.Errorf(msg, args...))
	return c
}

// message creates a composite message from all errors separated by a newline.
func (c CompositeError) message() string {
	qb := new(strings.Builder)
	for i, each := range c.list {
		if i > 0 {
			qb.WriteString("\n")
		}
		qb.WriteString(each.Error())
	}
	return qb.String()
}
