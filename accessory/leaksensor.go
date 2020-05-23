package accessory

import (
	"github.com/jgulick48/hc/service"
)

type LeakSensor struct {
	*Accessory
	LeakSensor *service.LeakSensor
}

// NewSwitch returns a switch which implements model.Switch.
func NewLeakSensor(info Info) *LeakSensor {
	acc := LeakSensor{}
	acc.Accessory = New(info, TypeSensor)
	acc.LeakSensor = service.NewLeakSensor()
	acc.AddService(acc.LeakSensor.Service)
	return &acc
}
