package main

import (
	Ascii "Ascii/functions"
	//"Ascii/functions"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

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

	wordSplit := Ascii.WordSplit(args[0])
	fmt.Println(wordSplit)
	storeInMap := make(map[rune]int)

	ind := 1
	for i := ' '; i <= '~'; i++ {
		storeInMap[i] = ind
		ind += 9
	}

	for i := 0; i < 8; i++ {
		for _, v := range args[0] {
			fmt.Printf(strarr[storeInMap[v]+i])
		}
		if i != 7 {
			fmt.Println()
		}
	}
	fmt.Println()
}
