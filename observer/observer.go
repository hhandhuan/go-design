package observer

import "fmt"

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

type IObserver interface {
	Update(msg string)
}

type Subject struct {
	observers []IObserver
}

// Register 注册观察者
func (sub *Subject) Register(observer IObserver) {
	sub.observers = append(sub.observers, observer)
}

// Remove 删除观察者
func (sub *Subject) Remove(observer IObserver) {
	for i, ob := range sub.observers {
		if ob == observer {
			sub.observers = append(sub.observers[:i], sub.observers[i+1:]...)
		}
	}
}

// Notify 遍历调用观察者的update()方法进行通知，不关心其具体实现方式
func (sub *Subject) Notify(msg string) {
	for _, o := range sub.observers {
		o.Update(msg)
	}
}

type ObserverOne struct {}

func (observer ObserverOne) Update(msg string) {
	fmt.Println("Observer one update")
}

type ObserverTwo struct {}

func (observer ObserverTwo) Update(msg string) {
	fmt.Println("Observer two update")
}
