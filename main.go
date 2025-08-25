package main

import (
	AsciiArt "AsciiArt/functions"
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
	if !AsciiArt.CheckFileHashing(bannerName) {
		fmt.Println("ERROR!: the file was changed!")
		return
	}

	//inputFile := Ascii.GetFileName(bannerName)

	bytes, _ := os.ReadFile(bannerName)

	strarr := strings.Split(string(bytes), "\n")

	wordSplit := AsciiArt.WordSplit(args[0])
	storeInMap := make(map[rune]int)

	ind := 1
	for i := ' '; i <= '~'; i++ {
		storeInMap[i] = ind
		ind += 9
	}

	for _, word := range wordSplit {
		if word == "\n" {
			fmt.Println()
		} else {
			for i := 0; i < 8; i++ {
				for _, v := range word {
					fmt.Printf(strarr[storeInMap[v]+i])
				}
				if i != 7 {
					fmt.Println()
				}
			}
		}
	}
	if wordSplit[len(wordSplit)-1] == "\n" {
		fmt.Println()
	} else {
		fmt.Println()
	}
}
