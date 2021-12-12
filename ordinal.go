package wordnumber

import (
	"strings"
)

// IntToOrdinal returns the given number as a ordinal word, eg 13 => thirteenth
func IntToOrdinal(i int) (string, error) {
	if i == 0 {
		return zeroth, nil
	} else if i < 0 {
		return "", ErrNegative
	}

	groups := GetGroups(i, 3)
	l := groups.lastNonZero()

	var w []string
	for x := len(groups) - 1; x >= 0; x-- {
		var g []string
		if x == 0 {
			g = constructPositionalGroup(groups[x])
		} else {
			g = constructWordGroup(groups[x])
		}
		if len(g) > 0 {
			if x == 0 && len(groups) > 1 && groups[x]/100 == 0 {
				w = append(w, and)
			}
			w = append(w, g...)

			if x != 0 && x == l {
				w = append(w, thousandsOrdinal[x])
			} else if x != 0 {
				w = append(w, thousands[x])
			}
		}
	}

	return strings.Join(w, " "), nil
}

// OrdinalToInt converts a word description for an ordinal number into an int
//
// eg. "one hundred and fourth" => 104
func OrdinalToInt(s string) (int, error) {
	return CardinalToInt(s)
}

// constructPositionalGroup is a helper method that converts a number < 1000 into words
func constructPositionalGroup(i int) []string {
	if i == 0 {
		return nil
	}

	o := i % 10
	t := (i / 10) % 10
	h := (i / 100) % 10

	var w []string
	if h != 0 {
		if o+t != 0 {
			w = append(w, digits[h], hundred, and)
		} else {
			w = append(w, digits[h], hundredth)
		}
	}

	if t == 1 {
		w = append(w, teensOrdinal[o])
	} else {
		if t != 0 {
			if o == 0 {
				w = append(w, tensOrdinal[t])
			} else {
				w = append(w, tens[t])
			}
		}
		if o != 0 {
			w = append(w, digitsOrdinal[o])
		}
	}

	return w
}
