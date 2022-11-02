package main

import (
	"fmt"
	"sync"
)

// 线程安全的单例模式实现
// 含义:
// “懒汉式”实现是非线程安全的设计方式，也就是如果多个线程或者协程同时首次调用GetInstance()方法有概率导致多个实例被创建，则违背了单例的设计初衷。
// 那么在上面的基础上进行修改，可以利用Sync.Mutex进行加锁，保证线程安全

//定义锁
var lock sync.Mutex

type singleton struct{}

var XcInstance *singleton

func GetXcInstance() *singleton {
	lock.Lock()
	defer lock.Unlock()
	if XcInstance == nil {
		return new(singleton)
	} else {
		return XcInstance
	}
}

func (s *singleton) SomeThing() {
	fmt.Println("线程安全单例对象的某方法")
}
