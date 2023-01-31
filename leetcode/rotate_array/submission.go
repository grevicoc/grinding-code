func rotate(nums []int, k int)  {
    //itung len
    //modulo len dengan k
    //bikin array kosong
    //masukin elemen tiap array kosong berdasarkan array sebelumnya
    lengthNums := len(nums)
    rotate := k % lengthNums
    // fmt.Println(rotate)
    
    var newNums []int
    i := 0
    for i < lengthNums {
        oldIndex := i - rotate
        if oldIndex < 0 {
            oldIndex += lengthNums
        }
        // fmt.Println(nums[oldIndex])
        newNums = append(newNums, nums[oldIndex])
        i++
    }

    i = 0
    for i < lengthNums {
        nums[i] = newNums[i]
        i++
    }
}   