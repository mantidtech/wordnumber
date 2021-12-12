package wordnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIntToOrdinalShort(t *testing.T) {
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
			want: "0th",
		},
		{
			name: "1",
			num:  1,
			want: "1st",
		},
		{
			name: "814",
			num:  814,
			want: "814th",
		},
		{
			name: "1,230",
			num:  1_230,
			want: "1230th",
		},
		{
			name: "1,234",
			num:  1_234,
			want: "1234th",
		},
		{
			name: "4,002",
			num:  4_002,
			want: "4002nd",
		},
		{
			name: "90,100",
			num:  90_100,
			want: "90100th",
		},
		{
			name: "10,000,000,000,000,000",
			num:  10000000000000000,
			want: "10000000000000000th",
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := IntToOrdinalShort(tt.num)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

// TestOrdinalShortToInt provides unit test coverage for OrdinalShortToInt
// OrdinalShortToInt converts a word description for a short ordinal number into an int
//
// eg. "104th" => 104
func TestOrdinalShortToInt(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		wantInt   int
		wantError bool
	}{
		{
			name:      "empty",
			s:         "",
			wantError: true,
		},
		{
			name:    "first",
			s:       "1st",
			wantInt: 1,
		},
		{
			name:      "not ordinal",
			s:         "something else",
			wantError: true,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			gotInt, gotError := OrdinalShortToInt(tt.s)
			if tt.wantError {
				require.Error(t, gotError)
				return
			}
			require.NoError(t, gotError)
			assert.Equal(t, tt.wantInt, gotInt)
		})
	}
}
