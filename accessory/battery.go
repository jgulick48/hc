package accessory

import (
	"github.com/jgulick48/hc/service"
)

type Battery struct {
	*Accessory
	Battery *service.BatteryService
}

// NewSwitch returns a switch which implements model.Switch.
func NewBattery(info Info) *Battery {
	acc := Battery{}
	acc.Accessory = New(info, TypeSensor)
	acc.Battery = service.NewBatteryService()
	acc.AddService(acc.Battery.Service)

	return &acc
}
