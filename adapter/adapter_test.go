package adapter

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	client := &Client{}

	windowsMachineAdapter := &WindowsAdapter{
		WindowMachine: &Windows{},
	}
	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)

	macMachineAdapter := &MacAdapter{
		MacMachine: &Mac{},
	}
	client.InsertLightningConnectorIntoComputer(macMachineAdapter)

	t.Log("success")
}