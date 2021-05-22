package leetcode

import (
	"sort"
	"strings"
)

func sortString(str string) string {
	letters := strings.Split(str, "")
	sort.Strings(letters)

	return strings.Join(letters, "")
}

// LC 49
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	anagramMap := map[string][]string{}

	for _, str := range strs {
		sortedString := sortString(str)
		anagramMap[sortedString] = append(anagramMap[sortedString], str)
	}

	for _, v := range anagramMap {
		res = append(res, v)
	}

	return res
}
