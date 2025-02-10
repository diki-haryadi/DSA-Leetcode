func findKthPositive(arr []int, k int) int {
  j := 1
  for i := 0; i<len(arr); j++ {
    if arr[i] != j {
        k--
    } else {
        i++
    }
    if k == 0 {
        return j
    }
  }
  return j + k -1
}