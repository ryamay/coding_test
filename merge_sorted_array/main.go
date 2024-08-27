package mergesortedarray

import "log"

const (
	// 制約条件
	_M_MIN              = 0
	_M_MAX              = 200
	_N_MIN              = 0
	_N_MAX              = 200
	_SUM_OF_M_AND_N_MIN = 1
	_SUM_OF_M_AND_N_MAX = 200
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if !matchPrecondition(nums1, m, nums2, n) {
		return // 制約条件違反のため何もしない
	}
	//note: 自分で作成した回答は先頭から比較しているが、最善の回答は末尾から比較してnon-zeroのnums1要素を設定していくもの。
	copy(nums1, mergeSub(nums1, m, nums2, n))
}

// 再帰関数
// nums1, nums2が昇順ソート済みの配列のため、先頭の要素を比較して小さい方を先頭にする。
func mergeSub(nums1 []int, m int, nums2 []int, n int) (result []int) {
	if n == 0 { // n=0の場合、nums1がそのまま答え
		return nums1
	}
	if m == 0 { // m=0の場合、nums2がそのまま答え
		return nums2
	}

	// 昇順ソート済みのため、先頭を順番に比較して答えをつくる
	if nums1[0] <= nums2[0] {
		return append([]int{nums1[0]}, mergeSub(nums1[1:], m-1, nums2, n)...)
	} else {
		return append([]int{nums2[0]}, mergeSub(nums1, m, nums2[1:], n-1)...)
	}
}

func matchPrecondition(nums1 []int, m int, nums2 []int, n int) bool {
	if m+n != len(nums1) {
		log.Fatal("m+n = len(nums1)に違反")
		return false // 制約条件違反: m+n = len(nums1)
	}
	if n != len(nums2) {
		log.Fatal("n = len(nums2)に違反")
		return false // 制約条件違反: n = len(nums2)
	}
	if !(_M_MIN <= m && m <= _M_MAX && _N_MIN <= n && n <= _N_MAX) {
		log.Fatal("0 <= m, n <= 200に違反")
		return false // 0 <= m, n <= 200
	}
	if !(_SUM_OF_M_AND_N_MIN <= m+n && m+n <= _SUM_OF_M_AND_N_MAX) {
		log.Fatal("1 <= m + n <= 200")
		return false // 1 <= m + n <= 200
	}
	// -10^9 <= nums1[i], nums2[j] <= 10^9 の制約条件チェックは省略する
	// （2^31-1より10^9が小さい = 32ビット符号付き整数で表現可能のため）
	return true
}
