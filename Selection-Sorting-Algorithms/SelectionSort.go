package main 
import (
	"fmt"
)
func selectionSort(arr []int){
	n:= len(arr)
	for i:=0; i < n-1; i++ {
		minIndex := i
		for j:= i+1; j < n; j++ {
			if arr[j] < arr[minIndex]{
				minIndex = j
			}
		}  
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
func main(){
	arr:= []int{1,3,5,7,9,2,4,6,8,0}
	fmt.Printf("Siralamadan once: %d\n", arr)
	selectionSort(arr)
	fmt.Printf("Siralamadan sonra: %d", arr)
	}

// Selection Sort algoritması, bir dizideki elemanları sıralamak için kullanılan bir algoritmadır. Algoritma, dizinin en küçük elemanını bulur ve onu başa alır, sonra dizinin kalanını sıralar. Bu işlem, dizinin tüm elemanlarını sıralamak için tekrar edilir.
// Selection sort algoritmasına göre bir dizinin en kötü durumu ters sıralanması durumudur. O(n^2)