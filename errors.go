package wordnumber

import "fmt"

// Error is the error type returned from this package
type Error string

// Error implements the Error interface
func (e Error) Error() string {
	return string(e)
}

// With supplies parameters to error messages
func (e Error) With(p ...interface{}) Error {
	return Error(fmt.Sprintf(string(e), p...))
}

// Possible errors returned from this package
const (
	ErrNotImplemented    Error = "method '%s' is not yet implemented"
	ErrOutOfRange        Error = "number out of range"
	ErrNegative          Error = "cannot represent negative numbers"
	ErrFmtOverflow       Error = "cannot represent numbers greater than %d"
	ErrUnknownConversion Error = "unknown conversion method %s"
	ErrParseCardinal     Error = "unknown word '%s' in expression, can't parse as a cardinal number"
	ErrParseOrdinal      Error = "unknown word '%s' in expression, can't parse as an ordinal number"
	ErrParseShortOrdinal Error = "can't parse '%s' as a short ordinal number"
	ErrBadRoman          Error = "unexpected character parsing roman numerals"
	ErrParseInit         Error = "could not process '%s' as an integer"
	ErrNoInput           Error = "no input was provided"
)
