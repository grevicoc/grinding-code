func lengthOfLongestSubstring(s string) int {
    maxNumber := 0
    currMaxNumber := 0
    mapCurrChars := make(map[byte]int)
    i := 0
    for i < len(s){
        val, ok := mapCurrChars[s[i]]
        if ok{
            if currMaxNumber > maxNumber{
                maxNumber = currMaxNumber
            }
            currMaxNumber = 0
            i = val + 1
            mapCurrChars = make(map[byte]int)
        }else{
            currMaxNumber++
            mapCurrChars[s[i]] = i
            i++
        } 
    }

    if currMaxNumber > maxNumber{
        maxNumber = currMaxNumber
    }

    return maxNumber
}