// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/jgulick48/hc/characteristic"
)

const TypeThermostat = "4A"

type Thermostat struct {
	*Service

	CurrentHeatingCoolingState  *characteristic.CurrentHeatingCoolingState
	TargetHeatingCoolingState   *characteristic.TargetHeatingCoolingState
	CurrentTemperature          *characteristic.CurrentTemperature
	TargetTemperature           *characteristic.TargetTemperature
	TemperatureDisplayUnits     *characteristic.TemperatureDisplayUnits
	CoolingThresholdTemperature *characteristic.CoolingThresholdTemperature
	HeatingThresholdTemperature *characteristic.HeatingThresholdTemperature
}

func NewThermostat() *Thermostat {
	svc := Thermostat{}
	svc.Service = New(TypeThermostat)

	svc.CurrentHeatingCoolingState = characteristic.NewCurrentHeatingCoolingState()
	svc.AddCharacteristic(svc.CurrentHeatingCoolingState.Characteristic)

	svc.TargetHeatingCoolingState = characteristic.NewTargetHeatingCoolingState()
	svc.AddCharacteristic(svc.TargetHeatingCoolingState.Characteristic)

	svc.CurrentTemperature = characteristic.NewCurrentTemperature()
	svc.AddCharacteristic(svc.CurrentTemperature.Characteristic)

	svc.TargetTemperature = characteristic.NewTargetTemperature()
	svc.AddCharacteristic(svc.TargetTemperature.Characteristic)

	svc.TemperatureDisplayUnits = characteristic.NewTemperatureDisplayUnits()
	svc.AddCharacteristic(svc.TemperatureDisplayUnits.Characteristic)

	svc.CoolingThresholdTemperature = characteristic.NewCoolingThresholdTemperature()
	svc.AddCharacteristic(svc.CoolingThresholdTemperature.Characteristic)

	svc.HeatingThresholdTemperature = characteristic.NewHeatingThresholdTemperature()
	svc.AddCharacteristic(svc.HeatingThresholdTemperature.Characteristic)

	return &svc
}
