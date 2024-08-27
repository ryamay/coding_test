package removeelement

// - valと等しい要素を探し、nums上のindexを記録する
// -
func removeElement(nums []int, val int) int {
	valIndexes := []int{}
	for i, n := range nums {
		if n == val {
			valIndexes = append(valIndexes, i)
		}
	}

	// numsは、nonValsCount=len(nums) - len(valIndexes)個まで、valと等しくない要素に詰め替える必要がある.
	// numsのうち、nonValsCountを超えるindexの要素を対象に、
	// valと等しくないものを探して、valIndexesに保持した場所へ入れていく。
	nonValsCount := len(nums) - len(valIndexes)
	valIndexesAt := 0 // valと等しくない要素を配置する先
	for i := len(nums) - 1; i >= len(nums)-len(valIndexes); i-- {
		if nums[i] == val {
			continue
		}

		nums[valIndexes[valIndexesAt]] = nums[i]
		valIndexesAt++
	}

	return nonValsCount
}
