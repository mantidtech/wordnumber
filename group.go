package wordnumber

import "math"

// Group is a division of a number into sets of numerals.
// different translations of numbers are easiest when batched into certain group sizes
type Group []int

// GetGroups divides a number into groups of a given size (number of digits)
// note: don't use s < 1, that's bad
func GetGroups(i int, s int) (g Group) {
	exp := int(math.Pow10(s))
	for d := i; d != 0; d = d / exp {
		g = append(g, d%exp)
	}
	return g
}

// lastNonZero returns the index to the last non-zero item in the group
func (g Group) lastNonZero() int {
	for i, t := range g {
		if t != 0 {
			return i
		}
	}
	return -1
}
