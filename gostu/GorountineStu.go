package main

import (
	"fmt"
	"sync"
)

/**
 * 学习协程
 */

func sum(list []int64, leftIdx int64, rightIdx int64, wg *sync.WaitGroup, taskC chan int64) {
	if leftIdx == rightIdx {
		taskC <- list[leftIdx]
		wg.Done()
		return
	}
	mid := leftIdx + (rightIdx-leftIdx)/2
	leftList := list[leftIdx:mid]
	rightList := list[mid : rightIdx+1]
	fmt.Println(leftList)
	fmt.Println(rightList)

	nwg := sync.WaitGroup{}
	childTaskC := make(chan int64)
	nwg.Add(1)
	go sum(leftList, leftIdx, mid, &nwg, childTaskC)
	nwg.Add(1)
	go sum(rightList, mid, rightIdx, &nwg, childTaskC)
	nwg.Wait()

	wg.Done()
	return
}
func main() {
	var ls []int64 = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	sum(ls, 0, int64(len(ls)-1), &sync.WaitGroup{}, make(chan int64))
}
