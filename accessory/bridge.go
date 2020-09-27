package accessory

import (
	"github.com/jgulick48/hc/service"
)

type Bridge struct {
	*Accessory
	Battery *service.BatteryService
}

// NewBridge returns a bridge which implements model.Bridge.
func NewBridge(info Info) *Bridge {
	acc := Bridge{}
	acc.Accessory = New(info, TypeBridge)
	acc.Battery = service.NewBatteryService()
	acc.AddService(acc.Battery.Service)

	return &acc
}
