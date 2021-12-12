package wordnumber

import (
	"strings"
	"unicode"
)

// IntToRoman returns the given number in roman numerals
//
// eg. 54 => "LIV"
func IntToRoman(i int) (string, error) {
	if i == 0 {
		return zeroRoman, nil
	} else if i < 0 {
		return "", ErrNegative
	} else if i > MaxRoman {
		return "", ErrFmtOverflow.With(MaxRoman)
	}

	groups := GetGroups(i, 1)

	var w []string
	for x := len(groups) - 1; x >= 0; x-- {
		g := constructRomanFragment(groups[x], x)
		w = append(w, g)
	}

	return strings.Join(w, ""), nil
}

// constructRomanFragment builds a roman set of numerals corresponding to a single decimal numeral in a given decimal place
func constructRomanFragment(numeral int, place int) string {
	var res string
	idx := place * 2

	if numeral == 4 {
		res = romanLetters[idx] + romanLetters[idx+1]
	} else if numeral == 9 {
		res = romanLetters[idx] + romanLetters[idx+2]
	} else {
		if numeral >= 5 {
			numeral -= 5
			res = romanLetters[idx+1]
		}
		for y := 0; y < numeral; y++ {
			res += romanLetters[idx]
		}
	}

	return res
}

// RomanToInt returns the decimal value of the given roman numeral
func RomanToInt(s string) (res int, err error) {
	r := []rune(s)
	var c string

	last := 0
	for i := len(r) - 1; i >= 0; i-- {
		if r[i] == macronCombining {
			if i == 0 {
				err = ErrBadRoman
				break
			}
			c = bar(unicode.ToUpper(r[i-1]))
			i--
		} else {
			c = string(unicode.ToUpper(r[i]))
		}

		cur, found := romanValues[c]
		if !found {
			err = ErrBadRoman
			break
		}

		if last > cur {
			res -= cur
		} else {
			res += cur
		}

		last = cur
	}

	return
}

// PrintedLen returns the printed length of a string, ie the number of cells on a terminal it uses
//
// this method is useful because the extended roman numerals use combining unicode characters to render the over-bar.
// That means
//  len("M̄") == 2 // ie the number of runes in the string
// and
//  PrintedLen("M̄") == 1
func PrintedLen(s string) (res int) {
	for _, c := range s {
		if c != macronCombining {
			res++
		}
	}
	return
}
