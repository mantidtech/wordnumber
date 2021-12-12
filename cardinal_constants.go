package wordnumber

const and = "and"

const negative = "negative"

const zero = "zero"

var digits = []string{
	"",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

var teens = []string{
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tens = []string{
	"",
	"ten",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

const hundred = "hundred"

var thousands = []string{
	"",
	"thousand",
	"million",
	"billion",
	"trillion",
	"quadrillion",
	"quintillion", //MaxInt64 gets to here
	"sextillion",
	"septillion",
	"octillion",
	"nonillion", // 10^30
}

// https://lcn2.github.io/mersenne-english-name/tenpower/tenpower.html

var digitsIndexed = map[string]int{}
var teensIndexed = map[string]int{}
var tensIndexed = map[string]int{}
var thousandsIndexed = map[string]int{}

func init() {
	for i, d := range digits {
		digitsIndexed[d] = i
	}
	for i, d := range teens {
		teensIndexed[d] = i
	}
	for i, d := range tens {
		tensIndexed[d] = i
	}
	for i, d := range thousands {
		thousandsIndexed[d] = i
	}
}

// allowedFillerWords are words that are allowed to appear in a number expression, but don't impact its value
var allowedFillerWords = map[string]int{
	and: 0,
}
