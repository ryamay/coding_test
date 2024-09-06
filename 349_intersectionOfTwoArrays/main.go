package intersectionoftwoarrays

// nums1, nums2両方に含まれる要素を重複なく返す。
// mapのKeyを使って重複が削除できることを利用して回答を作成する
func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]bool)
	for _, n1 := range nums1 {
		m1[n1] = true
	}
	// m1のKeysはnums1から重複を削除した一覧になる

	mAns := make(map[int]bool) //このmapのKeysを回答とする
	for _, n2 := range nums2 {
		if _, ok := m1[n2]; ok {
			mAns[n2] = true // n2が重複排除したnums1に含まれるので、回答となる
		}
	}

	ans := make([]int, 0, len(mAns))
	// for n := range maps.Keys(mAns) {
	// 	ans = append(ans, n)
	// }
	for n, _ := range mAns {
		ans = append(ans, n)
	}

	return ans
}
