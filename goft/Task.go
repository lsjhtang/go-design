package goft

import "sync"

type TaskFunc func(params ...interface{})

var taskList chan *TaskExecutor
var One sync.Once //单例模式

func init()  {//初始化的时候执行队列任务
	chanList := getTaskList()
	go func() {
		for f := range chanList{
			f.Exec()
		}
	}()
}

func getTaskList() chan *TaskExecutor  {
	One.Do(func() {
		taskList = make(chan *TaskExecutor)
	})
	return taskList
}

type TaskExecutor struct {
	f TaskFunc
	p []interface{}
}

func(this TaskExecutor) Exec()  {
	this.f(this.p...)
}

func NewTaskExecutor(f TaskFunc, p []interface{}) *TaskExecutor {
	return &TaskExecutor{f: f, p: p}
}

func Task(f TaskFunc, params ...interface{}) {
	getTaskList() <- NewTaskExecutor(f,params) //把任务丢入到队列
}
