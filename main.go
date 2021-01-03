package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
var count int

var mutex sync.RWMutex

func main(){

	// 1.声明互斥锁
	//var mutex sync.Mutex
	//count := 0
	//for r := 0; r < 20; r++ {
	//	go func() {
	//		mutex.Lock()
	//		count+=1
	//		defer mutex.Unlock()
	//	}()
	//	time.Sleep(time.Second)
	//	fmt.Println("the count is : ", count)
	//}


	for i := 0; i < 10; i++ {
		go read(i + 1)
	}

	for i := 0; i < 10; i++ {
		go write(i + 1)
	}
	time.Sleep(time.Second*5)
}

func write(n int) {
	rand.Seed(time.Now().UnixNano())
	//fmt.Errorf("写 goroutine %d 正在写数据...\n", n)
	mutex.Lock()
	num := rand.Intn(500)
	count = num
	fmt.Printf("写 goroutine %d 写数据结束，写入新值 %d\n", n, num)
	mutex.Unlock()

}
func read(n int) {
	mutex.RLock()
	//fmt.Printf("读 goroutine %d 正在读取数据...\n", n)
	num := count
	fmt.Printf("读 goroutine %d 读取数据结束，读到 %d\n", n, num)
	mutex.RUnlock()
}
