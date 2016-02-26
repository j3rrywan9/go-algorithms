package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const words_path = "/usr/share/dict/words"

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false;
	}

	ss1 := strings.Split(s1, "")
	sort.Strings(ss1)

	ss2 := strings.Split(s2, "")
	sort.Strings(ss2)

	return strings.Join(ss1, "") == strings.Join(ss2, "")
}

func buildAnagramMap() map[string][]string {
	anagramMap := make(map[string][]string)

	words, _ := os.Open(words_path)
	defer words.Close()

	scanner := bufio.NewScanner(words)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		word := scanner.Text()

		chars_in_word := strings.Split(word, "")
		sort.Strings(chars_in_word)
		sorted_word := strings.Join(chars_in_word, "")

		if _, ok := anagramMap[sorted_word]; ok == false {
			anagramMap[sorted_word] = []string{}
		}
		anagramMap[sorted_word] = append(anagramMap[sorted_word], word)
	}

	return anagramMap
}

func getAnagrams(word string, anagramMap map[string][]string) []string {
	chars_in_word := strings.Split(word, "")
	sort.Strings(chars_in_word)
	sorted_word := strings.Join(chars_in_word, "")

	if anagramList, ok := anagramMap[sorted_word]; ok {
		return anagramList
	} else {
		return []string{}
	}
}

func main() {
	var word string = "care"

	anagramMap := buildAnagramMap()
	anagramList := getAnagrams(word, anagramMap)
	fmt.Printf("Word: %s\nAnagrams: %v\n", word, anagramList)
}
