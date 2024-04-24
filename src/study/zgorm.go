package main

//数据表对应结构体
//数据行对应结构体实例
//字段对应结构体字段

//gorm进行安全操作可以进行增加等，不会进行删除或改变
import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Test struct { //可以在里面直接内嵌别的结构体
	ID     uint   `gorm:"unique;not null"` //默认作为主键
	Name   string `gorm:"default:'名字'"`    //设置默认值，如果传入空(0)值直接设定为默认值
	Gender string //1.可以传new(string)（空字符串）来避免使用默认值
	//gorm.Model    //相当于主键和增删改操作时间的集合包
}

func maino() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:11575@(127.0.0.1:3306)/myku?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Test{}) //如果表不存在则自动创建
	//创建
	t1 := Test{1, "乔岳", "man"}
	db.Create(&t1) //创建记录，如果存在则先删除再创建，创建后Test自动变为tests

	//查询
	var t Test                                  //第一条数据实例
	db.First(&t, 1)                             //查询第一条记录，将查询到的数据用指针存到t中
	db.Debug().Where("name = ?", "乔岳").Find(&t) //Debug()函数可以查询原sql语句
	fmt.Println(t)
	db.Raw("select name from tests") //用于执行原sql语句

	//更新
	db.Model(&t).Update("gender", "feman")
	//删除
	db.Delete(&t)

	//链式操作，Find()等函数为立即执行函数，会直接翻译成sql语句，而where()等并不会，
	//可以将多个where查询付给一条语句然后再进行执行翻译为sql语句

}
