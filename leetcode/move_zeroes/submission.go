func moveZeroes(nums []int)  {
    // lastZeroIdx := 0
    i := 0
    for i < len(nums){
        if nums[i] == 0 {
            j := i
            for j < len(nums) && nums[j] == 0{
                j++
            }
            
            if j == len(nums){
                break
            }
            
            nums[i] = nums[j]
            nums[j] = 0
        }
        i++
    }
}