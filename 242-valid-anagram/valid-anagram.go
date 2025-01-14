func isAnagram(s string, t string) bool {
    // Check if lengths are different
    if len(s) != len(t) {
        return false
    }
    
    // Create an array for character counts (26 for lowercase letters)
    counts := [26]int{}
    
    // Count characters in both strings
    for i := 0; i < len(s); i++ {
        counts[s[i]-'a']++
        counts[t[i]-'a']--
    }
    
    // Check if all counts are zero
    for _, count := range counts {
        if count != 0 {
            return false
        }
    }
    
    return true
}