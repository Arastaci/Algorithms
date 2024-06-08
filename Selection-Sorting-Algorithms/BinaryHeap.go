package main

import (
    "fmt"
)

type MinHeap struct {
    arr []int
}

func NewMinHeap() *MinHeap {
    return &MinHeap{
        arr: []int{},
    }
}

func (h *MinHeap) Insert(key int) {
    h.arr = append(h.arr, key)
    h.heapifyUp(len(h.arr) - 1)
}

func (h *MinHeap) ExtractMin() int {
    if len(h.arr) == 0 {
        fmt.Println("Heap bos")
        return -1
    }
    root := h.arr[0]
    h.arr[0] = h.arr[len(h.arr)-1]
    h.arr = h.arr[:len(h.arr)-1]
    h.heapifyDown(0)
    return root
}

func (h *MinHeap) heapifyDown(i int) {
    smallest := i
    left := 2*i + 1
    right := 2*i + 2

    if left < len(h.arr) && h.arr[left] < h.arr[smallest] {
        smallest = left
    }
    if right < len(h.arr) && h.arr[right] < h.arr[smallest] {
        smallest = right
    }
    if smallest != i {
        h.arr[i], h.arr[smallest] = h.arr[smallest], h.arr[i]
        h.heapifyDown(smallest)
    }
}

func (h *MinHeap) heapifyUp(i int) {
    parent := (i - 1) / 2
    if i > 0 && h.arr[i] < h.arr[parent] {
        h.arr[i], h.arr[parent] = h.arr[parent], h.arr[i]
        h.heapifyUp(parent)
    }
}

func (h *MinHeap) Print() {
    fmt.Println(h.arr)
}

func main() {
    heap := NewMinHeap()
    heap.Insert(1)
    heap.Insert(3)
    heap.Insert(5)
    heap.Insert(7)
    heap.Insert(9)
    heap.Insert(2)
	heap.Insert(4)
	heap.Insert(6)
	heap.Insert(8)
	heap.Insert(0)
	heap.Insert(1)


    fmt.Println("Heap icerigi:")
    heap.Print()

    fmt.Printf("Min eleman: %d\n", heap.ExtractMin())
    fmt.Println("Min eleman cikarildiktan sonra heap icerigi:")
    heap.Print()
}


// Binary heap, bir tam ikili ağaç olarak organize edilmiş bir veri yapısıdır. İki tür binary heap vardır: min heap ve max heap. Min heap'te her düğüm, alt düğümlerinden daha küçük veya onlara eşit bir değere sahiptir. Max heap'te ise her düğüm, alt düğümlerinden daha büyük veya onlara eşit bir değere sahiptir. Binary heap, genellikle bir dizi veya liste kullanılarak uygulanır.
// Binary heap'in yüksekliği h, log n seviyesindedir (n, düğüm sayısıdır). Bu nedenle, ekleme işleminin zaman karmaşıklığı O(log n)'dir.