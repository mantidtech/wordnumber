package wordnumber

const zeroth = "zeroth"

const hundredth = "hundredth"

var digitsOrdinal = []string{
	"",
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
}

var teensOrdinal = []string{
	"tenth",
	"eleventh",
	"twelfth",
	"thirteenth",
	"fourteenth",
	"fifteenth",
	"sixteenth",
	"seventeenth",
	"eighteenth",
	"nineteenth",
}

var tensOrdinal = []string{
	"",
	"tenth",
	"twentieth",
	"thirtieth",
	"fortieth",
	"fiftieth",
	"sixtieth",
	"seventieth",
	"eightieth",
	"ninetieth",
}

var thousandsOrdinal = []string{
	"",
	"thousandth",
	"millionth",
	"billionth",
	"trillionth",
	"quadrillionth",
	"quintillionth",
	"sextillionth",
	"septillionth",
	"octillionth",
	"nonillionth", // 10^30
}

func init() {
	for i, d := range digitsOrdinal {
		digitsIndexed[d] = i
	}
	for i, d := range teensOrdinal {
		teensIndexed[d] = i
	}
	for i, d := range tensOrdinal {
		tensIndexed[d] = i
	}
	for i, d := range thousandsOrdinal {
		thousandsIndexed[d] = i
	}
}
