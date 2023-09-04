package main

import (
	"fmt"
	"io"
)

func outputRunes(rd io.RuneReader) error {
	for {
		r, _, err := rd.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("%#c", r)
	}
	fmt.Println("")
	return nil
}

func outputCodePoints(rd io.RuneReader) error {
	sep := ""
	for {
		r, _, err := rd.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("%s%d", sep, r)
		sep = " "
	}
	fmt.Println("")
	return nil
}
