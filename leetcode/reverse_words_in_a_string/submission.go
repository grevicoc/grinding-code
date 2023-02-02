func reverseWords(s string) string {
    var retval string
    length := len(s)
    start := 0
    for start < length{
        end := start
        for end < length && s[end] != ' '{
            end++
        }

        retval += reverseString(s[start:end])
        start = end+1
        if start < length{
            retval += " "
        }
    }
    return retval
}

func reverseString(s string) string {
    tempBytes := []byte(s)
    length := len(tempBytes)
    r := length-1
    l := 0
    for l < length/2{
        temp := tempBytes[l]
        tempBytes[l] = tempBytes[r]
        tempBytes[r] = temp
        l++
        r--
    }
    return string(tempBytes)
}