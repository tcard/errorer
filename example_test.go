package errorer_test

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/tcard/errorer"
)

func ExampleNumbered() {
	doSomething := func() error { return nil }
	doFailing := func() (int, error) { return 0, errors.New("I tried :(") }
	wontExecute := func() (int, error) { return 0, nil }

	// This could be e. g. a http.ResponseWriter.
	var publicResponse io.Writer = ioutil.Discard

	// isNil is a function that calls the given handler when needed.
	isNil := errorer.Numbered(func(err error, n int) {
		fmt.Fprintln(publicResponse, "Unexpected error. Sorry!")
		fmt.Println("ERROR mypackage.MyFunction", n, err)
	})

	// isNil(err) replaces the usual if err == nil.
	if err := doSomething(); !isNil(err) {
		return
	}

	fmt.Println("Got first result OK!")

	otherResult, err := doFailing()
	if !isNil(err) {
		// No need to put handling code here; just return early.
		return
	}

	fmt.Println("Got other result result OK:", otherResult)

	if nevermind, err := wontExecute(); isNil(err) {
		_ = nevermind
	} else {
		return
	}

	// Output:
	// Got first result OK!
	// ERROR mypackage.MyFunction 2 I tried :(
}
