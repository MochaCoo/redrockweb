package main

import (
	"fmt"
	"sync"
)
/*
* 在main启动的匿名函数协程运行mu.Lock()代码前，main函数中的mu.Unlock()可能已经被执行
*
* 此时mu并未加锁，却要被main中的mu.Unlock()给解锁，所以报错
*/
func main() {
	var mu sync.Mutex

	go func() {
		fmt.Println("有点强人锁男")
		mu.Lock()
	}()

	mu.Unlock()
}
