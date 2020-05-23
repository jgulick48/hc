package accessory

import (
	"github.com/jgulick48/hc/characteristic"
	"github.com/jgulick48/hc/service"
)

type LevelSensor struct {
	*Accessory
	Source *service.InputSource
}

// NewSwitch returns a switch which implements model.Switch.
func NewLevelSensor(info Info) *LevelSensor {
	acc := LevelSensor{}
	acc.Accessory = New(info, TypeSensor)
	acc.Source = service.NewInputSource()
	acc.AddService(acc.Source.Service)
	level := characteristic.NewWaterLevel()
	acc.Source.Service.AddCharacteristic(level.Characteristic)

	return &acc
}
