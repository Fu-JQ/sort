package main

import "fmt"

// ********************冒泡排序********************

func BubbleSort(list []int) {
	length := len(list)
	for i := 0; i < length-1; i++ {
		didSwap := false
		for j := 0; j < length-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				didSwap = true
			}
		}
		if !didSwap {
			return
		}
	}
}

// ********************选择排序********************

func SelectSort(list []int) {
	length := len(list)
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}

func SelectSort1(list []int) {
	length := len(list)
	for i := 0; i < length-1; i++ {
		min := list[i]
		index := i
		for j := i + 1; j < length; j++ {
			if list[j] < min {
				min = list[j]
				index = j
			}
		}
		if i != index {
			list[i], list[index] = list[index], list[i]
		}
	}
}

// ********************选择排序改进********************

func SelectSort2(list []int) {
	length := len(list)
	for i := 0; i < length/2; i++ {
		minIndex := i
		maxIndex := i
		for j := i + 1; j < length-i; j++ {
			if list[j] > list[maxIndex] {
				maxIndex = j
				continue
			}
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}
		if minIndex == length-i-1 && maxIndex == i {
			list[minIndex], list[maxIndex] = list[maxIndex], list[minIndex]
		} else if maxIndex == i && minIndex != length-i-1 {
			list[maxIndex], list[length-i-1] = list[length-i-1], list[maxIndex]
			list[i], list[minIndex] = list[minIndex], list[i]
		} else {
			list[i], list[minIndex] = list[minIndex], list[i]
			list[maxIndex], list[length-i-1] = list[length-i-1], list[maxIndex]
		}
	}
}

// ********************直接插入排序********************

func InsertSort(list []int) {
	length := len(list)
	for i := 1; i < length; i++ {
		num := list[i]
		j := i - 1
		if num < list[j] {
			for ; j >= 0 && num < list[j]; j-- {
				list[j+1] = list[j]
			}
			list[j+1] = num
		}
	}
}

func InsertSort1(list []int) {
	length := len(list)
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if list[j+1] < list[j] {
				list[j+1], list[j] = list[j], list[j+1]
			}
		}
	}
}

// ********************希尔排序********************

func ShellSort(list []int) {
	length := len(list)
	for step := length / 2; step >= 1; step /= 2 {
		for i := step; i < length; i += step {
			for j := i - step; j >= 0; j -= step {
				if list[j+step] < list[j] {
					list[j+step], list[j] = list[j], list[j+step]
					continue
				}
				break
			}
		}
	}
}

// ********************归并排序********************

func MergeSort(list []int, begin int, end int) {
	if end-begin > 1 {
		mid := (end-begin+1)/2 + begin

		MergeSort(list, begin, mid)
		MergeSort(list, mid, end)
		Merge(list, begin, mid, end)
	}
}

func MergeSort1(list []int, begin int, end int) {
	step := 1
	for end-begin > step {
		for i := begin; i < end; i += step << 1 {
			var lo = i
			var mid = lo + step
			var hi = lo + (step << 1)
			if mid > end {
				return
			}
			if hi > end {
				hi = end
			}
			Merge(list, lo, mid, hi)
		}
		step <<= 1
	}
}

func Merge(list []int, begin int, mid int, end int) {
	leftSize := mid - begin
	rightSize := end - mid
	newSize := leftSize + rightSize
	result := make([]int, 0, newSize)

	l, r := 0, 0
	for l < leftSize && r < rightSize {
		lValue := list[begin+l]
		rValue := list[mid+r]
		if lValue > rValue {
			result = append(result, rValue)
			r++
		} else {
			result = append(result, lValue)
			l++
		}
	}
	result = append(result, list[begin+l:mid]...)
	result = append(result, list[mid+r:end]...)

	for i := 0; i < newSize; i++ {
		list[begin+i] = result[i]
	}
	return
}

// ********************堆排序********************

type Heap struct {
	Size  int
	Array []int
}

func NewHeap(array []int) *Heap {
	h := new(Heap)
	h.Array = array
	return h
}

func (h *Heap) Push(x int) {
	if h.Size == 0 {
		h.Array[0] = x
		h.Size++
		return
	}
	i := h.Size
	for i > 0 {
		parent := (i - 1) / 2
		if x <= h.Array[parent] {
			break
		}
		h.Array[i] = h.Array[parent]
		i = parent
	}
	h.Array[i] = x
	h.Size++
}

func (h *Heap) Pop() int {
	if h.Size == 0 {
		return -1
	}
	ret := h.Array[0]
	h.Size--
	x := h.Array[h.Size]
	h.Array[h.Size] = ret
	i := 0
	for {
		lChild := 2*i + 1
		rChild := 2*i + 2
		if lChild >= h.Size {
			break
		}
		if rChild < h.Size && h.Array[rChild] > h.Array[lChild] {
			lChild = rChild
		}
		if x >= h.Array[lChild] {
			break
		}
		h.Array[i] = h.Array[lChild]
		i = lChild
	}
	h.Array[i] = x
	return ret
}


// ********************快速排序********************

func QuickSort(list []int, begin int, end int) {
	if begin < end {
		if end-begin <= 4 {
			InsertSort1(list[begin : end+1])
			return
		}
		loc := Partition(list, begin, end)
		QuickSort(list, begin, loc-1)
		QuickSort(list, loc+1, end)
	}
}

func Partition(list []int, begin int, end int) int {
	i := begin + 1
	j := end
	for i < j {
		if list[i] > list[begin] {
			list[i], list[j] = list[j], list[i]
			j--
		} else {
			i++
		}
	}
	if list[i] >= list[begin] {
		i--
	}
	list[i], list[begin] = list[begin], list[i]
	return i
}

func main() {
	list := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	QuickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
