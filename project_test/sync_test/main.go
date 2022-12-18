package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

type token struct{}

type Pool struct {
	sem chan token
}

func NewPoolWithLimit(limit int) Pool {
	if limit <= 0 {
		return Pool{}
	}
	return Pool{
		sem: make(chan token, limit),
	}
}

func (p *Pool) RunFunc(f func()) {
	if p.sem != nil {
		p.sem <- token{}
	}
	go func() {
		defer func() {
			<-p.sem
		}()
		f()
	}()
}

type Person struct {
	Name  string
	Age   int
	Child *Person
}

func main() {
	testWithContextCancel()
	pool := NewPoolWithLimit(10)
	for i := 0; i < 15; i++ {
		i := i
		pool.RunFunc(func() {
			time.Sleep(100 * time.Millisecond)
			log.Printf("协程%v执行任务完成", i)
		})
	}
}

func testSetLimitTryGo() {
	var group errgroup.Group
	// 设置10个协程
	group.SetLimit(10)
	// 启动11个任务
	for i := 1; i <= 11; i++ {
		i := i
		fn := func() error {
			time.Sleep(100 * time.Millisecond)
			return nil
		}
		if ok := group.TryGo(fn); !ok {
			log.Printf("tryGo false, goroutine no = %v", i)
		} else {
			log.Printf("tryGo true, goroutine no = %v", i)
		}
	}
	group.Wait()
	log.Printf("group task finished")
}

func testWithContextCancel() {
	group, ctx := errgroup.WithContext(context.Background())
	// 设置10个协程
	group.SetLimit(10)
	// 启动10个任务，在第5个任务生成错误
	for i := 1; i <= 10; i++ {
		i := i
		fn := func() error {
			time.Sleep(100 * time.Millisecond)
			if i == 5 {
				return errors.New("task 5 is fail")
			}
			// 当某个任务错误时，终止当前任务
			select {
			case <-ctx.Done():
				if errors.Is(ctx.Err(), context.Canceled) {
					log.Printf("ctx Cancel, all task cancel, goroutine no = %v", i)
				} else {
					log.Printf("ctx Done, all task done, goroutine no = %v", i)
				}
			default:
				log.Printf("task Done, goroutine no = %v", i)
			}
			return nil
		}
		if ok := group.TryGo(fn); !ok {
			log.Printf("tryGo false, goroutine no = %v", i)
		} else {
			log.Printf("tryGo true, goroutine no = %v", i)
		}
	}
	if err := group.Wait(); err != nil {
		log.Printf("group.Wait err = %v", err)
		return
	}
	log.Printf("group task finished")
}
