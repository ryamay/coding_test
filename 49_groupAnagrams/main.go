package groupanagrams

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	// key: map[rune]int (key: strの文字, value: 文字が現れる回数)
	// value: string (strsの各要素)
	// のmapを作れば、anagramが等しいもののグルーピングができる

	group := map[string][]string{}

	for _, str := range strs {
		runes := []rune(str)
		slices.Sort(runes)
		sortedStr := string(runes)

		group[sortedStr] = append(group[sortedStr], str)
	}

	ans := make([][]string, 0, len(group))
	// for g := range maps.Values(group) {
	// 	ans = append(ans, g)
	// }
	for _, g := range group {
		ans = append(ans, g)
	}

	return ans
}
