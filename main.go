package main

import (
	//"Ascii/functions"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)

	if len(args) != 1 {
		fmt.Println("Please enter only one argument")
		return
	}

	for _, ch := range args[0] {
		if (ch < 32 || ch > 126) && ch != 10 {
			fmt.Println("Please use valid symbols to print")
			return
		}
	}
	bannerName := "standard.txt"
	//inputFile := Ascii.GetFileName(bannerName)

	bytes, _ := os.ReadFile(bannerName)

	strarr := strings.Split(string(bytes), "\n")
	fmt.Println(strarr)
}
