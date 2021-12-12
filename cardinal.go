package wordnumber

import (
	"math"
	"strings"
	"unicode"
)

// IntToCardinal returns the word form of the given number
//
// eg. 37 => "thirty seven"
func IntToCardinal(i int) string {
	var w []string

	if i == 0 {
		return zero
	} else if i < 0 {
		w = append(w, negative)
		i = i * -1
	}

	groups := GetGroups(i, 3)

	for x := len(groups) - 1; x >= 0; x-- {
		g := constructWordGroup(groups[x])
		if len(g) > 0 {
			if x == 0 && len(groups) > 1 && groups[x]/100 == 0 {
				w = append(w, and)
			}

			w = append(w, g...)

			if x != 0 {
				w = append(w, thousands[x])
			}
		}
	}

	return strings.Join(w, " ")
}

// CardinalToInt converts a word description for a number into an int
//
// eg. "one hundred and four" => 104
func CardinalToInt(s string) (int, error) {
	if s == zero {
		return 0, nil
	}
	n := normalise(s)

	if len(n) == 0 {
		return 0, ErrNoInput
	}

	words := strings.Split(n, " ")

	multiplier := 1
	if words[0] == negative {
		multiplier = -1
		words = words[1:]
	}

	var acc = 0
	var a = 0
	for _, w := range words {
		i, isDigit := digitsIndexed[w]
		if isDigit {
			a += i
			continue
		}

		e, isTeens := teensIndexed[w]
		if isTeens {
			a += e + 10
			continue
		}

		t, isTens := tensIndexed[w]
		if isTens {
			a += t * 10
			continue
		}

		if w == hundred {
			a = a * 100
			continue
		}

		k, isThousands := thousandsIndexed[w]
		if isThousands {
			acc += a * int(math.Pow10(k*3))
			a = 0
			continue
		}

		_, isFiller := allowedFillerWords[w]
		if isFiller {
			continue
		}

		return 0, ErrParseCardinal.With(w)
	}
	acc += a
	acc *= multiplier

	return acc, nil
}

// constructWordGroup is a helper method that converts a number < 1000 into words
func constructWordGroup(i int) []string {
	if i == 0 {
		return nil
	}

	o := i % 10
	t := (i / 10) % 10
	h := (i / 100) % 10

	var w []string
	if h != 0 {
		w = append(w, digits[h], hundred)
		if o+t != 0 {
			w = append(w, and)
		}
	}

	if t == 1 {
		w = append(w, teens[o])
	} else {
		if t != 0 {
			w = append(w, tens[t])
		}

		if o != 0 {
			w = append(w, digits[o])
		}
	}

	return w
}

func normalise(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsSpace(c) {
			// #nosec G104 - WriteRune has an error param but only ever returns nil, so we don't care
			_, _ = sb.WriteRune(unicode.ToLower(c))
		}
	}
	return sb.String()
}
