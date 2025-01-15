func maxProfitBruteForce(prices []int) int {
    maxProfit := 0

    // Loop through each possible buying day
    for i := 0; i < len(prices); i++ {
        for j := i + 1; j < len(prices); j++ {
            profit := prices[j] - prices[i]
            if profit > maxProfit {
                maxProfit = profit
            }
        }
    }
    return maxProfit
}

// Optimal solution
func maxProfit(prices []int) int {
    if len(prices) < 2 {
        return 0
    }

    minPrice := prices[0]
    maxProfit := 0

    for i := 1; i < len(prices); i++ {
        if prices[i] < minPrice {
            minPrice = prices[i]
        } else {
            currentProfit := prices[i] - minPrice 
            if currentProfit > maxProfit {
                maxProfit = currentProfit
            }
        }
    }
    return maxProfit
}