import "sort"
func longestConsecutiveBruteforce(nums []int) int {
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

// Optimal Solution using Hash Set  / Solusi Optimal menggunakan Hash Set
func longestConsecutive(nums []int) int {
    // Create a hash set / Buat hash set
    numSet := make(map[int]bool)
    for _, num := range nums {
        numSet[num] = true
    }

    longestStreak := 0

    // Check each number / Periksa setiap angka
    for num := range numSet {
        // Only start counting if it's the start of a sequence
        // Hanya mulai menghitung jika ini adalah awal dari urutan
        if !numSet[num-1] {
            currentNum := num
            currentStreak := 1

            // Count consecutive numbers / Hitung angka berurutan
            for numSet[currentNum+1] {
                currentNum++
                currentStreak++
            }

            if currentStreak > longestStreak {
                longestStreak = currentStreak
            }
        }
    }
    return longestStreak
}