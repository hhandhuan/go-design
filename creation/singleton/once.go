package main

import (
	"fmt"
	"sync"
)

// Once 单例模式

var once sync.Once

type onceSingleton struct{}

var instance *onceSingleton

func GetOnceInstance() *onceSingleton {
	once.Do(func() {
		instance = new(onceSingleton)
	})
	return instance
}

func (s *onceSingleton) SomeThing() {
	fmt.Println("单例对象的某方法")
}
