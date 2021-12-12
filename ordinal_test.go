package wordnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntToOrdinal(t *testing.T) {
	tests := []struct {
		name    string
		num     int
		want    string
		wantErr bool
	}{
		{
			name:    "negative",
			num:     -1,
			want:    "",
			wantErr: true,
		},
		{
			name: "0",
			num:  0,
			want: "zeroth",
		},
		{
			name: "1",
			num:  1,
			want: "first",
		},
		{
			name: "814",
			num:  814,
			want: "eight hundred and fourteenth",
		},
		{
			name: "1230",
			num:  1230,
			want: "one thousand two hundred and thirtieth",
		},
		{
			name: "1234",
			num:  1234,
			want: "one thousand two hundred and thirty fourth",
		},
		{
			name: "4002",
			num:  4002,
			want: "four thousand and second",
		},
		{
			name: "90100",
			num:  90100,
			want: "ninety thousand one hundredth",
		},
		{
			name: "10,000,000,000,000,000",
			num:  10000000000000000,
			want: "ten quadrillionth",
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := IntToOrdinal(tt.num)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

// TestOrdinalToInt provides unit test coverage for OrdinalToInt
// OrdinalToInt converts a word description for an ordinal number into an int
//
// eg. "one hundred and fourth" => 104
func TestOrdinalToInt(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		wantInt   int
		wantError bool
	}{
		{
			name:    "first",
			s:       "first",
			wantInt: 1,
		},
		{
			name:    "eighty second",
			s:       "eighty second",
			wantInt: 82,
		},
		{
			name:      "not a number",
			s:         "foo",
			wantError: true,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotInt, gotError := OrdinalToInt(tt.s)
			if tt.wantError {
				require.Error(t, gotError)
				return
			}
			require.NoError(t, gotError)
			assert.Equal(t, tt.wantInt, gotInt)
		})
	}
}

// TestInternal_constructPositionalGroup provides unit test coverage for constructPositionalGroup
// constructPositionalGroup is a helper method that converts a number < 1000 into words
func TestInternal_constructPositionalGroup(t *testing.T) {
	tests := []struct {
		name string
		i    int
		want []string
	}{
		{},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := constructPositionalGroup(tt.i)
			assert.Equal(t, tt.want, got)
		})
	}
}
