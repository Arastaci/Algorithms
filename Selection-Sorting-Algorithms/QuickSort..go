package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}
	for i, val := range arr {
		if i != len(arr)/2 {
			if val < pivot {
				left = append(left, val)
			} else {
				right = append(right, val)
			}
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	fmt.Printf("Orijinal dizi: %d\n", arr)
	sortedArr := quickSort(arr)
	fmt.Printf("Siralanmis dizi: %d\n", sortedArr)
}

// Quick Sort, Divide and Conquer prensibini kullanan başka bir etkili sıralama algoritmasıdır. 
//Quick Sortun ortalama zaman karmaşıklığı O(nlogn)'dir. Eğer pivot her seferinde en küçük veya en büyük eleman olarak seçilirse, zaman karmaşıklığı O(n^2) olur. Ancak iyi bir pivot seçimi bu durumu büyük ölçüde önler.