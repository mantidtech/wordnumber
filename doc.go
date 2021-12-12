// Package wordnumber provides a set of methods to convert numbers into various 'word' representations
//
// Currently that means cardinal and ordinal numbers in british english using the short scale, plus roman numerals
package wordnumber

import "fmt"

//go:generate go fmt ./...
//go:generate golint -set_exit_status ./...
//go:generate go vet ./...
//go:generate errcheck -ignoretests ./...
//go:generate staticcheck ./...
//go:generate gosec -quiet ./...
//go:generate go test ./... -short -test.timeout 60s -cover -coverprofile=coverage.out
//go:generate go tool cover -html coverage.out -o coverage.html
//go:generate go tool cover -func coverage.out
//go:generate ./build.sh

var version = "unknown"
var build = "snapshot"

// Version returns the compiled version information for this library
func Version() string {
	return fmt.Sprintf("%s-%s", version, build)
}
