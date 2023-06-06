package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func merge(arr []int, low, mid, high int) {
	var (
		i, j, k int
		n1, n2  int
	)
	n1 = mid - low + 1
	n2 = high - mid
	l := make([]int, n1)
	r := make([]int, n2)
	for i := 0; i < n1; i++ {
		l[i] = arr[low+i]
	}
	for j := 0; j < n2; j++ {
		r[j] = arr[mid+1+j]
	}
	i = 0
	j = 0
	k = low
	for i < n1 && j < n2 {
		if l[i] <= r[j] {
			arr[k] = l[i]
			i++
		} else {
			arr[k] = r[j]
			j++
		}
		k++
	}
	for i < n1 {
		arr[k] = l[i]
		i++
		k++
	}
	for j < n2 {
		arr[k] = r[j]
		j++
		k++
	}
}

func parasort(arr []int, low, high int) {
	var wg sync.WaitGroup
	var mut sync.Mutex
	if low < high {
		mid := (low + high) / 2
		wg.Add(2)
		go func() {
			defer wg.Done()
			mut.Lock()
			parasort(arr, low, mid)
			mut.Unlock()
		}()
		go func() {
			defer wg.Done()
			mut.Lock()
			parasort(arr, mid+1, high)
			mut.Unlock()
		}()
		wg.Wait()
		merge(arr, low, mid, high)
	}
}

func main() {
	start := time.Now()
	oriList := genRdSlice(100)
	list1 := make([]int, len(oriList))
	list2 := make([]int, len(oriList))
	copy(list1, oriList)
	copy(list2, oriList)
	fmt.Printf("Init time: %fs\n", time.Since(start).Seconds())
	fmt.Println(oriList)
	start = time.Now()
	parasort(list1, 0, len(list1)-1)
	fmt.Printf("Parallel Merge Sort: %fs\n", time.Since(start).Seconds())
	fmt.Println(list1)
	start = time.Now()
	sort.Ints(list2)
	fmt.Printf("STD Merge Sort: %fs\n", time.Since(start).Seconds())
	fmt.Println(list2)
}

func genRdSlice(n int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}
