package main

import (
	"fmt"
)

func knapsack(W int, wt, val []int, n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}
	for i := 0; i <= n; i++ {
		for w := 0; w <= W; w++ {
			if i == 0 || w == 0 {
				dp[i][w] = 0
			} else if wt[i-1] <= w {
				dp[i][w] = max(val[i-1]+dp[i-1][w-wt[i-1]], dp[i-1][w])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	return dp[n][W]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	val := []int{50, 75, 100, 125, 150, 175}
	wt := []int{10, 20, 30, 40, 50, 60}
	W := 50
	n := len(val)
	fmt.Printf("Sırt çantasındaki maksimum değer = %d", knapsack(W, wt, val, n))
}

// Knapsack problemi, verilen ağırlık ve değerlerle belirli bir kapasiteye sahip bir sırt çantasını en yüksek değeri elde edecek şekilde doldurma problemidir.
// Dinamik Programlama Yaklaşımı ile Zaman Karmaşıklığı: O(nW), n = öğe sayısı, W = sırt çantasının kapasitesi
