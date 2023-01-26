func search(nums []int, target int) int {
    //cari index tengah
    //cek angka tersebut apakh sesuai target or not
    //kalo sesuai ya itu jawbaannya. Kalo ngga cek angka tersebut lebih kecil or lebih besar
    // kalo lebih besar ke kanan kl lebih kecil ke kiri
    start := 0
    end := len(nums)-1

    answer := -1
    for {
        middle := ((end-start)/2) + start

        if (nums[middle]==target){
            answer = middle
        }else if (nums[middle] < target){
            start = middle+1
        }else{
            end = middle-1
        }
        if (answer != -1 || (start > end || end < start)){
            break
        }

    }
    
    return answer
}