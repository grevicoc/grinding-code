func searchInsert(nums []int, target int) int {
    start := 0
    end := len(nums)-1

    found := false
    answer := 0
    for !found && (start <= end) {
        middle := (end-start)/2 + start

        if (nums[middle] == target){
            found = true
            answer = middle
        } else if (nums[middle] < target){
            start = middle+1
        } else {
            end = middle-1
        }
    }
    
    if (!found){
        answer = start
    }
    return answer
}