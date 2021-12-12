package wordnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestError_Error provides unit test coverage for Error.Error
// Error implements the Error interface
func TestError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  Error
		want Error
	}{
		{
			name: "out of range",
			err:  ErrOutOfRange,
			want: "number out of range",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.err.Error()
			assert.Equal(t, tt.want.Error(), got)
		})
	}
}

// TestError_With provides unit test coverage for Error.With
// With supplies parameters to error messages
func TestError_With(t *testing.T) {
	tests := []struct {
		name string
		err  Error
		args []interface{}
		want Error
	}{
		{
			name: "out of range",
			err:  ErrOutOfRange,
			args: []interface{}{},
			want: "number out of range",
		},
		{
			name: "overflow",
			err:  ErrFmtOverflow,
			args: []interface{}{5},
			want: "cannot represent numbers greater than 5",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.err.With(tt.args...)
			assert.Equal(t, tt.want, got)
		})
	}
}
