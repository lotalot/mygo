package main

import (
	"fmt"
	bag "mygo/bag"
	//bag "study/bag" //打开包bag,在前面可以添加标识符来引用
)

//————————————接口————————————

type jkname interface { //接口的声明
	f1() //方法，参数列表和返回值列表
	f2()
	//条件一 接口的方法与实现接口的类型方法格式一致
	//条件二 接口中所有方法都被实现
}

// 一个类型可以实现多个接口，多个类型可以实现相同的接口
// 接口可以进行嵌套
type T interface{} //空接口，可以保存所有类型的值
func jiekou() {
	var i jkname
	t, ok := i.(T) //接口断言，将接口变为另一个接口
	//i接口变量 T转换后的目标类型 t转换后的变量 接口未实现时防止宕机ok保存为false
	fmt.Println(t, ok)

	var v T = 10
	switch v.(type) { //switch语句可以判断空接口保存的类型
	case string:
		fmt.Println(v, "is string")
	case int:
		fmt.Println(v, "is int")
	default:
	}
}

//————————————包————————————

func b() {
	bag.Bag()
}

//————————————并发（线程）————————————

func run(c chan int) {

	a := <-c       //接收通道传递的数据10，每次只接收一个数据元素
	fmt.Println(a) //将接收的数据打印出来

	//c <- 0         //将数据放到通道中

}
func main3() {
	b()
	c := make(chan int) //创建一个通道（如果声明为 <-chan为只能接收数据，chan<-相反）
	go run(c)           //创建另一个线程，main函数默认为一个线程
	//线程之间通过通道进行传输，通道类似于队列，按顺序传输数据
	c <- 10 //把数据10通过通道进行传输
	//发送方或接受方未得到数据进行阻塞
	<-c //此语句将发生阻塞，直到接收数据然后忽略以实现并发

	c2 := make(chan interface{}, 20) //带缓冲的通道，缓冲区为20个元素，相当于建一个数据柜
	//缓冲区为满或空也会引起阻塞
	close(c2) //关闭通道
	_, ok := <-c2
	if !ok { //c2关闭时会令ok为false
		fmt.Println("c2 is closed")
	}
	//使用select实现多路复用，类似于Switch，那个case最先实现便执行，可以与时间函数并用
}
