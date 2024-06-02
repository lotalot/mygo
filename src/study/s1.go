//初始化：go mod init xxx  xxx为你起项目的名字，外部包引用所使用的名称

package main //打开包main，主入口
import (
	"container/list"
	"fmt" //该包提供格式化，输入输出的函数
	"sort"
)

func main1() { //执行入口 main函数，{ 要求与函数同一行

	//go补充: 1 严格区分大小，2 每行自动加 ;，3 一行就写一条语句，4 如果声明则就需要使用

	//————————————语法————————————

	var a int  //声明变量，自动初始化
	var d = 10 //自动识别类型
	e := 100   //短变量声明
	var (      //快速声明
		b string
		c []float32
	)
	d, e = e, d //多重赋值，从左到右交换值
	//匿名变量，如果不需要接收值用 _ 替换变量

	var a1 int8 //相当于short和long形 uint代表无符号整型
	var a2 int64
	var b1 byte = 'a' //ASCII字符和复合字符
	var b2 rune = '你'

	const f = `第一行
	第二行
	\n\t
	` //换行字符，在之间的转义字符无效 f为常量形字符串
	g := &e //将e地址付给g，h取地址g
	h := *g
	i := new(string) //创建指针
	*i = "sb"
	//ASCII字符串长度用 len( )函数, 复合用utf8.RuneCountInString( )
	//复合字符串遍历使用 for range
	fmt.Println(a, b, c, d, h, f, a1, a2, b1, b2)

	//————————————容器————————————

	var team = [...]string{"sb", "op"} //创建数组并自动确定大小
	for k, v := range team {
		fmt.Println(k, v) //打印键值及数组
	}
	var qp []string                                 //声明一个切片，和创建类似,[]中有明确大小既为数组
	var cheeses = make([]string, 3)                 //创建一个大小为3的切片，类似vector
	cheeses = append(cheeses, "sb")                 //向容器中添加元素，可以添加多个值
	cheeses = append(cheeses, qp...)                //将qp切片整个添加到后面
	copy(cheeses, qp)                               //复制qp到cheeses1
	cheeses = append(cheeses[:1], cheeses[1+1:]...) //切片删除一位（该编号为1）的操作（包括前不包括后）
	fmt.Println(team[1:2])                          //切，打印数组(或切片)1到2之间不包括2的内容，空为到一端
	//当切片 a = b 操作时会让a和b指向同一块内存，改变a也会改变b，所以要使用copy函数，同时使用append()函数时
	//若超出切片容积时，会自动扩容，但会重新分配内存和地址，所以不要在循环中append，否则会频繁分配内存

	scene := make(map[string]int) //创建一个以String形为键 int为值的映射
	//map可以通过将键值设置为结构体的方式实现多键值映射
	scene["op"] = 666         //写入数据，若打印键输出值，若无则输出默认值
	for k, v := range scene { //for range语句来执行遍历，k为键值，v为数据不需要打印时用_
		fmt.Println(k, v)
	}
	//由于键值无法按序打印，需要将数据复制到切片中进行排序再输出该切片
	qp = append(qp, "op") //在上面，String形切片
	sort.Strings(qp)
	delete(scene, "op") //删除一组键值，由于go会自动清理垃圾不用手动清空
	//并发时使用sync.Map

	l := list.New()                              //初始化一个链表，不限制元素类型
	la := l.PushFront("one")                     //队前插入元素
	l.PushBack(1)                                //队后插入...
	l.Remove(la)                                 //删除la， （）内为操作名(实际为返回的地址型)
	for i := l.Front(); i != nil; i = i.Next() { //遍历列表方法
		println(i.Value)
	}

	//————————————流程————————————

	//if语句除无括号外，}else{都在同一行
	//for range各种打印：K键 V值
	//获得数组的下标和每一位
	//获得字符串的下标和每一位
	//获得map的键和值

	//switch语句不需要使用break跳出，case为独立部分
	switch a {
	case 2, 3:
		fmt.Println("ssll")
	}

	//goto break continue与其他类似
}
