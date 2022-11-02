package main

import "log"

/**

练习：

设计一个电脑主板架构，电脑包括（显卡，内存，CPU）3个固定的插口，显卡具有显示功能（display，功能实现只要打印出意义即可），内存具有存储功能（storage），cpu具有计算功能（calculate）。
现有Intel厂商，nvidia厂商，Kingston厂商，均会生产以上三种硬件。
要求组装两台电脑，
	1台（Intel的CPU，Intel的显卡，Intel的内存）
	1台（Intel的CPU， nvidia的显卡，Kingston的内存）
*/

type AbstractGPU interface {
	Display()
}

type AbstractRAM interface {
	Storage()
}

type AbstractCPU interface {
	Calculate()
}

type AbstractComputerFactory interface {
	CreateGPU() AbstractGPU
	CreateRAM() AbstractRAM
	CreateCPU() AbstractCPU
}

type IntelGPU struct{}

func (i IntelGPU) Display() {
	log.Println("intel GPU display.")
}

type IntelRAW struct{}

func (i IntelRAW) Storage() {
	log.Println("intel RAM storage.")
}

type IntelCPU struct{}

func (i IntelCPU) Calculate() {
	log.Println("intel CUP calculate.")
}

type NvidiaGPU struct{}

func (i NvidiaGPU) Display() {
	log.Println("nvidia GPU display.")
}

type NvidiaRAM struct{}

func (i NvidiaRAM) Storage() {
	log.Println("nvidia RAM storage.")
}

type NvidiaCPU struct{}

func (i NvidiaCPU) Calculate() {
	log.Println("nvidia CPU calculate.")
}

type KingstonGPU struct{}

func (i KingstonGPU) Display() {
	log.Println("Kingston GPU display.")
}

type KingstonRAM struct{}

func (i KingstonRAM) Storage() {
	log.Println("Kingston RAM storage.")
}

type KingstonCPU struct{}

func (i KingstonCPU) Calculate() {
	log.Println("Kingston CPU calculate.")
}

type IntelComputerFactory struct{}

func (f IntelComputerFactory) CreateGPU() AbstractGPU {
	var gpu AbstractGPU
	gpu = new(IntelGPU)
	return gpu
}

func (f IntelComputerFactory) CreateRAM() AbstractRAM {
	var ram AbstractRAM
	ram = new(IntelRAW)
	return ram
}

func (f IntelComputerFactory) CreateCPU() AbstractCPU {
	var cpu AbstractCPU
	cpu = new(IntelCPU)
	return cpu
}

type NvidiaComputerFactory struct{}

func (f NvidiaComputerFactory) CreateGPU() AbstractGPU {
	var gpu AbstractGPU
	gpu = new(NvidiaGPU)
	return gpu
}

func (f NvidiaComputerFactory) CreateRAM() AbstractRAM {
	var ram AbstractRAM
	ram = new(NvidiaRAM)
	return ram
}

func (f NvidiaComputerFactory) CreateCPU() AbstractCPU {
	var cpu AbstractCPU
	cpu = new(NvidiaCPU)
	return cpu
}

type KingstonComputerFactory struct{}

func (f KingstonComputerFactory) CreateGPU() AbstractGPU {
	var gpu AbstractGPU
	gpu = new(KingstonGPU)
	return gpu
}

func (f KingstonComputerFactory) CreateRAM() AbstractRAM {
	var ram AbstractRAM
	ram = new(KingstonRAM)
	return ram
}

func (f KingstonComputerFactory) CreateCPU() AbstractCPU {
	var cpu AbstractCPU
	cpu = new(KingstonCPU)
	return cpu
}

type Computer struct {
	GPU AbstractGPU
	CPU AbstractCPU
	RAM AbstractRAM
}

func (c Computer) Show() {
	c.GPU.Display()
	c.CPU.Calculate()
	c.RAM.Storage()
}

func Test() {

	intelFactory := new(IntelComputerFactory)

	kingstonFactory := new(IntelComputerFactory)

	nvidiaFactory := new(NvidiaComputerFactory)

	c1 := Computer{
		GPU: intelFactory.CreateGPU(),
		RAM: intelFactory.CreateRAM(),
		CPU: intelFactory.CreateCPU(),
	}

	c2 := Computer{
		GPU: nvidiaFactory.CreateGPU(),
		RAM: kingstonFactory.CreateRAM(),
		CPU: intelFactory.CreateCPU(),
	}

	c1.Show()
	c2.Show()
}
