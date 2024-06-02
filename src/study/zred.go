package main

import (
	"context"
	"errors"
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

func r2() { //集群
	ctx := context.Background()
	// Cluster:集群 Client:对象
	rdb1 := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})

	err := rdb1.ForEachShard(ctx, func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	}) //遍历每一个节点，只遍历主节点：ForEachMaster，只遍历从节点：ForEachSlave
	if err != nil {
		panic(err)
	}

	rdb2 := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master-name", //Failover:哨兵
		SentinelAddrs: []string{":9126", ":9127", ":9128"},
	})
	rdb2.Close()

	rdb3 := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{":6379"},
	})
	rdb3.Close()
	//通用客户端将其他客户端整合到一起，自动判断使用那种客户端
	//如果指定了 MasterName 选项，则返回 FailoverClient 哨兵客户端。
	//如果 Addrs 是 2 个以上的地址，则返回 ClusterClient 集群客户端。
	//其他情况，返回 Client 单节点客户端。
}

const maxRetries = 200 //代表发生错误时最大重试次数
func r3(key string) error { //管道和事务
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	cmds, _ := rdb.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 10; i++ { //使用管道来进行多次Get操作
			pipe.Get(ctx, fmt.Sprintf("k%d", i))
		} //Pipelined会自动执行Exec
		return nil
	})
	for _, cmd := range cmds {
		fmt.Println(cmd.(*redis.StringCmd).Val())
	}

	// 事务函数，官方示例
	txf := func(tx *redis.Tx) error { //此函数在此声明但并未使用
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}

		n++

		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n, 0) //0代表键值对的过期时间
			return nil
		})
		return err
	}

	for i := 0; i < maxRetries; i++ {
		err := rdb.Watch(ctx, txf, key)
		//使用watch来监听,如果此时key被修改则会返回一个错误并重新执行txf直到成功或达到次数上限
		//如果没被修改则正常执行事务函数
		if err == nil {
			// Success.
			return nil
		}
		if err == redis.TxFailedErr {
			// 乐观锁失败
			continue
		}
		return err
	}
	return errors.New("increment reached maximum number of retries")
}

func r4() { //发布  订阅
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{ //连接到数据库
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})
	err := rdb.Publish(ctx, "mychannel1", "payload").Err() //发布
	if err != nil {
		panic(err)
	}
	pubsub := rdb.Subscribe(ctx, "mychannel1") //订阅
	defer pubsub.Close()                       // 使用完毕，记得关闭

	ch := pubsub.Channel() //读取信息
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}
func mainr() {

	rdb := redis.NewClient(&redis.Options{ //连接到数据库
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	r1(rdb)
}
