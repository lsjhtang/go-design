package goft

import (
	"github.com/robfig/cron/v3"
	"sync"
)

type TaskFunc func(params ...interface{})

var (
	taskList chan *TaskExecutor
	one      sync.Once  //队列任务单例模式
	oneCron  sync.Once  //定时任务单例模式
	taskCron *cron.Cron //定时任务
)

func init() { //初始化的时候执行队列任务
	chanList := getTaskList()
	go func() {
		for f := range chanList {
			DoTask(f)
		}
	}()
}

func DoTask(f *TaskExecutor) {
	go func() {
		defer func() { //执行回调函数
			f.callback()
		}()

		f.Exec()
	}()
}

func getCronTask() *cron.Cron {
	oneCron.Do(func() {
		taskCron = cron.New(cron.WithSeconds())
	})
	return taskCron
}

func getTaskList() chan *TaskExecutor {
	one.Do(func() {
		taskList = make(chan *TaskExecutor)
	})
	return taskList
}

type TaskExecutor struct {
	f        TaskFunc
	callback func()
	p        []interface{}
}

func NewTaskExecutor(f TaskFunc, callback func(), p []interface{}) *TaskExecutor {
	return &TaskExecutor{f: f, callback: callback, p: p}
}

func (this TaskExecutor) Exec() {
	this.f(this.p...)
}

func Task(f TaskFunc, callback func(), params ...interface{}) {
	if f == nil {
		return
	}
	go func() {
		getTaskList() <- NewTaskExecutor(f, callback, params) //把任务丢入到队列
	}()
}
