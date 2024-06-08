package main
import(
	"fmt"
)

func merge(left, right []int)[]int{
	result := []int{}
	i, j := 0,0
	for i < len(left) && j < len(right){
		if left[i] < right[j]{
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

result = append(result, left[i:]...)
result = append(result, right[j:]...)
return result
}

func mergeSort(arr []int)[]int {
	if len(arr) <= 1{
		return arr
	}
	mid := len(arr) / 2 
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right) 
}
func main(){
	arr:= []int{1,3,5,7,9,2,4,6,8,0}
	fmt.Printf("Orijinal dizi: %d\n", arr)
	sortedArr:= mergeSort(arr)
	fmt.Printf("Siralanmis dizi: %d\n", sortedArr)
}


// Merge Sort, Divide and Conquer prensibini kullanan bir sıralama algoritmasıdır.
// Divide: Diziyi ortadan ikiye bölerek iki alt dizi oluşturur. Dizi her seferinde ikiye bölündüğü için O(log n).
// Conquer: Alt diziler sıralanana kadar Merge Sort algoritması rekürsif olarak uygulanır. Her alt dizi birleştirildiğinde tüm elemanlar üzerinden geçildiği için O(n).
// Combine: Sıralı alt diziler birleştirilerek tek bir sıralı dizi oluşturulur. Tüm elemanların birleşimi yine O(n) zaman alır.
// Toplam zaman karmaşıklığı O(n log n)'dir.