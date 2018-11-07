package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func printUsage() {
	fmt.Println(`Usage:
$ runa 33 90
> !"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ

$ runa -c a
> 97`)
}

func parseInt(s string) (int64, error) {
	if strings.HasPrefix(s, "0x") {
		return strconv.ParseInt(s, 0, 64)
	}
	return strconv.ParseInt(s, 10, 64)
}

func main() {
	if os.Args[1] == "-c" {
		if len(os.Args) < 3 {
			printUsage()
			return
		}

		sep := ""
		for _, r := range []rune(os.Args[2]) {
			fmt.Print(sep, r)
			sep = " "
		}
		fmt.Println("")
		return
	}

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	fromArg := os.Args[1]

	from, err := parseInt(fromArg)
	if err != nil {
		log.Fatalf("Invalid from: %v", fromArg)
		printUsage()
		return
	}

	if len(os.Args) == 2 {
		fmt.Printf("%#c\n", from)
		return
	}

	toArg := os.Args[2]

	to, err := parseInt(toArg)
	if err != nil {
		log.Fatalf("Invalid to: %v", toArg)
		printUsage()
		return
	}

	for i := from; i <= to; i++ {
		fmt.Printf("%#c", i)
	}
	fmt.Println("")
}
