// Package errorer avoids copypasting complex, repeated error handling.
package errorer

type NumberedErrorer func(err error) (isNil bool)

// Numbered makes a function that should then be called with every error
// to be handled by the provided handler. Those calls return true if the
// the error is nil and false otherwise. The calls will be assigned a number,
// increasing from 1, so the handler knows which call "caught" the error. Only if the error
// is not nil the provided handler is called, with the error and the assigned number as arguments.
func Numbered(handler func(err error, n int)) NumberedErrorer {
	n := 0
	return NumberedErrorer(func(err error) bool {
		n++
		if err == nil {
			return true
		} else {
			handler(err, n)
			return false
		}
	})
}
