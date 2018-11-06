package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func printUsage() {
	fmt.Println("Please input `from` `to` as arguments. Usage: runa 33 126")
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
