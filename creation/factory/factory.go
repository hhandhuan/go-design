package main

import "log"

// Fruit 抽奖类
type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CreateFruit() Fruit
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

// ========= 工厂模块  =========

// AppleFactory 具体的苹果工厂
type AppleFactory struct {
	AbstractFactory
}

func (f *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit

	//生产一个具体的苹果
	fruit = new(Apple)

	return fruit
}

// BananaFactory 具体的香蕉工厂
type BananaFactory struct {
	AbstractFactory
}

func (f *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit

	//生产一个具体的香蕉
	fruit = new(Banana)

	return fruit
}

// PearFactory 具体的梨工厂
type PearFactory struct {
	AbstractFactory
}

func (f *PearFactory) CreateFruit() Fruit {
	var fruit Fruit

	//生产一个具体的梨
	fruit = new(Pear)

	return fruit
}

func main() {

	//需求1：需要一个具体的苹果对象
	//1-先要一个具体的苹果工厂
	var appleFac AbstractFactory
	appleFac = new(AppleFactory)
	//2-生产相对应的具体水果
	var apple Fruit
	apple = appleFac.CreateFruit()

	apple.Show()

	//需求2：需要一个具体的香蕉对象
	//1-先要一个具体的香蕉工厂
	var bananaFac AbstractFactory
	bananaFac = new(BananaFactory)
	//2-生产相对应的具体水果
	var banana Fruit
	banana = bananaFac.CreateFruit()

	banana.Show()

	//需求3：需要一个具体的梨对象
	//1-先要一个具体的梨工厂
	var pearFac AbstractFactory
	pearFac = new(PearFactory)
	//2-生产相对应的具体水果
	var pear Fruit
	pear = pearFac.CreateFruit()

	pear.Show()
}
