import "sort"
func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    // Sort the array first / Urutkan array terlebih dahulu
    sort.Ints(nums)
    
    currentStreak := 1
    longestStreak := 1
    
    // Iterate through sorted array / Iterasi melalui array yang sudah diurutkan
    for i := 1; i < len(nums); i++ {
        // Skip duplicates / Lewati angka yang sama
        if nums[i] == nums[i-1] {
            continue
        }
        
        // If consecutive / Jika berurutan
        if nums[i] == nums[i-1] + 1 {
            currentStreak++
            if currentStreak > longestStreak {
                longestStreak = currentStreak
            }
        } else {
            currentStreak = 1
        }
    }
    
    return longestStreak
}

// Optimal
func longestConsecutiveOptimal(nums []int) int {
return 0
}