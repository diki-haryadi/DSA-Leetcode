func stringMatching(words []string) []string {
    result := make([]string, 0)
    
    // Sort words by length to process shorter strings first
    sort.Slice(words, func(i, j int) bool {
        return len(words[i]) < len(words[j])
    })
    
    // For each word, check if it's a substring of other words
    for i := 0; i < len(words); i++ {
        for j := i + 1; j < len(words); j++ {
            if strings.Contains(words[j], words[i]) {
                result = append(result, words[i])
                break  // Break once we find a match for current word
            }
        }
    }
    
    return result
}