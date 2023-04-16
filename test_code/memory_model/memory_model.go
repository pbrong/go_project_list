package memorymodel

import (
	"fmt"
	"runtime"
)

func Run() {
	// test1()
	test2()
}

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func test1() {
	go setup()
	for !done {
	}
	fmt.Printf("a = %v", a)
}

func test2() {
	// 分配100个int类型的空间
	nums := make([]int, 100)
	fmt.Println("Before:", runtime.NumGoroutine())

	// 开启100个goroutine，并将nums传递给它们
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println("In Goroutine:", runtime.NumGoroutine())
			fmt.Println(nums)
		}()
	}

	// 主goroutine睡眠5s
	runtime.Gosched()

	// 手动释放nums占用的内存
	nums = nil
	fmt.Println("After:", runtime.NumGoroutine())
}
