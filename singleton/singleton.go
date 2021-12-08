package singleton

import (
	"sync"
)

var once = &sync.Once{}

type single struct {
	Name string
}

var instance *single

func GetInstance() *single {
	// 防止多线程程序并发读取
	// 当然也可以使用:
	// lock := sync.Mutex | lock.lock()  | defer lock.unlock()
	if instance == nil {
		once.Do(func() {
			instance = &single{}
		})
	}
	return instance
}
