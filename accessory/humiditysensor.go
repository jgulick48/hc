package accessory

import (
	"github.com/jgulick48/hc/service"
)

type HumiditySensor struct {
	*Accessory
	HumiditySensor *service.HumiditySensor
}

// NewSwitch returns a switch which implements model.Switch.
func NewHumiditySensor(info Info) *HumiditySensor {
	acc := HumiditySensor{}
	acc.Accessory = New(info, TypeSensor)
	acc.HumiditySensor = service.NewHumiditySensor()
	acc.AddService(acc.HumiditySensor.Service)
	return &acc
}
