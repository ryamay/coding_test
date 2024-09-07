package firstuniquecharacter

func firstUniqChar(s string) int {
	runeCount := make(map[rune]int) // key: 文字、 value: 出現回数

	for _, ch := range s {
		runeCount[ch]++
	}

	for i, ch := range s {
		if runeCount[ch] == 1 {
			return i
		}
	}

	return -1
}
