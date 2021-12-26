package main

import (
	"fmt"
)

var t chan int

func main(){
	t=make(chan int)
	const num =20
	const thread =10
	size:=num/thread

	for i:=1;i<=thread;i++{
		//fmt.Printf("%d %d\n",(i-1)*size,i*size)
		go out((i-1)*size+1,i*size)
	}

	for i:=1;i<=thread;i++{
		<-t
	}
}
func out(start,end int){
	for i:=start;i<=end;i++{
		if isPrime(i){
			fmt.Printf("%d\n",i)
		}
	}
	t<-0
}
func isPrime(v int) bool{
	for i:=2;i<=v/2;i++{
		if v%i==0{
			return false
		}
	}
	return true
}
