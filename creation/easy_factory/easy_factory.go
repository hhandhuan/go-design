package main

import "log"

// Fruit 抽奖类
type Fruit interface {
	Show()
}

type Apple struct{}

func (s Apple) Show() {
	log.Println("apple")
}

type Banana struct{}

func (s Banana) Show() {
	log.Println("banana")
}

type Pear struct{}

func (s Pear) Show() {
	log.Println("banana")
}

type Factory struct{}

func (f *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit
	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}

	return fruit
}

func main() {
	factory := new(Factory)

	factory.CreateFruit("apple").Show()

	factory.CreateFruit("banana").Show()

	factory.CreateFruit("pear").Show()
}
