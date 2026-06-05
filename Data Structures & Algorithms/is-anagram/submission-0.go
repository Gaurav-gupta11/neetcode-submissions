func isAnagram(s string, t string) bool {
    if len(s)!= len(t) {
        return false
    }

    freq := make(map[rune]int)

    for _,ch := range s {
        freq[ch]++
    }
    for _,ch := range t {
        freq[ch]--
    }
    for _, count := range freq {
        if count != 0 {
        return false
        }
    }
    return true
}
