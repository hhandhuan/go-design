package adapter

import "fmt"

type Client struct{}

func (c *Client)InsertLightningConnectorIntoComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

type computer interface {
	InsertIntoLightningPort()
}


// Mac 实现 computer 接口
type Mac struct {}
func (m *Mac) InsertIntoUSBPort()  {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

// Windows 实现 computer 接口
type Windows struct {}
func (w *Windows) InsertIntoUSBPort()  {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

type WindowsAdapter struct {
	WindowMachine *Windows
}
func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.WindowMachine.InsertIntoUSBPort()
}

type MacAdapter struct {
	MacMachine *Mac
}
func (m *MacAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	m.MacMachine.InsertIntoUSBPort()
}