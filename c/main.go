package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var mu sync.Mutex
	var num int
	t:=make(chan int,10)
	for i:=1;i<=10;i++ {
		go func() {
			mu.Lock()
			num++
			fmt.Println("协程"+strconv.Itoa(num))
			mu.Unlock()
			t<-0
		}()
	}

	for i:=1;i<=10;i++{
		<-t
	}
}
