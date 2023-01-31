func sortedSquares(nums []int) []int {
    lengthNums := len(nums)
    if lengthNums == 0 {
        return nums
    }
    if lengthNums == 1 {
        var newNums []int
        newNums = append(newNums, int(math.Pow(float64(nums[0]), 2)))
        return newNums
    }

    lowest := 0
    for (lowest < lengthNums-1) && (math.Abs(float64(nums[lowest])) >= math.Abs(float64(nums[lowest+1]))) {
        lowest += 1
    }

    left := lowest-1
    right := lowest+1
    var answer []int

    answer = append(answer, int(math.Pow(float64(nums[lowest]), 2)))
    for left >= 0 && right <= len(nums)-1 {
        leftSquare := int(math.Pow(float64(nums[left]), 2))
        rightSquare := int(math.Pow(float64(nums[right]), 2))
        if leftSquare <= rightSquare {
            answer = append(answer, leftSquare)
            left -= 1
        }else {
            answer = append(answer, rightSquare)
            right += 1
        }
    }

    if left >= 0 {
        for left >= 0 {
            leftSquare := int(math.Pow(float64(nums[left]), 2))
            answer = append(answer, leftSquare)
            left -= 1
        }
    }else {
        for right <= len(nums)-1{
            rightSquare := int(math.Pow(float64(nums[right]), 2))
            answer = append(answer, rightSquare)
            right += 1
        }
    }

    return answer
}