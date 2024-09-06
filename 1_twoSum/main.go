package twosum

// 総当たりで確認することもできるが、
// 確認対象の数字に対応するindexを保持するMapを作成して、numsを1loopするだけで解がわかるようにしている
func twoSum(nums []int, target int) []int {
	ansMap := map[int]int{} // key: 足す数, value: 回答となる数字へのindex

	for i, num := range nums {
		idx, ok := ansMap[num] // mapにペアとなる数字が存在するか確認

		if ok {
			return []int{i, idx}
		} else {
			key := target - num // 確認対象の数字とペアになる数字を計算する
			ansMap[key] = i     // 存在しない場合は、足す数->現在のindexを追加
		}
	}

	return []int{} // 必ず答えがあるとの前提なので、ここには到達しない
}
