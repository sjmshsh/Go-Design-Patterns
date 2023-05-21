package main

import "fmt"

/*
多个观察者同时监听同一个主题对象
*/

// Observer 1. 定义观察者接口
type Observer interface {
	Update(data interface{})
}

// Subject 2. 定义主题对象接口
type Subject interface {
	Attach(observer Observer) // 添加观察者
	Detach(observer Observer) // 解除观察者
	Notify(data interface{})  // 通知观察者
}

// ConcreteSubject 3. 实现具体的主题
type ConcreteSubject struct {
	observers []Observer
	data      interface{}
}

func (s *ConcreteSubject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) Detach(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) Notify(data interface{}) {
	s.data = data
	for _, observer := range s.observers {
		observer.Update(data)
	}
}

// ConcreteObserver 4. 实现具体的观察者
type ConcreteObserver struct {
	id int
}

// Update 数据更新之后观察者应该调用的方法
func (o *ConcreteObserver) Update(data interface{}) {
	fmt.Printf("Observer: %d received: %v\n", o.id, data)
}

// 5 使用观察者模式
func main() {
	subject := &ConcreteSubject{}
	observe1 := &ConcreteObserver{id: 1}
	observe2 := &ConcreteObserver{id: 2}

	subject.Attach(observe1)
	subject.Attach(observe2)

	subject.Notify("data 1")

	subject.Detach(observe1)

	subject.Notify("data 2")
}
