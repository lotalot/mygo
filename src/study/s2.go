package main

import (
	"errors"
	"fmt"
)

//————————————函数————————————

func hanshu(a, b int) (c, d int) { //func 函数名（参数列表）（返回参数列表）
	//c, d为命名的返回值
	c = 1
	d = 2
	return //直接return将按照列表进行返回，此函数为命名返回，只写类型为类型返回，不可同时使用
}
func bibao(name string) func() (string, int) { //返回参数为一个匿名函数
	//闭包的使用演示
	hp := 150
	return func() (string, int) {
		return name, hp //hp的值被引用到匿名函数中形成闭包
	}
}

var err = errors.New("发现一个错误")

func kebian(v ...int) error { //可变参数（固定参数列表，v ... T）（返回参数）{
	//v相当于T类型的数组，T为可变参数的类似，interface{}时为任意类型
	//多个可变参数中传递参数在v后加...进行完整传递

	//延迟执行语句，先延迟的后释放，打印2 1，此语句会延迟到函数结束时执行
	defer fmt.Println(1)
	defer fmt.Println(2)

	if 0 > 1 {
		return err //自定义一个错误，如果发生返回上面的语句，没有需返回nil
	}
	return nil
}

func main2() {
	e, f := hanshu(1, 1)
	player := bibao("sb") //创建一个角色，相当于一个模版的使用
	name, hp := player()

	defer fmt.Println("宕机后执行1")
	defer fmt.Println("2")
	//panic("宕机") //用于触发宕机，recover宕机恢复机制类似于try/catch
	fmt.Println(e, f, name, hp)
}

//————————————结构体————————————

type JieGou struct { //结构体的声明，struct为类型名的类型
	x, y int
	z    *JieGou `ne:"sth"` //此为验证器，为结构等添加限制条件
}

func a(i int) *JieGou { //由于go语言中没有构造函数和继承功能，所以要进行模拟
	return &JieGou{ //模拟构造函数
		x: i,
	}
}
func (j *JieGou) jsq() { //接收器，每个方法只能有一个，类似于this指针
	//func（建议结构小写字母头 结构或其指针）函数名（）...
	//指针形接收器无需复制适合大对象，非指针适合小对象
}

type Son struct { //模拟派生，进行嵌入然后再通过上面方法发现构造
	JieGou //结构体内嵌，使Son可以直接调用Jiegou，初始化使将类型名作为字段名，如下面
	a      int
}

func JG() {
	var jg1 JieGou //实例化结构体
	jg1.y = 10
	jg2 := new(JieGou) //创建成指针类型,其为指针类型
	jg2.x = 10
	jg3 := &JieGou{ //取地址实例化，使用最广
		x: 10, //使用键值对来填充结构体
		y: 10,
		z: &JieGou{
			x: 20,
		},
	}
	jg3.x = 10
	//使用键值对填充结构体时，将“ x: ”省略，必须初始化所有，声明顺序一致，不能混用
	//匿名结构体与函数类似但较少使用

	a := Son{
		JieGou: JieGou{
			x: 1,
			y: 2,
		},
		a: 3,
	}
	fmt.Println(a.x)
}
