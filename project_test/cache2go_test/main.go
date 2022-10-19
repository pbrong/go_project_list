package main

import (
	"github.com/muesli/cache2go"
	"log"
	"strings"
	"time"
)

type Item struct {
	Name   string `json:"name"`
	Prices int64  `json:"prices"`
	Stocks int64  `json:"stocks"`
}

func main() {
	//basicOpTest()
	//callBackTest()
	dataLoaderTest()
}

func dataLoaderTest() {
	// 初始化itemCache本地缓存
	redisItemCache := cache2go.Cache("redisItemCache")

	// 设置自定义的cache加载逻辑
	redisItemCache.SetDataLoader(func(key interface{}, args ...interface{}) *cache2go.CacheItem {
		// 如果是redis开头的key，先从redis中获取
		if strings.HasPrefix(key.(string), "redis") {
			return cache2go.NewCacheItem(key, 0, Item{
				Name: "redis_item",
			})
		}
		return nil
	})

	// 写入一条数据
	redisItemCache.Add("item1", 0, Item{
		Name: "item1",
	})

	item1, _ := redisItemCache.Value("item1")
	log.Printf("item1 = %#v", item1)

	redisItem, _ := redisItemCache.Value("redis_item")
	log.Printf("redisItem = %#v", redisItem)
}

func callBackTest() {
	// 初始化itemCache本地缓存
	itemCache := cache2go.Cache("itemCache")

	// 设置各操作回调函数
	item := itemCache.Add("expire_item", 1*time.Second, Item{
		Name:   "expire_item",
		Prices: 1,
		Stocks: 1,
	})
	item.AddAboutToExpireCallback(func(item interface{}) {
		log.Printf("expired callback, item = %#v", item)
	})
	itemCache.AddAddedItemCallback(func(item *cache2go.CacheItem) {
		log.Printf("added callback, item = %#v", item)
	})
	itemCache.AddAboutToDeleteItemCallback(func(item *cache2go.CacheItem) {
		log.Printf("deleted callback, item = %#v", item)
	})
	// 执行基本操作
	basicOpTest()
}

func basicOpTest() {
	// 初始化itemCache本地缓存
	itemCache := cache2go.Cache("itemCache")
	item := &Item{
		Name:   "MacBookPro",
		Prices: 10000,
		Stocks: 1,
	}

	// 添加item1缓存，过期时间为5秒钟
	itemCache.Add("item1", 5*time.Second, item)

	// 读取item1缓存
	if v, err := itemCache.Value("item1"); err != nil {
		log.Printf("item1 err = %v", err)
	} else {
		log.Printf("读取item1缓存：%#v", v.Data())
	}

	// 睡眠6s后读取
	time.Sleep(6 * time.Second)
	if v, err := itemCache.Value("item1"); err != nil {
		log.Printf("item1 err = %v", err)
	} else {
		log.Printf("6s后读取item1缓存：%#v", v.Data())
	}

	// 添加item2，不设置过期时间
	itemCache.Add("item2", 0, item)

	// 读取item2缓存
	if v, err := itemCache.Value("item2"); err != nil {
		log.Printf("item2 err = %v", err)
	} else {
		log.Printf("读取item2缓存：%#v", v.Data())
	}

	// 删除掉item2缓存
	itemCache.Delete("item2")

	// 再读取item2缓存
	if v, err := itemCache.Value("item2"); err != nil {
		log.Printf("item2 err = %v", err)
	} else {
		log.Printf("读取item2缓存：%#v", v.Data())
	}

	// 添加item3缓存，并删除所有缓存
	itemCache.Add("item3", 0, item)
	itemCache.Flush()

	// 读取item3缓存
	if v, err := itemCache.Value("item3"); err != nil {
		log.Printf("item3 err = %v", err)
	} else {
		log.Printf("读取item3缓存：%#v", v.Data())
	}
}
