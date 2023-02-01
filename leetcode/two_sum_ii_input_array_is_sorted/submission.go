func twoSum(numbers []int, target int) []int {
    // cari batas atas dari angka kedua dengan b-search
    // hitung batas angka pertama dari target-batas atas
    // iterasi dari index awal, kalo batas angka pertamanya lewat, index batas atasnya dikurangin
    // ulang terus sampe nemu yang pas.
    var upperBound int
    l := 0
    r := len(numbers) - 1
    m := (r-l)/2
    for l != r{
        if numbers[m] <= target{
            l = m+1
        }else if numbers[m] > target{
            r = m
        }
        m = (r-l)/2 + l
    }

    upperBound = l
    if upperBound == 0 {
        upperBound++
    }

    i := 0
    for i < len(numbers){
        if numbers[i]+numbers[upperBound] == target{
            break
        }else if numbers[i]+numbers[upperBound] < target{
            i++
        }else {
            upperBound--
        }
    }

    retval := []int{i+1, upperBound+1}

    return retval
}