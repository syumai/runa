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
> 97
`)
}

func parseInt(s string) (int64, error) {
	if strings.HasPrefix(s, "0x") {
		return strconv.ParseInt(s, 0, 64)
	}
	return strconv.ParseInt(s, 10, 64)
}

func main() {
	if len(os.Args) < 3 {
		printUsage()
		return
	}

	fromArg, toArg := os.Args[1], os.Args[2]

	if fromArg == "-c" {
		fmt.Println([]rune(toArg)[0])
		return
	}

	from, err := parseInt(fromArg)
	if err != nil {
		log.Fatalf("Invalid from: %v", fromArg)
		printUsage()
		return
	}

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
