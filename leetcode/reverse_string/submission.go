func reverseString(s []byte)  {
    length := len(s)
    r := length-1
    l := 0
    for l < length/2{
        temp := s[l]
        s[l] = s[r]
        s[r] = temp
        l++
        r--
    }
}