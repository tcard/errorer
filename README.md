# errorer

Package errorer avoids copypasting complex, repeated error handling.

http://godoc.org/github.com/tcard/errorer

```go
http.HandlerFunc("/", func(w http.ResponseWriter, req *http.Request) {
	isNil := errorer.Numbered(func(err error, n int) {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ERROR", n, req.URL)
	})

	result, err := something()
	if !isNil(err) {
		// No boilerplate here.
		return
	}

	otherResult, err := otherSomething()
	if !isNil(err) {
		// Plus, uniform handling in all calls you want.
		return
	}

	...
})
```

## Mandatory idiomaticity-awareness caveat

I wouldn't recomment using this other than for small, personal programs. The idiomatic way of handling errors inside a function is to just return the (possibly wrapped) error, and let side effects (like logging or setting error HTTP headers) be caused by the caller of the function. Like this:

```go
http.HandlerFunc("/", func(w http.ResponseWriter, req *http.Request) {
	err := rootHandler(req.URL.Values.Get("example_arg"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ERROR", req.URL, err)
	}
})

func rootHandler(exampleArg string) {
	result, err := something()
	if err != nil {
		return fmt.Errorf("doing something: %v", err)
	}

	otherResult, err := otherSomething()
	if err != nil {
		return fmt.Errorf("doing otherSomething: %v", err)
	}

	...
}
```

It will also make clean coders in the room happier (or a bit less angry, given that you are using Go), so you won't have to listen to their haughty, generalizing silver bullet doctrine yet again.
