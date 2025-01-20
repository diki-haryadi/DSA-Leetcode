import "sort"
// Bruteforce solution
func containsDuplicateBruteforce(nums []int) bool {
    // Check each number with all other numbers
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            if nums[i]==nums[j] {
                return true
            }
        }
    }
    return false
}

// Sorting solution
func containsSuplicateSort(nums []int) bool {
    sort.Ints(nums)

    for i := 1; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            return true
        }
    }
    return false
}

// Optimal Solution using Hash Set
func containsDuplicate(nums []int) bool {
    seen := make(map[int]bool)

    for _, num := range nums {
        if seen[num] {
            return true
        }
        seen[num]= true
    }
    return false
}