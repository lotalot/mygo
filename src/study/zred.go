package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

// https://redis.uptrace.dev/zh/guide/go-redis.html

func r1(rdb *redis.Client) { //执行命令
	ctx := context.Background()

	val1, _ := rdb.Get(ctx, "key").Result()
	fmt.Println(val1)

	//你也可以分别访问值和错误：
	get := rdb.Get(ctx, "key")
	fmt.Println(get.Val(), get.Err())

	//使用Do 方法执行任意命令返回一个 cmd 指（包含已筛选或未导出的字段）
	val2, err := rdb.Do(ctx, "get", "key").Result()
	if err != nil {
		if err == redis.Nil { //判断key是否存在方法
			fmt.Println("key does not exists")
			return
		}
		panic(err)
	}
	fmt.Println(val2.(string)) //转为string
}

func r2() {

}

func mainr() {

	rdb := redis.NewClient(&redis.Options{ //连接到数据库
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	r1(rdb)
}
