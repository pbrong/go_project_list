package main

import (
	"fmt"
	"sync"
)

func main() {
	//testAccumulator()
	//testBuildFileSuffix()
	//testDelayBidding()
	//forRange()
	//TestForRangeDelayBidding()
	//TestForRange()
	//TestForRangeGoRuntineDelayBidding()
	//TestForRange1()
	TestForRange2()
}

// 提供一个值, 每次调用函数会指定对值进行累加
func Accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}
func testAccumulator() {
	// 创建一个累加器, 初始值为1
	accumulator := Accumulate(1)
	// 累加1并打印
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator)
	// 创建一个累加器, 初始值为10
	accumulator2 := Accumulate(10)
	// 累加1并打印
	fmt.Println(accumulator2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2)
}

func BuildFileSuffix(suffix string) func(fileName string) string {
	return func(fileName string) string {
		// 闭包绑定环境变量suffix
		// 闭包自身输入fileName，组成完成文件名
		return fmt.Sprintf("%v.%v", fileName, suffix)
	}
}

func testBuildFileSuffix() {
	// 后缀名java
	javaFunc := BuildFileSuffix("java")
	fmt.Println(javaFunc("file1")) // file1.java
	fmt.Println(javaFunc("file2")) // file2.java
	// 后缀名golang
	golangFunc := BuildFileSuffix("golang")
	fmt.Println(golangFunc("file3")) // file3.golang
	fmt.Println(golangFunc("file4")) // file4.golang
}

func DelayBidding() func() {
	x := 1
	f := func() {
		fmt.Println(x)
	}
	x = 2
	return f
}

func testDelayBidding() {
	DelayBidding()() // 2
}

func TestForRange1() {
	list := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	for _, v := range list {
		wg.Add(1)
		// 创建一个新值v2，v2不会被复用
		v2 := v
		go func() {
			defer wg.Done()
			// 闭包引用v2
			fmt.Println(v2)
		}()
	}
	wg.Wait()
}

func TestForRange2() {
	list := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	for _, v := range list {
		wg.Add(1)
		// 将值拷贝传入闭包
		go func(v2 int) {
			defer wg.Done()
			fmt.Println(v2)
		}(v)
	}
	wg.Wait()
}

func TestForRangeDelayBidding() {
	list := []int{1, 2, 3, 4, 5}
	var funcList []func()
	for _, v := range list {
		f := func() {
			fmt.Println(v)
		}
		funcList = append(funcList, f)
	}
	for _, f := range funcList {
		f()
	}
}

func forRange() {
	list := []int{1, 2, 3, 4, 5}
	var i, v int
	for i, v = range list {
		fmt.Printf("%v-%v\n", &i, &v)
	}
}

func TestForRangeGoRuntineDelayBidding() {
	list := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	for _, v := range list {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(v)
		}()
	}
	wg.Wait()
}
