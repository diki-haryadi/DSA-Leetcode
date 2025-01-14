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