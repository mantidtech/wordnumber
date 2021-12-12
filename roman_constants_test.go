package wordnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_macron(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want string
	}{
		{
			name: "M",
			r:    'M',
			want: "M̄",
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := bar(tt.r)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestInternal_bar provides unit test coverage for bar
// const macronCombining = '\u035E' // Combining Double Macron (combines over 2 chars, so not suitable)
func TestInternal_bar(t *testing.T) {
	type Args struct {
		r rune
	}

	tests := []struct {
		name string
		args Args
		want string
	}{
		// tests go here
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := bar(tt.args.r)
			assert.Equal(t, tt.want, got)
		})
	}
}
