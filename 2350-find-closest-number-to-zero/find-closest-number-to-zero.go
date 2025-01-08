func findClosestNumber(nums []int) int {
    closestNum := nums[0]
    minDistance := abs(nums[0])

    for _, num := range nums {
        distance := abs(num)

        if distance < minDistance {
            closestNum = num
            minDistance = distance
        }else if distance == minDistance && num > closestNum {
            closestNum = num
        }
    }
    return closestNum
}

func abs(x int) int{
    if x < 0 {
        return -x
    }
    return x
}