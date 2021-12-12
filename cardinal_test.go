package wordnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntToCardinal(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{
			name: "negative",
			num:  -1,
			want: "negative one",
		},
		{
			name: "zero",
			num:  0,
			want: "zero",
		},
		{
			name: "one",
			num:  1,
			want: "one",
		},
		{
			name: "ten",
			num:  10,
			want: "ten",
		},
		{
			name: "one-hundred",
			num:  100,
			want: "one hundred",
		},
		{
			name: "one-hundred-and-ten",
			num:  110,
			want: "one hundred and ten",
		},
		{
			name: "one thousand and two",
			num:  1002,
			want: "one thousand and two",
		},
		{
			name: "one thousand and thirty",
			num:  1030,
			want: "one thousand and thirty",
		},
		{
			name: "one-million-seven-hundred-and-forty-three-thousand",
			num:  1743000,
			want: "one million seven hundred and forty three thousand",
		},
		{
			name: "four-million",
			num:  4000000,
			want: "four million",
		},
		{
			name: "one-million-two-hundred-and-thirty-four-thousand-five-hundred-and-sixty-seven",
			num:  1234567,
			want: "one million two hundred and thirty four thousand five hundred and sixty seven",
		},
		{
			name: "fourteen-million",
			num:  14000000,
			want: "fourteen million",
		},
		{
			name: "one-hundred-and-twenty-three-million-four-hundred-and-fifty-six-thousand-seven-hundred-and-eighty-nine",
			num:  123456789,
			want: "one hundred and twenty three million four hundred and fifty six thousand seven hundred and eighty nine",
		},
		{
			name: "one-billion-three-hundred-and-forty-two-thousand-nine-hundred-and-ninety-nine",
			num:  1000342999,
			want: "one billion three hundred and forty two thousand nine hundred and ninety nine",
		},
		{
			name: "ten-quadrillion",
			num:  10000000000000000,
			want: "ten quadrillion",
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := IntToCardinal(tt.num)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCardinalToInt(t *testing.T) {
	tests := []struct {
		name    string
		expr    string
		want    int
		wantErr bool
	}{
		{
			name:    "empty",
			expr:    "",
			wantErr: true,
		},
		{
			name:    "not a number",
			expr:    "dooker",
			wantErr: true,
		},
		{
			name:    "negative",
			expr:    "negative one",
			want:    -1,
			wantErr: false,
		},
		{
			name:    "zero",
			expr:    "zero",
			want:    0,
			wantErr: false,
		},
		{
			name:    "one",
			expr:    "one",
			want:    1,
			wantErr: false,
		},
		{
			name:    "ten",
			expr:    "ten",
			want:    10,
			wantErr: false,
		},
		{
			name:    "one-hundred",
			expr:    "one hundred",
			want:    100,
			wantErr: false,
		},
		{
			name:    "one-hundred-and-ten",
			expr:    "one hundred and ten",
			want:    110,
			wantErr: false,
		},
		{
			name:    "one thousand and two",
			expr:    "one thousand and two",
			want:    1002,
			wantErr: false,
		},
		{
			name:    "one thousand and thirty",
			expr:    "one thousand and thirty",
			want:    1030,
			wantErr: false,
		},
		{
			name:    "one-million-seven-hundred-and-forty-three-thousand",
			expr:    "one million seven hundred and forty three thousand",
			want:    1743000,
			wantErr: false,
		},
		{
			name:    "four-million",
			expr:    "four million",
			want:    4000000,
			wantErr: false,
		},
		{
			name:    "one-million-two-hundred-and-thirty-four-thousand-five-hundred-and-sixty-seven",
			expr:    "one million two hundred and thirty four thousand five hundred and sixty seven",
			want:    1234567,
			wantErr: false,
		},
		{
			name:    "fourteen-million",
			expr:    "fourteen million",
			want:    14000000,
			wantErr: false,
		},
		{
			name:    "one-hundred-and-twenty-three-million-four-hundred-and-fifty-six-thousand-seven-hundred-and-eighty-nine",
			expr:    "one hundred and twenty three million four hundred and fifty six thousand seven hundred and eighty nine",
			want:    123456789,
			wantErr: false,
		},
		{
			name:    "one-billion-three-hundred-and-forty-two-thousand-nine-hundred-and-ninety-nine",
			expr:    "one billion three hundred and forty two thousand nine hundred and ninety nine",
			want:    1000342999,
			wantErr: false,
		},
		{
			name:    "ten-quadrillion",
			expr:    "ten quadrillion",
			want:    10000000000000000,
			wantErr: false,
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := CardinalToInt(tt.expr)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

// TestInternal_constructWordGroup provides unit test coverage for constructWordGroup
// constructWordGroup is a helper method that converts a number < 1000 into words
func TestInternal_constructWordGroup(t *testing.T) {
	tests := []struct {
		name string
		i    int
		want []string
	}{
		{
			name: "five",
			i:    5,
			want: []string{"five"},
		},
		{
			name: "three five",
			i:    35,
			want: []string{"thirty", "five"},
		},
		{
			name: "nine three five",
			i:    935,
			want: []string{"nine", "hundred", "and", "thirty", "five"},
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := constructWordGroup(tt.i)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestInternal_normalise provides unit test coverage for normalise
func TestInternal_normalise(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Upper",
			s:    "TEN",
			want: "ten",
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := normalise(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
