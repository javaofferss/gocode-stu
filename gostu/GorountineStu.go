package main

import (
	"fmt"
	"sync"
)

/**
 * 学习协程
 */

func sum(list []int64, wg *sync.WaitGroup, taskC chan int64) {

	if len(list) == 0 {
		taskC <- 0
		wg.Done()
		return
	} else if len(list) == 1 {
		taskC <- list[0]
		wg.Done()
		return
	}

	leftList := list[0 : len(list)/2]
	rightList := list[len(list)/2:]
	fmt.Println(leftList)
	fmt.Println(rightList)

	nwg := sync.WaitGroup{}
	childTaskC := make(chan int64, 2)
	nwg.Add(1)
	nwg.Add(1)

	go sum(leftList, &nwg, childTaskC)
	go sum(rightList, &nwg, childTaskC)
	nwg.Wait()

	wg.Done()

	var a = <-childTaskC
	var b = <-childTaskC

	c := a + b

	taskC <- c
	return
}
func main() {
	var ls []int64 = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	nwg := sync.WaitGroup{}
	nwg.Add(1)
	int64s := make(chan int64, 1)
	sum(ls, &nwg, int64s)
	nwg.Wait()
	var result = <-int64s
	fmt.Println(result)
}
