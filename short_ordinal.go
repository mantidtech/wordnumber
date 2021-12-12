package wordnumber

import (
	"fmt"
	"strconv"
)

// IntToOrdinalShort returns the given number as an ordinal number in short form, eg 13 => 13th
func IntToOrdinalShort(i int) (string, error) {
	if i < 0 {
		return "", ErrNegative
	}
	o := i % 10
	t := (i / 10) % 10
	if t == 1 || o > 3 {
		return fmt.Sprintf("%d%s", i, ordinalShortSuffix[0]), nil
	}
	return fmt.Sprintf("%d%s", i, ordinalShortSuffix[o]), nil
}

// OrdinalShortToInt converts a word description for a short ordinal number into an int
//
// eg. "104th" => 104
func OrdinalShortToInt(s string) (int, error) {
	if len(s) < 3 {
		return 0, ErrParseShortOrdinal.With(s)
	}

	num := s[:len(s)-2]
	// suffix := s[len(s)-2:]

	n, err := strconv.ParseInt(num, 0, 64)
	if err != nil {
		return 0, ErrParseShortOrdinal.With(s)
	}

	return int(n), nil
}
