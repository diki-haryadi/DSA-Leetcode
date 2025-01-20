// Bruteforce Solution / Solusi Bruteforce
func summaryRangesBruteForce(nums []int) []string {
    if len(nums) == 0 {
        return []string{}
    }

    result := []string{}
    start := nums[0]

    for i := 0; i < len(nums); i++ {
        if i == len(nums)-1 || nums[i]+1 != nums[i+1] {
            if start == nums[i] {
                result = append(result, fmt.Sprintf("%d", start))
            } else {
                result = append(result, fmt.Sprintf("%d->%d", start, nums[i]))
            }

            if i != len(nums)-1 {
                start = nums[i+1]
            }
        }
    }
    return result
}

// Optimal Solution
func summaryRanges(nums []int) []string {
    result := []string{}
    i := 0

    for i < len(nums) {
        start := nums[i]

        // find the end of current range
        for i+1 < len(nums) && nums[i+1] == nums[i]+1 {
            i++
        }

        // Create range string
        if start == nums[i] {
            result = append(result, fmt.Sprintf("%d", start))
        } else {
            result = append(result, fmt.Sprintf("%d->%d", start, nums[i]))
        }
        i++
    }
    return result
}