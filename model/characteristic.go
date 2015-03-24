package model

// A Characteristic is identifiable and has a (observeable) value.
type Characteristic interface {
	Compareable

	// GetID returns the characteristic's id
	GetID() int64

	// GetValue returns the raw value
	GetValue() interface{}

	// SetValueFromRemote sets the value
	// Only call this method when a client (e.g. iOS device) changes the value
	// Otherwise use the provided setter methods ( e.g. `switch.SetOn(true)`)
	SetValueFromRemote(interface{})

	// SetEventsEnabled dis-/enables events for the characteristic
	SetEventsEnabled(enable bool)

	// EventsEnabled returns true when events for this characteristic are enabled, otherwise false
	EventsEnabled() bool
}
