func isAnagram(s string, t string) bool {
    // Check if lengths are different
    if len(s) != len(t) {
        return false
    }
    
    // Create a map to store character counts
    charCount := make(map[rune]int)
    
    // Count characters in first string
    for _, char := range s {
        charCount[char]++
    }
    
    // Decrease count for second string
    for _, char := range t {
        charCount[char]--
        // If count becomes negative, strings aren't anagrams
        if charCount[char] < 0 {
            return false
        }
    }
    
    // Check if all counts are zero
    for _, count := range charCount {
        if count != 0 {
            return false
        }
    }
    
    return true
}

// Using an array (more efficient for ASCII)
func isAnagram2(s string, t string) bool {
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