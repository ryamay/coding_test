package subarraysumequalsk

// 素朴に考えたバージョン。O(n^2)の計算量。
func subarraySum_(nums []int, k int) int {
	kToArraysMap := make(map[int]int) // key: 和, value: 和となるsubArrayの数

	for i, _ := range nums {
		// i番目、i-1番目、i-2番目……を順番に足していき、その和に対応するsubArrayの数を増やしていく
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			kToArraysMap[sum]++
		}
	}

	return kToArraysMap[k] // kに対応する和となるsubArrayがない場合、ゼロ値=0を返す
}

// prefix Sumを使った改良版。O(n)の計算量。
func subarraySum(nums []int, k int) int {
	// 要素iまでの和(0, 0-1, 0-2,……):累積和を保持するmap
	// ※1要素のみのsubArrayがkに等しい場合に備えて、最初に0を追加しておく
	sumCountMap := map[int]int{0: 1}

	count := 0 // kに等しいsubArrayの個数
	sum := 0   // 要素iまでの和

	for _, n := range nums {
		sum += n

		// 累積和の間の差が、連続する要素を抜き出した場合の和となる。
		// 今回追加する和といままでの累積和の差がk となるものはカウント対象となるので、その分を増やす
		count += sumCountMap[sum-k]

		sumCountMap[sum]++ // 累積和を追加
	}

	return count
}
