// example printing of a selection of numbers in formats provided by the wordnumber library
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/mantidtech/wordnumber"
	"golang.org/x/sys/unix"
)

func main() {
	numbers := []int{
		-1,
		0,
		1,
		2,
		10,
		21,
		12,
		40,
		87,
		123,
		101,
		100,
		110,
		120,
		130,
		111,
		1000,
		1002,
		1030,
		1200,
		1230,
		1234,
		90100,
		99000,
		100000,
		1000000,
		1743000,
		4000000,
		1234567,
		14000000,
		123456789,
		1000342999,
		10000000000000000,
		math.MaxInt64,
	}

	fmt.Println("\nCardinal")
	for _, n := range numbers {
		fmt.Printf("%20d = %s\n", n, wordnumber.IntToCardinal(n))
	}

	fmt.Println("\nOrdinal")
	for _, n := range numbers {
		w, err := wordnumber.IntToOrdinal(n)
		if err != nil {
			w = err.Error()
		}
		s, err := wordnumber.IntToOrdinalShort(n)
		if err != nil {
			s = err.Error()
		}
		fmt.Printf("%20d = %20s = %s\n", n, s, w)
	}

	fmt.Print("\n\nRoman\n\n")
	printRoman(-1, 3, 1)
	fmt.Println("")
	printRoman(1, 55, 5)
	fmt.Println("")
	printRoman(973, 55, 5)
	fmt.Println("")
	printRoman(999950, 50, 5)
	fmt.Println("")
	printRoman(wordnumber.MaxRoman-1, 3, 1)

	r, err := wordnumber.IntToRoman(52010)
	if err != nil {
		r = err.Error()
	}
	fmt.Printf("%s -> printed length = %d, string length = %d\n", r, wordnumber.PrintedLen(r), len(r))

	var pi = "C̄M̄X̄C̄MX̄CMLXX"
	po, err := wordnumber.RomanToInt(pi)
	if err != nil {
		fmt.Printf("Have error: %s\n", err)
	} else {
		fmt.Printf("%s -> %d\n", pi, po)
	}
}

func printRoman(start, number, maxColumns int) {
	rowWidth := getTerminalColumns() - 1

	sample := sampleRoman(start, number)
	colWidth := largestString(sample)

	numCols := rowWidth / (colWidth + numberSize(start+number) + 5)
	if numCols > maxColumns && maxColumns > 0 {
		numCols = maxColumns
	}

	post := 1
	if number%numCols == 0 {
		post = 0
	}
	numRows := post + number/numCols

	format := fmt.Sprintf("%%%dd = %%%ds  ", numberSize(start+number), -colWidth)

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			x := j*numRows + i
			if x >= len(sample) {
				continue
			}
			fmt.Printf(format, x+start, sample[x])
		}
		fmt.Println()
	}
}

func sampleRoman(start, number int) (res []string) {
	for i := start; i < number+start; i++ {
		r, err := wordnumber.IntToRoman(i)
		if err != nil {
			r = err.Error()
		}
		res = append(res, r)
	}
	return
}

func getTerminalColumns() int {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil || ws.Col == 0 {
		return 80
	}
	return int(ws.Col)
}

func largestString(s []string) int {
	n := 0
	for _, x := range s {
		c := len(x)
		if c > n {
			n = c
		}
	}
	return n
}

func numberSize(i int) int {
	return len(strconv.Itoa(i)) // int(math.Log10(i)) + 1?
}
