package main

import (
	"fmt"
)

func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	target := 8
	result := linearSearch(arr, target)
	if result != -1 {
		fmt.Printf("Eleman bulundu: %d\n", result)
	} else {
		fmt.Printf("Eleman bulunamadi.\n ")
	}
}

//  Linear search (doğrusal arama) algoritması, bir dizide bir öğeyi aramak için kullanılan basit bir algoritmadır. Bu algoritma, aranan öğeyi bulana veya dizinin sonuna ulaşana kadar dizinin her bir elemanını sırayla kontrol eder.
//En kötü durumda O(n) zaman karmaşıklığına sahiptir çünkü aranan eleman dizinin sonunda veya dizide bulunmayabilir ve bu durumda bütün diziyi kontrol etmek zorunda kalır. Ancak basit ve küçük veri kümeleri için oldukça etkilidir
