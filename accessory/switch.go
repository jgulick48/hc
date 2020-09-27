package accessory

import (
	"github.com/jgulick48/hc/service"
)

type Switch struct {
	*Accessory
	Switch  *service.Switch
	Battery *service.BatteryService
}

// NewSwitch returns a switch which implements model.Switch.
func NewSwitch(info Info) *Switch {
	acc := Switch{}
	acc.Accessory = New(info, TypeSwitch)
	acc.Switch = service.NewSwitch()
	acc.AddService(acc.Switch.Service)
	acc.Battery = service.NewBatteryService()
	acc.AddService(acc.Battery.Service)

	return &acc
}
