package main

import (
	addon "gateway-addon-golang"
	"yeelight_adapter/lib"
)

type YeeLightAdapter struct {
	*addon.Adapter
}

type YeeLightDevice struct {
	*addon.Device
}

type YeeLightProperty struct {
	addon.IProperty
}

func main() {

	var adapter = YeeLightAdapter{
		addon.NewAdapter("yeelight-adapter", "yeelight-adapter"),
	}
	adapter.StartPairing(3000)
}

func (adapter *YeeLightAdapter) StartPairing(timeout int) {
	adapter.Adapter.StartPairing(timeout)
	lights := lib.Discover()
	devs := NewDevice(adapter.Adapter, lights)
	for _, dev := range devs {
		adapter.HandleDeviceAdded(dev)
	}
}

func NewDevice(adapter *addon.Adapter, lights []lib.Light) []*addon.Device {
	var devs []*addon.Device
	for _, light := range lights {
		dev := addon.NewDevice(adapter)
		dev.Properties[]
		devs = append(devs, dev)
	}
}
