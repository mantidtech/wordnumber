package wordnumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getGroups(t *testing.T) {

	tests := []struct {
		name   string
		number int
		size   int
		want   Group
	}{
		{
			name:   "0-one",
			number: 0,
			size:   1,
			want:   nil,
		},
		{
			name:   "0-four",
			number: 0,
			size:   4,
			want:   nil,
		},
		{
			name:   "1234-one",
			number: 1234,
			size:   1,
			want:   Group{4, 3, 2, 1},
		},
		{
			name:   "1234-two",
			number: 1234,
			size:   2,
			want:   Group{34, 12},
		},
		{
			name:   "1234-three",
			number: 1234,
			size:   3,
			want:   Group{234, 1},
		},
		{
			name:   "1234-four",
			number: 1234,
			size:   4,
			want:   Group{1234},
		},
		{
			name:   "1234-five",
			number: 1234,
			size:   5,
			want:   Group{1234},
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := GetGroups(tt.number, tt.size)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_group_lastNonZero(t *testing.T) {
	tests := []struct {
		name string
		g    Group
		want int
	}{
		{
			name: "zero-as-nil",
			g:    nil,
			want: -1,
		},
		{
			name: "zero-as-empty",
			g:    Group{},
			want: -1,
		},
		{
			name: "zero",
			g:    Group{0},
			want: -1,
		},
		{
			name: "ten",
			g:    Group{0, 1},
			want: 1,
		},
		{
			name: "one",
			g:    Group{1},
			want: 0,
		},
		{
			name: "eight-hundred-and-four",
			g:    Group{8, 0, 4},
			want: 0,
		},
	}
	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.g.lastNonZero()
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestGetGroups provides unit test coverage for GetGroups
// GetGroups divides a number into groups of a given size (number of digits)
// note: don't use s < 1, that's bad
func TestGetGroups(t *testing.T) {
	type Args struct {
		i int
		s int
	}

	tests := []struct {
		name string
		args Args
		want Group
	}{
		{
			name: "two twos",
			args: Args{
				i: 1234,
				s: 2,
			},
			want: Group{34, 12},
		},
		{
			name: "three ones",
			args: Args{
				i: 123,
				s: 1,
			},
			want: Group{3, 2, 1},
		},
		{
			name: "odd size",
			args: Args{
				i: 12345,
				s: 3,
			},
			want: Group{345, 12},
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := GetGroups(tt.args.i, tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}

// TestInternal_Group_lastNonZero provides unit test coverage for Group.lastNonZero
func TestInternal_Group_lastNonZero(t *testing.T) {
	tests := []struct {
		name string
		obj  Group
		want int
	}{
		{
			name: "empty",
			obj:  Group{},
			want: -1,
		},
		{
			name: "none",
			obj:  Group{1, 2, 3},
			want: 0,
		},
		{
			name: "first",
			obj:  Group{0, 2, 3},
			want: 1,
		},
		{
			name: "last",
			obj:  Group{1, 2, 0},
			want: 0,
		},
		{
			name: "middle",
			obj:  Group{1, 0, 3},
			want: 0,
		},
		{
			name: "all",
			obj:  Group{0, 0, 0},
			want: -1,
		},
		{
			name: "first of 2",
			obj:  Group{1, 0, 0},
			want: 0,
		},
		{
			name: "middle of 2",
			obj:  Group{0, 3, 0},
			want: 1,
		},
		{
			name: "last of 2",
			obj:  Group{0, 0, 1},
			want: 2,
		},
	}

	for _, st := range tests {
		tt := st
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.obj.lastNonZero()
			assert.Equal(t, tt.want, got)
		})
	}
}
