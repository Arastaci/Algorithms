package main

import (
	"fmt"
)

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	fmt.Printf("Unsorted array: %d\n", arr)
	bubbleSort(arr)
	fmt.Printf("Sorted array: %d\n", arr)
}

// Bubble sort algoritması adım adım iki bitişik öğeyi karşılaştırarak ve gerektiğinde yerlerini değiştirerek çalışır. Bu işlemi, dizinin sıralı olduğunu kontrol edene kadar tekrarlar.
// Bubble Sort algoritmasına göre en iyi durum dizinin sıralı olduğu durumdur ve zaman karmaşıklığı n'dir
// En kötü ve ortalama durum O(n^2)
