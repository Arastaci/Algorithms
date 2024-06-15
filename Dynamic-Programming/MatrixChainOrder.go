package main

import (
	"fmt"
	"math"
)

func MatrixChainOrder(p []int) ([][]int, [][]int) {
	n := len(p) - 1
	m := make([][]int, n)
	s := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
		s[i] = make([]int, n)
	}

	for l := 2; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			m[i][j] = math.MaxInt32
			for k := i; k < j; k++ {
				q := m[i][k] + m[k+1][j] + p[i]*p[k+1]*p[j+1]
				if q < m[i][j] {
					m[i][j] = q
					s[i][j] = k
				}
			}
		}
	}
	return m, s
}

func PrintOptimalParens(s [][]int, i, j int) {
	if i == j {
		fmt.Printf("A%d", i+1)
	} else {
		fmt.Print("(")
		PrintOptimalParens(s, i, s[i][j])
		PrintOptimalParens(s, s[i][j]+1, j)
		fmt.Print(")")
	}
}

func main() {
	p := []int{10, 20, 5, 30, 45, 40, 15}
	m, s := MatrixChainOrder(p)

	fmt.Println("Minimum number of multiplications:", m[0][len(p)-2])
	fmt.Print("Optimal Parenthesization: ")
	PrintOptimalParens(s, 0, len(p)-2)
	fmt.Println()
}

// Matris zincir çarpma, bilgisayar bilimlerinde ve matematikte önemli bir optimizasyon problemidir. Bu problem, belirli bir matris çarpma sırasını bulup toplam çarpma işlemi sayısını minimize etmeyi amaçlar. Matris çarpma işlemi, büyük hesaplama yükü gerektirdiğinden, matris zincir çarpma algoritması performansı artırmak için kullanılır.
// Matris zincir çarpma algoritmasının zaman karmaşıklığı O(n3) olarak ifade edilir. İç içe üç tane for döngüsüsü içermesi nedeniyle ortaya çıkar.
