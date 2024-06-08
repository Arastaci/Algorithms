package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	fmt.Printf("Siralamadan once: %d\n", arr)
	insertionSort(arr)
	fmt.Printf("Siralamadan sonra: %d", arr)
}


// Insertion sort (ekleme sıralaması) algoritması, küçük veri kümelerini sıralamak için kullanılan basit ve etkili bir algoritmadır. Bu algoritma, diziyi sol alt dizinin sıralı ve sağ alt dizinin sıralanmamış olduğu iki parçaya ayırarak çalışır. Her iterasyonda, sıralanmamış bölümden bir eleman alır ve bunu sıralı bölüme yerleştirir.
// Insertion Sort algoritmasına göre bir dizinin en kötü durumu ters sıralanması durumudur. O(n^2)