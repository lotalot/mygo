package bag

import "fmt"

func Bag() { //将函数名称首字母大写以实现其他文件里的调用
	fmt.Println("open bag")
}

type MyStruct struct { //结构体首字母大写实现外部调用
	Ex    int //大写供外部使用
	priva int //小写内部使用
}

func init() { //函数不可被其他函数引用
	//init()函数会在程序执行前自动调用，根据包的引用关系，后引用的先调用
	fmt.Println("this is inint()")
}
