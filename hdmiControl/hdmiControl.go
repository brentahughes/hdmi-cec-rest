package hdmiControl

import (
	"fmt"

	"github.com/chbmuc/cec"
)

type DeviceInfo struct {
	OsdName string `json:"osdName"`
	PhysicalAddress string `json:"physicalAddress"`
	VendorId uint64 `json:"vendorId"`
}

var hdmi *cec.Connection
var hdmiPort = 0

func init() {
	var err error

	hdmi, err = cec.Open("", "cec.go")
	if err != nil {
		fmt.Println(err)
	}
}

func SetPort(port int) {
	hdmiPort = port
}

func GetActiveDeviceList() map[string]cec.Device {
	return hdmi.List()
}

func GetPowerStatus() string {
	return hdmi.GetDevicePowerStatus(hdmiPort)
}

func Power(state string) {
	if state == "on" {
		hdmi.PowerOn(hdmiPort)
	} else {
		hdmi.Standby(hdmiPort)
	}
}

func SetVolume(state string) {
	switch state {
		case "up":
			hdmi.VolumeUp()
		case "down":
			hdmi.VolumeDown()
		case "mute":
			hdmi.Mute()
	}
}
