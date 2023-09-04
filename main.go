package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func printUsage() {
	fmt.Println(`Usage:

# show characters

$ runa 33 90
> ! Z


# show characters in a range

$ runa -r 33 90
> !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ


# show code points of characters

$ runa -c ! Z
> 33 90


# show code points from input

$ runa -i abcde
> 97 98 99 100 101`)
}

func parseInt(s string) (int64, error) {
	if strings.HasPrefix(s, "0x") {
		return strconv.ParseInt(s, 0, 64)
	}
	return strconv.ParseInt(s, 10, 64)
}

func main() {
	if os.Args[1] == "-r" {
		if len(os.Args) != 4 {
			printUsage()
			os.Exit(1)
		}

		fromArg, toArg := os.Args[2], os.Args[3]

		from, err := parseInt(fromArg)
		if err != nil {
			printUsage()
			log.Fatalf("invalid from: %v", fromArg)
		}

		to, err := parseInt(toArg)
		if err != nil {
			printUsage()
			log.Fatalf("invalid to: %v", toArg)
		}

		rd := NewRangeStream(from, to)
		err = outputRunes(rd)
		if err != nil {
			log.Fatalf("unexpected err in -r command: %v", err)
		}
		return
	}

	if os.Args[1] == "-s" {
		if len(os.Args) < 3 {
			printUsage()
			return
		}

		for _, arg := range os.Args[2:] {
			rd := strings.NewReader(arg)
			err := outputCodePoints(rd)
			if err != nil {
				log.Fatalf("unexpected err in -s command: %v", err)
			}
		}
		return
	}

	if os.Args[1] == "-c" {
		if len(os.Args) < 3 {
			printUsage()
			return
		}

		sep := ""
		for _, arg := range os.Args[2:] {
			if utf8.RuneCountInString(arg) > 1 {
				log.Fatalf("-c command doesn't allow words")
			}

			fmt.Printf("%s%d", sep, []rune(arg)[0])
			sep = " "
		}
		fmt.Println("")
		return
	}

	sep := ""
	for _, s := range os.Args[1:] {
		i, err := parseInt(s)
		if err != nil {
			log.Fatal("invalid number: %#v", i)
		}

		fmt.Printf("%s%#c", sep, i)
		sep = " "
	}
	fmt.Println("")
}
