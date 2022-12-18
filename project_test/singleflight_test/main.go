package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"
	"time"
)

func main() {
	testDo()
	//testDoChan()
}

func testDoChan() {
	var flight singleflight.Group
	var errGroup errgroup.Group

	// 模拟并发获取数据缓存
	start := time.Now()
	for i := 0; i < 10; i++ {
		i := i
		errGroup.Go(func() error {
			fmt.Printf("协程%v准备获取缓存\n", i)
			ch := flight.DoChan("getCache", func() (interface{}, error) {
				// 模拟获取缓存操作
				fmt.Printf("协程%v正在读数据库获取缓存\n", i)
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("协程%v读取数据库获取缓存成功\n", i)
				return "mockCache", nil
			})
			res := <-ch
			if res.Err != nil {
				fmt.Printf("err = %v", res.Err)
				return res.Err
			}
			fmt.Printf("协程%v获取缓存成功, v = %v, shared = %v\n", i, res.Val, res.Shared)
			return nil
		})
	}
	if err := errGroup.Wait(); err != nil {
		fmt.Printf("errGroup wait err = %v", err)
	}
	milliseconds := time.Now().Sub(start).Milliseconds()
	fmt.Printf("执行耗时 = %v毫秒", milliseconds)
}

func testDo() {
	var flight singleflight.Group
	var errGroup errgroup.Group

	start := time.Now()
	// 模拟并发获取数据缓存
	for i := 0; i < 10; i++ {
		i := i
		errGroup.Go(func() error {
			fmt.Printf("协程%v准备获取缓存\n", i)
			v, err, shared := flight.Do("getCache", func() (interface{}, error) {
				// 模拟获取缓存操作
				fmt.Printf("协程%v正在读数据库获取缓存\n", i)
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("协程%v读取数据库获取缓存成功\n", i)
				return "mockCache", nil
			})
			if err != nil {
				fmt.Printf("err = %v", err)
				return err
			}
			fmt.Printf("协程%v获取缓存成功, v = %v, shared = %v\n", i, v, shared)
			return nil
		})
	}
	if err := errGroup.Wait(); err != nil {
		fmt.Printf("errGroup wait err = %v", err)
	}
	milliseconds := time.Now().Sub(start).Milliseconds()
	fmt.Printf("执行耗时 = %v毫秒", milliseconds)
}
