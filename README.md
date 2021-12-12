# Word-Number

This package provides a methods for converting numbers to/from various 'word' representations.

Currently, that means cardinal and ordinal numbers in british english using the short scale.

Roman numerals are extended using the Vinculum system (allowing numbers from 4k to 4M-1).

You can use wordnumber as a package in your own project, or as a command line application.

## Installation

For use as a library:
```
  go get github.com/mantidtech/wordnumber
```

Install the command line application:
```
  go install github.com/mantidtech/wordnumber/cmd/wordnumber@latest
```
Or:
```
  go clone ssh://sooty@github.com/mantidtech/wordnumber.git
  cd wordnumber
  ./build.sh install
```

## Command-Line

```
wordnumber converts an integer to various 'word' formats (and back)

Usage: wordnumber [-i] [-c|-o|-s|-r] <int>
       wordnumber -v
       wordnumber -h
  -h, --help            this message
  -c, --cardinal        display the number given in cardinal form (eg 12 -> 'twelve') (default true)
  -o, --ordinal         display the number given in ordinal form (eg 12 -> 'twelfth')
  -s, --short-ordinal   display the number given in short ordinal form (eg 12 -> '12th')
  -r, --roman           display the number given in roman form (eg 12 -> 'XII')
  -i, --inverse         perform the inverse operation (eg with -i -c 'twelve' -> 12)
  -v, --version         display the version
```


## Types

### Cardinal

The names for numbers

eg
```
    1 -> one
  812 -> eight hundred and twelve
-1000 -> negative one thousand
```
#### Limits

For now the numbers used for input are taken as ints, so the following limits apply:

* Low - `MinInt64`
* High - `MaxInt64`


### Roman

Standard Roman numerals cover numbers in the range 1-3999 (I to MMMCMXCIX).

The Vinculum system uses bars (a vinculum) above numerals to denote x1000 multiplier to numerals, 
thus extending the range to 3,999,999 (M̄M̄M̄C̄M̄X̄C̄MX̄CMXCIX)

Fractions are not yet supported.

eg
```
    1 -> I
   45 -> XLV
 3999 -> MMMCMXCIX
52010 -> L̄MMX
```

* Low - `0` (Roman numerals start at one, though this library provides `nulla` for `0`)
* High - `4M-1`


### Ordinal

The names for a position represented by a number

eg
```
    1 -> first
    2 -> second
  812 -> eight hundred and twelfth
```

For now the numbers used for input are taken as ints, so the following limits apply:

* Low - `0` (negative ordinal numbers don't make sense)
* High - `MaxInt64`


### Short Ordinal

The names for a position represented by a number, using numerals and suffix

eg
```
    1 -> 1st
    2 -> 2nd
  812 -> 812th
```

For now the numbers used for input are taken as ints, so the following limits apply:

* Low - `0` (negative ordinal numbers don't make sense)
* High - `MaxInt64`


## Library

The following methods are defined to perform the operation detailed above.

* `RomanToInt(s string) (res int, err error)`
* `IntToRoman(i int) (string, error)`
* `CardinalToInt(s string) (int, error)`
* `IntToCardinal(i int) string`
* `OrdinalShortToInt(s string) (int, error)`
* `IntToOrdinalShort(i int) (string, error)`
* `OrdinalToInt(s string) (int, error)`
* `IntToOrdinal(i int) (string, error)`


## Known Bugs

* Roman 4000 should maybe/probably be ĪV̄ rather than MV̄

## Todo

* magnitude style numbers, eg giga- milli-
* fractional values
* have a selectable way to represent roman numerals between 4,000->3,999,999
  (not all fonts deal with the unicode combining bar (U+0304) well)

## Copyright

(c) 2021 - Julian Peterson `<code@mantid.org>` - [MIT Licensed](LICENSE.md)

