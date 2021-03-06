package dberror

import (
	"errors"
	"fmt"
	"sort"
)

var (
	// ErrValidationFailed is the error returned when validation failed. This
	// should always be considered user error.
	ErrValidationFailed = errors.New("validation failed")
)

// Errorable defines an embeddable struct for managing errors on models.
type Errorable struct {
	// errors is the list of errors on the model, usually from validation. The
	// string key is the column name (or virtual column name) of the field that
	// has errors.
	errors map[string][]string
}

// AddError adds a new error to the list.
func (e *Errorable) AddError(key, err string) {
	e.init()
	e.errors[key] = append(e.errors[key], err)
}

// Errors returns the list of errors.
func (e *Errorable) Errors() map[string][]string {
	e.init()
	return e.errors
}

// ErrorMessages returns the list of error messages.
func (e *Errorable) ErrorMessages() []string {
	e.init()

	// Sort keys so the response is in predictable ordering.
	keys := make([]string, 0, len(e.errors))
	for k := range e.errors {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	l := make([]string, 0, len(e.errors))
	for _, k := range keys {
		v := e.errors[k]
		for _, msg := range v {
			l = append(l, fmt.Sprintf("%s %s", k, msg))
		}
	}
	return l
}

// ErrorsFor returns the list of errors for the key
func (e *Errorable) ErrorsFor(key string) []string {
	e.init()
	return e.errors[key]
}

// ErrorOrNil returns ErrValidationFailed if there are any errors, or nil if
// there are none.
func (e *Errorable) ErrorOrNil() error {
	e.init()
	if len(e.errors) == 0 {
		return nil
	}
	return ErrValidationFailed
}

func (e *Errorable) init() {
	if e.errors == nil {
		e.errors = make(map[string][]string)
	}
}
