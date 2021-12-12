package wordnumber

const zeroRoman = "nulla" // not a counting number, but lets us print something for nothing

const macronCombining = '\u0304' // U+0304 = combining bar, U+0305 = combining overline
// const macronCombining = '\u035E' // Combining Double Macron (combines over 2 chars, so not suitable)
func bar(r rune) string {
	return string([]rune{r, macronCombining})
}

var romanLetters = []string{
	"I",
	"V",
	"X",
	"L",
	"C",
	"D",
	"M",
	bar('V'),
	bar('X'),
	bar('L'),
	bar('C'),
	bar('D'),
	bar('M'),
}

var romanValues = map[string]int{
	"I":      1,
	"V":      5,
	"X":      10,
	"L":      50,
	"C":      100,
	"D":      500,
	"M":      1_000,
	bar('V'): 5_000,
	bar('X'): 10_000,
	bar('L'): 50_000,
	bar('C'): 100_000,
	bar('D'): 500_000,
	bar('M'): 1_000_000,
}

// MaxRoman is the largest roman numeral that can be represented
const MaxRoman = 1_000_000*4 - 1

// can use the below if expr auto-calculated non-const
// var MaxRoman = romanValues[romanLetters[len(romanLetters)-1]]*4 - 1
