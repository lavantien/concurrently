package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func merge_(arr []int, low, mid, high int) {
	var (
		i, j, k int
		n1, n2  int
	)
	n1 = mid - low + 1
	n2 = high - mid
	L := make([]int, n1)
	R := make([]int, n2)
	for i := 0; i < n1; i++ {
		L[i] = arr[low+i]
	}
	for j := 0; j < n2; j++ {
		R[j] = arr[mid+1+j]
	}
	i = 0
	j = 0
	k = low
	for i < n1 && j < n2 {
		if L[i] <= R[j] {
			arr[k] = L[i]
			i++
		} else {
			arr[k] = R[j]
			j++
		}
		k++
	}
	for i < n1 {
		arr[k] = L[i]
		i++
		k++
	}
	for j < n2 {
		arr[k] = R[j]
		j++
		k++
	}
}

func parasort_(arr []int, low, high int) {
	var wg sync.WaitGroup
	ch := make(chan int)
	if low < high {
		mid := (low + high) / 2
		wg.Add(2)
		go func() {
			defer wg.Done()
			parasort_(arr, low, mid)
			ch <- 1
		}()
		go func() {
			defer wg.Done()
			parasort_(arr, mid+1, high)
			ch <- 1
		}()
		go func() {
			defer wg.Done()
			merge_(arr, low, mid, high)
		}()
		for i := 0; i < 2; i++ {
			<-ch
		}
	}
}

func main_() {
	start := time.Now()
	oriList := genRdSlice(100_000_000)
	list1 := make([]int, len(oriList))
	list2 := make([]int, len(oriList))
	copy(list1, oriList)
	copy(list2, oriList)
	fmt.Printf("Init time: %fs\n", time.Since(start).Seconds())
	start = time.Now()
	parasort(list1, 0, len(list1)-1)
	fmt.Printf("Parallel Merge Sort: %fs\n", time.Since(start).Seconds())
	start = time.Now()
	sort.Ints(list2)
	fmt.Printf("STD Merge Sort: %fs\n", time.Since(start).Seconds())
}

func genRdSlice(n int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}
