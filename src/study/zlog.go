package main

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

/*
日志的七个级别:
PanicLevel  // 会抛一个异常
FatalLevel  // 打印日志之后就会退出
ErrorLevel
WarnLevel
InfoLevel
DebugLevel
TraceLevel  // 低级别
*/
func l1() {
	logrus.SetLevel(logrus.DebugLevel) //设置日志等级,输出debug以上的级别
	logrus.SetReportCaller(true)       //打印出日志所在的文件名和行号
	logrus.Debugln("Debugln")
	logrus.Infoln("Infoln")
	logrus.Warnln("Warnln")
	logrus.Errorln("Errorln")
	logrus.Println("Println") //默认的日志输出等级是 info

	log := logrus.WithField("k", "v")                             //在输出中添加一些字段
	l := logrus.WithFields(logrus.Fields{"k1": "v1", "k2": "v2"}) //添加多个字段
	//接受一个logrus.Fields类型的参数，其底层实际上为MAP
	l.Info("msg")

	logrus.SetFormatter(&nested.Formatter{
		//在这里设置日志格式，比如显示颜色，隐藏键等
		HideKeys: true,
	})
	log.Errorf("h1")
} //可以通过io操作将日志输出到文件

type Hook interface { // 使用钩子hook,需要实现该接口
	Levels() []logrus.Level
	Fire(*logrus.Entry) error
}
type MyHook struct {
}

func (hook *MyHook) Levels() []logrus.Level {
	return logrus.AllLevels //无论日志级别是什么，这个钩子都会被触发。
}
func (hook *MyHook) Fire(entry *logrus.Entry) error { //接收一个当前日志条目
	entry.Data["sth"] = "hello"
	return nil
} //接口中两个函数的具体实现
func mainl() {
	l1()

	logrus.AddHook(&MyHook{}) //添加钩子
	logrus.Info("hook")
}
