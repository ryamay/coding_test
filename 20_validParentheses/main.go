package validparentheses

var parens = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := []rune{}
	for _, r := range s {
		if _, ok := parens[r]; ok {
			stack = append(stack, r) // 左カッコなら、stackに入れる
		} else if isMatchParen(stack, r) {
			// stack先頭の要素と対象の文字が有効なカッコのペアであれば、
			// stackからpopできたとして末尾の要素を削除する
			stack = stack[:len(stack)-1]
		} else {
			// 有効なカッコのペアでなかったため、falseを返す
			return false
		}
	}

	// stackに何も残っていなければ、文字列全体として有効な括弧文字列だった
	return len(stack) == 0
}

func isMatchParen(stack []rune, r rune) bool {
	if len(stack) == 0 {
		return false // stackが空
	}

	popped := stack[len(stack)-1]
	// 左括弧なら、parensから対応する右括弧を取得できる
	rightParen := parens[popped]

	// 対応する右括弧文字と、チェック対象の文字が等しければカッコの対応がついている
	return rightParen == r
}
