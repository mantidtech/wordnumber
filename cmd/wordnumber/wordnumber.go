// wordnumber converts an integer to various 'word' formats (and back)
//
// Usage: wordnumber [-i] [-c|-o|-s|-r] <int>
//        wordnumber -v
//        wordnumber -h
//   -h, --help            this message
//   -c, --cardinal        display the number given in cardinal form (eg 12 -> 'twelve') (default true)
//   -o, --ordinal         display the number given in ordinal form (eg 12 -> 'twelfth')
//   -s, --short-ordinal   display the number given in short ordinal form (eg 12 -> '12th')
//   -r, --roman           display the number given in roman form (eg 12 -> 'XII')
//   -i, --inverse         perform the inverse operation (eg with -i -c 'twelve' -> 12)
//   -v, --version         display the version
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/mantidtech/wordnumber"
)

var showVersion = false

var method string
var help bool

var methodCardinal, methodOrdinal, methodShortOrdinal, methodRoman bool
var inverse bool

var fs *flag.FlagSet

var helpFd = os.Stderr

func init() {
	fs = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fs.SetOutput(helpFd)

	fs.BoolVarP(&help, "help", "h", false, "this message")
	//fs.StringVarP(&method, "method", "m", "cardinal", "cardinal|ordinal|short-ordinal|roman")

	fs.BoolVarP(&methodCardinal, "cardinal", "c", true, "display the number given in cardinal form (eg 12 -> 'twelve')")
	fs.BoolVarP(&methodOrdinal, "ordinal", "o", false, "display the number given in ordinal form (eg 12 -> 'twelfth')")
	fs.BoolVarP(&methodShortOrdinal, "short-ordinal", "s", false, "display the number given in short ordinal form (eg 12 -> '12th')")
	fs.BoolVarP(&methodRoman, "roman", "r", false, "display the number given in roman form (eg 12 -> 'XII')")

	fs.BoolVarP(&inverse, "inverse", "i", false, "perform the inverse operation (eg with -i -c 'twelve' -> 12)")

	fs.BoolVarP(&showVersion, "version", "v", false, "display the version")

	fs.SortFlags = false
}

func main() {
	var err error

	err = fs.Parse(os.Args[1:])
	if err != nil {
		showHelpAndExit(err)
	}

	if showVersion {
		printConsole("%s\n", wordnumber.Version())
		return
	}

	if fs.NArg() == 0 {
		showHelpAndExit()
	}

	if methodCardinal {
		method = "cardinal"
	}
	if methodOrdinal {
		method = "ordinal"
	}
	if methodShortOrdinal {
		method = "short-ordinal"
	}
	if methodRoman {
		method = "roman"
	}

	var res string
	if !inverse {
		res, err = backwardConversion(method, fs.Arg(0))
	} else {
		res, err = forwardConversion(method, strings.Join(fs.Args(), " "))
	}

	if err != nil {
		showHelpAndExit(err)
	}
	fmt.Println(res)
}

func forwardConversion(method string, words string) (string, error) {
	var err error
	var n int
	switch method {
	case "cardinal":
		n, err = wordnumber.CardinalToInt(words)
	case "ordinal":
		n, err = wordnumber.OrdinalToInt(words)
	case "short-ordinal":
		n, err = wordnumber.OrdinalShortToInt(words)
	case "roman":
		n, err = wordnumber.RomanToInt(words)
	default:
		showHelpAndExit(wordnumber.ErrUnknownConversion.With(method))
	}
	return fmt.Sprintf("%d", n), err
}

func backwardConversion(method string, value string) (string, error) {
	i, err := strconv.Atoi(value)
	if err != nil {
		showHelpAndExit(wordnumber.ErrParseInit.With(fs.Arg(0)))
	}

	var s string
	switch method {
	case "cardinal":
		s = wordnumber.IntToCardinal(i)
	case "ordinal":
		s, err = wordnumber.IntToOrdinal(i)
	case "short-ordinal":
		s, err = wordnumber.IntToOrdinalShort(i)
	case "roman":
		s, err = wordnumber.IntToRoman(i)
	default:
		showHelpAndExit(wordnumber.ErrUnknownConversion.With(method))
	}
	return s, err
}

func showHelpAndExit(desc ...error) {
	for _, d := range desc {
		printConsole("%s\n", d.Error())
	}
	if len(desc) > 0 {
		printConsole("\n")
	}

	printConsole("wordnumber converts an integer to various 'word' formats (and back)\n\n")
	printConsole("Usage: wordnumber [-i] [-c|-o|-s|-r] <int>\n")
	printConsole("       wordnumber -v\n")
	printConsole("       wordnumber -h\n")

	fs.PrintDefaults()
	os.Exit(0)
}

func printConsole(f string, p ...interface{}) {
	_, _ = fmt.Fprintf(helpFd, f, p...)
}
