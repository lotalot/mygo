package main

//实现并发上下文调用控制
import (
	"context"
	"fmt"
	"time"
)

func c1(c chan string) {
	b := <-c
	for i := 1; i != 5; i++ {
		fmt.Println(b)
		time.Sleep(1 * time.Second)
	}
	c <- "done"
}

func s1(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "k", "v") //创建子上下文(父辈，k,v)
	return child
}
func s2(ctx context.Context) context.Context {
	child := context.WithValue(ctx, "age", 18) //创建子上下文(父辈，k,v)
	//fmt.Println(ctx.Value("k"))
	return child //在context内封装数据,相当于一个map
}
func s3(ctx context.Context) {
	if ctx.Done() == nil {
		fmt.Println("V:", ctx.Value("k"))
		fmt.Println("age:", ctx.Value("age"))
	} else {
		fmt.Println("context is already closed")
	}
}
func mainc() { //调用Backgroud()或TODO()来作为初始父辈context来形成继承链
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//带超时的context,超时时间2s，2s后管道超时关闭
	//如果创建子ctx会按时间短的为主
	defer cancel() //取消带有超时的上下文
	//context.WithCancel(),如果调用cancel()则管道关闭

	grandpa := context.TODO() //value()的使用
	father := s1(grandpa)
	son := s2(father)
	s3(son)

	c := make(chan string)
	go c1(c)
	c <- "多线程开始"
	fmt.Println("这里是主线程")

	select { //select相当于通信switch 有case可以执行时随机执行，如果没有case则阻塞直到有数据到来
	//而每个case执行条件为可以进行读或写,若超时则执行超时的case
	case <-ctx.Done(): //Done()是只读类型的chan,由于只读其中没有数据，所以会阻塞到2s后超时
		err := ctx.Err() //输出超时
		fmt.Println(err)
		close(c)
	}

	d := <-c       //这么做是为了防止主线程结束后直接结束程序,让其发生阻塞等待管道传来数据
	fmt.Println(d) //由于上面2s超时后关闭了线程所以c1并未完成输出done

}
