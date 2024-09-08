package wordladder

import "slices"

// queueを使った幅優先探索を使う。
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if !slices.Contains(wordList, endWord) {
		return 0
	}

	queue := queue{}                 //幅優先探索用につかうキュー
	wordMap := make(map[string]bool) // 存在確認用に使うmap
	for _, w := range wordList {
		wordMap[w] = true
	}

	// beginWordからスタート
	queue.push(beginWord)
	phase := 0

	for len(queue) > 0 {
		phase++
		isFoundEndWord := queue.processSinglePhase(endWord, wordMap)
		if isFoundEndWord {
			ladderLength := phase + 1
			return ladderLength
		}
	}

	// キューをすべて使い切ってもendWordに至らなかった場合、経路が存在しないため0を返す
	return 0
}

type queue []string

func (q *queue) push(s string) {
	*q = append(*q, s)
}

func (q *queue) pop() string {
	// 先頭をqから除いて返す
	s := (*q)[0]
	*q = (*q)[1:]
	return s
}

func (q *queue) processSinglePhase(endWord string, wordMap map[string]bool) (isFoundEndWord bool) {
	// この時点でキューに存在する要素分だけ、処理を行う
	for range *q {
		current := q.pop()

		// 現在の文字列の先頭から、1文字ずつ別のアルファベットに置き換えて、wordListに含まれるかチェックする。
		// wordListに含まれる場合、そのwordをqueueにpushする
		// queueにpushしたwordは、wordMapから削除して次回以降の検索対象外とする
		for i := range current {
			for ch := 'a'; ch <= 'z'; ch++ {
				nextWord := current[:i] + string(ch) + current[i+1:]
				if _, ok := wordMap[nextWord]; ok {
					if nextWord == endWord {
						return true
					}
					q.push(nextWord)
					delete(wordMap, nextWord)
				}
			}
		}
	}
	return false
}
