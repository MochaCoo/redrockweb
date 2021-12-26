package main

import "fmt"

func main() {
	t:=make(chan int)
	go func() {
		fmt.Println("下山的路又堵起了")
		t<-0
	}()
	<-t
}