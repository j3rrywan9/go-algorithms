package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const words_path = "/usr/share/dict/words"

func sortString(word string) string {
	letters := strings.Split(word, "")
	sort.Strings(letters)
	return strings.Join(letters, "")
}

func buildAnagramMap() map[string][]string {
	anagramMap := make(map[string][]string)

	words, _ := os.Open(words_path)
	defer words.Close()

	scanner := bufio.NewScanner(words)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		word := scanner.Text()
		sorted_word := sortString(word)

		if _, ok := anagramMap[sorted_word]; ok == false {
			anagramMap[sorted_word] = []string{}
		}
		anagramMap[sorted_word] = append(anagramMap[sorted_word], word)
	}

	return anagramMap
}

func main() {
	var word string

	anagramMap := buildAnagramMap()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		word = strings.TrimSpace(scanner.Text())

		if anagramList, ok := anagramMap[sortString(word)]; ok {
			fmt.Printf("%v\n", anagramList)
		} else {
			fmt.Println("No matches")
		}
	}
}
