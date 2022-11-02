package main

import (
	"fmt"
	"log"
)

// “懒汉式”实现是非线程安全的设计方式，也就是如果多个线程或者协程同时首次调用GetInstance()方法有概率导致多个实例被创建，则违背了单例的设计初衷。
// 那么在上面的基础上进行修改，可以利用Sync.Mutex进行加锁，保证线程安全。

type lhsSingleton struct{}

var lhsInstance *lhsSingleton

func GetLhsInstance() *lhsSingleton {
	if lhsInstance != nil {
		return lhsInstance
	}
	log.Println("new GetLhsInstance")
	lhsInstance = new(lhsSingleton)
	return lhsInstance
}

func (s *lhsSingleton) SomeThing() {
	fmt.Println("懒汉式单例对象的某方法")
}
