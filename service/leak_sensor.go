// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/jgulick48/hc/characteristic"
)

const TypeLeakSensor = "83"

type LeakSensor struct {
	*Service

	LeakDetected *characteristic.LeakDetected
	WaterLevel   *characteristic.WaterLevel
}

func NewLeakSensor() *LeakSensor {
	svc := LeakSensor{}
	svc.Service = New(TypeLeakSensor)

	svc.LeakDetected = characteristic.NewLeakDetected()
	svc.AddCharacteristic(svc.LeakDetected.Characteristic)
	svc.WaterLevel = characteristic.NewWaterLevel()
	svc.AddCharacteristic(svc.WaterLevel.Characteristic)

	return &svc
}
