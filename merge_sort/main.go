package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"sync"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func concurrent_sort(arr []int) []int {
	min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	threads := previousPowerOfTwo(runtime.NumCPU())
	parts := make([][]int, threads)
	increment := len(arr) / threads
	for i := 0; i < len(arr); i += increment {
		parts[i/increment] = arr[i:min(len(arr), i+increment)]
	}
	var wg sync.WaitGroup
	for _, part := range parts {
		wg.Add(1)
		go func(part []int) {
			sort.Ints(part)
			wg.Done()
		}(part)
	}
	wg.Wait()
	var merged []int
	for _, part := range parts {
		merged = merge(merged, part)
	}
	return merged
}

func previousPowerOfTwo(x int) int {
	a := uint64(x)
	if a < 3 {
		return int(a)
	}
	a |= a >> 1
	a |= a >> 2
	a |= a >> 4
	a |= a >> 8
	a |= a >> 16
	a |= a >> 32
	return int(a - (a >> 1))
}

func merge(a, b []int) []int {
	merged := make([]int, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			merged[i+j] = a[i]
			i++
		} else {
			merged[i+j] = b[j]
			j++
		}
	}
	for i < len(a) {
		merged[i+j] = a[i]
		i++
	}
	for j < len(b) {
		merged[i+j] = b[j]
		j++
	}
	return merged
}

func main() {
	start := time.Now()
	fmt.Println("Init records ...")
	records := 500_000_000
	lista := genRdSlice(records)
	listb := make([]int, len(lista))
	copy(listb, lista)
	prt := message.NewPrinter(language.English)
	prt.Printf("Number of records: %d; ", records)
	fmt.Printf("Init time: %fs\n", time.Since(start).Seconds())
	// fmt.Println(oriList)
	start = time.Now()
	fmt.Println("\nRunning STD Merge Sort ...")
	sort.Ints(lista)
	stamp1 := time.Since(start).Seconds()
	fmt.Printf("STD Merge Sort: %fs\n", stamp1)
	// fmt.Println(list1)
	start = time.Now()
	fmt.Printf("\nRunning Concurrent Multiway Merge Sort; Threads use: %d/%d ...\n", previousPowerOfTwo(runtime.NumCPU()), runtime.NumCPU())
	concurrent_sort(listb)
	stamp2 := time.Since(start).Seconds()
	fmt.Printf("Concurrent Mulitway Merge Sort: %fs\n", stamp2)
	if stamp1 < stamp2 {
		fmt.Printf("\nSlower: %.1fx times\n", stamp2/stamp1)
	} else {
		fmt.Printf("\nFaster: %.1fx times\n", stamp1/stamp2)
	}
	// fmt.Println(list2)
}

func genRdSlice(n int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(n)
	}
	return arr
}
