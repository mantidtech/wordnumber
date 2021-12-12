package wordnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		name    string
		num     int
		want    string
		wantErr bool
	}{
		{
			name:    "zero",
			num:     0,
			want:    "nulla",
			wantErr: false,
		},
		{
			name:    "one",
			num:     1,
			want:    "I",
			wantErr: false,
		},
		{
			name:    "three",
			num:     3,
			want:    "III",
			wantErr: false,
		},
		{
			name:    "four",
			num:     4,
			want:    "IV",
			wantErr: false,
		},
		{
			name:    "eight",
			num:     8,
			want:    "VIII",
			wantErr: false,
		},
		{
			name:    "nine",
			num:     9,
			want:    "IX",
			wantErr: false,
		},
		{
			name:    "minus-one",
			num:     -1,
			want:    "",
			wantErr: true,
		},
		{
			name:    "really big",
			num:     5000000,
			want:    "",
			wantErr: true,
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := IntToRoman(tt.num)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name     string
		numerals string
		wantRes  int
		wantErr  bool
	}{
		{
			name:     "one",
			numerals: "I",
			wantRes:  1,
			wantErr:  false,
		},
		{
			name:     "lowercase",
			numerals: "xviii",
			wantRes:  18,
			wantErr:  false,
		},
		{
			name:     "not roman",
			numerals: "DOOKER",
			wantRes:  0,
			wantErr:  true,
		},
		{
			name:     "bad sequence",
			numerals: string([]rune{macronCombining, 'X'}),
			wantRes:  0,
			wantErr:  true,
		},
		{
			name:     "basic combination",
			numerals: "XVI",
			wantRes:  16,
			wantErr:  false,
		},
		{
			name:     "combination with subtraction",
			numerals: "XIV",
			wantRes:  14,
			wantErr:  false,
		},
		{
			name:     "invalid subtraction",
			numerals: "IM",
			wantRes:  999,
			wantErr:  false, // if we're being strict, this is an error
		},
		{
			name:     "bar numerals",
			numerals: "C̄M̄X̄C̄MX̄CMXCIX",
			wantRes:  999999,
			wantErr:  false,
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotRes, err := RomanToInt(tt.numerals)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.wantRes, gotRes)
			}
		})
	}
}

func TestPrintedLen(t *testing.T) {
	tests := []struct {
		name     string
		numerals string
		wantRes  int
	}{
		{
			name:     "blank",
			numerals: "",
			wantRes:  0,
		},
		{
			name:     "I",
			numerals: "I",
			wantRes:  1,
		},
		{
			name:     "III",
			numerals: "III",
			wantRes:  3,
		},
		{
			name:     "bar",
			numerals: "C̄M̄X̄C̄MX̄",
			wantRes:  6,
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotRes := PrintedLen(tt.numerals)
			assert.Equal(t, tt.wantRes, gotRes)
		})
	}
}

// TestInternal_constructRomanFragment provides unit test coverage for constructRomanFragment.
// constructRomanFragment builds a roman set of digits corresponding to a single precision number
func TestInternal_constructRomanFragment(t *testing.T) {
	tests := []struct {
		name  string
		num   int
		place int
		want  string
	}{
		{
			name:  "zero & zero",
			num:   0,
			place: 0,
			want:  "",
		},
		{
			name:  "zero & one",
			num:   0,
			place: 1,
			want:  "",
		},
		{
			name:  "one & zero",
			num:   1,
			place: 0,
			want:  "I",
		},
		{
			name:  "four & zero",
			num:   4,
			place: 0,
			want:  "IV",
		},
		{
			name:  "five & zero",
			num:   5,
			place: 0,
			want:  "V",
		},
		{
			name:  "six & zero",
			num:   6,
			place: 0,
			want:  "VI",
		},
		{
			name:  "seven & one",
			num:   7,
			place: 1,
			want:  "LXX",
		},
		{
			name:  "eight & two",
			num:   8,
			place: 2,
			want:  "DCCC",
		},
		{
			name:  "nine & two",
			num:   9,
			place: 2,
			want:  "CM",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := constructRomanFragment(tt.num, tt.place)
			assert.Equal(t, tt.want, got)
		})
	}
}
